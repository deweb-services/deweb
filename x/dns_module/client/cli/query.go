package cli

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/deweb-services/deweb/x/dns_module/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                types.ModuleName,
		Short:              "Querying commands for the domains module",
		DisableFlagParsing: true,
	}

	queryCmd.AddCommand(
		GetCmdQueryParams(),
		GetCmdQueryDomain(),
		GetCmdQueryOwnersDomain(),
	)

	return queryCmd
}

// GetCmdQueryDomain queries a single Domains from a collection
func GetCmdQueryDomain() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "domain [domain]",
		Long:    "Query a single domain NFT ",
		Example: fmt.Sprintf("$ %s query domain <domain>", version.AppName),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			if err := types.ValidateTokenID(args[0]); err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			resp, err := queryClient.Domain(context.Background(), &types.QueryDomainRequest{
				DomainName: args[0],
			})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(resp.Domain)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryOwnersDomain queries a list of domains owned by iser
func GetCmdQueryOwnersDomain() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "address [address] [offset] [limit]",
		Long:    "Query a list of domains owned by user ",
		Example: fmt.Sprintf("$ %s query address <domain>", version.AppName),
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			var offset int
			if len(args) > 1 {
				offset, err = strconv.Atoi(args[1])
				if err != nil {
					return fmt.Errorf("invalid offset value: %s: %w", args[1], err)
				}
			}

			count := 10
			if len(args) > 2 {
				count, err = strconv.Atoi(args[2])
				if err != nil {
					return fmt.Errorf("invalid count value: %s: %w", args[1], err)
				}
			}

			queryClient := types.NewQueryClient(clientCtx)
			resp, err := queryClient.OwnedDomains(context.Background(), &types.QueryOwnedDomainsRequest{
				Address: args[0],
				Offset:  int64(offset),
				Count:   int64(count),
			})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(resp)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "params",
		Long:    "Query module parameters ",
		Example: fmt.Sprintf("$ %s query nft params", version.AppName),
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			resMsg := &types.Params{}
			return clientCtx.PrintProto(resMsg)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
