package subspaces

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"

	storetypes "cosmossdk.io/store/types"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"

	modulev1 "github.com/desmos-labs/desmos/v6/api/desmos/subspaces/module/v1"

	v2 "github.com/desmos-labs/desmos/v6/x/subspaces/legacy/v2"

	"github.com/desmos-labs/desmos/v6/x/subspaces/authz"
	"github.com/desmos-labs/desmos/v6/x/subspaces/client/cli"
	"github.com/desmos-labs/desmos/v6/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/v6/x/subspaces/simulation"
	"github.com/desmos-labs/desmos/v6/x/subspaces/types"
)

const (
	consensusVersion = 7
)

// type check to ensure the interface is properly implemented
var (
	_ module.AppModule           = AppModule{}
	_ module.AppModuleBasic      = AppModuleBasic{}
	_ module.AppModuleSimulation = AppModule{}
	_ appmodule.HasBeginBlocker  = AppModule{}
	_ appmodule.AppModule        = AppModule{}
)

// AppModuleBasic defines the basic application module used by the subspaces module.
type AppModuleBasic struct {
	cdc codec.Codec
}

// Name returns the subspaces module's name.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec registers the subspaces module's types for the given codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(cdc)
}

// DefaultGenesis returns default genesis state as raw bytes for the subspaces module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesisState())
}

// ValidateGenesis performs genesis state validation for the subspaces module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	var data types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return types.ValidateGenesis(&data)
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the subspaces module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
}

// GetTxCmd returns the root tx command for the subspaces module.
func (ab AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.NewTxCmd()
}

// GetQueryCmd returns the root query command for the subspaces module.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

// RegisterInterfaces registers interfaces and implementations of the subspaces module.
func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	v2.RegisterInterfaces(registry)
	types.RegisterInterfaces(registry)
	authz.RegisterInterfaces(registry)
}

// --------------------------------------------------------------------------------------------------------------------

// AppModule implements an application module for the subspaces module.
type AppModule struct {
	AppModuleBasic

	// To ensure setting hooks properly, keeper must be a reference as DesmosApp
	keeper *keeper.Keeper

	authzk authzkeeper.Keeper
	ak     authkeeper.AccountKeeper
	bk     bankkeeper.Keeper
}

// RegisterServices registers module services.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)

	m := keeper.NewMigrator(am.keeper, am.authzk, am.ak)
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
	cdc codec.Codec, keeper *keeper.Keeper, authzKeeper authzkeeper.Keeper,
	ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{cdc: cdc},
		keeper:         keeper,
		authzk:         authzKeeper,
		ak:             ak,
		bk:             bk,
	}
}

// Name returns the subspaces module's name.
func (AppModule) Name() string {
	return types.ModuleName
}

// RegisterInvariants performs a no-op.
func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {
	keeper.RegisterInvariants(ir, am.keeper)
}

// QuerierRoute returns the subspaces module's querier route name.
func (am AppModule) QuerierRoute() string {
	return types.QuerierRoute
}

// InitGenesis performs genesis initialization for the subspaces module.
// It returns no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState types.GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)
	am.keeper.InitGenesis(ctx, genesisState)
	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the exported genesis state as raw bytes for the
// subspaces module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	gs := am.keeper.ExportGenesis(ctx)
	return cdc.MustMarshalJSON(gs)
}

// ConsensusVersion implements AppModule.
func (AppModule) ConsensusVersion() uint64 {
	return consensusVersion
}

// BeginBlock returns the begin blocker for the subspaces module.
func (am AppModule) BeginBlock(ctx context.Context) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	BeginBlocker(sdkCtx, am.keeper)
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// AppModuleSimulation defines the module simulation functions used by the subspaces module.
type AppModuleSimulation struct{}

// GenerateGenesisState creates a randomized GenState of the bank module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	simulation.RandomizeGenState(simState)
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return simulation.ProposalMsgs(am.keeper)
}

// RegisterStoreDecoder performs a no-op.
func (am AppModule) RegisterStoreDecoder(sdr sdk.StoreDecoderRegistry) {
	sdr[types.ModuleName] = simulation.NewDecodeStore(am.cdc)
}

// WeightedOperations returns the all the subspaces module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	return simulation.WeightedOperations(simState.AppParams, simState.Cdc, am.keeper, am.ak, am.bk, am.authzk)
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
			InvokeSetSubspacesHooks,
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
	AuthzKeeper   authzkeeper.Keeper
}

type ModuleOutputs struct {
	depinject.Out

	SubspacesKeeper *keeper.Keeper
	Module          appmodule.AppModule
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
		in.AccountKeeper,
		in.AuthzKeeper,
		authority.String(),
	)

	m := NewAppModule(
		in.Cdc,
		k,
		in.AuthzKeeper,
		in.AccountKeeper,
		in.BankKeeper,
	)

	return ModuleOutputs{SubspacesKeeper: k, Module: m}
}

func InvokeSetSubspacesHooks(
	config *modulev1.Module,
	keeper *keeper.Keeper,
	wrappers map[string]types.SubspacesHooksWrapper,
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

	var multiHooks types.MultiSubspacesHooks
	for _, modName := range order {
		wrapper, ok := wrappers[modName]
		if !ok {
			return fmt.Errorf("can't find subspaces hooks for module %s", modName)
		}

		multiHooks = append(multiHooks, wrapper.Hooks)
	}

	keeper.SetHooks(multiHooks)
	return nil
}
