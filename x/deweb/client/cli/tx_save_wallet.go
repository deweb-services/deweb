package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/deweb-services/deweb/x/deweb/types"
)

var _ = strconv.Itoa(0)

func CmdSaveWallet() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "save-wallet [address] [encrypted_key] [chain]",
		Short: "Broadcast message save_wallet",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAddress := args[0]
			argEncKey := args[1]
			argChain := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSaveUser(
				clientCtx.GetFromAddress().String(),
				argAddress,
				argEncKey,
				argChain,
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
