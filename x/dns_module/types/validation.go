package types

import (
	"fmt"
	"regexp"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	DoNotModify  = "[do-not-modify]"
	MinDenomLen  = 3
	MaxDenomLen  = 128
	MinDomainLen = 1
	MaxDomainLen = 255

	MaxTokenURILen = 256

	ReservedPeg  = "peg"
	ReservedIBC  = "ibc"
	ReservedHTLT = "htlt"
	ReservedTIBC = "tibc"
)

var (
	// IsAlphaNumeric only accepts [a-z0-9]
	IsAlphaNumeric = regexp.MustCompile(`^[a-z0-9]+$`).MatchString
	// IsDomainValidChars only accepts [a-z0-9\.]
	IsDomainValidChars = regexp.MustCompile(`^[a-z0-9\.-]+$`).MatchString
	// IsBeginWithAlpha only begin with [a-z]
	IsBeginWithAlpha = regexp.MustCompile(`^[a-z].*`).MatchString

	keywords          = strings.Join([]string{ReservedPeg, ReservedIBC, ReservedHTLT, ReservedTIBC}, "|")
	regexpKeywordsFmt = fmt.Sprintf("^(%s).*", keywords)
	regexpKeyword     = regexp.MustCompile(regexpKeywordsFmt).MatchString
)

// ValidateDenomID verifies whether the  parameters are legal
func ValidateDenomID(denomID string) error {
	if len(denomID) < MinDenomLen || len(denomID) > MaxDenomLen {
		return sdkerrors.Wrapf(ErrInvalidDenom, "the length of denom(%s) only accepts value [%d, %d]", denomID, MinDenomLen, MaxDenomLen)
	}
	boolPrifix := strings.HasPrefix(denomID, "tibc-")
	if !IsBeginWithAlpha(denomID) || !IsAlphaNumeric(denomID) && !boolPrifix {
		return sdkerrors.Wrapf(ErrInvalidDenom, "the denom(%s) only accepts alphanumeric characters, and begin with an english letter", denomID)
	}
	return nil
}

// ValidateTokenID verify that the tokenID is legal
func ValidateTokenID(tokenID string) error {
	if len(tokenID) < MinDomainLen || len(tokenID) > MaxDomainLen {
		return sdkerrors.Wrapf(ErrInvalidTokenID, "the length of domain name (%s) only accepts value [%d, %d]", tokenID, MinDenomLen, MaxDomainLen)
	}
	if !IsBeginWithAlpha(tokenID) || !IsDomainValidChars(tokenID) {
		return sdkerrors.Wrapf(ErrInvalidTokenID, "nft domain name (%s) only accepts alphanumeric characters, dots, dashes, and begin with an english letter", tokenID)
	}
	return nil
}

// Modified returns whether the field is modified
func Modified(target string) bool {
	return target != DoNotModify
}

// ValidateKeywords checks if the given denomId begins with `DenomKeywords`
func ValidateKeywords(denomId string) error {
	if regexpKeyword(denomId) {
		return sdkerrors.Wrapf(ErrInvalidDenom, "invalid denomId: %s, can not begin with keyword: (%s)", denomId, keywords)
	}
	return nil
}
