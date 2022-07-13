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
	parentDomainOwner, err := k.getParentDomainOwner(ctx, dnsName)
	if err != nil {
		return errors.Wrap(err, "cannot get parent domain")
	}
	if !parentDomainOwner.Equals(creatorAddress) {
		return &DomainNotOwned{domainName: dnsName}
	}
	return nil
}

func (k Keeper) getParentDomainOwner(ctx types.Context, dnsName string) (types.AccAddress, error) {
	domainParts := strings.Split(dnsName, ".")
	if len(domainParts) == 1 {
		domainRec, err := k.GetDomain(ctx, domainParts[0])
		if err != nil {
			return nil, &DomainDoesntExist{domainName: domainParts[0]}
		}
		return domainRec.GetOwner(), nil
	}
	parentDomainParts := domainParts[1:]
	parentDomain := strings.Join(parentDomainParts, ".")
	domainRec, err := k.GetDomain(ctx, parentDomain)
	if err != nil {
		return nil, &DomainDoesntExist{domainName: dnsName}
	}
	return domainRec.GetOwner(), nil
}
