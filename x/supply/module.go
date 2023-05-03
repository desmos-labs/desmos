package supply

import (
	"context"
	"encoding/json"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	distributionkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	"github.com/gorilla/mux"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	modulev1 "github.com/desmos-labs/desmos/v5/api/desmos/supply/module/v1"

	"github.com/desmos-labs/desmos/v5/x/supply/client/cli"
	"github.com/desmos-labs/desmos/v5/x/supply/keeper"
	"github.com/desmos-labs/desmos/v5/x/supply/types"
)

var (
	_ module.AppModule           = AppModule{}
	_ module.AppModuleBasic      = AppModuleBasic{}
	_ module.AppModuleSimulation = AppModule{}
	_ appmodule.AppModule        = AppModule{}
	_ depinject.OnePerModuleType = AppModule{}
)

// AppModuleBasic defines the basic application module used by the supply module.
type AppModuleBasic struct {
	cdc codec.Codec
}

var _ module.AppModuleBasic = AppModuleBasic{}

// Name returns the supply module's name.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec registers the supply module's types on the given LegacyAmino codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(_ *codec.LegacyAmino) {}

// RegisterInterfaces registers the module's interface types
func (b AppModuleBasic) RegisterInterfaces(_ cdctypes.InterfaceRegistry) {}

// DefaultGenesis returns default genesis state as raw bytes for the supply
// module.
func (AppModuleBasic) DefaultGenesis(_ codec.JSONCodec) json.RawMessage {
	return nil
}

// ValidateGenesis performs genesis state validation for the supply module.
func (AppModuleBasic) ValidateGenesis(_ codec.JSONCodec, _ client.TxEncodingConfig, _ json.RawMessage) error {
	return nil
}

// RegisterRESTRoutes registers the REST routes for the supply module.
func (AppModuleBasic) RegisterRESTRoutes(clientCtx client.Context, rtr *mux.Router) {
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the supply module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
}

// GetTxCmd returns the root tx command for the supply module.
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return nil
}

// GetQueryCmd returns no root query command for the supply module.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

// AppModule implements an application module for the supply module.
type AppModule struct {
	AppModuleBasic
	keeper keeper.Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(cdc codec.Codec, keeper keeper.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{cdc: cdc},
		keeper:         keeper,
	}
}

// RegisterServices registers module services.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

// Name returns the supply module's name.
func (AppModule) Name() string {
	return types.ModuleName
}

// RegisterInvariants registers the supply module invariants.
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// QuerierRoute returns the supply module's querier route name.
func (AppModule) QuerierRoute() string {
	return types.QuerierRoute
}

// InitGenesis performs genesis initialization for the supply module. It returns
// no validator updates.
func (am AppModule) InitGenesis(_ sdk.Context, _ codec.JSONCodec, _ json.RawMessage) []abci.ValidatorUpdate {
	return nil
}

// ExportGenesis returns the exported genesis state as raw bytes for the supply
// module.
func (am AppModule) ExportGenesis(_ sdk.Context, _ codec.JSONCodec) json.RawMessage {
	return nil
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return 1 }

// BeginBlock returns the begin blocker for the supply module.
func (am AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}

// EndBlock returns the end blocker for the supply module. It returns no validator
// updates.
func (am AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// AppModuleSimulation functions

// GenerateGenesisState creates a randomized GenState of the supply module.
func (AppModule) GenerateGenesisState(_ *module.SimulationState) {}

// RegisterStoreDecoder registers a decoder for supply module's types
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the supply module operations with their respective weights.
func (am AppModule) WeightedOperations(_ module.SimulationState) []simtypes.WeightedOperation {
	return nil
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

	AccountKeeper      authkeeper.AccountKeeper
	BankKeeper         bankkeeper.Keeper
	DistributionKeeper distributionkeeper.Keeper
}

type ModuleOutputs struct {
	depinject.Out

	SupplyKeeper keeper.Keeper
	Module       appmodule.AppModule
}

func provideModule(in ModuleInputs) ModuleOutputs {

	k := keeper.NewKeeper(
		in.Cdc,
		in.AccountKeeper,
		in.BankKeeper,
		in.DistributionKeeper,
	)

	m := NewAppModule(in.Cdc, k)

	return ModuleOutputs{SupplyKeeper: k, Module: m}
}
