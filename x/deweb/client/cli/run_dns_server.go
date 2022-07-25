package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/deweb-services/deweb/x/dns_server"
	"github.com/spf13/cobra"
	"strconv"
)

func CmdRunDNSServer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run-dns-server [port]",
		Short: "Run DNS server for DNS NFT",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqPort, err := strconv.ParseInt(args[0], 10, 32)
			if err != nil {
				return fmt.Errorf("cannot parse parameter limit: %w", err)
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			serverResolver := dns_server.NewDNSResolverService(clientCtx)
			serverResolver.RunServer(int(reqPort))

			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
