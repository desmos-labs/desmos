package cli

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// GetCmdGrantUserAllowance returns the command used to grant a user fee allowance
func GetCmdGrantUserAllowance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "grant-user-allowance [subspace-id] [grantee]",
		Short: "Grant a fee allowance to an address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			allowance, err := createAllowanceFromFlags(cmd.Flags())
			if err != nil {
				return err
			}
			subspaceID, err := types.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}
			msg := types.NewMsgGrantUserAllowance(subspaceID, clientCtx.FromAddress.String(), args[1], allowance)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().StringSlice(FlagAllowedMsgs, []string{}, "Set of allowed messages for fee allowance")
	cmd.Flags().String(FlagExpiration, "", "The RFC 3339 timestamp after which the grant expires for the user")
	cmd.Flags().String(FlagSpendLimit, "", "Spend limit specifies the max limit can be used, if not mentioned there is no limit")
	cmd.Flags().Int64(FlagPeriod, 0, "Period specifies the time duration in which period_spend_limit coins can be spent before that allowance is reset")
	cmd.Flags().String(FlagPeriodLimit, "", "Period limit specifies the maximum number of coins that can be spent in the period")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdRevokeUserAllowance returns the command used to revoke a user fee allowance
func GetCmdRevokeUserAllowance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke-user-allowance [subspace-id] [user]",
		Short: "Revoke a fee allowance from an address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			subspaceID, err := types.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}
			msg := types.NewMsgRevokeUserAllowance(subspaceID, clientCtx.FromAddress.String(), args[1])
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdGrantGroupAllowance returns the command used to grant a group fee allowance
func GetCmdGrantGroupAllowance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "grant-group-allowance [subspace-id] [group-id]",
		Short: "Grant a fee allowance to a group",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			allowance, err := createAllowanceFromFlags(cmd.Flags())
			if err != nil {
				return err
			}

			subspaceID, err := types.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}
			groupID, err := types.ParseGroupID(args[1])
			if err != nil {
				return err
			}
			msg := types.NewMsgGrantGroupAllowance(subspaceID, clientCtx.FromAddress.String(), groupID, allowance)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().StringSlice(FlagAllowedMsgs, []string{}, "Set of allowed messages for fee allowance")
	cmd.Flags().String(FlagExpiration, "", "The RFC 3339 timestamp after which the grant expires for the user")
	cmd.Flags().String(FlagSpendLimit, "", "Spend limit specifies the max limit can be used, if not mentioned there is no limit")
	cmd.Flags().Int64(FlagPeriod, 0, "Period specifies the time duration in which period_spend_limit coins can be spent before that allowance is reset")
	cmd.Flags().String(FlagPeriodLimit, "", "Period limit specifies the maximum number of coins that can be spent in the period")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdRevokeUserFeeAllowance returns the command used to revoke a group fee allowance
func GetCmdRevokeGroupAllowance() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke-group-allowance [subspace-id] [group-id]",
		Short: "Revoke a fee allowance from a group",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			subspaceID, err := types.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}
			groupID, err := types.ParseGroupID(args[1])
			if err != nil {
				return err
			}
			msg := types.NewMsgRevokeGroupAllowance(subspaceID, clientCtx.FromAddress.String(), groupID)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// createAllowanceFromFlags create a allowance from flags
func createAllowanceFromFlags(flags *pflag.FlagSet) (feegrant.FeeAllowanceI, error) {
	spendLimit, err := flags.GetString(FlagSpendLimit)
	if err != nil {
		return nil, err
	}
	// if `FlagSpendLimit` isn't set, limit will be nil
	limit, err := sdk.ParseCoinsNormalized(spendLimit)
	if err != nil {
		return nil, err
	}
	expired, err := flags.GetString(FlagExpiration)
	if err != nil {
		return nil, err
	}
	periodClock, err := flags.GetInt64(FlagPeriod)
	if err != nil {
		return nil, err
	}
	periodLimit, err := flags.GetString(FlagPeriodLimit)
	if err != nil {
		return nil, err
	}
	allowedMsgs, err := flags.GetStringSlice(FlagAllowedMsgs)
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
