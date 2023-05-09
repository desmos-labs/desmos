package profiles

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
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/spf13/cobra"

	modulev1 "github.com/desmos-labs/desmos/v5/api/desmos/profiles/module/v1"

	relationshipskeeper "github.com/desmos-labs/desmos/v5/x/relationships/keeper"

	v4 "github.com/desmos-labs/desmos/v5/x/profiles/legacy/v4/types"
	v5 "github.com/desmos-labs/desmos/v5/x/profiles/legacy/v5/types"

	"github.com/desmos-labs/desmos/v5/x/profiles/client/cli"
	"github.com/desmos-labs/desmos/v5/x/profiles/keeper"
	"github.com/desmos-labs/desmos/v5/x/profiles/simulation"
	"github.com/desmos-labs/desmos/v5/x/profiles/types"
)

const (
	consensusVersion = 10
)

// type check to ensure the interface is properly implemented
var (
	_ module.AppModule           = AppModule{}
	_ module.AppModuleBasic      = AppModuleBasic{}
	_ module.AppModuleSimulation = AppModule{}
	_ appmodule.AppModule        = AppModule{}
	_ depinject.OnePerModuleType = AppModule{}
)

// AppModuleBasic defines the basic application module used by the profiles module.
type AppModuleBasic struct {
	cdc         codec.Codec
	legacyAmino *codec.LegacyAmino
}

// Name returns the profiles module's name.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec registers the profiles module's types for the given codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(cdc)
}

// DefaultGenesis returns default genesis state as raw bytes for the profiles module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesisState())
}

// ValidateGenesis performs genesis state validation for the profiles module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	var data types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return types.ValidateGenesis(&data)
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the profiles module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
}

// GetTxCmd returns the root tx command for the profiles module.
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.NewTxCmd()
}

// GetQueryCmd returns the root query command for the profiles module.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

// RegisterInterfaces registers interfaces and implementations of the profiles module.
func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	v4.RegisterInterfaces(registry)
	v5.RegisterInterfaces(registry)
	types.RegisterInterfaces(registry)
}

// --------------------------------------------------------------------------------------------------------------------

// AppModule implements an application module for the profiles module.
type AppModule struct {
	AppModuleBasic
	keeper *keeper.Keeper
	ak     authkeeper.AccountKeeper
	bk     bankkeeper.Keeper

	// legacySubspace is used solely for migration of x/params managed parameters
	legacySubspace types.ParamsSubspace
}

// RegisterServices registers module services.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(*am.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)

	m := keeper.NewMigrator(am.ak, *am.keeper, am.legacySubspace)
	err := cfg.RegisterMigration(types.ModuleName, 4, m.Migrate4to5)
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
	err = cfg.RegisterMigration(types.ModuleName, 7, m.Migrate7to8)
	if err != nil {
		panic(err)
	}
	err = cfg.RegisterMigration(types.ModuleName, 8, m.Migrate8to9)
	if err != nil {
		panic(err)
	}
	err = cfg.RegisterMigration(types.ModuleName, 9, m.Migrate9to10)
	if err != nil {
		panic(err)
	}
}

// NewAppModule creates a new AppModule Object
func NewAppModule(
	cdc codec.Codec, legacyAmino *codec.LegacyAmino,
	k *keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, legacySubspace types.ParamsSubspace,
) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{cdc: cdc, legacyAmino: legacyAmino},
		keeper:         k,
		ak:             ak,
		bk:             bk,

		legacySubspace: legacySubspace,
	}
}

// Name returns the profiles module's name.
func (AppModule) Name() string {
	return types.ModuleName
}

// RegisterInvariants performs a no-op.
func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {
	keeper.RegisterInvariants(ir, *am.keeper)
}

// QuerierRoute returns the profiles module's querier route name.
func (am AppModule) QuerierRoute() string {
	return types.RouterKey
}

// InitGenesis performs genesis initialization for the profiles module.
// It returns no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState types.GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)
	am.keeper.InitGenesis(ctx, genesisState)
	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the exported genesis state as raw bytes for the
// profiles module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	gs := am.keeper.ExportGenesis(ctx)
	return cdc.MustMarshalJSON(gs)
}

// ConsensusVersion implements AppModule.
func (AppModule) ConsensusVersion() uint64 {
	return consensusVersion
}

// BeginBlock returns the begin blocker for the profiles module.
func (am AppModule) BeginBlock(ctx sdk.Context, _ abci.RequestBeginBlock) {
	BeginBlocker(ctx, *am.keeper)
}

// EndBlock returns the end blocker for the profiles module. It returns no validator
// updates.
func (am AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

// --------------------------------------------------------------------------------------------------------------------

// AppModuleSimulation defines the module simulation functions used by the profiles module.
type AppModuleSimulation struct{}

// GenerateGenesisState creates a randomized GenState of the bank module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	simulation.RandomizedGenState(simState)
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return simulation.ProposalMsgs()
}

// RegisterStoreDecoder performs a no-op.
func (am AppModule) RegisterStoreDecoder(sdr sdk.StoreDecoderRegistry) {
	sdr[types.ModuleName] = simulation.NewDecodeStore(am.cdc)
}

// WeightedOperations returns the all the profiles module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	return simulation.WeightedOperations(simState.AppParams, simState.Cdc, *am.keeper, am.ak, am.bk)
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
	)
}

type ModuleInputs struct {
	depinject.In

	Config    *modulev1.Module
	Cdc       codec.Codec
	LegacyCdc *codec.LegacyAmino
	Key       *storetypes.KVStoreKey

	AccountKeeper authkeeper.AccountKeeper
	BankKeeper    bankkeeper.Keeper

	FeesKeeper          feeskeeper.Keeper
	RelationshipsKeeper relationshipskeeper.Keeper

	// LegacySubspace is used solely for migration of x/params managed parameters
	LegacySubspace types.ParamsSubspace `optional:"true"`
}

type ModuleOutputs struct {
	depinject.Out

	ProfilesKeeper *keeper.Keeper
	Module         appmodule.AppModule
}

func ProvideModule(in ModuleInputs) ModuleOutputs {

	// default to governance authority if not provided
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)
	if in.Config.Authority != "" {
		authority = authtypes.NewModuleAddressOrBech32Address(in.Config.Authority)
	}

	k := keeper.NewKeeper(
		in.Cdc,
		in.LegacyCdc,
		in.Key,
		in.AccountKeeper,
		in.RelationshipsKeeper,
		nil,
		nil,
		nil,
		authority.String(),
	)

	m := NewAppModule(
		in.Cdc,
		in.LegacyCdc,
		&k,
		in.AccountKeeper,
		in.BankKeeper,
		in.FeesKeeper,
		in.LegacySubspace,
	)

	return ModuleOutputs{ProfilesKeeper: &k, Module: m}
}
