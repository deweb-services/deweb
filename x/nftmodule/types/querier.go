package types

// DONTCOVER

// query endpoints supported by the NFT Querier
const (
	QueryParams = "params"
	QueryNFT    = "nft"
)

// QueryNFTParams params for query 'custom/nfts/nft'
type QueryNFTParams struct {
	TokenID string
}

// NewQueryNFTParams creates a new instance of QueryNFTParams
func NewQueryNFTParams(id string) QueryNFTParams {
	return QueryNFTParams{
		TokenID: id,
	}
}
