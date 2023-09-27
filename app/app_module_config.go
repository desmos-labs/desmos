package app

import (
	"cosmossdk.io/depinject"

	runtimev1alpha1 "cosmossdk.io/api/cosmos/app/runtime/v1alpha1"
	appv1alpha1 "cosmossdk.io/api/cosmos/app/v1alpha1"
	authmodulev1 "cosmossdk.io/api/cosmos/auth/module/v1"
	authzmodulev1 "cosmossdk.io/api/cosmos/authz/module/v1"
	bankmodulev1 "cosmossdk.io/api/cosmos/bank/module/v1"
	consensusmodulev1 "cosmossdk.io/api/cosmos/consensus/module/v1"
	crisismodulev1 "cosmossdk.io/api/cosmos/crisis/module/v1"
	distrmodulev1 "cosmossdk.io/api/cosmos/distribution/module/v1"
	evidencemodulev1 "cosmossdk.io/api/cosmos/evidence/module/v1"
	feegrantmodulev1 "cosmossdk.io/api/cosmos/feegrant/module/v1"
	genutilmodulev1 "cosmossdk.io/api/cosmos/genutil/module/v1"
	govmodulev1 "cosmossdk.io/api/cosmos/gov/module/v1"
	mintmodulev1 "cosmossdk.io/api/cosmos/mint/module/v1"
	paramsmodulev1 "cosmossdk.io/api/cosmos/params/module/v1"
	slashingmodulev1 "cosmossdk.io/api/cosmos/slashing/module/v1"
	stakingmodulev1 "cosmossdk.io/api/cosmos/staking/module/v1"
	txconfigv1 "cosmossdk.io/api/cosmos/tx/config/v1"
	upgrademodulev1 "cosmossdk.io/api/cosmos/upgrade/module/v1"
	vestingmodulev1 "cosmossdk.io/api/cosmos/vesting/module/v1"

	"cosmossdk.io/core/appconfig"
	evidencetypes "cosmossdk.io/x/evidence/types"
	"cosmossdk.io/x/feegrant"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	solomachine "github.com/cosmos/ibc-go/v8/modules/light-clients/06-solomachine"
	ibctm "github.com/cosmos/ibc-go/v8/modules/light-clients/07-tendermint"

	_ "cosmossdk.io/x/evidence"                       // import for side-effects
	_ "cosmossdk.io/x/feegrant/module"                // import for side-effects
	_ "cosmossdk.io/x/upgrade"                        // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/auth/tx/config" // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/auth/vesting"   // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/authz/module"   // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/bank"           // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/consensus"      // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/crisis"         // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/distribution"   // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/mint"           // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/params"         // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/slashing"       // import for side-effects
	_ "github.com/cosmos/cosmos-sdk/x/staking"        // import for side-effects

	icatypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"

	postsmodulev1 "github.com/desmos-labs/desmos/v6/api/desmos/posts/module/v1"
	profilesmodulev1 "github.com/desmos-labs/desmos/v6/api/desmos/profiles/module/v1"
	reactionsmodulev1 "github.com/desmos-labs/desmos/v6/api/desmos/reactions/module/v1"
	relationshipsmodulev1 "github.com/desmos-labs/desmos/v6/api/desmos/relationships/module/v1"
	reportsmodulev1 "github.com/desmos-labs/desmos/v6/api/desmos/reports/module/v1"
	subspacesmodulev1 "github.com/desmos-labs/desmos/v6/api/desmos/subspaces/module/v1"
	supplymodulev1 "github.com/desmos-labs/desmos/v6/api/desmos/supply/module/v1"
	tokenfactorymodulev1 "github.com/desmos-labs/desmos/v6/api/desmos/tokenfactory/module/v1"

	_ "github.com/desmos-labs/desmos/v6/x/posts"         // import for side-effects
	_ "github.com/desmos-labs/desmos/v6/x/profiles"      // import for side-effects
	_ "github.com/desmos-labs/desmos/v6/x/reactions"     // import for side-effects
	_ "github.com/desmos-labs/desmos/v6/x/relationships" // import for side-effects
	_ "github.com/desmos-labs/desmos/v6/x/reports"       // import for side-effects
	_ "github.com/desmos-labs/desmos/v6/x/subspaces"     // import for side-effects
	_ "github.com/desmos-labs/desmos/v6/x/supply"        // import for side-effects
	_ "github.com/desmos-labs/desmos/v6/x/tokenfactory"  // import for side-effects

	poststypes "github.com/desmos-labs/desmos/v6/x/posts/types"
	profilestypes "github.com/desmos-labs/desmos/v6/x/profiles/types"
	reactionstypes "github.com/desmos-labs/desmos/v6/x/reactions/types"
	relationshipstypes "github.com/desmos-labs/desmos/v6/x/relationships/types"
	reportstypes "github.com/desmos-labs/desmos/v6/x/reports/types"
	subspacestypes "github.com/desmos-labs/desmos/v6/x/subspaces/types"
	supplytypes "github.com/desmos-labs/desmos/v6/x/supply/types"
	tokenfactorytypes "github.com/desmos-labs/desmos/v6/x/tokenfactory/types"
)

var (
	// blocked account addresses
	blockAccAddrs = []string{
		authtypes.FeeCollectorName,
		distrtypes.ModuleName,
		minttypes.ModuleName,
		stakingtypes.BondedPoolName,
		stakingtypes.NotBondedPoolName,
		ibctransfertypes.ModuleName,
		ibcfeetypes.ModuleName,
		icatypes.ModuleName,
		wasmtypes.ModuleName,
		tokenfactorytypes.ModuleName,

		// We allow the following module accounts to receive funds:
		// govtypes.ModuleName
	}

	// application configuration (used by depinject)
	AppConfig = depinject.Configs(appconfig.Compose(&appv1alpha1.Config{
		Modules: []*appv1alpha1.ModuleConfig{
			// SDK modules
			{
				Name: runtime.ModuleName,
				Config: appconfig.WrapAny(&runtimev1alpha1.Module{
					AppName:       appName,
					BeginBlockers: beginBlockerOrder,
					EndBlockers:   endBlockerOrder,
					OverrideStoreKeys: []*runtimev1alpha1.StoreKeyConfig{
						{
							ModuleName: authtypes.ModuleName,
							KvStoreKey: "acc",
						},
					},
					InitGenesis:     genesisModuleOrder,
					OrderMigrations: migrationModuleOrder,
				}),
			},
			{
				Name: authtypes.ModuleName,
				Config: appconfig.WrapAny(&authmodulev1.Module{
					Bech32Prefix:             sdk.Bech32MainPrefix,
					ModuleAccountPermissions: maccPerms,
					// By default modules authority is the governance module. This is configurable with the following:
					// Authority: "group", // A custom module authority can be set using a module name
					// Authority: "cosmos1cwwv22j5ca08ggdv9c2uky355k908694z577tv", // or a specific address
				}),
			},
			{
				Name:   vestingtypes.ModuleName,
				Config: appconfig.WrapAny(&vestingmodulev1.Module{}),
			},
			{
				Name: banktypes.ModuleName,
				Config: appconfig.WrapAny(&bankmodulev1.Module{
					BlockedModuleAccountsOverride: blockAccAddrs,
				}),
			},
			{
				Name:   stakingtypes.ModuleName,
				Config: appconfig.WrapAny(&stakingmodulev1.Module{
					// NOTE: specifying a prefix is only necessary when using bech32 addresses
					// If not specfied, the auth Bech32Prefix appended with "valoper" and "valcons" is used by default
					// Bech32PrefixValidator: "desmosvaloper",
					// Bech32PrefixConsensus: "desmosvalcons",
				}),
			},
			{
				Name:   slashingtypes.ModuleName,
				Config: appconfig.WrapAny(&slashingmodulev1.Module{}),
			},
			{
				Name:   paramstypes.ModuleName,
				Config: appconfig.WrapAny(&paramsmodulev1.Module{}),
			},
			{
				Name:   "tx",
				Config: appconfig.WrapAny(&txconfigv1.Config{}),
			},
			{
				Name:   genutiltypes.ModuleName,
				Config: appconfig.WrapAny(&genutilmodulev1.Module{}),
			},
			{
				Name:   authz.ModuleName,
				Config: appconfig.WrapAny(&authzmodulev1.Module{}),
			},
			{
				Name:   upgradetypes.ModuleName,
				Config: appconfig.WrapAny(&upgrademodulev1.Module{}),
			},
			{
				Name:   distrtypes.ModuleName,
				Config: appconfig.WrapAny(&distrmodulev1.Module{}),
			},
			{
				Name:   evidencetypes.ModuleName,
				Config: appconfig.WrapAny(&evidencemodulev1.Module{}),
			},
			{
				Name:   minttypes.ModuleName,
				Config: appconfig.WrapAny(&mintmodulev1.Module{}),
			},
			{
				Name:   feegrant.ModuleName,
				Config: appconfig.WrapAny(&feegrantmodulev1.Module{}),
			},
			{
				Name:   govtypes.ModuleName,
				Config: appconfig.WrapAny(&govmodulev1.Module{}),
			},
			{
				Name:   crisistypes.ModuleName,
				Config: appconfig.WrapAny(&crisismodulev1.Module{}),
			},
			{
				Name:   consensustypes.ModuleName,
				Config: appconfig.WrapAny(&consensusmodulev1.Module{}),
			},

			// Desmos modules
			{
				Name:   poststypes.ModuleName,
				Config: appconfig.WrapAny(&postsmodulev1.Module{}),
			},
			{
				Name:   profilestypes.ModuleName,
				Config: appconfig.WrapAny(&profilesmodulev1.Module{}),
			},
			{
				Name:   reactionstypes.ModuleName,
				Config: appconfig.WrapAny(&reactionsmodulev1.Module{}),
			},
			{
				Name:   relationshipstypes.ModuleName,
				Config: appconfig.WrapAny(&relationshipsmodulev1.Module{}),
			},
			{
				Name:   reportstypes.ModuleName,
				Config: appconfig.WrapAny(&reportsmodulev1.Module{}),
			},
			{
				Name:   subspacestypes.ModuleName,
				Config: appconfig.WrapAny(&subspacesmodulev1.Module{}),
			},
			{
				Name:   supplytypes.ModuleName,
				Config: appconfig.WrapAny(&supplymodulev1.Module{}),
			},
			{
				Name:   tokenfactorytypes.ModuleName,
				Config: appconfig.WrapAny(&tokenfactorymodulev1.Module{}),
			},
		},
	}), depinject.Supply(
		// supply custom module basics
		map[string]module.AppModuleBasic{
			genutiltypes.ModuleName: genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),
			govtypes.ModuleName: gov.NewAppModuleBasic(
				[]govclient.ProposalHandler{
					paramsclient.ProposalHandler,
				},
			),
			ibctm.ModuleName:       ibctm.AppModuleBasic{},
			solomachine.ModuleName: solomachine.AppModuleBasic{},
		},
	))
)
