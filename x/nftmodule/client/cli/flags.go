package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagTokenName   = "name"
	FlagTokenPrice  = "price"
	FlagURI         = "uri"
	FlagURIHash     = "uri-hash"
	FlagDescription = "description"
	FlagRecipient   = "recipient"
	FlagOwner       = "owner"
	FlagData        = "data"
	FlagCancel      = "cancel"

	FlagDenomName        = "name"
	FlagDenomID          = "denom-id"
	FlagSchema           = "schema"
	FlagSymbol           = "symbol"
	FlagMintRestricted   = "mint-restricted"
	FlagUpdateRestricted = "update-restricted"
)

var (
	FsIssueDenom  = flag.NewFlagSet("", flag.ContinueOnError)
	FsMintNFT     = flag.NewFlagSet("", flag.ContinueOnError)
	FsEditNFT     = flag.NewFlagSet("", flag.ContinueOnError)
	FsTransferNFT = flag.NewFlagSet("", flag.ContinueOnError)
	FsQuerySupply = flag.NewFlagSet("", flag.ContinueOnError)
	FsQueryOwner  = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsIssueDenom.String(FlagSchema, "", "Denom data structure definition")
	FsIssueDenom.String(FlagURI, "", "The uri for the class metadata stored off chain. It can define schema for Class and NFT `Data` attributes. Optional")
	FsIssueDenom.String(FlagURIHash, "", "The uri-hash is a hash of the document pointed by uri. Optional")
	FsIssueDenom.String(FlagDescription, "", "The description is a brief description of nft classification. Optional")
	FsIssueDenom.String(FlagDenomName, "", "The name of the denom")
	FsIssueDenom.String(FlagSymbol, "", "The symbol of the denom")
	FsIssueDenom.String(FlagData, "", "The data is the app specific metadata of the NFT class. Optional")
	FsIssueDenom.Bool(FlagMintRestricted, false, "mint restricted of nft under denom")
	FsIssueDenom.Bool(FlagUpdateRestricted, false, "update restricted of nft under denom")

	FsMintNFT.String(FlagRecipient, "", "The receiver of the nft, if not filled, the default is the sender of the transaction")
	FsMintNFT.String(FlagData, "", "The origin data of the nft")

	FsEditNFT.String(FlagData, "[do-not-modify]", "The token data of the nft")

	FsTransferNFT.String(FlagRecipient, "", "Domain NFT recipient")
	FsTransferNFT.String(FlagTokenPrice, "0", "Price for NFT in uDWS")
	FsTransferNFT.String(FlagCancel, "false", "Cancel created transfer for NFT")

	FsQuerySupply.String(FlagOwner, "", "The owner of the nft")

	FsQueryOwner.String(FlagDenomID, "", "The name of the collection")
}
