//go:build !app_v1

package app

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/desmos-labs/desmos/v5/app/upgrades"

	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/server"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	postskeeper "github.com/desmos-labs/desmos/v5/x/posts/keeper"

	"cosmossdk.io/depinject"

	consensusparamkeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"

	ibcclient "github.com/cosmos/ibc-go/v7/modules/core/02-client"

	"github.com/cosmos/cosmos-sdk/client/flags"

	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"

	dbm "github.com/cometbft/cometbft-db"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	tmos "github.com/cometbft/cometbft/libs/os"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/streaming"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	_ "github.com/cosmos/cosmos-sdk/x/auth/tx/config" // import for side-effects
	"github.com/cosmos/cosmos-sdk/x/params"

	"github.com/cosmos/cosmos-sdk/x/auth/ante"

	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"

	ibcfee "github.com/cosmos/ibc-go/v7/modules/apps/29-fee"
	ibcfeekeeper "github.com/cosmos/ibc-go/v7/modules/apps/29-fee/keeper"
	ibcfeetypes "github.com/cosmos/ibc-go/v7/modules/apps/29-fee/types"
	ibctransfer "github.com/cosmos/ibc-go/v7/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v7/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v7/modules/core"
	ibcclienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	porttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"
	ibckeeper "github.com/cosmos/ibc-go/v7/modules/core/keeper"
	solomachine "github.com/cosmos/ibc-go/v7/modules/light-clients/06-solomachine"
	ibctm "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"

	ica "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts"
	icacontroller "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/keeper"
	icacontrollertypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/types"
	icahost "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host"
	icahostkeeper "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host/keeper"
	icahosttypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host/types"

	"github.com/cosmos/cosmos-sdk/x/upgrade"

	"github.com/desmos-labs/desmos/v5/x/profiles"
	profileskeeper "github.com/desmos-labs/desmos/v5/x/profiles/keeper"
	profilestypes "github.com/desmos-labs/desmos/v5/x/profiles/types"
	reactionskeeper "github.com/desmos-labs/desmos/v5/x/reactions/keeper"
	relationshipskeeper "github.com/desmos-labs/desmos/v5/x/relationships/keeper"
	reportskeeper "github.com/desmos-labs/desmos/v5/x/reports/keeper"
	subspaceskeeper "github.com/desmos-labs/desmos/v5/x/subspaces/keeper"

	supplykeeper "github.com/desmos-labs/desmos/v5/x/supply/keeper"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"

	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
)

var (
	_ runtime.AppI            = (*DesmosApp)(nil)
	_ servertypes.Application = (*DesmosApp)(nil)
)

// DesmosApp extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type DesmosApp struct {
	*runtime.App
	legacyAmino       *codec.LegacyAmino
	appCodec          codec.Codec
	txConfig          client.TxConfig
	interfaceRegistry types.InterfaceRegistry

	// non depinject support modules store keys
	keys map[string]*storetypes.KVStoreKey

	// Keepers
	AccountKeeper         authkeeper.AccountKeeper
	BankKeeper            bankkeeper.Keeper
	CapabilityKeeper      *capabilitykeeper.Keeper
	StakingKeeper         *stakingkeeper.Keeper
	SlashingKeeper        slashingkeeper.Keeper
	MintKeeper            mintkeeper.Keeper
	DistrKeeper           distrkeeper.Keeper
	GovKeeper             *govkeeper.Keeper
	CrisisKeeper          *crisiskeeper.Keeper
	UpgradeKeeper         *upgradekeeper.Keeper
	ParamsKeeper          paramskeeper.Keeper
	AuthzKeeper           authzkeeper.Keeper
	EvidenceKeeper        evidencekeeper.Keeper
	FeeGrantKeeper        feegrantkeeper.Keeper
	ConsensusParamsKeeper consensusparamkeeper.Keeper

	WasmKeeper wasm.Keeper

	// IBC modules
	TransferKeeper      ibctransferkeeper.Keeper
	IBCKeeper           *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	IBCFeeKeeper        ibcfeekeeper.Keeper
	ICAControllerKeeper icacontrollerkeeper.Keeper
	ICAHostKeeper       icahostkeeper.Keeper

	// make scoped keepers public for test purposes
	ScopedIBCKeeper           capabilitykeeper.ScopedKeeper
	ScopedIBCTransferKeeper   capabilitykeeper.ScopedKeeper
	ScopedProfilesKeeper      capabilitykeeper.ScopedKeeper
	ScopedWasmKeeper          capabilitykeeper.ScopedKeeper
	ScopedICAHostKeeper       capabilitykeeper.ScopedKeeper
	ScopedICAControllerKeeper capabilitykeeper.ScopedKeeper

	// Custom modules
	SubspacesKeeper     *subspaceskeeper.Keeper
	ProfilesKeeper      *profileskeeper.Keeper
	RelationshipsKeeper relationshipskeeper.Keeper
	PostsKeeper         *postskeeper.Keeper
	ReportsKeeper       reportskeeper.Keeper
	ReactionsKeeper     reactionskeeper.Keeper
	SupplyKeeper        supplykeeper.Keeper

	// Simulation manager
	sm *module.SimulationManager
}

func NewDesmosApp(
	logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool,
	appOpts servertypes.AppOptions, wasmEnabledProposals []wasm.ProposalType,
	baseAppOptions ...func(*baseapp.BaseApp),
) *DesmosApp {
	var (
		app        = &DesmosApp{}
		appBuilder *runtime.AppBuilder

		// merge the AppConfig and other configuration in one config
		appConfig = depinject.Configs(
			AppConfig,
			depinject.Supply(
				// supply the application options
				appOpts,

				// ADVANCED CONFIGURATION

				//
				// AUTH
				//
				// For providing a custom function required in auth to generate custom account types
				// add it below. By default the auth module uses simulation.RandomGenesisAccounts.
				//
				// authtypes.RandomGenesisAccountsFn(simulation.RandomGenesisAccounts),

				// For providing a custom a base account type add it below.
				// By default the auth module uses authtypes.ProtoBaseAccount().
				//
				// func() authtypes.AccountI { return authtypes.ProtoBaseAccount() },

				//
				// MINT
				//

				// For providing a custom inflation function for x/mint add here your
				// custom function that implements the minttypes.InflationCalculationFn
				// interface.
			),
		)
	)

	if err := depinject.Inject(appConfig,
		&appBuilder,
		&app.appCodec,
		&app.legacyAmino,
		&app.txConfig,
		&app.interfaceRegistry,
		&app.AccountKeeper,
		&app.BankKeeper,
		&app.CapabilityKeeper,
		&app.StakingKeeper,
		&app.SlashingKeeper,
		&app.MintKeeper,
		&app.DistrKeeper,
		&app.GovKeeper,
		&app.CrisisKeeper,
		&app.UpgradeKeeper,
		&app.ParamsKeeper,
		&app.AuthzKeeper,
		&app.EvidenceKeeper,
		&app.FeeGrantKeeper,
		&app.ConsensusParamsKeeper,

		// Desmos modules
		&app.SubspacesKeeper,
		&app.ProfilesKeeper,
		&app.RelationshipsKeeper,
		&app.PostsKeeper,
		&app.ReportsKeeper,
		&app.ReactionsKeeper,
		&app.SupplyKeeper,
	); err != nil {
		panic(err)
	}

	// Below we could construct and set an application specific mempool and
	// ABCI 1.0 PrepareProposal and ProcessProposal handlers. These defaults are
	// already set in the SDK's BaseApp, this shows an example of how to override
	// them.
	//
	// Example:
	//
	// app.App = appBuilder.Build(...)
	// nonceMempool := mempool.NewSenderNonceMempool()
	// abciPropHandler := NewDefaultProposalHandler(nonceMempool, app.App.BaseApp)
	//
	// app.App.BaseApp.SetMempool(nonceMempool)
	// app.App.BaseApp.SetPrepareProposal(abciPropHandler.PrepareProposalHandler())
	// app.App.BaseApp.SetProcessProposal(abciPropHandler.ProcessProposalHandler())
	//
	// Alternatively, you can construct BaseApp options, append those to
	// baseAppOptions and pass them to the appBuilder.
	//
	// Example:
	//
	// prepareOpt = func(app *baseapp.BaseApp) {
	// 	abciPropHandler := baseapp.NewDefaultProposalHandler(nonceMempool, app)
	// 	app.SetPrepareProposal(abciPropHandler.PrepareProposalHandler())
	// }
	// baseAppOptions = append(baseAppOptions, prepareOpt)

	app.App = appBuilder.Build(logger, db, traceStore, baseAppOptions...)

	// set up non depinject support modules store keys
	app.keys = sdk.NewKVStoreKeys(
		ibcexported.StoreKey, ibctransfertypes.StoreKey, ibcfeetypes.StoreKey,
		icahosttypes.StoreKey, icacontrollertypes.StoreKey, wasmtypes.StoreKey,
	)
	app.MountKVStores(app.keys)

	// load state streaming if enabled
	if _, _, err := streaming.LoadStreamingServices(app.App.BaseApp, appOpts, app.appCodec, logger, app.kvStoreKeys()); err != nil {
		logger.Error("failed to load state streaming", "err", err)
		os.Exit(1)
	}

	/****  Module Options ****/

	// set params subspaces
	for _, m := range []string{ibctransfertypes.ModuleName, ibcexported.ModuleName, icahosttypes.SubModuleName, icacontrollertypes.SubModuleName} {
		app.ParamsKeeper.Subspace(m)
	}

	// add capability keeper and ScopeToModule for ibc module
	scopedIBCKeeper := app.CapabilityKeeper.ScopeToModule(ibcexported.ModuleName)
	scopedIBCTransferKeeper := app.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	scopedProfilesKeeper := app.CapabilityKeeper.ScopeToModule(profilestypes.ModuleName)
	scopedWasmKeeper := app.CapabilityKeeper.ScopeToModule(wasm.ModuleName)
	scopedICAControllerKeeper := app.CapabilityKeeper.ScopeToModule(icacontrollertypes.SubModuleName)
	scopedICAHostKeeper := app.CapabilityKeeper.ScopeToModule(icahosttypes.SubModuleName)

	// Create IBC keeper
	app.IBCKeeper = ibckeeper.NewKeeper(
		app.appCodec, app.GetKey(ibcexported.StoreKey), app.GetSubspace(ibcexported.ModuleName), app.StakingKeeper, app.UpgradeKeeper, scopedIBCKeeper,
	)

	// Register the proposal types
	// Deprecated: Avoid adding new handlers, instead use the new proposal flow
	// by granting the governance module the right to execute the message.
	// See: https://docs.cosmos.network/main/modules/gov#proposal-messages
	govRouter := govv1beta1.NewRouter()
	govRouter.AddRoute(govtypes.RouterKey, govv1beta1.ProposalHandler).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(app.ParamsKeeper)). // This should be removed. It is still in place to avoid failures of modules that have not yet been upgraded.
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.UpgradeKeeper)).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(app.IBCKeeper.ClientKeeper))

	app.IBCFeeKeeper = ibcfeekeeper.NewKeeper(
		app.appCodec, app.GetKey(ibcfeetypes.StoreKey),
		app.IBCKeeper.ChannelKeeper, // may be replaced with IBC middleware
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper, app.AccountKeeper, app.BankKeeper,
	)

	// Create IBC transfer keeper
	app.TransferKeeper = ibctransferkeeper.NewKeeper(
		app.appCodec,
		app.GetKey(ibctransfertypes.StoreKey),
		app.GetSubspace(ibctransfertypes.ModuleName),
		app.IBCKeeper.ChannelKeeper,
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		scopedIBCTransferKeeper,
	)

	// Create interchain account keepers
	app.ICAHostKeeper = icahostkeeper.NewKeeper(
		app.appCodec,
		app.GetKey(icahosttypes.StoreKey),
		app.GetSubspace(icahosttypes.SubModuleName),
		app.IBCFeeKeeper, // use ics29 fee as ics4Wrapper in middleware stack
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		app.AccountKeeper,
		scopedICAHostKeeper,
		app.MsgServiceRouter(),
	)
	app.ICAControllerKeeper = icacontrollerkeeper.NewKeeper(
		app.appCodec,
		app.GetKey(icacontrollertypes.StoreKey),
		app.GetSubspace(icacontrollertypes.SubModuleName),
		app.IBCFeeKeeper, // use ics29 fee as ics4Wrapper in middleware stack
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		scopedICAControllerKeeper,
		app.MsgServiceRouter(),
	)

	// Setup profiles IBC keepers manually
	app.ProfilesKeeper.SetIBCKeepers(
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		scopedProfilesKeeper,
	)

	// ------ CosmWasm setup ------

	// var wasmRouter = bApp.Router()
	homePath := cast.ToString(appOpts.Get(flags.FlagHome))
	wasmDir := filepath.Join(homePath, "wasm")

	wasmConfig, err := wasm.ReadWasmConfig(appOpts)
	if err != nil {
		panic("error while reading wasm config: " + err.Error())
	}

	wasmOpts := GetWasmOpts(
		appOpts,
		app.GRPCQueryRouter(),
		app.appCodec,
		*app.ProfilesKeeper,
		*app.SubspacesKeeper,
		app.RelationshipsKeeper,
		*app.PostsKeeper,
		app.ReportsKeeper,
		app.ReactionsKeeper,
	)

	// The last arguments can contain custom message handlers, and custom query handlers,
	// if we want to allow any custom callbacks
	// See https://github.com/CosmWasm/cosmwasm/blob/main/docs/CAPABILITIES-BUILT-IN.md
	availableCapabilities := "iterator,staking,stargate,cosmwasm_1_1,cosmwasm_1_2"
	app.WasmKeeper = wasm.NewKeeper(
		app.appCodec,
		app.GetKey(wasm.StoreKey),
		app.AccountKeeper,
		app.BankKeeper,
		app.StakingKeeper,
		distrkeeper.NewQuerier(app.DistrKeeper),
		app.IBCFeeKeeper, // ICS4 Wrapper: fee IBC middleware
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		scopedWasmKeeper,
		app.TransferKeeper,
		app.MsgServiceRouter(),
		app.GRPCQueryRouter(),
		wasmDir,
		wasmConfig,
		availableCapabilities,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		wasmOpts...,
	)

	// The gov proposal types can be individually enabled
	if len(wasmEnabledProposals) != 0 {
		govRouter.AddRoute(wasm.RouterKey, wasm.NewWasmProposalHandler(app.WasmKeeper, wasmEnabledProposals))
	}

	// Set legacy router for backwards compatibility with gov v1beta1
	app.GovKeeper.SetLegacyRouter(govRouter)

	// Create IBC modules with ibcfee middleware
	transferIBCModule := ibcfee.NewIBCMiddleware(ibctransfer.NewIBCModule(app.TransferKeeper), app.IBCFeeKeeper)
	profilesIBCModule := ibcfee.NewIBCMiddleware(profiles.NewIBCModule(app.appCodec, app.ProfilesKeeper), app.IBCFeeKeeper)
	wasmIBCModule := ibcfee.NewIBCMiddleware(wasm.NewIBCHandler(app.WasmKeeper, app.IBCKeeper.ChannelKeeper, app.IBCFeeKeeper), app.IBCFeeKeeper)

	// integration point for custom authentication modules
	// see https://medium.com/the-interchain-foundation/ibc-go-v6-changes-to-interchain-accounts-and-how-it-impacts-your-chain-806c185300d7
	// TODO: Once a public ICA auth module is released, replace nil with the module
	var noAuthzModule porttypes.IBCModule
	icaControllerIBCModule := ibcfee.NewIBCMiddleware(
		icacontroller.NewIBCMiddleware(noAuthzModule, app.ICAControllerKeeper),
		app.IBCFeeKeeper,
	)

	icaHostIBCModule := ibcfee.NewIBCMiddleware(icahost.NewIBCModule(app.ICAHostKeeper), app.IBCFeeKeeper)

	// Create static IBC router, add transfer route, then set and seal it
	ibcRouter := porttypes.NewRouter().
		AddRoute(ibctransfertypes.ModuleName, transferIBCModule).
		AddRoute(profilestypes.ModuleName, profilesIBCModule).
		AddRoute(wasm.ModuleName, wasmIBCModule).
		AddRoute(icacontrollertypes.SubModuleName, icaControllerIBCModule).
		AddRoute(icahosttypes.SubModuleName, icaHostIBCModule)

	app.IBCKeeper.SetRouter(ibcRouter)

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.
	legacyModules := []module.AppModule{
		ibc.NewAppModule(app.IBCKeeper),
		ibctransfer.NewAppModule(app.TransferKeeper),
		ibcfee.NewAppModule(app.IBCFeeKeeper),
		ica.NewAppModule(&app.ICAControllerKeeper, &app.ICAHostKeeper),
		wasm.NewAppModule(app.appCodec, &app.WasmKeeper, app.StakingKeeper, app.AccountKeeper, app.BankKeeper, app.MsgServiceRouter(), app.GetSubspace(wasmtypes.ModuleName)),
	}

	app.RegisterModules(legacyModules...)
	cfg := app.Configurator()
	for _, m := range legacyModules {
		if s, ok := m.(module.HasServices); ok {
			s.RegisterServices(cfg)
		}
	}

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: this is not required apps that don't use the simulator for fuzz testing
	// transactions
	overrideModules := map[string]module.AppModuleSimulation{
		authtypes.ModuleName: auth.NewAppModule(app.appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts, app.GetSubspace(authtypes.ModuleName)),
	}

	// NOTE: Simulation manager, invariants and upgrade handlers must be after all the modules are registered
	app.sm = module.NewSimulationManagerFromAppModules(app.ModuleManager.Modules, overrideModules)

	app.sm.RegisterStoreDecoders()
	app.ModuleManager.RegisterInvariants(app.CrisisKeeper)
	app.registerUpgradeHandlers()

	// register additional types
	ibctm.AppModuleBasic{}.RegisterInterfaces(app.interfaceRegistry)
	solomachine.AppModuleBasic{}.RegisterInterfaces(app.interfaceRegistry)

	app.ScopedIBCKeeper = scopedIBCKeeper
	app.ScopedIBCTransferKeeper = scopedIBCTransferKeeper
	app.ScopedProfilesKeeper = scopedProfilesKeeper
	app.ScopedWasmKeeper = scopedWasmKeeper
	app.ScopedICAHostKeeper = scopedICAHostKeeper
	app.ScopedICAControllerKeeper = scopedICAControllerKeeper

	anteHandler, err := NewAnteHandler(
		HandlerOptions{
			HandlerOptions: ante.HandlerOptions{
				AccountKeeper:   app.AccountKeeper,
				BankKeeper:      app.BankKeeper,
				SignModeHandler: app.txConfig.SignModeHandler(),
				FeegrantKeeper:  app.FeeGrantKeeper,
				SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
			},
			IBCkeeper:         app.IBCKeeper,
			TxCounterStoreKey: app.GetKey(wasm.StoreKey),
			WasmConfig:        &wasmConfig,
			SubspacesKeeper:   *app.SubspacesKeeper,
		},
	)
	if err != nil {
		panic(err)
	}

	app.SetAnteHandler(anteHandler)

	// In v0.46, the SDK introduces _postHandlers_. PostHandlers are like
	// antehandlers, but are run _after_ the `runMsgs` execution. They are also
	// defined as a chain, and have the same signature as antehandlers.
	//
	// In baseapp, postHandlers are run in the same store branch as `runMsgs`,
	// meaning that both `runMsgs` and `postHandler` state will be committed if
	// both are successful, and both will be reverted if any of the two fails.
	//
	// The SDK exposes a default postHandlers chain, which comprises of only
	// one decorator: the Transaction Tips decorator. However, some chains do
	// not need it by default, so feel free to comment the next line if you do
	// not need tips.
	// To read more about tips:
	// https://docs.cosmos.network/main/core/tips.html
	//
	// Please note that changing any of the anteHandler or postHandler chain is
	// likely to be a state-machine breaking change, which needs a coordinated
	// upgrade.
	postHandler, err := NewPostHandler()
	if err != nil {
		panic(err)
	}
	app.SetPostHandler(postHandler)

	// must be before Loading version
	// requires the snapshot store to be created and registered as a BaseAppOption
	// see cmd/wasmd/root.go: 206 - 214 approx
	if manager := app.SnapshotManager(); manager != nil {
		err := manager.RegisterExtensions(
			wasmkeeper.NewWasmSnapshotter(app.CommitMultiStore(), &app.WasmKeeper),
		)
		if err != nil {
			panic(fmt.Errorf("failed to register snapshot extension: %s", err))
		}
	}

	// A custom InitChainer can be set if extra pre-init-genesis logic is required.
	// By default, when using app wiring enabled module, this is not required.
	// For instance, the upgrade module will set automatically the module version map in its init genesis thanks to app wiring.
	// However, when registering a module manually (i.e. that does not support app wiring), the module version map
	// must be set manually as follow. The upgrade module will de-duplicate the module version map.
	app.SetInitChainer(func(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
		app.UpgradeKeeper.SetModuleVersionMap(ctx, app.ModuleManager.GetVersionMap())
		return app.App.InitChainer(ctx, req)
	})

	if err := app.Load(loadLatest); err != nil {
		panic(err)
	}

	if loadLatest {
		ctx := app.BaseApp.NewUncachedContext(true, tmproto.Header{})

		// Initialize pinned codes in wasmvm as they are not persisted there
		if err := app.WasmKeeper.InitializePinnedCodes(ctx); err != nil {
			tmos.Exit(fmt.Sprintf("failed initialize pinned codes %s", err))
		}
	}

	return app
}

// Name returns the name of the App
func (app *DesmosApp) Name() string { return app.BaseApp.Name() }

// LegacyAmino returns DesmosApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *DesmosApp) LegacyAmino() *codec.LegacyAmino {
	return app.legacyAmino
}

// AppCodec returns DesmosApp's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *DesmosApp) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns DesmosApp's InterfaceRegistry
func (app *DesmosApp) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// TxConfig returns DesmosApp's TxConfig
func (app *DesmosApp) TxConfig() client.TxConfig {
	return app.txConfig
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *DesmosApp) GetKey(storeKey string) *storetypes.KVStoreKey {
	if key, ok := app.keys[storeKey]; ok {
		return key
	}

	sk := app.UnsafeFindStoreKey(storeKey)
	kvStoreKey, ok := sk.(*storetypes.KVStoreKey)
	if !ok {
		return nil
	}
	return kvStoreKey
}

// kvStoreKeys returns all the kv store keys registered inside DesmosApp
func (app *DesmosApp) kvStoreKeys() map[string]*storetypes.KVStoreKey {
	keys := make(map[string]*storetypes.KVStoreKey)
	for _, k := range app.GetStoreKeys() {
		if kv, ok := k.(*storetypes.KVStoreKey); ok {
			keys[kv.Name()] = kv
		}
	}

	for _, kv := range app.keys {
		keys[kv.Name()] = kv
	}

	return keys
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *DesmosApp) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// SimulationManager implements the SimulationApp interface
func (app *DesmosApp) SimulationManager() *module.SimulationManager {
	return app.sm
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *DesmosApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	app.App.RegisterAPIRoutes(apiSvr, apiConfig)
	// register swagger API in app.go so that other applications can override easily
	if err := server.RegisterSwaggerAPI(apiSvr.ClientCtx, apiSvr.Router, apiConfig.Swagger); err != nil {
		panic(err)
	}
}

// registerUpgrade registers the given upgrade to be supported by the app
func (app *DesmosApp) registerUpgrade(upgrade upgrades.Upgrade) {
	app.UpgradeKeeper.SetUpgradeHandler(upgrade.Name(), upgrade.Handler())

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}

	if upgradeInfo.Name == upgrade.Name() && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		// Configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, upgrade.StoreUpgrades()))
	}
}

// BlockedAddresses returns all the app's blocked account addresses.
func BlockedAddresses() map[string]bool {
	result := make(map[string]bool)

	if len(blockAccAddrs) > 0 {
		for _, addr := range blockAccAddrs {
			result[addr] = true
		}
	} else {
		for addr := range GetMaccPerms() {
			result[addr] = true
		}
	}

	return result
}
