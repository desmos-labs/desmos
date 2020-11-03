package app

import (
	"io"
	"os"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authvesting "github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/desmos-labs/desmos/x/fees"
	"github.com/desmos-labs/desmos/x/fees/ante"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"

	"github.com/desmos-labs/desmos/x/relationships"

	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"

	feesKeeper "github.com/desmos-labs/desmos/x/fees/keeper"
	feesTypes "github.com/desmos-labs/desmos/x/fees/types"
	"github.com/desmos-labs/desmos/x/magpie"
	magpieKeeper "github.com/desmos-labs/desmos/x/magpie/keeper"
	magpieTypes "github.com/desmos-labs/desmos/x/magpie/types"
	"github.com/desmos-labs/desmos/x/posts"
	postsKeeper "github.com/desmos-labs/desmos/x/posts/keeper"
	postsTypes "github.com/desmos-labs/desmos/x/posts/types"
	"github.com/desmos-labs/desmos/x/profiles"
	profilesKeeper "github.com/desmos-labs/desmos/x/profiles/keeper"
	profilesTypes "github.com/desmos-labs/desmos/x/profiles/types"
	relationshipsKeeper "github.com/desmos-labs/desmos/x/relationships/keeper"
	relationshipsTypes "github.com/desmos-labs/desmos/x/relationships/types"
	"github.com/desmos-labs/desmos/x/reports"
	reportsKeeper "github.com/desmos-labs/desmos/x/reports/keeper"
	reportsTypes "github.com/desmos-labs/desmos/x/reports/types"
)

const (
	appName          = "desmos"
	Bech32MainPrefix = "desmos"
)

var (
	// DefaultCLIHome represents the default home directory for the application CLI
	DefaultCLIHome = os.ExpandEnv("$HOME/.desmoscli")

	// DefaultNodeHome sets the folder where the application data and configuration will be stored
	DefaultNodeHome = os.ExpandEnv("$HOME/.desmosd")

	// ModuleBasics is in charge of setting up basic module elements
	ModuleBasics = module.NewBasicManager(
		genutil.AppModuleBasic{},
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(
			paramsclient.ProposalHandler,
			distr.ProposalHandler,
			upgradeclient.ProposalHandler,
		),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		supply.AppModuleBasic{},
		evidence.AppModuleBasic{},

		// Custom modules
		fees.AppModuleBasic{},
		magpie.AppModuleBasic{},
		posts.AppModuleBasic{},
		profiles.AppModuleBasic{},
		reports.AppModuleBasic{},
		relationships.AppModuleBasic{},
	)

	// Module account permissions
	maccPerms = map[string][]string{
		auth.FeeCollectorName:     nil,
		distr.ModuleName:          nil,
		staking.BondedPoolName:    {supply.Burner, supply.Staking},
		staking.NotBondedPoolName: {supply.Burner, supply.Staking},
		gov.ModuleName:            {supply.Burner},
	}

	// module accounts that are allowed to receive tokens
	allowedReceivingModAcc = map[string]bool{
		distr.ModuleName: true,
	}
)

// MakeCodec generates the necessary codecs for Amino
func MakeCodec() *codec.Codec {
	var cdc = codec.New()

	ModuleBasics.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	codec.RegisterEvidences(cdc)
	authvesting.RegisterCodec(cdc)

	return cdc.Seal()
}

// DesmosApp extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type DesmosApp struct {
	*bam.BaseApp
	cdc *codec.Codec

	invCheckPeriod uint

	// sdk keys to access the substores
	keys  map[string]*sdk.KVStoreKey
	tkeys map[string]*sdk.TransientStoreKey

	// subspaces
	subspaces map[string]params.Subspace

	// Keepers
	AccountKeeper  auth.AccountKeeper
	BankKeeper     bank.Keeper
	SupplyKeeper   supply.Keeper
	StakingKeeper  staking.Keeper
	SlashingKeeper slashing.Keeper
	DistrKeeper    distr.Keeper
	GovKeeper      gov.Keeper
	CrisisKeeper   crisis.Keeper
	UpgradeKeeper  upgrade.Keeper
	ParamsKeeper   params.Keeper
	EvidenceKeeper evidence.Keeper

	// Custom modules
	FeesKeeper          feesKeeper.Keeper
	MagpieKeeper        magpieKeeper.Keeper
	PostsKeeper         postsKeeper.Keeper
	ProfileKeeper       profilesKeeper.Keeper
	ReportsKeeper       reportsKeeper.Keeper
	RelationshipsKeeper relationships.Keeper

	// Module Manager
	mm *module.Manager

	// Simulation manager
	sm *module.SimulationManager
}

// NewDesmosApp is a constructor function for DesmosApp
func NewDesmosApp(logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool,
	skipUpgradeHeights map[int64]bool, invCheckPeriod uint, baseAppOptions ...func(*bam.BaseApp),
) *DesmosApp {
	// First define the top level codec that will be shared by the different modules
	cdc := MakeCodec()

	// BaseApp handles interactions with Tendermint through the ABCI protocol
	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetAppVersion(version.Version)
	keys := sdk.NewKVStoreKeys(
		bam.MainStoreKey, auth.StoreKey, staking.StoreKey,
		supply.StoreKey, distr.StoreKey, slashing.StoreKey,
		gov.StoreKey, upgrade.StoreKey, params.StoreKey, evidence.StoreKey,

		// Custom modules
		magpieTypes.StoreKey, postsTypes.StoreKey, reportsTypes.StoreKey,
		profilesTypes.StoreKey, relationshipsTypes.StoreKey,
	)
	tkeys := sdk.NewTransientStoreKeys(params.TStoreKey)

	// Here you initialize your application with the store keys it requires
	var app = &DesmosApp{
		BaseApp:        bApp,
		cdc:            cdc,
		invCheckPeriod: invCheckPeriod,
		keys:           keys,
		tkeys:          tkeys,
		subspaces:      make(map[string]params.Subspace),
	}

	// Init params keeper and subspaces
	app.ParamsKeeper = params.NewKeeper(app.cdc, keys[params.StoreKey], tkeys[params.TStoreKey])
	app.subspaces[auth.ModuleName] = app.ParamsKeeper.Subspace(auth.DefaultParamspace)
	app.subspaces[bank.ModuleName] = app.ParamsKeeper.Subspace(bank.DefaultParamspace)
	app.subspaces[staking.ModuleName] = app.ParamsKeeper.Subspace(staking.DefaultParamspace)
	app.subspaces[distr.ModuleName] = app.ParamsKeeper.Subspace(distr.DefaultParamspace)
	app.subspaces[slashing.ModuleName] = app.ParamsKeeper.Subspace(slashing.DefaultParamspace)
	app.subspaces[gov.ModuleName] = app.ParamsKeeper.Subspace(gov.DefaultParamspace).WithKeyTable(gov.ParamKeyTable())
	app.subspaces[evidence.ModuleName] = app.ParamsKeeper.Subspace(evidence.DefaultParamspace)
	app.subspaces[crisis.ModuleName] = app.ParamsKeeper.Subspace(crisis.DefaultParamspace)
	app.subspaces[feesTypes.ModuleName] = app.ParamsKeeper.Subspace(feesTypes.DefaultParamspace)
	app.subspaces[postsTypes.ModuleName] = app.ParamsKeeper.Subspace(postsTypes.DefaultParamspace)
	app.subspaces[profilesTypes.ModuleName] = app.ParamsKeeper.Subspace(profilesTypes.DefaultParamspace)

	// Add keepers
	app.AccountKeeper = auth.NewAccountKeeper(
		app.cdc,
		keys[auth.StoreKey],
		app.subspaces[auth.ModuleName],
		auth.ProtoBaseAccount,
	)
	app.BankKeeper = bank.NewBaseKeeper(
		app.AccountKeeper,
		app.subspaces[bank.ModuleName],
		app.BlacklistedAccAddrs(),
	)
	app.SupplyKeeper = supply.NewKeeper(
		app.cdc,
		keys[supply.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
		maccPerms,
	)
	stakingKeeper := staking.NewKeeper(
		app.cdc,
		keys[staking.StoreKey],
		app.SupplyKeeper,
		app.subspaces[staking.ModuleName],
	)
	app.DistrKeeper = distr.NewKeeper(
		app.cdc,
		keys[distr.StoreKey],
		app.subspaces[distr.ModuleName],
		&stakingKeeper,
		app.SupplyKeeper,
		auth.FeeCollectorName,
		app.ModuleAccountAddrs(),
	)
	app.SlashingKeeper = slashing.NewKeeper(
		app.cdc,
		keys[slashing.StoreKey],
		&stakingKeeper,
		app.subspaces[slashing.ModuleName],
	)
	app.CrisisKeeper = crisis.NewKeeper(
		app.subspaces[crisis.ModuleName],
		app.invCheckPeriod,
		app.SupplyKeeper,
		auth.FeeCollectorName,
	)
	app.UpgradeKeeper = upgrade.NewKeeper(
		skipUpgradeHeights,
		keys[upgrade.StoreKey],
		app.cdc,
	)

	// Create evidence keeper with router
	evidenceKeeper := evidence.NewKeeper(
		app.cdc,
		keys[evidence.StoreKey],
		app.subspaces[evidence.ModuleName],
		&app.StakingKeeper,
		app.SlashingKeeper,
	)
	evidenceRouter := evidence.NewRouter()
	evidenceKeeper.SetRouter(evidenceRouter)
	app.EvidenceKeeper = *evidenceKeeper

	// Create gov keeper with router
	govRouter := gov.NewRouter()
	govRouter.
		AddRoute(gov.RouterKey, gov.ProposalHandler).
		AddRoute(params.RouterKey, params.NewParamChangeProposalHandler(app.ParamsKeeper)).
		AddRoute(distr.RouterKey, distr.NewCommunityPoolSpendProposalHandler(app.DistrKeeper)).
		AddRoute(upgrade.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.UpgradeKeeper))

	app.GovKeeper = gov.NewKeeper(
		app.cdc,
		keys[gov.StoreKey],
		app.subspaces[gov.ModuleName],
		app.SupplyKeeper,
		&stakingKeeper,
		govRouter,
	)

	// Register custom modules
	app.FeesKeeper = feesKeeper.NewKeeper(
		app.cdc,
		app.subspaces[feesTypes.ModuleName],
	)
	app.MagpieKeeper = magpieKeeper.NewKeeper(
		app.cdc,
		keys[magpieTypes.StoreKey],
	)
	app.RelationshipsKeeper = relationshipsKeeper.NewKeeper(
		app.cdc,
		keys[relationshipsTypes.StoreKey],
	)
	app.PostsKeeper = postsKeeper.NewKeeper(
		app.cdc,
		keys[postsTypes.StoreKey],
		app.subspaces[postsTypes.ModuleName],
		app.RelationshipsKeeper,
	)
	app.ProfileKeeper = profilesKeeper.NewKeeper(
		app.cdc,
		keys[profilesTypes.StoreKey],
		app.subspaces[profilesTypes.ModuleName],
		app.RelationshipsKeeper,
	)
	app.ReportsKeeper = reportsKeeper.NewKeeper(
		app.PostsKeeper,
		app.cdc,
		keys[reportsTypes.StoreKey],
	)

	// Register the staking hooks
	// NOTE: StakingKeeper above is passed by reference, so that it will contain these hooks
	app.StakingKeeper = *stakingKeeper.SetHooks(
		staking.NewMultiStakingHooks(
			app.DistrKeeper.Hooks(),
			app.SlashingKeeper.Hooks()),
	)

	// Create the module manager
	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.
	app.mm = module.NewManager(
		genutil.NewAppModule(app.AccountKeeper, app.StakingKeeper, app.BaseApp.DeliverTx),
		auth.NewAppModule(app.AccountKeeper),
		bank.NewAppModule(app.BankKeeper, app.AccountKeeper),
		crisis.NewAppModule(&app.CrisisKeeper),
		supply.NewAppModule(app.SupplyKeeper, app.AccountKeeper),
		gov.NewAppModule(app.GovKeeper, app.AccountKeeper, app.SupplyKeeper),
		slashing.NewAppModule(app.SlashingKeeper, app.AccountKeeper, app.StakingKeeper),
		distr.NewAppModule(app.DistrKeeper, app.AccountKeeper, app.SupplyKeeper, app.StakingKeeper),
		staking.NewAppModule(app.StakingKeeper, app.AccountKeeper, app.SupplyKeeper),
		upgrade.NewAppModule(app.UpgradeKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),

		// Custom modules
		fees.NewAppModule(app.FeesKeeper, app.AccountKeeper),
		magpie.NewAppModule(app.MagpieKeeper, app.AccountKeeper),
		posts.NewAppModule(app.PostsKeeper, app.AccountKeeper),
		profiles.NewAppModule(app.ProfileKeeper, app.AccountKeeper),
		reports.NewAppModule(app.ReportsKeeper, app.AccountKeeper, app.PostsKeeper),
		relationships.NewAppModule(app.RelationshipsKeeper, app.AccountKeeper),
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	app.mm.SetOrderBeginBlockers(
		upgrade.ModuleName, distr.ModuleName, slashing.ModuleName,
		evidence.ModuleName, staking.ModuleName,
	)
	app.mm.SetOrderEndBlockers(crisis.ModuleName, gov.ModuleName, staking.ModuleName)

	app.mm.SetOrderInitGenesis(
		auth.ModuleName, // loads all accounts - should run before any module with a module account
		distr.ModuleName,
		staking.ModuleName, bank.ModuleName, slashing.ModuleName,
		gov.ModuleName, evidence.ModuleName,

		// custom modules
		feesTypes.ModuleName, magpieTypes.ModuleName,
		profilesTypes.ModuleName, relationshipsTypes.ModuleName,
		postsTypes.ModuleName, reportsTypes.ModuleName,

		supply.ModuleName,  // calculates the total supply from account - should run after modules that modify accounts in genesis
		crisis.ModuleName,  // runs the invariants at genesis - should run after other modules
		genutil.ModuleName, // genutils must occur after staking so that pools are properly initialized with tokens from genesis accounts.
	)

	app.mm.RegisterInvariants(&app.CrisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter())

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: this is not required apps that don't use the simulator for fuzz testing
	// transactions
	app.sm = module.NewSimulationManager(
		auth.NewAppModule(app.AccountKeeper),
		bank.NewAppModule(app.BankKeeper, app.AccountKeeper),
		supply.NewAppModule(app.SupplyKeeper, app.AccountKeeper),
		gov.NewAppModule(app.GovKeeper, app.AccountKeeper, app.SupplyKeeper),
		distr.NewAppModule(app.DistrKeeper, app.AccountKeeper, app.SupplyKeeper, app.StakingKeeper),
		staking.NewAppModule(app.StakingKeeper, app.AccountKeeper, app.SupplyKeeper),
		slashing.NewAppModule(app.SlashingKeeper, app.AccountKeeper, app.StakingKeeper),

		// Custom modules
		fees.NewAppModule(app.FeesKeeper, app.AccountKeeper),
		posts.NewAppModule(app.PostsKeeper, app.AccountKeeper),
		magpie.NewAppModule(app.MagpieKeeper, app.AccountKeeper),
		profiles.NewAppModule(app.ProfileKeeper, app.AccountKeeper),
		reports.NewAppModule(app.ReportsKeeper, app.AccountKeeper, app.PostsKeeper),
		relationships.NewAppModule(app.RelationshipsKeeper, app.AccountKeeper),
	)

	app.sm.RegisterStoreDecoders()

	// Initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)

	// Initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetAnteHandler(ante.NewAnteHandler(
		app.AccountKeeper, app.SupplyKeeper, auth.DefaultSigVerificationGasConsumer, app.FeesKeeper),
	)
	app.SetEndBlocker(app.EndBlocker)

	if loadLatest {
		err := app.LoadLatestVersion(app.keys[bam.MainStoreKey])
		if err != nil {
			tmos.Exit(err.Error())
		}
	}

	return app
}

// SetupConfig sets up the given config as it should be for Desmos
func SetupConfig(config *sdk.Config) {
	config.SetBech32PrefixForAccount(
		Bech32MainPrefix,
		Bech32MainPrefix+sdk.PrefixPublic,
	)
	config.SetBech32PrefixForValidator(
		Bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixOperator,
		Bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixOperator+sdk.PrefixPublic,
	)
	config.SetBech32PrefixForConsensusNode(
		Bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixConsensus,
		Bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixConsensus+sdk.PrefixPublic,
	)

	// 852 is the international dialing code of Hong Kong
	// Following the coin type registered at https://github.com/satoshilabs/slips/blob/master/slip-0044.md
	config.SetCoinType(852)
	config.SetFullFundraiserPath("44'/852'/0'/0/0")
}

// BeginBlocker application updates every begin block
func (app *DesmosApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block
func (app *DesmosApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// InitChainer application update.md at chain initialization
func (app *DesmosApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	app.cdc.MustUnmarshalJSON(req.AppStateBytes, &genesisState)
	return app.mm.InitGenesis(ctx, genesisState)
}

// LoadHeight loads a particular height
func (app *DesmosApp) LoadHeight(height int64) error {
	return app.LoadVersion(height, app.keys[bam.MainStoreKey])
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *DesmosApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[supply.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// BlacklistedAccAddrs returns all the app's module account addresses black listed for receiving tokens.
func (app *DesmosApp) BlacklistedAccAddrs() map[string]bool {
	blacklistedAddrs := make(map[string]bool)
	for acc := range maccPerms {
		blacklistedAddrs[supply.NewModuleAddress(acc).String()] = !allowedReceivingModAcc[acc]
	}

	return blacklistedAddrs
}

// Codec returns the application's sealed codec.
func (app *DesmosApp) Codec() *codec.Codec {
	return app.cdc
}

// SimulationManager implements the SimulationApp interface
func (app *DesmosApp) SimulationManager() *module.SimulationManager {
	return app.sm
}
