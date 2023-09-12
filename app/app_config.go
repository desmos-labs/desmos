package app

import (
	"os"
	"path/filepath"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"

	"github.com/cosmos/cosmos-sdk/x/consensus"

	servertypes "github.com/cosmos/cosmos-sdk/server/types"

	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"

	authmodulev1 "cosmossdk.io/api/cosmos/auth/module/v1"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"

	"github.com/cosmos/cosmos-sdk/x/capability"

	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/mint"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	ibcfee "github.com/cosmos/ibc-go/v7/modules/apps/29-fee"
	ibcfeetypes "github.com/cosmos/ibc-go/v7/modules/apps/29-fee/types"
	ibctransfer "github.com/cosmos/ibc-go/v7/modules/apps/transfer"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v7/modules/core"
	ibcclientclient "github.com/cosmos/ibc-go/v7/modules/core/02-client/client"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"
	solomachine "github.com/cosmos/ibc-go/v7/modules/light-clients/06-solomachine"
	ibctm "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"

	ica "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts"
	icatypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/types"

	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"github.com/desmos-labs/desmos/v6/x/posts"
	postskeeper "github.com/desmos-labs/desmos/v6/x/posts/keeper"
	poststypes "github.com/desmos-labs/desmos/v6/x/posts/types"
	"github.com/desmos-labs/desmos/v6/x/profiles"
	profileskeeper "github.com/desmos-labs/desmos/v6/x/profiles/keeper"
	profilestypes "github.com/desmos-labs/desmos/v6/x/profiles/types"
	"github.com/desmos-labs/desmos/v6/x/reactions"
	reactionskeeper "github.com/desmos-labs/desmos/v6/x/reactions/keeper"
	reactionstypes "github.com/desmos-labs/desmos/v6/x/reactions/types"
	"github.com/desmos-labs/desmos/v6/x/relationships"
	relationshipskeeper "github.com/desmos-labs/desmos/v6/x/relationships/keeper"
	relationshipstypes "github.com/desmos-labs/desmos/v6/x/relationships/types"
	"github.com/desmos-labs/desmos/v6/x/reports"
	reportskeeper "github.com/desmos-labs/desmos/v6/x/reports/keeper"
	reportstypes "github.com/desmos-labs/desmos/v6/x/reports/types"
	"github.com/desmos-labs/desmos/v6/x/subspaces"
	subspaceskeeper "github.com/desmos-labs/desmos/v6/x/subspaces/keeper"
	subspacestypes "github.com/desmos-labs/desmos/v6/x/subspaces/types"
	supplytypes "github.com/desmos-labs/desmos/v6/x/supply/types"
	"github.com/desmos-labs/desmos/v6/x/tokenfactory"
	tokenfactorytypes "github.com/desmos-labs/desmos/v6/x/tokenfactory/types"

	"github.com/desmos-labs/desmos/v6/x/supply"

	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
)

const (
	appName          = "Desmos"
	Bech32MainPrefix = "desmos"

	CoinType           = 852
	FullFundraiserPath = "44'/852'/0'/0/0"
)

// GetWasmOpts parses appOpts and add wasm opt to the given options array.
// if telemetry is enabled, the wasmVM cache metrics are activated.
func GetWasmOpts(
	appOpts servertypes.AppOptions,
	grpcQueryRouter *baseapp.GRPCQueryRouter,
	cdc codec.Codec,
	profilesKeeper *profileskeeper.Keeper,
	subspacesKeeper subspaceskeeper.Keeper,
	relationshipsKeeper relationshipskeeper.Keeper,
	postsKeeper postskeeper.Keeper,
	reportsKeeper reportskeeper.Keeper,
	reactionsKeeper reactionskeeper.Keeper,
) []wasmkeeper.Option {
	var wasmOpts []wasmkeeper.Option
	// FIXME (wasmd-1575): This is commented out temporarily because it causes panics in telemetry server
	// due to the bug fixed here: https://github.com/CosmWasm/wasmd/pull/1575.
	// This will be released with CosmWasm v0.42, so we will un-comment this once we upgrade to that version.
	// if cast.ToBool(appOpts.Get("telemetry.enabled")) {
	// wasmOpts = append(wasmOpts, wasmkeeper.WithVMCacheMetrics(prometheus.DefaultRegisterer))
	// }

	customQueryPlugin := NewDesmosCustomQueryPlugin(
		cdc,
		grpcQueryRouter,
		profilesKeeper,
		subspacesKeeper,
		relationshipsKeeper,
		postsKeeper,
		reportsKeeper,
		reactionsKeeper,
	)
	customMessageEncoder := NewDesmosCustomMessageEncoder(cdc)

	wasmOpts = append(wasmOpts, wasmkeeper.WithGasRegister(NewDesmosWasmGasRegister()))
	wasmOpts = append(wasmOpts, wasmkeeper.WithQueryPlugins(&customQueryPlugin))
	wasmOpts = append(wasmOpts, wasmkeeper.WithMessageEncoders(&customMessageEncoder))

	return wasmOpts
}

var (
	// DefaultNodeHome sets the folder where the application data and configuration will be stored
	DefaultNodeHome string

	// ModuleBasics is in charge of setting up basic module elements
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(
			[]govclient.ProposalHandler{
				paramsclient.ProposalHandler,
				upgradeclient.LegacyProposalHandler,
				upgradeclient.LegacyCancelProposalHandler,
				ibcclientclient.UpdateClientProposalHandler,
				ibcclientclient.UpgradeProposalHandler,
			},
		),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		authzmodule.AppModuleBasic{},
		vesting.AppModuleBasic{},
		consensus.AppModuleBasic{},

		wasm.AppModuleBasic{},

		// IBC modules
		ibc.AppModuleBasic{},
		solomachine.AppModuleBasic{},
		ibctm.AppModuleBasic{},
		ibctransfer.AppModuleBasic{},
		ibcfee.AppModuleBasic{},
		ica.AppModuleBasic{},

		// Custom modules
		profiles.AppModuleBasic{},
		relationships.AppModuleBasic{},
		subspaces.AppModuleBasic{},
		posts.AppModuleBasic{},
		reports.AppModuleBasic{},
		reactions.AppModuleBasic{},
		supply.AppModuleBasic{},
		tokenfactory.AppModuleBasic{},
	)

	// Module account permissions
	maccPerms = []*authmodulev1.ModuleAccountPermission{
		{Account: authtypes.FeeCollectorName},
		{Account: distrtypes.ModuleName},
		{Account: minttypes.ModuleName, Permissions: []string{authtypes.Minter}},
		{Account: stakingtypes.BondedPoolName, Permissions: []string{authtypes.Burner, authtypes.Staking}},
		{Account: stakingtypes.NotBondedPoolName, Permissions: []string{authtypes.Burner, authtypes.Staking}},
		{Account: govtypes.ModuleName, Permissions: []string{authtypes.Burner}},
		{Account: ibctransfertypes.ModuleName, Permissions: []string{authtypes.Minter, authtypes.Burner}},
		{Account: wasmtypes.ModuleName, Permissions: []string{authtypes.Burner}},
		{Account: ibcfeetypes.ModuleName},
		{Account: icatypes.ModuleName},
		{Account: tokenfactorytypes.ModuleName, Permissions: []string{authtypes.Minter, authtypes.Burner}},
	}

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	// NOTE: capability module's beginblocker must come before any modules using capabilities (e.g. IBC)
	beginBlockerOrder = []string{
		upgradetypes.ModuleName,
		capabilitytypes.ModuleName,
		minttypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		stakingtypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		govtypes.ModuleName,
		crisistypes.ModuleName,
		genutiltypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		vestingtypes.ModuleName,
		consensusparamtypes.ModuleName,

		// IBC modules
		ibcexported.ModuleName,
		ibctransfertypes.ModuleName,
		ibcfeetypes.ModuleName,
		icatypes.ModuleName,

		// Custom modules
		subspacestypes.ModuleName,
		relationshipstypes.ModuleName,
		profilestypes.ModuleName,
		poststypes.ModuleName,
		reportstypes.ModuleName,
		reactionstypes.ModuleName,
		supplytypes.ModuleName,
		tokenfactorytypes.ModuleName,

		wasmtypes.ModuleName,
	}

	endBlockerOrder = []string{
		crisistypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		minttypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
		consensusparamtypes.ModuleName,

		// IBC modules
		ibcexported.ModuleName,
		ibctransfertypes.ModuleName,
		ibcfeetypes.ModuleName,
		icatypes.ModuleName,

		// Custom modules
		subspacestypes.ModuleName,
		relationshipstypes.ModuleName,
		profilestypes.ModuleName,
		poststypes.ModuleName,
		reportstypes.ModuleName,
		reactionstypes.ModuleName,
		supplytypes.ModuleName,
		tokenfactorytypes.ModuleName,

		wasmtypes.ModuleName,
	}

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: The genutils module must also occur after auth so that it can access the params from auth.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	// NOTE: wasm module should be at the end as it can call other module functionality direct or via message dispatching during
	// genesis phase. For example bank transfer, auth account check, staking, ...
	genesisModuleOrder = []string{
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		minttypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
		consensusparamtypes.ModuleName,

		// IBC modules
		ibcexported.ModuleName,
		ibctransfertypes.ModuleName,
		ibcfeetypes.ModuleName,
		icatypes.ModuleName,

		// Custom modules
		subspacestypes.ModuleName,
		profilestypes.ModuleName,
		relationshipstypes.ModuleName,
		poststypes.ModuleName,
		reportstypes.ModuleName,
		reactionstypes.ModuleName,
		supplytypes.ModuleName,
		tokenfactorytypes.ModuleName,

		// wasm module should be at the end of app modules
		wasmtypes.ModuleName,
		crisistypes.ModuleName,
	}

	// NOTE: The auth module must occur before everyone else. All other modules can be sorted
	// alphabetically (default order)
	// NOTE: The relationships module must occur before the profiles module, or all relationships will be deleted
	migrationModuleOrder = []string{
		authtypes.ModuleName,
		authz.ModuleName,
		banktypes.ModuleName,
		capabilitytypes.ModuleName,
		distrtypes.ModuleName,
		evidencetypes.ModuleName,
		feegrant.ModuleName,
		genutiltypes.ModuleName,
		govtypes.ModuleName,
		minttypes.ModuleName,
		slashingtypes.ModuleName,
		stakingtypes.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
		consensusparamtypes.ModuleName,

		// IBC modules
		ibcexported.ModuleName,
		ibctransfertypes.ModuleName,
		ibcfeetypes.ModuleName,
		icatypes.ModuleName,

		// Custom modules
		subspacestypes.ModuleName,
		relationshipstypes.ModuleName,
		profilestypes.ModuleName,
		poststypes.ModuleName,
		reportstypes.ModuleName,
		reactionstypes.ModuleName,
		supplytypes.ModuleName,
		tokenfactorytypes.ModuleName,

		wasmtypes.ModuleName,
		crisistypes.ModuleName,
	}
)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, ".desmos")
}

// SetupConfig sets up the given config as it should be for Desmos
func SetupConfig(config *sdk.Config) {
	config.SetBech32PrefixForAccount(Bech32MainPrefix, Bech32MainPrefix+sdk.PrefixPublic)
	config.SetBech32PrefixForValidator(Bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixOperator, Bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixOperator+sdk.PrefixPublic)
	config.SetBech32PrefixForConsensusNode(Bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixConsensus, Bech32MainPrefix+sdk.PrefixValidator+sdk.PrefixConsensus+sdk.PrefixPublic)

	// 852 is the international dialing code of Hong Kong
	// Following the coin type registered at https://github.com/satoshilabs/slips/blob/master/slip-0044.md
	config.SetCoinType(CoinType)
}

// MakeCodecs constructs the *std.Codec and *codec.LegacyAmino instances used by
// DesmosApp. It is useful for tests and clients who do not want to construct the
// full DesmosApp
func MakeCodecs() (codec.Codec, *codec.LegacyAmino) {
	cfg := MakeEncodingConfig()
	return cfg.Codec, cfg.Amino
}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for _, perms := range maccPerms {
		dupMaccPerms[perms.Account] = perms.Permissions
	}
	return dupMaccPerms
}
