package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"strings"
)

type DomainAlreadyCreated struct {
	domainName string
}

func (e DomainAlreadyCreated) Error() string {
	return fmt.Sprintf("domain %s already created", e.domainName)
}

type DomainNotOwned struct {
	domainName string
}

func (e DomainNotOwned) Error() string {
	return fmt.Sprintf("domain %s does not belong to this user", e.domainName)
}

type DomainDoesntExist struct {
	domainName string
}

func (e DomainDoesntExist) Error() string {
	return fmt.Sprintf("domain %s was not registered", e.domainName)
}

var DomainsBlocklist = []string{"com", "org", "edu"}

func CheckDomainBlocked(domain string) bool {
	for _, blockedLD := range DomainsBlocklist {
		if blockedLD == domain {
			return true
		}
	}
	return false
}

func (m msgServer) CheckAllowedForAddress(ctx types.Context, dnsName string, creatorAddress types.AccAddress) error {
	domainParts := strings.Split(dnsName, ".")
	if len(domainParts) == 1 {
		_, err := m.Keeper.GetNFT(ctx, m.dnsDenomName, domainParts[0])
		if err == nil {
			return DomainAlreadyCreated{domainName: domainParts[0]}
		}
		return nil
	}
	parentDomainParts := domainParts[1:len(domainParts)]
	parentDomain := strings.Join(parentDomainParts, ".")
	chErr := m.checkUserOwnsDomain(ctx, parentDomain, creatorAddress)
	if chErr != nil {
		return errors.Wrap(chErr, "parent domain check error")
	}
	return nil
}

func (m msgServer) checkUserOwnsDomain(ctx types.Context, dnsName string, creatorAddress types.AccAddress) error {
	nftRec, err := m.Keeper.GetNFT(ctx, m.dnsDenomName, dnsName)
	if err != nil {
		return DomainDoesntExist{domainName: dnsName}
	}
	if !nftRec.GetOwner().Equals(creatorAddress) {
		return DomainNotOwned{domainName: dnsName}
	}
	return nil
}
