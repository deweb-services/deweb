package types

import (
	"fmt"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

const (
	domainBasePriceDWS             = 100000000 //100 DWS = 100 * 10^6 uDWS
	domainDefaultValidityMinutes   = 20
	domainOwnerProlongationMinutes = 10
)

var (
	KeyDomainPrice                                        = []byte("DomainPrice")
	KeyDomainExpirationMinutes                            = []byte("DomainExpiration")
	KeyDomainOwnerProlongationMinutes                     = []byte("DomainOwnerProlongation")
	KeyBlockedTLDs                                        = []byte("BlockedTLDs")
	_                                 paramtypes.ParamSet = (*Params)(nil)
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(domainPrice uint64, domainExpirationMinutes int64, domainOwnerProlong int64, blockTLDs []string) Params {
	return Params{
		DomainPrice:                    domainPrice,
		DomainExpirationMinutes:        domainExpirationMinutes,
		DomainOwnerProlongationMinutes: domainOwnerProlong,
		BlockTLDs:                      blockTLDs,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(domainBasePriceDWS, domainDefaultValidityMinutes, domainOwnerProlongationMinutes, DefaultDomainsBlockList)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyDomainPrice, &p.DomainPrice, validateDomainPrice),
		paramtypes.NewParamSetPair(KeyDomainExpirationMinutes, &p.DomainExpirationMinutes, validateDomainExpirationMinutes),
		paramtypes.NewParamSetPair(KeyDomainOwnerProlongationMinutes, &p.DomainOwnerProlongationMinutes, validateDomainOwnerProlongationMinutes),
		paramtypes.NewParamSetPair(KeyBlockedTLDs, &p.BlockTLDs, validateBlockedTLDs),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateDomainPrice(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("domain price must be positive: %d", v)
	}

	return nil
}

func validateDomainExpirationMinutes(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("domain expiration period minutes must be positive: %d", v)
	}

	return nil
}

func validateDomainOwnerProlongationMinutes(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("domain owner prolongation before expiration minutes must be positive: %d", v)
	}

	return nil
}

func validateBlockedTLDs(i interface{}) error {
	_, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}
