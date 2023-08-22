package posts

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"

	modulev1 "github.com/desmos-labs/desmos/v6/api/desmos/posts/module/v1"

	profileskeeper "github.com/desmos-labs/desmos/v6/x/profiles/keeper"
	relationshipskeeper "github.com/desmos-labs/desmos/v6/x/relationships/keeper"
	subspaceskeeper "github.com/desmos-labs/desmos/v6/x/subspaces/keeper"
	subspacestypes "github.com/desmos-labs/desmos/v6/x/subspaces/types"

	v2 "github.com/desmos-labs/desmos/v6/x/posts/legacy/v2"
	v4 "github.com/desmos-labs/desmos/v6/x/posts/legacy/v4"

	"github.com/desmos-labs/desmos/v6/x/posts/client/cli"
	"github.com/desmos-labs/desmos/v6/x/posts/keeper"
	"github.com/desmos-labs/desmos/v6/x/posts/simulation"
	"github.com/desmos-labs/desmos/v6/x/posts/types"
)

const (
	consensusVersion = 7
)

// type check to ensure the interface is properly implemented
var (
	_ module.AppModule           = AppModule{}
	_ module.AppModuleBasic      = AppModuleBasic{}
	_ module.AppModuleSimulation = AppModule{}
	_ appmodule.AppModule        = AppModule{}
	_ depinject.OnePerModuleType = AppModule{}
)

// AppModuleBasic defines the basic application module used by the posts module.
type AppModuleBasic struct {
	cdc codec.Codec
}

// Name returns the posts module's name.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec registers the posts module's types for the given codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(cdc)
}

// DefaultGenesis returns default genesis state as raw bytes for the posts module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesisState())
}

// ValidateGenesis performs genesis state validation for the posts module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	var data types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return types.ValidateGenesis(&data)
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the posts module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
}

// GetTxCmd returns the root tx command for the posts module.
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.NewTxCmd()
}

// GetQueryCmd returns the root query command for the posts module.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

// RegisterInterfaces registers interfaces and implementations of the posts module.
func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
	v2.RegisterInterfaces(registry)
	v4.RegisterInterfaces(registry)
}

// --------------------------------------------------------------------------------------------------------------------

// AppModule implements an application module for the posts module.
type AppModule struct {
	AppModuleBasic

	// To ensure setting hooks properly, keeper must be a reference as DesmosApp
	keeper *keeper.Keeper

	ak authkeeper.AccountKeeper
	bk bankkeeper.Keeper
	sk types.SubspacesKeeper

	// legacySubspace is used solely for migration of x/params managed parameters
	legacySubspace types.ParamsSubspace
}

// RegisterServices registers module services.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(*am.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)

	m := keeper.NewMigrator(*am.keeper, am.sk, am.legacySubspace)
	err := cfg.RegisterMigration(types.ModuleName, 1, m.Migrate1to2)
	if err != nil {
		panic(err)
	}
	err = cfg.RegisterMigration(types.ModuleName, 2, m.Migrate2to3)
	if err != nil {
		panic(err)
	}
	err = cfg.RegisterMigration(types.ModuleName, 3, m.Migrate3to4)
	if err != nil {
		panic(err)
	}
	err = cfg.RegisterMigration(types.ModuleName, 4, m.Migrate4to5)
	if err != nil {
		panic(err)
	}
	err = cfg.RegisterMigration(types.ModuleName, 5, m.Migrate5to6)
	if err != nil {
		panic(err)
	}
	err = cfg.RegisterMigration(types.ModuleName, 6, m.Migrate6to7)
	if err != nil {
		panic(err)
	}
}

// NewAppModule creates a new AppModule Object
func NewAppModule(
	cdc codec.Codec, keeper *keeper.Keeper, sk types.SubspacesKeeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, legacySubspace types.ParamsSubspace,
) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{cdc: cdc},
		keeper:         keeper,
		ak:             ak,
		bk:             bk,
		sk:             sk,

		legacySubspace: legacySubspace,
	}
}

// Name returns the posts module's name.
func (AppModule) Name() string {
	return types.ModuleName
}

// RegisterInvariants registers the module invariants
func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {
	keeper.RegisterInvariants(ir, *am.keeper)
}

// QuerierRoute returns the posts module's querier route name.
func (am AppModule) QuerierRoute() string {
	return types.QuerierRoute
}

// InitGenesis performs genesis initialization for the posts module.
// It returns no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState types.GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)
	am.keeper.InitGenesis(ctx, genesisState)
	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the exported genesis state as raw bytes for the
// posts module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	gs := am.keeper.ExportGenesis(ctx)
	return cdc.MustMarshalJSON(gs)
}

// ConsensusVersion implements AppModule.
func (AppModule) ConsensusVersion() uint64 {
	return consensusVersion
}

// BeginBlock returns the begin blocker for the posts module.
func (am AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}

// EndBlock returns the end blocker for the posts module. It returns no validator
// updates.
func (am AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	EndBlocker(ctx, *am.keeper)
	return []abci.ValidatorUpdate{}
}

// --------------------------------------------------------------------------------------------------------------------

// AppModuleSimulation defines the module simulation functions used by the posts module.
type AppModuleSimulation struct{}

// GenerateGenesisState creates a randomized GenState of the bank module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	simulation.RandomizeGenState(simState)
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return simulation.ProposalMsgs()
}

// RegisterStoreDecoder performs a no-op.
func (am AppModule) RegisterStoreDecoder(sdr sdk.StoreDecoderRegistry) {
	sdr[types.ModuleName] = simulation.NewDecodeStore(am.cdc)
}

// WeightedOperations returns the all the posts module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	return simulation.WeightedOperations(simState.AppParams, simState.Cdc, *am.keeper, am.sk, am.ak, am.bk)
}

// --------------------------------------------------------------------------------------------------------------------

// App Wiring Setup

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

func init() {
	appmodule.Register(
		&modulev1.Module{},
		appmodule.Provide(
			ProvideModule,
		),
		appmodule.Invoke(
			InvokeSetPostsHooks,
		),
	)
}

type ModuleInputs struct {
	depinject.In

	Config *modulev1.Module
	Cdc    codec.Codec
	Key    *storetypes.KVStoreKey

	AccountKeeper authkeeper.AccountKeeper
	BankKeeper    bankkeeper.Keeper

	ProfilesKeeper      *profileskeeper.Keeper
	SubspacesKeeper     *subspaceskeeper.Keeper
	RelationshipsKeeper relationshipskeeper.Keeper

	// LegacySubspace is used solely for migration of x/params managed parameters
	LegacySubspace types.ParamsSubspace `optional:"true"`
}

type ModuleOutputs struct {
	depinject.Out

	PostsKeeper *keeper.Keeper
	Module      appmodule.AppModule

	SubspacesHooks subspacestypes.SubspacesHooksWrapper
}

func ProvideModule(in ModuleInputs) ModuleOutputs {

	// default to governance authority if not provided
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)
	if in.Config.Authority != "" {
		authority = authtypes.NewModuleAddressOrBech32Address(in.Config.Authority)
	}

	k := keeper.NewKeeper(
		in.Cdc,
		in.Key,
		in.ProfilesKeeper,
		in.SubspacesKeeper,
		in.RelationshipsKeeper,
		authority.String(),
	)

	m := NewAppModule(
		in.Cdc,
		&k,
		in.SubspacesKeeper,
		in.AccountKeeper,
		in.BankKeeper,
		in.LegacySubspace,
	)

	return ModuleOutputs{
		PostsKeeper:    &k,
		Module:         m,
		SubspacesHooks: subspacestypes.SubspacesHooksWrapper{Hooks: k.Hooks()},
	}
}

func InvokeSetPostsHooks(
	config *modulev1.Module,
	keeper *keeper.Keeper,
	wrappers map[string]types.PostsHooksWrapper,
) error {
	// all arguments to invokers are optional
	if keeper == nil || config == nil {
		return nil
	}

	modNames := maps.Keys(wrappers)
	order := config.HooksOrder
	if len(order) == 0 {
		order = modNames
		sort.Strings(order)
	}

	if len(order) != len(modNames) {
		return fmt.Errorf("len(hooks_order: %v) != len(hooks modules: %v)", order, modNames)
	}

	if len(modNames) == 0 {
		return nil
	}

	var multiHooks types.MultiPostsHooks
	for _, modName := range order {
		wrapper, ok := wrappers[modName]
		if !ok {
			return fmt.Errorf("can't find posts hooks for module %s", modName)
		}

		multiHooks = append(multiHooks, wrapper.Hooks)
	}

	keeper.SetHooks(multiHooks)
	return nil
}
