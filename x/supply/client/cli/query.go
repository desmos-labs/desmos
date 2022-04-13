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

	"github.com/desmos-labs/desmos/v3/x/supply/types"
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
		Use:   "total-supply [denom] [[divider_exponent]]",
		Short: "Query the total supply of the given denom. It can be converted with an optional 10 divider powered with the given divider_exponent",
		Example: fmt.Sprintf(`%s query supply total-supply "stake"
%s query supply total-supply "stake" 5`, version.AppName, version.AppName),
		Long: `Get the total supply of a token with the given denom. If a divider exponent is given,the returned result 
will be divided by 10^(divider_exponent).`,
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
			res, err := queryClient.TotalSupply(context.Background(), types.NewQueryTotalSupplyRequest(denom, divider))
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
		Use:   "circulating-supply [denom] [[divider_exponent]]",
		Short: "Query the circulating supply of the given denom. It can be converted with an optional 10 divider powered with the given divider_exponent",
		Example: fmt.Sprintf(`%s query supply circulating-supply "stake"
%s query supply circulating-supply "stake" 5`, version.AppName, version.AppName),
		Long: fmt.Sprintf(`Get the circulating supply of a token with the given denom. The result can be converted with an optional 10 divider powered by the given divider_exponent. If the default value is kept, the result will be displayed in millionth (the common way with which tokens' amount are displayed on cosmo-SDK chains. Otherwise it will be converted according to the divider_exponent).

1. Without divider
%s desmos query supply circulating-supply udsm

2. With divider
%s desmos query supply circulating-supply udsm 6

6 means 10^6 = 1_000_000 divider
`, version.AppName, version.AppName,
		),
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
			res, err := queryClient.CirculatingSupply(context.Background(), types.NewQueryCirculatingSupplyRequest(denom, divider))
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
