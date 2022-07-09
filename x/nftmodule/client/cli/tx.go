package cli

import (
	"fmt"
	"io/ioutil"
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
		GetCmdIssueDenom(),
		GetCmdMintNFT(),
		GetCmdEditNFT(),
		GetCmdTransferNFT(),
		GetCmdBurnNFT(),
	)

	return txCmd
}

// GetCmdIssueDenom is the CLI command for an IssueDenom transaction
func GetCmdIssueDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "issue [denom-id]",
		Long: "Issue a new denom.",
		Example: fmt.Sprintf(
			"$ %s tx nft issue <denom-id> "+
				"--from=<key-name> "+
				"--name=<denom-name> "+
				"--symbol=<denom-symbol> "+
				"--mint-restricted=<mint-restricted> "+
				"--update-restricted=<update-restricted> "+
				"--schema=<schema-content or path to schema.json> "+
				"--description=<description> "+
				"--uri=<uri> "+
				"--uri-hash=<uri-hash> "+
				"--data=<data> "+
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

			denomName, err := cmd.Flags().GetString(FlagDenomName)
			if err != nil {
				return err
			}
			schema, err := cmd.Flags().GetString(FlagSchema)
			if err != nil {
				return err
			}
			symbol, err := cmd.Flags().GetString(FlagSymbol)
			if err != nil {
				return err
			}
			mintRestricted, err := cmd.Flags().GetBool(FlagMintRestricted)
			if err != nil {
				return err
			}
			updateRestricted, err := cmd.Flags().GetBool(FlagUpdateRestricted)
			if err != nil {
				return err
			}
			uri, err := cmd.Flags().GetString(FlagURI)
			if err != nil {
				return err
			}
			uriHash, err := cmd.Flags().GetString(FlagURIHash)
			if err != nil {
				return err
			}
			description, err := cmd.Flags().GetString(FlagDescription)
			if err != nil {
				return err
			}
			data, err := cmd.Flags().GetString(FlagData)
			if err != nil {
				return err
			}
			optionsContent, err := ioutil.ReadFile(schema)
			if err == nil {
				schema = string(optionsContent)
			}

			msg := types.NewMsgIssueDenom(
				args[0],
				denomName,
				schema,
				clientCtx.GetFromAddress().String(),
				symbol,
				mintRestricted,
				updateRestricted,
				description,
				uri,
				uriHash,
				data,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().AddFlagSet(FsIssueDenom)
	_ = cmd.MarkFlagRequired(FlagMintRestricted)
	_ = cmd.MarkFlagRequired(FlagUpdateRestricted)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdMintNFT is the CLI command for a MintNFT transaction
func GetCmdMintNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "mint [domain]",
		Long: "Register an NFT domain and set the owner to the recipient.",
		Example: fmt.Sprintf(
			"$ %s tx nft mint <domain> "+
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

			msg := types.NewMsgMintNFT(
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
			msg := types.NewMsgEditNFT(args[0], tokenData, clientCtx.GetFromAddress().String())
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

			msg := types.NewMsgTransferNFT(
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

// GetCmdBurnNFT is the CLI command for sending a BurnNFT transaction
func GetCmdBurnNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "burn [domain]",
		Long: "Burn a domain NFT.",
		Example: fmt.Sprintf(
			"$ %s tx nft burn <domain> "+
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

			msg := types.NewMsgBurnNFT(clientCtx.GetFromAddress().String(), args[0])
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
