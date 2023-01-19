package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantcli "github.com/cosmos/cosmos-sdk/x/feegrant/client/cli"

	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

// DONTCOVER

const (
	FlagUserGrantee  = "user"
	FlagGroupGrantee = "group"
)

// GetAllowanceTxCmd returns a new command to perform subspaces treasury transactions
func GetAllowanceTxCmd() *cobra.Command {
	treasuryTxCmd := &cobra.Command{
		Use:                        "allowances",
		Short:                      "Tx commands for subspace treasury",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	treasuryTxCmd.AddCommand(
		GetCmdGrantAllowance(),
		GetCmdRevokeAllowance(),
	)

	return treasuryTxCmd
}

// GetCmdGrantAllowance returns the command used to grant a fee allowance
func GetCmdGrantAllowance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "grant [subspace-id]",
		Short: "Grant a fee allowance to an address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			subspaceID, err := types.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}
			grantee, err := getGranteeFromFlags(cmd.Flags())
			if err != nil {
				return err
			}
			allowance, err := getAllowanceFromFlags(cmd.Flags())
			if err != nil {
				return err
			}
			msg := types.NewMsgGrantAllowance(subspaceID, clientCtx.FromAddress.String(), grantee, allowance)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagUserGrantee, "", "Address of the user being the allowance grantee")
	cmd.Flags().Uint32(FlagGroupGrantee, 0, "Id of group being the allowance grantee")
	cmd.Flags().StringSlice(feegrantcli.FlagAllowedMsgs, []string{}, "Set of allowed messages for fee allowance")
	cmd.Flags().String(feegrantcli.FlagExpiration, "", "The RFC 3339 timestamp after which the grant expires for the user")
	cmd.Flags().String(feegrantcli.FlagSpendLimit, "", "Spend limit specifies the max limit can be used, if not mentioned there is no limit")
	cmd.Flags().Int64(feegrantcli.FlagPeriod, 0, "Period specifies the time duration in which period_spend_limit coins can be spent before that allowance is reset")
	cmd.Flags().String(feegrantcli.FlagPeriodLimit, "", "Period limit specifies the maximum number of coins that can be spent in the period")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdRevokeAllowance returns the command used to revoke a fee allowance
func GetCmdRevokeAllowance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke [subspace-id]",
		Short: "Revoke a fee allowance from an address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			subspaceID, err := types.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}
			grantee, err := getGranteeFromFlags(cmd.Flags())
			if err != nil {
				return err
			}
			msg := types.NewMsgRevokeAllowance(subspaceID, clientCtx.FromAddress.String(), grantee)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().String(FlagUserGrantee, "", "Address of the user being the allowance grantee")
	cmd.Flags().Uint32(FlagGroupGrantee, 0, "Id of group being the allowance grantee")
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// getGranteeFromFlags creates a grantee from flags
func getGranteeFromFlags(flags *pflag.FlagSet) (types.Grantee, error) {
	userGrantee, err := flags.GetString(FlagUserGrantee)
	if err != nil {
		return nil, err
	}

	groupGrantee, err := flags.GetUint32(FlagGroupGrantee)
	if err != nil {
		return nil, err
	}

	switch {
	case userGrantee != "" && groupGrantee != 0:
		return nil, fmt.Errorf("only one of --%s or --%s must be used", FlagUserGrantee, FlagGroupGrantee)

	case userGrantee != "":
		return types.NewUserGrantee(userGrantee), nil

	case groupGrantee != 0:
		return types.NewGroupGrantee(groupGrantee), nil
	}

	return nil, fmt.Errorf("one of --%s or --%s must be used", FlagUserGrantee, FlagGroupGrantee)
}

// getAllowanceFromFlags create a allowance from flags
func getAllowanceFromFlags(flags *pflag.FlagSet) (feegrant.FeeAllowanceI, error) {
	spendLimit, err := flags.GetString(feegrantcli.FlagSpendLimit)
	if err != nil {
		return nil, err
	}
	// if `FlagSpendLimit` isn't set, limit will be nil
	limit, err := sdk.ParseCoinsNormalized(spendLimit)
	if err != nil {
		return nil, err
	}
	expired, err := flags.GetString(feegrantcli.FlagExpiration)
	if err != nil {
		return nil, err
	}
	periodClock, err := flags.GetInt64(feegrantcli.FlagPeriod)
	if err != nil {
		return nil, err
	}
	periodLimit, err := flags.GetString(feegrantcli.FlagPeriodLimit)
	if err != nil {
		return nil, err
	}
	allowedMsgs, err := flags.GetStringSlice(feegrantcli.FlagAllowedMsgs)
	if err != nil {
		return nil, err
	}

	var allowance feegrant.FeeAllowanceI
	basic := feegrant.BasicAllowance{
		SpendLimit: limit,
	}
	var expiresAtTime time.Time
	if expired != "" {
		expiresAtTime, err = time.Parse(time.RFC3339, expired)
		if err != nil {
			return nil, err
		}
		basic.Expiration = &expiresAtTime
	}
	allowance = &basic
	// Check any of period or periodLimit flags set, If set consider it as periodic fee allowance
	if periodClock > 0 || periodLimit != "" {
		periodLimit, err := sdk.ParseCoinsNormalized(periodLimit)
		if err != nil {
			return nil, err
		}
		if periodClock <= 0 {
			return nil, fmt.Errorf("period clock was not set")
		}
		if periodLimit == nil {
			return nil, fmt.Errorf("period limit was not set")
		}
		periodReset := getPeriodReset(periodClock)
		if basic.Expiration != nil && periodReset.Sub(expiresAtTime) > 0 {
			return nil, fmt.Errorf("period (%d) cannot reset after expiration (%v)", periodClock, expired)
		}
		periodAllowance := &feegrant.PeriodicAllowance{
			Basic:            basic,
			Period:           getPeriod(periodClock),
			PeriodReset:      periodReset,
			PeriodSpendLimit: periodLimit,
			PeriodCanSpend:   periodLimit,
		}
		allowance = periodAllowance
	}
	if len(allowedMsgs) > 0 {
		filteredAllowance, err := feegrant.NewAllowedMsgAllowance(allowance, allowedMsgs)
		if err != nil {
			return nil, err
		}
		allowance = filteredAllowance
	}
	return allowance, nil
}

// getPeriodReset generates a next period reset time from a duration
func getPeriodReset(duration int64) time.Time {
	return time.Now().Add(getPeriod(duration))
}

// getPeriod turns duration type from int64 into time.Duration
func getPeriod(duration int64) time.Duration {
	return time.Duration(duration) * time.Second
}
