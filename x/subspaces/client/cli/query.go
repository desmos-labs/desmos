package cli

import (
	"context"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/x/subspaces/types"
)

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
		GetCmdQuerySubspaceAdmins(),
		GetCmdQuerySubspaceRegisteredUsers(),
		GetCmdQuerySubspaceBannedUsers(),
		GetCmdQueryTokenomics(),
		GetCmdQueryAllTokenomics(),
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

			res, err := queryClient.Subspace(context.Background(), types.NewQuerySubspaceRequest(args[0]))
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
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query for paginated subspaces:

Example:
$ %s query subspaces subspaces
$ %s query subspaces subspaces --page=2 --limit=100
`, version.AppName, version.AppName)),
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

// GetCmdQuerySubspaceAdmins returns the command to query the admins of a given subspace
func GetCmdQuerySubspaceAdmins() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "admins [subspace-id]",
		Short: "Query subspace admins with optional pagination",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query for paginated admins:

Example:
$ %s query subspaces admins [subspace-id]
$ %s query subspaces admins [subspace-id] --page=2 --limit=100
`, version.AppName, version.AppName)),
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

			res, err := queryClient.Admins(
				context.Background(),
				types.NewQueryAdminsRequest(args[0], pageReq),
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "admins")

	return cmd
}

// GetCmdQuerySubspaceRegisteredUsers returns the command to query the registered users of a given subspace
func GetCmdQuerySubspaceRegisteredUsers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "registered-users [subspace-id]",
		Short: "Query subspace registered users with optional pagination",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query for paginated users:

Example:
$ %s query subspaces registered-users [subspace-id]
$ %s query subspaces registered-users [subspace-id] --page=2 --limit=100
`, version.AppName, version.AppName)),
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

			res, err := queryClient.RegisteredUsers(
				context.Background(),
				types.NewQueryRegisteredUsersRequest(args[0], pageReq),
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "registered users")

	return cmd
}

// GetCmdQuerySubspaceBannedUsers returns the command to query the banned users of a given subspace
func GetCmdQuerySubspaceBannedUsers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "banned-users [subspace-id]",
		Short: "Query subspace banned users with optional pagination",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query for paginated users:

Example:
$ %s query subspaces banned-users [subspace-id]
$ %s query subspaces banned-users [subspace-id] --page=2 --limit=100
`, version.AppName, version.AppName)),
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

			res, err := queryClient.BannedUsers(
				context.Background(),
				types.NewQueryBannedUsersRequest(args[0], pageReq),
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "banned users")

	return cmd
}

// GetCmdQueryTokenomics returns the command to query the tokenomics of a subspace
func GetCmdQueryTokenomics() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tokenomics [subspace-id]",
		Short: "Query subspace tokenomics",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Tokenomics(context.Background(), types.NewQueryTokenomicsRequest(args[0]))
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryAllTokenomics returns the command to query all the tokenomics
func GetCmdQueryAllTokenomics() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-tokenomics",
		Short: "Query all the tokenomics inside the current context",
		Args:  cobra.NoArgs,
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

			res, err := queryClient.AllTokenomics(
				context.Background(),
				types.NewQueryAllTokenomicsRequest(pageReq),
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "tokenomics pairs")

	return cmd
}
