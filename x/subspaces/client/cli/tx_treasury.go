package cli

import (
	"fmt"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzcli "github.com/cosmos/cosmos-sdk/x/authz/client/cli"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

// DONTCOVER

const (
	delegate   = "delegate"
	redelegate = "redelegate"
	unbond     = "unbond"
)

// GetTreasuryTxCmd returns a new command to perform subspaces treasury transactions
func GetTreasuryTxCmd() *cobra.Command {
	treasuryTxCmd := &cobra.Command{
		Use:                        "treasury",
		Short:                      "Tx commands for subspace treasury",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	treasuryTxCmd.AddCommand(
		GetCmdGrantTreasuryAuthorization(),
		GetCmdRevokeTreasuryAuthorization(),
	)

	return treasuryTxCmd
}

// GetCmdGrantTreasuryAuthorization returns the command used to grant a treasury authorization
func GetCmdGrantTreasuryAuthorization() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "grant [subspace-id] [grantee] [authorization_type=\"send\"|\"generic\"|\"delegate\"|\"unbond\"|\"redelegate\"] --from [granter]",
		Short: "Grant a treasury authorization to a user",
		Long: strings.TrimSpace(
			fmt.Sprintf(`grant treasury authorization to an address to execute a transaction on your behalf:
Examples:
 $ %[1]s tx %[2]s grant desmos1463vltcqk6ql6zpk0g6s595jjcrzk4804hyqw7 send --spend-limit=1000stake --from=desmos1463vltcqk6ql6zpk0g6s595jjcrzk4804hyqw7
 $ %[1]s tx %[2]s grant desmos1463vltcqk6ql6zpk0g6s595jjcrzk4804hyqw7 generic --msg-type=/cosmos.gov.v1beta1.MsgVote --from=desmos1463vltcqk6ql6zpk0g6s595jjcrzk4804hyqw7
	`, version.AppName, types.ModuleName),
		),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := types.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			exp, err := cmd.Flags().GetInt64(authzcli.FlagExpiration)
			if err != nil {
				return err
			}

			var authorization authz.Authorization
			switch args[2] {
			case "send":
				authorization, err = getSendAuthorization(cmd.Flags())
				if err != nil {
					return err
				}

			case "generic":
				authorization, err = getGenericAuthorization(cmd.Flags())
				if err != nil {
					return err
				}

			case delegate, unbond, redelegate:
				authorization, err = getStakeAuthorization(cmd.Flags(), args[2])
				if err != nil {
					return err
				}

			default:
				return fmt.Errorf("invalid authorization type, %s", args[2])
			}

			expiration := time.Unix(exp, 0)
			msg := types.NewMsgGrantTreasuryAuthorization(subspaceID, clientCtx.GetFromAddress().String(), args[1], authorization, &expiration)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(authzcli.FlagMsgType, "", "The Msg method name for which we are creating a GenericAuthorization")
	cmd.Flags().String(authzcli.FlagSpendLimit, "", "SpendLimit for Send Authorization, an array of Coins allowed spend")
	cmd.Flags().StringSlice(authzcli.FlagAllowedValidators, []string{}, "Allowed validators addresses separated by ,")
	cmd.Flags().StringSlice(authzcli.FlagDenyValidators, []string{}, "Deny validators addresses separated by ,")
	cmd.Flags().Int64(authzcli.FlagExpiration, time.Now().AddDate(1, 0, 0).Unix(), "The Unix timestamp. Default is one year.")
	cmd.Flags().StringSlice(authzcli.FlagAllowList, []string{}, "Allowed addresses grantee is allowed to send funds separated by ,")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// getStakeAuthorization returns a send authorization from the given command flags
func getSendAuthorization(flags *pflag.FlagSet) (*banktypes.SendAuthorization, error) {
	limit, err := flags.GetString(authzcli.FlagSpendLimit)
	if err != nil {
		return nil, err
	}

	spendLimit, err := sdk.ParseCoinsNormalized(limit)
	if err != nil {
		return nil, err
	}

	if !spendLimit.IsAllPositive() {
		return nil, fmt.Errorf("spend-limit should be greater than zero")
	}

	allowed, err := getAllowedListFromFlags(flags)
	if err != nil {
		return nil, err
	}

	return banktypes.NewSendAuthorization(spendLimit, allowed), nil
}

// getStakeAuthorization returns a generic authorization from the given command flags
func getGenericAuthorization(flags *pflag.FlagSet) (*authz.GenericAuthorization, error) {
	msgType, err := flags.GetString(authzcli.FlagMsgType)
	if err != nil {
		return nil, err
	}
	return authz.NewGenericAuthorization(msgType), nil
}

// getStakeAuthorization returns a stake authorization from the given command flags
func getStakeAuthorization(flags *pflag.FlagSet, stakingType string) (*stakingtypes.StakeAuthorization, error) {
	limit, err := flags.GetString(authzcli.FlagSpendLimit)
	if err != nil {
		return nil, err
	}

	var delegateLimit *sdk.Coin
	if limit != "" {
		spendLimit, err := sdk.ParseCoinNormalized(limit)
		if err != nil {
			return nil, err
		}

		if !spendLimit.IsPositive() {
			return nil, fmt.Errorf("spend-limit should be greater than zero")
		}
		delegateLimit = &spendLimit
	}

	allowed, err := getValidatorAddressesFromFlags(flags, authzcli.FlagAllowedValidators)
	if err != nil {
		return nil, err
	}

	denied, err := getValidatorAddressesFromFlags(flags, authzcli.FlagDenyValidators)
	if err != nil {
		return nil, err
	}

	var authorizationType stakingtypes.AuthorizationType
	switch stakingType {
	case delegate:
		authorizationType = stakingtypes.AuthorizationType_AUTHORIZATION_TYPE_DELEGATE
	case unbond:
		authorizationType = stakingtypes.AuthorizationType_AUTHORIZATION_TYPE_UNDELEGATE
	default:
		authorizationType = stakingtypes.AuthorizationType_AUTHORIZATION_TYPE_REDELEGATE
	}

	return stakingtypes.NewStakeAuthorization(allowed, denied, authorizationType, delegateLimit)
}

// getValidatorAddressesFromFlags returns validator addresses with type (allowed or deny) from flags
func getValidatorAddressesFromFlags(flags *pflag.FlagSet, typ string) ([]sdk.ValAddress, error) {
	validators, err := flags.GetStringSlice(typ)
	if err != nil {
		return nil, err
	}

	validatorAddrs := make([]sdk.ValAddress, len(validators))
	for i, validator := range validators {
		addr, err := sdk.ValAddressFromBech32(validator)
		if err != nil {
			return nil, err
		}
		validatorAddrs[i] = addr
	}
	return validatorAddrs, nil
}

// getAllowedListFromFlags returns addresses who will have send authorization from flags.
func getAllowedListFromFlags(flags *pflag.FlagSet) ([]sdk.AccAddress, error) {
	allowList, err := flags.GetStringSlice(authzcli.FlagAllowList)
	if err != nil {
		return nil, err
	}

	addrs := make([]sdk.AccAddress, len(allowList))
	for i, addr := range allowList {
		accAddr, err := sdk.AccAddressFromBech32(addr)
		if err != nil {
			return nil, err
		}
		addrs[i] = accAddr
	}
	return addrs, nil
}

// -------------------------------------------------------------------------------------------------------------------

// GetCmdGrantTreasuryAuthorization returns the command used to revoke a treasury authorization
func GetCmdRevokeTreasuryAuthorization() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke [subspace-id] [grantee] [msg_type] --from=[granter]",
		Short: "revoke a treasury authorization",
		Long: strings.TrimSpace(
			fmt.Sprintf(`revoke treasury authorization from a granter to a grantee:
Example:
 $ %s tx %s revoke desmos1463vltcqk6ql6zpk0g6s595jjcrzk4804hyqw7 %s --from=desmos1463vltcqk6ql6zpk0g6s595jjcrzk4804hyqw7
			`, version.AppName, authz.ModuleName, banktypes.SendAuthorization{}.MsgTypeURL()),
		),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			subspaceID, err := types.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgRevokeTreasuryAuthorization(subspaceID, clientCtx.GetFromAddress().String(), args[1], args[2])
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
