package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/deweb-services/deweb/x/deweb/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdConnectChain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "connect-chain [chain] [address]",
		Short: "Broadcast message connect_chain",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argChain := args[0]
			argAddress := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgConnectChain(
				clientCtx.GetFromAddress().String(),
				argChain,
				argAddress,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
