package types

import (
	"github.com/deweb-services/deweb/x/dns_module/exported"
)

// NewCollection creates a new Domain Collection
func NewCollection(denom Denom, domains []exported.Domain) (c Collection) {
	c.Denom = denom
	for _, nft := range domains {
		c = c.AddDomain(nft.(BaseDomain))
	}
	return c
}

// AddDomain adds an Domain to the collection
func (c Collection) AddDomain(nft BaseDomain) Collection {
	c.NFTs = append(c.NFTs, nft)
	return c
}

func (c Collection) Supply() int {
	return len(c.NFTs)
}

// NewCollection creates a new Domain Collection
func NewCollections(c ...Collection) []Collection {
	return append([]Collection{}, c...)
}
