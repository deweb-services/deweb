package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/deweb-services/deweb/x/deweb/types"
)

var _ = strconv.Itoa(0)

func CmdGetUserKeyRecords() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-user-key-records [address]",
		Short: "Query get_user_key_records",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqAddress := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetUserKeyRecordsRequest{

				Address: reqAddress,
			}

			res, err := queryClient.GetUserKeyRecords(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
