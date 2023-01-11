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
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
	"github.com/spf13/cobra"
)

const (
	delegate   = "delegate"
	redelegate = "redelegate"
	unbond     = "unbond"
)

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

func GetCmdGrantTreasuryAuthorization() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "grant [subspace-id] [grantee] [authorization_type=\"send\"|\"generic\"|\"delegate\"|\"unbond\"|\"redelegate\"] --from [granter]",
		Short: "Grant a treasury authorization to a user",
		Long: strings.TrimSpace(
			fmt.Sprintf(`grant treasury authorization to an address to execute a transaction on your behalf:
Examples:
 $ %[1]s tx %[2]s grant desmos1463vltcqk6ql6zpk0g6s595jjcrzk4804hyqw7 send --spend-limit=1000stake --from=cosmos1skl..
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
				limit, err := cmd.Flags().GetString(authzcli.FlagSpendLimit)
				if err != nil {
					return err
				}

				spendLimit, err := sdk.ParseCoinsNormalized(limit)
				if err != nil {
					return err
				}

				if !spendLimit.IsAllPositive() {
					return fmt.Errorf("spend-limit should be greater than zero")
				}

				authorization = banktypes.NewSendAuthorization(spendLimit)
			case "generic":
				msgType, err := cmd.Flags().GetString(authzcli.FlagMsgType)
				if err != nil {
					return err
				}

				authorization = authz.NewGenericAuthorization(msgType)
			case delegate, unbond, redelegate:
				limit, err := cmd.Flags().GetString(authzcli.FlagSpendLimit)
				if err != nil {
					return err
				}

				allowValidators, err := cmd.Flags().GetStringSlice(authzcli.FlagAllowedValidators)
				if err != nil {
					return err
				}

				denyValidators, err := cmd.Flags().GetStringSlice(authzcli.FlagDenyValidators)
				if err != nil {
					return err
				}

				var delegateLimit *sdk.Coin
				if limit != "" {
					spendLimit, err := sdk.ParseCoinsNormalized(limit)
					if err != nil {
						return err
					}

					if !spendLimit.IsAllPositive() {
						return fmt.Errorf("spend-limit should be greater than zero")
					}
					delegateLimit = &spendLimit[0]
				}

				allowed, err := bech32toValidatorAddresses(allowValidators)
				if err != nil {
					return err
				}

				denied, err := bech32toValidatorAddresses(denyValidators)
				if err != nil {
					return err
				}

				switch args[2] {
				case delegate:
					authorization, err = stakingtypes.NewStakeAuthorization(allowed, denied, stakingtypes.AuthorizationType_AUTHORIZATION_TYPE_DELEGATE, delegateLimit)
				case unbond:
					authorization, err = stakingtypes.NewStakeAuthorization(allowed, denied, stakingtypes.AuthorizationType_AUTHORIZATION_TYPE_UNDELEGATE, delegateLimit)
				default:
					authorization, err = stakingtypes.NewStakeAuthorization(allowed, denied, stakingtypes.AuthorizationType_AUTHORIZATION_TYPE_REDELEGATE, delegateLimit)
				}
				if err != nil {
					return err
				}

			default:
				return fmt.Errorf("invalid authorization type, %s", args[2])
			}

			msg := types.NewMsgGrantTreasuryAuthorization(subspaceID, clientCtx.GetFromAddress().String(), args[1], authorization, time.Unix(exp, 0))
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

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func bech32toValidatorAddresses(validators []string) ([]sdk.ValAddress, error) {
	vals := make([]sdk.ValAddress, len(validators))
	for i, validator := range validators {
		addr, err := sdk.ValAddressFromBech32(validator)
		if err != nil {
			return nil, err
		}
		vals[i] = addr
	}
	return vals, nil
}

// -------------------------------------------------------------------------------------------------------------------

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
