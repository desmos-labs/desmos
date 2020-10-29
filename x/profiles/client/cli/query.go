package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// GetQueryCmd returns the command allowing to perform queries
func GetQueryCmd() *cobra.Command {
	profileQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the profiles module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	profileQueryCmd.AddCommand(
		GetCmdQueryProfile(),
		GetCmdQueryProfiles(),
		GetCmdQueryProfileParams(),
		GetCmdQueryDTagRequests(),
	)
	return profileQueryCmd
}

// GetCmdQueryProfiles returns the command allowing to query all the profiles
func GetCmdQueryProfiles() *cobra.Command {
	return &cobra.Command{
		Use:   "all",
		Short: "Retrieve all the registered profiles.",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Profiles(context.Background(), &types.QueryProfilesRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}
}

// GetCmdQueryProfile returns the command that allows to query the profile of a specific user
func GetCmdQueryProfile() *cobra.Command {
	return &cobra.Command{
		Use:   "profile [address_or_dtag]",
		Short: "Retrieve the profile having the specified user address or profile dtag, if any.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Profile(
				context.Background(),
				&types.QueryProfileRequest{User: args[0]},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}
}

// GetCmdQueryDTagRequests returns the command allowing to query all the DTag transfer requests made towards a user
func GetCmdQueryDTagRequests() *cobra.Command {
	return &cobra.Command{
		Use:   "dtag-requests [address]",
		Short: "Retrieve the requests made to the given address to transfer its profile's dTag",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.DTagTransfers(
				context.Background(),
				&types.QueryDTagTransfersRequest{User: args[0]},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}
}

// GetCmdQueryProfileParams returns the command allowing to query the profiles module params
func GetCmdQueryProfileParams() *cobra.Command {
	return &cobra.Command{
		Use:   "parameters",
		Short: "Retrieve all the profile module parameters",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}
}
