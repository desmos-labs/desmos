package reports

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"

	subspaceskeeper "github.com/desmos-labs/desmos/v4/x/subspaces/keeper"

	feeskeeper "github.com/desmos-labs/desmos/v4/x/fees/keeper"
	postskeeper "github.com/desmos-labs/desmos/v4/x/posts/keeper"

	"github.com/desmos-labs/desmos/v4/x/reports/simulation"

	"github.com/desmos-labs/desmos/v4/x/reports/client/cli"

	"github.com/desmos-labs/desmos/v4/x/reports/keeper"
	"github.com/desmos-labs/desmos/v4/x/reports/types"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
)

const (
	consensusVersion = 2
)

// type check to ensure the interface is properly implemented
var (
	_ module.AppModule           = AppModule{}
	_ module.AppModuleBasic      = AppModuleBasic{}
	_ module.AppModuleSimulation = AppModule{}
)

// AppModuleBasic defines the basic application module used by the reports module.
type AppModuleBasic struct {
	cdc codec.Codec
}

// Name returns the reports module's name.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec registers the reports module's types for the given codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(cdc)
}

// DefaultGenesis returns default genesis state as raw bytes for the reports module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesisState())
}

// ValidateGenesis performs genesis state validation for the reports module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	var data types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return types.ValidateGenesis(&data)
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the reports module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
}

// GetTxCmd returns the root tx command for the reports module.
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.NewTxCmd()
}

// GetQueryCmd returns the root query command for the reports module.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

// RegisterInterfaces registers interfaces and implementations of the reports module.
func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

// --------------------------------------------------------------------------------------------------------------------

// AppModule implements an application module for the reports module.
type AppModule struct {
	AppModuleBasic
	keeper keeper.Keeper
	sk     subspaceskeeper.Keeper
	pk     postskeeper.Keeper
	ak     authkeeper.AccountKeeper
	bk     bankkeeper.Keeper
	fk     feeskeeper.Keeper

	// legacySubspace is used solely for migration of x/params managed parameters
	legacySubspace types.ParamsSubspace
}

// RegisterServices registers module services.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)

	m := keeper.NewMigrator(am.keeper, am.sk, am.legacySubspace)
	err := cfg.RegisterMigration(types.ModuleName, 1, m.Migrate1to2)
	if err != nil {
		panic(err)
	}
}

// NewAppModule creates a new AppModule Object
func NewAppModule(
	cdc codec.Codec, keeper keeper.Keeper, sk subspaceskeeper.Keeper, pk postskeeper.Keeper,
	ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper, legacySubspace types.ParamsSubspace,
) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{cdc: cdc},
		keeper:         keeper,
		sk:             sk,
		pk:             pk,
		ak:             ak,
		bk:             bk,
		fk:             fk,

		legacySubspace: legacySubspace,
	}
}

// Name returns the reports module's name.
func (AppModule) Name() string {
	return types.ModuleName
}

// RegisterInvariants performs a no-op.
func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {
	keeper.RegisterInvariants(ir, am.keeper)
}

// QuerierRoute returns the reports module's querier route name.
func (am AppModule) QuerierRoute() string {
	return types.QuerierRoute
}

// InitGenesis performs genesis initialization for the reports module.
// It returns no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState types.GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)
	am.keeper.InitGenesis(ctx, genesisState)
	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the exported genesis state as raw bytes for the
// reports module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	gs := am.keeper.ExportGenesis(ctx)
	return cdc.MustMarshalJSON(gs)
}

// ConsensusVersion implements AppModule.
func (AppModule) ConsensusVersion() uint64 {
	return consensusVersion
}

// BeginBlock returns the begin blocker for the reports module.
func (am AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {
}

// EndBlock returns the end blocker for the reports module. It returns no validator
// updates.
func (am AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

// --------------------------------------------------------------------------------------------------------------------

// AppModuleSimulation defines the module simulation functions used by the reports module.
type AppModuleSimulation struct{}

// GenerateGenesisState creates a randomized GenState of the bank module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	simulation.RandomizeGenState(simState)
}

// RandomizedParams creates randomized reports param changes for the simulator.
func (AppModule) RandomizedParams(r *rand.Rand) []simtypes.LegacyParamChange {
	return simulation.ParamChanges(r)
}

// RegisterStoreDecoder performs a no-op.
func (am AppModule) RegisterStoreDecoder(sdr sdk.StoreDecoderRegistry) {
	sdr[types.ModuleName] = simulation.NewDecodeStore(am.cdc)
}

// WeightedOperations returns the all the reports module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	return simulation.WeightedOperations(simState.AppParams, simState.Cdc, am.keeper, am.sk, am.pk, am.ak, am.bk, am.fk)
}
