package cli

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/desmos-labs/desmos/v3/x/supply/types"
	"github.com/spf13/cobra"
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
// converted with the given (optional) divider
func GetCmdQueryTotalSupply() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "total-supply [denom] [[divider]]",
		Short: "Query the total supply of the given denom. It can be converted with an optional divider",
		Example: fmt.Sprintf(
			"%s query supply total-supply [denom] [[divider]]", version.AppName,
		),
		Long: fmt.Sprintf(`
Get the total supply of a token with the given denom. The result can be converted with an optional divider which is set to
1 by default. If the default value is kept, the result will be displayed in millionth (the common way with which tokens' amount
are displayed on cosmo-SDK chains. Otherwise it will be converted according to the divider set.

1. Without divider
%s desmos query supply total-supply udsm

2. With divider
%s desmos query supply total-supply udsm 1000000
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

			divider := uint64(1)
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
		Use:   "circulating-supply [denom] [[divider]]",
		Short: "Query the circulating supply of the given denom. It can be converted with an optional divider",
		Example: fmt.Sprintf(
			"%s query supply circulating-supply [denom] [[divider]]", version.AppName,
		),
		Long: fmt.Sprintf(`
Get the circulating supply of a token with the given denom. The result can be converted with an optional divider which is set to
1 by default. If the default value is kept, the result will be displayed in millionth (the common way with which tokens' amount
are displayed on cosmo-SDK chains. Otherwise it will be converted according to the divider set.

1. Without divider
%s desmos query supply circulating-supply udsm

2. With divider
%s desmos query supply circulating-supply udsm 1000000
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

			divider := uint64(1)
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
