package cli

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/v7/x/supply/types"
)

// DONTCOVER

func GetQueryCmd() *cobra.Command {
	supplyQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the supply module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	supplyQueryCmd.AddCommand(
		GetCmdQueryTotalSupply(),
		GetCmdQueryCirculatingSupply(),
	)
	return supplyQueryCmd
}

// GetCmdQueryTotalSupply returns the command to query the total supply of the given denom
// converted with a divider powered with the given (optional) divider_exponent
func GetCmdQueryTotalSupply() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "total [denom] [[divider_exponent]]",
		Short: "Query the total supply of the given denom. It can be converted with an optional 10 divider powered with the given divider_exponent",
		Long: `Get the total supply of a token with the given denom. 
If a divider exponent is given,the returned result will be divided by 10^(divider_exponent).`,
		Example: fmt.Sprintf(`
%s query supply total "stake"
%s query supply total "stake" 5`,
			version.AppName, version.AppName),
		Args: cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			denom := strings.TrimSpace(args[0])
			if denom == "" {
				return fmt.Errorf("invalid denom given")
			}

			divider := uint64(0)
			if len(args) > 1 {
				divider, err = strconv.ParseUint(args[1], 10, 0)
				if err != nil {
					return err
				}
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Total(context.Background(), types.NewQueryTotalRequest(denom, divider))
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryCirculatingSupply returns the command to query the total supply of the given denom
// converted with the given (optional) divider
func GetCmdQueryCirculatingSupply() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "circulating [denom] [[divider_exponent]]",
		Short: "Query the circulating supply of the given denom. It can be converted with an optional 10 divider powered with the given divider_exponent",
		Long: `Get the circulating supply of a token with the given denom. 
If a divider exponent is given,the returned result will be divided by 10^(divider_exponent).`,
		Example: fmt.Sprintf(`
%s query supply circulating "stake"
%s query supply circulating "stake" 5`,
			version.AppName, version.AppName),
		Args: cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			denom := strings.TrimSpace(args[0])
			if denom == "" {
				return fmt.Errorf("invalid denom given")
			}

			divider := uint64(0)
			if len(args) > 1 {
				divider, err = strconv.ParseUint(args[1], 10, 0)
				if err != nil {
					return err
				}
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Circulating(context.Background(), types.NewQueryCirculatingRequest(denom, divider))
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
