package cli

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/deweb-services/deweb/x/deweb/types"
)

var _ = strconv.Itoa(0)

func CmdFilterUserKeyRecords() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "filter-user-key-records [address] [chain] [deleted]",
		Short: "Query filter_user_key_records",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqAddress := args[0]
			reqChain := args[1]
			reqDeleted, err := strconv.ParseBool(args[2])
			if err != nil {
				return fmt.Errorf("error processing parameter deleted, must be true/false")
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryFilterUserKeyRecordsRequest{

				Address: reqAddress,
				Chain:   reqChain,
				Deleted: reqDeleted,
			}

			res, err := queryClient.FilterUserKeyRecords(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
