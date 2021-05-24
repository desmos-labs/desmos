package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/x/ibc/applications/profiles/types"
)

// NewQueryUserConnectionsCmd returns a new Cobra command allowing to query the connections of a user
func NewQueryUserConnectionsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "connections [user]",
		Short: "Get all the connections for the given user",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			r, err := queryClient.Connections(context.Background(), &types.QueryUserConnectionsRequest{User: args[0]})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(r)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
