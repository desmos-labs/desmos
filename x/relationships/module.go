package relationships

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"

	feeskeeper "github.com/desmos-labs/desmos/v3/x/fees/keeper"

	profilesv4 "github.com/desmos-labs/desmos/v3/x/profiles/legacy/v4"

	subspaceskeeper "github.com/desmos-labs/desmos/v3/x/subspaces/keeper"

	"github.com/cosmos/cosmos-sdk/client"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/desmos-labs/desmos/v3/x/relationships/client/cli"
	"github.com/desmos-labs/desmos/v3/x/relationships/keeper"
	"github.com/desmos-labs/desmos/v3/x/relationships/simulation"
	"github.com/desmos-labs/desmos/v3/x/relationships/types"
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

// AppModuleBasic defines the basic application module used by the relationships module.
type AppModuleBasic struct {
	cdc codec.Codec
}

// Name returns the relationships module's name.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec registers the relationships module's types for the given codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(cdc)
}

// DefaultGenesis returns default genesis state as raw bytes for the relationships module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesisState())
}

// ValidateGenesis performs genesis state validation for the relationships module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	var data types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return types.ValidateGenesis(&data)
}

// RegisterRESTRoutes registers the REST routes for the relationships module.
func (AppModuleBasic) RegisterRESTRoutes(_ client.Context, _ *mux.Router) {}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the relationships module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
}

// GetTxCmd returns the root tx command for the relationships module.
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.NewTxCmd()
}

// GetQueryCmd returns the root query command for the relationships module.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

// RegisterInterfaces registers interfaces and implementations of the relationships module.
func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

// --------------------------------------------------------------------------------------------------------------------

// AppModule implements an application module for the relationships module.
type AppModule struct {
	AppModuleBasic
	keeper keeper.Keeper
	pk     profilesv4.Keeper
	sk     subspaceskeeper.Keeper
	ak     authkeeper.AccountKeeper
	bk     bankkeeper.Keeper
	fk     feeskeeper.Keeper
}

// RegisterServices registers module services.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)

	m := keeper.NewMigrator(am.keeper, am.pk)
	err := cfg.RegisterMigration(types.ModuleName, 1, m.Migrate1To2)
	if err != nil {
		panic(err)
	}
}

// NewAppModule creates a new AppModule Object
func NewAppModule(
	cdc codec.Codec,
	k keeper.Keeper, sk subspaceskeeper.Keeper, pk profilesv4.Keeper,
	ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{cdc: cdc},
		keeper:         k,
		ak:             ak,
		bk:             bk,
		sk:             sk,
		pk:             pk,
		fk:             fk,
	}
}

// Name returns the relationships module's name.
func (AppModule) Name() string {
	return types.ModuleName
}

// RegisterInvariants performs a no-op.
func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {
	keeper.RegisterInvariants(ir, am.keeper)
}

// Deprecated: Route returns the module's message router and handler.
func (am AppModule) Route() sdk.Route {
	return sdk.Route{}
}

// QuerierRoute returns the relationships module's querier route name.
func (am AppModule) QuerierRoute() string {
	return types.RouterKey
}

// LegacyQuerierHandler returns the relationships module sdk.Querier.
func (am AppModule) LegacyQuerierHandler(legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return nil
}

// InitGenesis performs genesis initialization for the relationships module.
// It returns no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState types.GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)
	am.keeper.InitGenesis(ctx, genesisState)
	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the exported genesis state as raw bytes for the
// relationships module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	gs := am.keeper.ExportGenesis(ctx)
	return cdc.MustMarshalJSON(gs)
}

// ConsensusVersion implements AppModule.
func (AppModule) ConsensusVersion() uint64 {
	return consensusVersion
}

// BeginBlock returns the begin blocker for the relationships module.
func (am AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {
}

// EndBlock returns the end blocker for the relationships module. It returns no validator
// updates.
func (am AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

// --------------------------------------------------------------------------------------------------------------------

// AppModuleSimulation defines the module simulation functions used by the relationships module.
type AppModuleSimulation struct{}

// GenerateGenesisState creates a randomized GenState of the bank module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	simulation.RandomizedGenState(simState)
}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized profile param changes for the simulator.
func (AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	return nil
}

// RegisterStoreDecoder performs a no-op.
func (am AppModule) RegisterStoreDecoder(sdr sdk.StoreDecoderRegistry) {
	sdr[types.ModuleName] = simulation.NewDecodeStore(am.cdc)
}

// WeightedOperations returns the all the relationships module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	return simulation.WeightedOperations(simState.AppParams, simState.Cdc, am.keeper, am.sk, am.ak, am.bk, am.fk)
}
