package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/v3/x/subspaces/types"
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

		GetGroupsQueryCmd(),
		GetCmdQueryUserPermissions(),
	)
	return subspaceQueryCmd
}

// GetCmdQuerySubspace returns the command to query the subspace with the given id
func GetCmdQuerySubspace() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "subspace [id]",
		Short:   "Query the subspace with the given id",
		Example: fmt.Sprintf(`%s query subspaces subspace 1`, version.AppName),
		Args:    cobra.ExactArgs(1),
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
		Use:     "subspaces",
		Short:   "Query subspaces with optional pagination",
		Example: fmt.Sprintf(`%s query subspaces subspaces --page=2 --limit=100`, version.AppName),
		Args:    cobra.NoArgs,
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

// -------------------------------------------------------------------------------------------------------------------

// GetGroupsQueryCmd returns a new command to perform queries for user groups
func GetGroupsQueryCmd() *cobra.Command {
	groupsQueryCmd := &cobra.Command{
		Use:                        "groups",
		Short:                      "Querying commands for subspace groups",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	groupsQueryCmd.AddCommand(
		GetCmdQueryUserGroups(),
		GetCmdQueryUserGroup(),
		GetCmdQueryUserGroupMembers(),
	)

	return groupsQueryCmd
}

// GetCmdQueryUserGroups returns the command to query the user groups of a subspace
func GetCmdQueryUserGroups() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list [subspace-id]",
		Short:   "Query groups in the given subspace with optional pagination",
		Example: fmt.Sprintf(`%s query subspaces groups list 1 --page=2 --limit=100`, version.AppName),
		Args:    cobra.ExactArgs(1),
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

// GetCmdQueryUserGroup returns the command to query a specific user group of a subspace
func GetCmdQueryUserGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "group [subspace-id] [group-id]",
		Short: "Query the group with the given id in the given subspace",
		Example: fmt.Sprintf(`
%s query subspaces groups group 1 2`,
			version.AppName),
		Args: cobra.ExactArgs(2),
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

			groupID, err := types.ParseGroupID(args[1])
			if err != nil {
				return err
			}

			res, err := queryClient.UserGroup(
				context.Background(),
				types.NewQueryUserGroupRequest(subspaceID, groupID),
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

// GetCmdQueryUserGroupMembers returns the command to query the members of a specific user group
func GetCmdQueryUserGroupMembers() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "members [subspace-id] [group-id]",
		Short:   "Query members in the given group with optional pagination",
		Example: fmt.Sprintf(`%s query subspaces groups memebers 1 1 --page=2 --limit=100`, version.AppName),
		Args:    cobra.ExactArgs(2),
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

			groupID, err := types.ParseGroupID(args[1])
			if err != nil {
				return err
			}

			res, err := queryClient.UserGroupMembers(
				context.Background(),
				types.NewQueryUserGroupMembersRequest(subspaceID, groupID, pageReq),
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

// GetCmdQueryUserPermissions returns the command to query the permissions of a specific user
func GetCmdQueryUserPermissions() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "permissions [subspace-id] [user]",
		Short:   "Query permissions of the given user",
		Example: fmt.Sprintf(`%s query subspaces permissions 1 desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud`, version.AppName),
		Args:    cobra.ExactArgs(2),
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

			res, err := queryClient.UserPermissions(
				context.Background(),
				types.NewQueryUserPermissionsRequest(subspaceID, args[1]),
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
