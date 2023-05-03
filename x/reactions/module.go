package reactions

import (
	"context"
	"encoding/json"
	"fmt"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"

	"github.com/cosmos/cosmos-sdk/client"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/spf13/cobra"

	modulev1 "github.com/desmos-labs/desmos/v5/api/desmos/reactions/module/v1"

	postskeeper "github.com/desmos-labs/desmos/v5/x/posts/keeper"
	subspaceskeeper "github.com/desmos-labs/desmos/v5/x/subspaces/keeper"

	"github.com/desmos-labs/desmos/v5/x/reactions/client/cli"
	"github.com/desmos-labs/desmos/v5/x/reactions/keeper"
	"github.com/desmos-labs/desmos/v5/x/reactions/simulation"
	"github.com/desmos-labs/desmos/v5/x/reactions/types"
)

const (
	consensusVersion = 3
)

// type check to ensure the interface is properly implemented
var (
	_ module.AppModule           = AppModule{}
	_ module.AppModuleBasic      = AppModuleBasic{}
	_ module.AppModuleSimulation = AppModule{}
	_ appmodule.AppModule        = AppModule{}
	_ depinject.OnePerModuleType = AppModule{}
)

// AppModuleBasic defines the basic application module used by the reactions module.
type AppModuleBasic struct {
	cdc codec.Codec
}

// Name returns the reactions module's name.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec registers the reactions module's types for the given codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(cdc)
}

// DefaultGenesis returns default genesis state as raw bytes for the reactions module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesisState())
}

// ValidateGenesis performs genesis state validation for the reactions module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	var data types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return types.ValidateGenesis(&data)
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the reactions module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
}

// GetTxCmd returns the root tx command for the reactions module.
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.NewTxCmd()
}

// GetQueryCmd returns the root query command for the reactions module.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

// RegisterInterfaces registers interfaces and implementations of the reactions module.
func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

// --------------------------------------------------------------------------------------------------------------------

// AppModule implements an application module for the reactions module.
type AppModule struct {
	AppModuleBasic
	keeper keeper.Keeper
	ak     authkeeper.AccountKeeper
	bk     bankkeeper.Keeper

	profilesKeeper types.ProfilesKeeper
	sk             subspaceskeeper.Keeper
	pk             postskeeper.Keeper
}

// RegisterServices registers module services.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)

	m := keeper.NewMigrator(am.keeper, am.sk, am.pk)
	err := cfg.RegisterMigration(types.ModuleName, 1, m.Migrate1to2)
	if err != nil {
		panic(err)
	}

	err = cfg.RegisterMigration(types.ModuleName, 2, m.Migrate2to3)
	if err != nil {
		panic(err)
	}
}

// NewAppModule creates a new AppModule Object
func NewAppModule(
	cdc codec.Codec,
	k keeper.Keeper, profilesKeeper types.ProfilesKeeper, sk subspaceskeeper.Keeper, pk postskeeper.Keeper,
	ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{cdc: cdc},
		keeper:         k,
		ak:             ak,
		bk:             bk,

		profilesKeeper: profilesKeeper,
		sk:             sk,
		pk:             pk,
	}
}

// Name returns the reactions module's name.
func (AppModule) Name() string {
	return types.ModuleName
}

// RegisterInvariants performs a no-op.
func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {
	keeper.RegisterInvariants(ir, am.keeper)
}

// QuerierRoute returns the reactions module's querier route name.
func (am AppModule) QuerierRoute() string {
	return types.RouterKey
}

// InitGenesis performs genesis initialization for the reactions module.
// It returns no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState types.GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)
	am.keeper.InitGenesis(ctx, genesisState)
	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the exported genesis state as raw bytes for the
// reactions module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	gs := am.keeper.ExportGenesis(ctx)
	return cdc.MustMarshalJSON(gs)
}

// ConsensusVersion implements AppModule.
func (AppModule) ConsensusVersion() uint64 {
	return consensusVersion
}

// BeginBlock returns the begin blocker for the reactions module.
func (am AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {
}

// EndBlock returns the end blocker for the reactions module. It returns no validator
// updates.
func (am AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

// --------------------------------------------------------------------------------------------------------------------

// AppModuleSimulation defines the module simulation functions used by the reactions module.
type AppModuleSimulation struct{}

// GenerateGenesisState creates a randomized GenState of the bank module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	simulation.RandomizedGenState(simState)
}

// RegisterStoreDecoder performs a no-op.
func (am AppModule) RegisterStoreDecoder(sdr sdk.StoreDecoderRegistry) {
	sdr[types.ModuleName] = simulation.NewDecodeStore(am.cdc)
}

// WeightedOperations returns the all the reactions module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	return simulation.WeightedOperations(simState.AppParams, simState.Cdc, am.keeper, am.profilesKeeper, am.sk, am.pk, am.ak, am.bk)
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
			provideModule,
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

	ProfilesKeeper      types.ProfilesKeeper
	SubspacesKeeper     subspaceskeeper.Keeper
	PostsKeeper         postskeeper.Keeper
	RelationshipsKeeper types.RelationshipsKeeper
}

type ModuleOutputs struct {
	depinject.Out

	ReactionsKeeper keeper.Keeper
	Module          appmodule.AppModule
}

func provideModule(in ModuleInputs) ModuleOutputs {

	k := keeper.NewKeeper(
		in.Cdc,
		in.Key,
		in.ProfilesKeeper,
		in.SubspacesKeeper,
		in.RelationshipsKeeper,
		in.PostsKeeper,
	)

	m := NewAppModule(
		in.Cdc,
		k,
		in.ProfilesKeeper,
		in.SubspacesKeeper,
		in.PostsKeeper,
		in.AccountKeeper,
		in.BankKeeper,
	)

	return ModuleOutputs{ReactionsKeeper: k, Module: m}
}
