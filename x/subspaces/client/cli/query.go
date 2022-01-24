package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/v2/x/subspaces/types"
)

// DONTCOVER

// GetQueryCmd returns the command allowing to perform queries
func GetQueryCmd() *cobra.Command {
	subspaceQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the subspaces module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	subspaceQueryCmd.AddCommand(
		GetCmdQuerySubspace(),
		GetCmdQuerySubspaces(),
		GetCmdQueryUserGroups(),
		GetCmdQueryUserGroupMembers(),
	)
	return subspaceQueryCmd
}

// GetCmdQuerySubspace returns the command to query the subspace with the given id
func GetCmdQuerySubspace() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subspace [id]",
		Short: "Query the subspace with the given id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			subspaceID, err := types.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			res, err := queryClient.Subspace(context.Background(), types.NewQuerySubspaceRequest(subspaceID))
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQuerySubspaces returns the command to query all the subspaces
func GetCmdQuerySubspaces() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subspaces",
		Short: "Query subspaces with optional pagination",
		Example: fmt.Sprintf(`
%s query subspaces subspaces --page=2 --limit=100`,
			version.AppName),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.Subspaces(context.Background(), types.NewQuerySubspacesRequest(pageReq))
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "subspaces")

	return cmd
}

// GetCmdQueryUserGroups returns the command to query the user groups of a subspace
func GetCmdQueryUserGroups() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user-groups [subspace-id]",
		Short: "Query subspaces with optional pagination",
		Example: fmt.Sprintf(`
%s query subspaces user-groups 1 --page=2 --limit=100`,
			version.AppName),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			subspaceID, err := types.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			res, err := queryClient.UserGroups(
				context.Background(),
				types.NewQueryUserGroupsRequest(subspaceID, pageReq),
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "user groups")

	return cmd
}

// GetCmdQueryUserGroupMembers returns the command to query the members of a specific user group
func GetCmdQueryUserGroupMembers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user-group-members [subspace-id] [group-name]",
		Short: "Query subspaces with optional pagination",
		Example: fmt.Sprintf(`
%s query subspaces user-group-members 1 "Admins" --page=2 --limit=100`,
			version.AppName),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			subspaceID, err := types.ParseSubspaceID(args[0])
			if err != nil {
				return err
			}

			groupName := args[1]

			res, err := queryClient.UserGroupMembers(
				context.Background(),
				types.NewQueryUserGroupMembersRequest(subspaceID, groupName, pageReq),
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "user group members")

	return cmd
}
