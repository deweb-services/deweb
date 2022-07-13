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

func (k Keeper) CheckDomainBlocked(ctx types.Context, domain string) bool {
	upperDomain := strings.ToUpper(domain)
	blockedDomains := k.BlockedTLDs(ctx)
	for _, blockedLD := range blockedDomains {
		if blockedLD == upperDomain {
			return true
		}
	}
	return false
}

func (k Keeper) CheckAllowedForAddress(ctx types.Context, dnsName string, creatorAddress types.AccAddress) error {
	domainParts := strings.Split(dnsName, ".")
	if len(domainParts) == 1 {
		_, err := k.GetDomain(ctx, domainParts[0])
		if err == nil {
			return DomainAlreadyCreated{domainName: domainParts[0]}
		}
		return nil
	}
	parentDomainParts := domainParts[1:len(domainParts)]
	parentDomain := strings.Join(parentDomainParts, ".")
	chErr := k.checkUserOwnsDomain(ctx, parentDomain, creatorAddress)
	if chErr != nil {
		return errors.Wrap(chErr, "parent domain check error")
	}
	return nil
}

func (k Keeper) checkUserOwnsDomain(ctx types.Context, dnsName string, creatorAddress types.AccAddress) error {
	domainRec, err := k.GetDomain(ctx, dnsName)
	if err != nil {
		return DomainDoesntExist{domainName: dnsName}
	}
	if !domainRec.GetOwner().Equals(creatorAddress) {
		return DomainNotOwned{domainName: dnsName}
	}
	return nil
}
