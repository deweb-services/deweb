package types

// DONTCOVER

// query endpoints supported by the NFT Querier
const (
	QueryParams = "params"
	QueryDomain = "domain"
)

// QueryDomainParams params for query 'custom/nfts/nft'
type QueryDomainParams struct {
	TokenID string
}

// NewQueryDomainsParams creates a new instance of QueryDomainParams
func NewQueryDomainsParams(id string) QueryDomainParams {
	return QueryDomainParams{
		TokenID: id,
	}
}
