package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/deweb-services/deweb/x/nftmodule/types"
)

// NewTxCmd returns the transaction commands for this module
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "NFT transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		GetCmdMintNFT(),
		GetCmdEditNFT(),
		GetCmdTransferNFT(),
		GetCmdRemoveDomain(),
	)

	return txCmd
}

// GetCmdMintNFT is the CLI command for a RegisterDomain transaction
func GetCmdMintNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "register [domain]",
		Long: "Register an NFT domain and set the owner to the recipient.",
		Example: fmt.Sprintf(
			"$ %s tx nft register <domain> "+
				"--data=<data> "+
				"--recipient=<recipient> "+
				"--from=<key-name> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>",
			version.AppName,
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var sender = clientCtx.GetFromAddress().String()

			recipient, err := cmd.Flags().GetString(FlagRecipient)
			if err != nil {
				return err
			}

			recipientStr := strings.TrimSpace(recipient)
			if len(recipientStr) > 0 {
				if _, err = sdk.AccAddressFromBech32(recipientStr); err != nil {
					return err
				}
			} else {
				recipient = sender
			}
			tokenData, err := cmd.Flags().GetString(FlagData)
			if err != nil {
				return err
			}

			msg := types.NewMsgRegisterDomain(
				args[0],
				tokenData,
				sender,
				recipient,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().AddFlagSet(FsMintNFT)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdEditNFT is the CLI command for sending an MsgEditNFT transaction
func GetCmdEditNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "edit [domain]",
		Long: "Edit the domain data of an NFT.",
		Example: fmt.Sprintf(
			"$ %s tx nft edit <domain> "+
				"--data=<data> "+
				"--from=<key-name> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>",
			version.AppName,
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			tokenData, err := cmd.Flags().GetString(FlagData)
			if err != nil {
				return err
			}
			msg := types.NewMsgEditDomain(args[0], tokenData, clientCtx.GetFromAddress().String())
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().AddFlagSet(FsEditNFT)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdTransferNFT is the CLI command for sending a TransferNFT transaction
func GetCmdTransferNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "transfer [domain]",
		Long: "Transfer an domain NFT to a recipient.",
		Example: fmt.Sprintf(
			"$ %s tx nft transfer <domain> " +
				"--recipient=<recipient> " +
				"--cancel=<cancel> " +
				"--price=<price_udws> " +
				"--uri-hash=<uri-hash> " +
				"--from=<key-name> " +
				"--chain-id=<chain-id> " +
				version.AppName,
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			tokenPriceStr, err := cmd.Flags().GetString(FlagTokenPrice)
			if err != nil {
				return err
			}
			priceInt, err := strconv.ParseUint(tokenPriceStr, 10, 64)
			if err != nil {
				return fmt.Errorf("cannot parse price value: %w", err)
			}
			cancelTransferStr, err := cmd.Flags().GetString(FlagCancel)
			if err != nil {
				return err
			}
			var cancelTransfer bool
			if len(cancelTransferStr) > 0 {
				cancelTransfer, err = strconv.ParseBool(cancelTransferStr)
				if err != nil {
					return fmt.Errorf("cannot parse cacnel flag: %w", err)
				}
			}

			tokenRecipient, err := cmd.Flags().GetString(FlagRecipient)
			if err != nil {
				return err
			}
			if len(tokenRecipient) > 0 {
				if _, err := sdk.AccAddressFromBech32(tokenRecipient); err != nil {
					return err
				}
			}

			msg := types.NewMsgTransferDomain(
				args[0],
				priceInt,
				cancelTransfer,
				clientCtx.GetFromAddress().String(),
				tokenRecipient,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			return err
		},
	}
	cmd.Flags().AddFlagSet(FsTransferNFT)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdRemoveDomain is the CLI command for sending a BurnNFT transaction
func GetCmdRemoveDomain() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "remove [domain]",
		Long: "Remove a domain NFT.",
		Example: fmt.Sprintf(
			"$ %s tx nft remove <domain> "+
				"--from=<key-name> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>",
			version.AppName,
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRemoveDomain(clientCtx.GetFromAddress().String(), args[0])
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
