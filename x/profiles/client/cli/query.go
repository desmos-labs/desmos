package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client/flags"

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
		GetCmdQueryDTagRequests(),
		GetCmdQueryUserRelationships(),
		GetCmdQueryUserBlocks(),
		GetCmdQueryParams(),
		GetCmdQueryLink(),
	)
	return profileQueryCmd
}

// GetCmdQueryProfile returns the command that allows to query the profile of a specific user
func GetCmdQueryProfile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "profile [address_or_dtag]",
		Short: "Retrieve the profile having the specified user address or profile dtag, if any.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
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

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryDTagRequests returns the command allowing to query all the DTag transfer requests made towards a user
func GetCmdQueryDTagRequests() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dtag-requests [address]",
		Short: "Retrieve the requests made to the given address to transfer its profile's dTag",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
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

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryParams returns the command allowing to query the profiles module params
func GetCmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "parameters",
		Short: "Retrieve all the profile module parameters",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryUserRelationships returns the command allowing to query all the relationships of a specific user
func GetCmdQueryUserRelationships() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "relationships [address]",
		Short: "Retrieve all the user's relationships",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.UserRelationships(
				context.Background(),
				&types.QueryUserRelationshipsRequest{User: args[0]},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryUserBlocks returns the command allowing to query all the blocks of a single user
func GetCmdQueryUserBlocks() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "blocklist [address]",
		Short: "Retrieve the list of all the blocked users of the given address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.UserBlocks(
				context.Background(),
				&types.QueryUserBlocksRequest{User: args[0]})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryLink returns the command allowing to query the link with chain id of a single user
func GetCmdQueryLink() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "link [chain-id] [address]",
		Short: "Retrieve the link of the given address and chain id",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Link(
				context.Background(),
				&types.QueryLinkRequest{User: args[1], ChainId: args[0]})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
