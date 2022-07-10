package keeper

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/bank/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"time"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/deweb-services/deweb/x/nftmodule/types"
)

const DNSDenomName = "domains"

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey     storetypes.StoreKey // Unexposed key to access store from sdk.Context
	bankKeeper   keeper.Keeper
	dnsDenomName string
	paramStore   paramtypes.Subspace
	cdc          codec.Codec
}

// NewKeeper creates a new instance of the NFT Keeper
func NewKeeper(cdc codec.Codec, storeKey storetypes.StoreKey, bankKeeper keeper.Keeper, ps paramtypes.Subspace) Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}
	return Keeper{
		storeKey:     storeKey,
		bankKeeper:   bankKeeper,
		dnsDenomName: DNSDenomName,
		paramStore:   ps,
		cdc:          cdc,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("domains/%s", types.ModuleName))
}

// IssueDenom issues a denom according to the given params
func (k Keeper) IssueDenom(ctx sdk.Context,
	id, name, schema, symbol string,
	creator sdk.AccAddress,
	mintRestricted, updateRestricted bool,
	description, uri, uriHash, data string,
) error {
	return k.SetDenom(ctx, types.Denom{
		Id:               id,
		Name:             name,
		Schema:           schema,
		Creator:          creator.String(),
		Symbol:           symbol,
		MintRestricted:   mintRestricted,
		UpdateRestricted: updateRestricted,
		Description:      description,
		Uri:              uri,
		UriHash:          uriHash,
		Data:             data,
	})
}

// RegisterDomain mints an NFT and manages the NFT's existence within Collections and Owners
func (k Keeper) RegisterDomain(ctx sdk.Context, tokenID, tokenData string, owner sdk.AccAddress, recipient sdk.AccAddress) error {
	if k.HasNFT(ctx, k.dnsDenomName, tokenID) {
		domain, err := k.Authorize(ctx, k.dnsDenomName, tokenID, owner)
		if err != nil {
			return sdkerrors.Wrapf(types.ErrNFTAlreadyExists, "Domain %s already exists in collection %s", tokenID, k.dnsDenomName)
		}
		return k.domainProlongation(ctx, domain, owner)
	}
	err := k.checkNFTValid(ctx, tokenID, owner)
	if err != nil {
		return sdkerrors.Wrapf(err, "validation check failed")
	}
	domainReceivedData, err := ParseNFTData([]byte(tokenData))
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal, "invalid data")
	}
	domainReceivedData.Issued = time.Now()
	domainExpirationHours := k.DomainExpirationHours(ctx)
	domainReceivedData.ValidTill = time.Now().Add(time.Duration(domainExpirationHours) * time.Hour)

	dataToSaveRaw, err := json.Marshal(domainReceivedData)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, "internal error marshal record: %w", err)
	}
	dataToSave := string(dataToSaveRaw)

	zeroAddress := []byte("0")
	domainPriceUDWS := k.DomainPriceUDWS(ctx)
	err = k.payForDomain(ctx, domainPriceUDWS, owner, zeroAddress)
	if err != nil {
		return sdkerrors.Wrapf(err, "cannot burn coins for domain")
	}
	k.registerDomain(
		ctx,
		types.NewBaseNFT(tokenID, recipient, dataToSave),
	)
	k.setOwner(ctx, k.dnsDenomName, tokenID, recipient)
	k.increaseSupply(ctx, k.dnsDenomName)

	return nil
}

func (k Keeper) domainProlongation(ctx sdk.Context, domain types.BaseNFT, owner sdk.AccAddress) error {
	storedData := domain.GetData()
	domainRecordData, err := ParseNFTData([]byte(storedData))
	if err != nil {
		return sdkerrors.Wrapf(err, "cannot process record for domain %s", domain.GetID())
	}
	now := time.Now()
	prologHourss := k.DomainOwnerProlongationHours(ctx)
	prologDuration := time.Duration(prologHourss) * time.Hour
	allowedProlongationAfter := domainRecordData.ValidTill.Add(-prologDuration)
	if allowedProlongationAfter.After(now) {
		return sdkerrors.Wrapf(types.ErrNFTAlreadyExists,
			"Domain %s prolongation not allowed not, try after %s",
			domain.GetID(), allowedProlongationAfter.Format("2006-01-02T15:04:05"))
	}
	domainExpirationHours := k.DomainExpirationHours(ctx)
	domainRecordData.ValidTill = domainRecordData.ValidTill.Add(time.Duration(domainExpirationHours) * time.Hour)
	dataToSaveRaw, err := json.Marshal(domainRecordData)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, "internal error marshal record: %w", err)
	}
	dataToSave := string(dataToSaveRaw)
	zeroAddress := []byte("0")
	domainPriceUDWS := k.DomainPriceUDWS(ctx)
	err = k.payForDomain(ctx, domainPriceUDWS, owner, zeroAddress)
	if err != nil {
		return sdkerrors.Wrapf(err, "cannot burn coins for domain")
	}
	domain.Data = dataToSave
	k.registerDomain(ctx, domain)
	return nil
}

func (k Keeper) checkNFTValid(ctx sdk.Context, tokenID string, owner sdk.AccAddress) error {
	// check domain in blacklist
	domainBlocked := k.CheckDomainBlocked(ctx, tokenID)
	if domainBlocked {
		return sdkerrors.Wrapf(types.ErrInvalidDenom, "domain %s in block list", tokenID)
	}

	// Check ownership of upper-level domain
	chErr := k.CheckAllowedForAddress(ctx, tokenID, owner)
	if chErr != nil {
		return sdkerrors.Wrapf(types.ErrDomainPermissionDenied, chErr.Error())
	}

	return nil
}

func (k Keeper) ProcessDomainsExpiration(ctx sdk.Context) {
	allDomains := k.GetNFTs(ctx, k.dnsDenomName)
	now := time.Now()
	for _, domain := range allDomains {
		storedData := domain.GetData()
		domainRecord, err := ParseNFTData([]byte(storedData))
		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("cannot process record for domain %s", domain.GetID()), err)
			continue
		}
		if domainRecord.ValidTill.Before(now) {
			err := k.BurnNFT(ctx, domain.GetID(), domain.GetOwner())
			if err != nil {
				k.Logger(ctx).Error(fmt.Sprintf("cannot burn expired domain %s", domain.GetID()), err)
				continue
			}
		}
	}
}

// EditDomain updates an already existing NFT
func (k Keeper) EditDomain(ctx sdk.Context, tokenID, tokenData string, owner sdk.AccAddress) error {
	// just the owner of NFT can edit
	domain, err := k.Authorize(ctx, k.dnsDenomName, tokenID, owner)
	if err != nil {
		return err
	}

	if types.Modified(tokenData) {
		domainRecordData, err := ParseNFTData([]byte(domain.Data))
		if err != nil {
			return sdkerrors.Wrapf(err, "cannot process record for domain %s", domain.GetID())
		}
		receivedUpdateData, err := ParseNFTData([]byte(tokenData))
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal, "invalid data")
		}
		domainRecordData.Records = receivedUpdateData.Records
		dataToSaveRaw, err := json.Marshal(domainRecordData)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, "internal error marshal record: %w", err)
		}
		dataToSave := string(dataToSaveRaw)
		domain.Data = dataToSave
	}

	k.registerDomain(ctx, domain)

	return nil
}

// TransferDomainOwner performs 2-step transfer
func (k Keeper) TransferDomainOwner(ctx sdk.Context, tokenID string, cancelTransfer bool, price uint64, trxSender sdk.AccAddress, recipientAddr string) error {
	denom, found := k.GetDenom(ctx, k.dnsDenomName)
	if !found {
		return sdkerrors.Wrapf(types.ErrInvalidDenom, "denom ID %s not exists", k.dnsDenomName)
	}

	if denom.UpdateRestricted {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "It is restricted to update NFT under this denom %s", denom.Id)
	}

	domainRec, err := k.GetNFT(ctx, tokenID)
	if err != nil {
		return err
	}
	domainRecordData, err := ParseNFTData([]byte(domainRec.GetData()))
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal,
			fmt.Sprintf("cannot process record for domain %s", domainRec.GetID()), err)
	}
	currentOwner := domainRec.GetOwner()
	domainRecUpdate := domainRec.(types.BaseNFT)
	if currentOwner.Equals(trxSender) {
		// Create of cancel transfer offer by owner
		if cancelTransfer {
			domainRecordData.TransferOffer = nil
		} else {
			transferOffer := &DNSTransferOffer{
				Price:                price,
				ExpectedOwnerAddress: recipientAddr,
			}
			domainRecordData.TransferOffer = transferOffer
		}
	} else {
		// process request to buy
		if domainRecordData.TransferOffer == nil {
			return sdkerrors.Wrap(types.ErrUnauthorized, trxSender.String())
		}
		transferOffer := domainRecordData.TransferOffer
		expectedReceiver := transferOffer.ExpectedOwnerAddress
		if expectedReceiver != "" {
			receivedAddr, err := sdk.AccAddressFromBech32(expectedReceiver)
			if err != nil {
				return sdkerrors.Wrap(err, "cannot process expected receiver address")
			}
			if !receivedAddr.Equals(trxSender) {
				return sdkerrors.Wrap(types.ErrUnauthorized,
					fmt.Sprintf("expected domain receiver address is %s", expectedReceiver))
			}
		}
		err = k.payForDomain(ctx, transferOffer.Price, domainRec.GetOwner(), trxSender)
		if err != nil {
			return sdkerrors.Wrapf(err, "cannot send coins to pay for domain")
		}
		k.swapOwner(ctx, k.dnsDenomName, tokenID, domainRec.GetOwner(), trxSender)
		domainRecUpdate.Owner = trxSender.String()
		domainRecordData.TransferOffer = nil
	}

	dataToSaveRaw, err := json.Marshal(domainRecordData)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, "internal error marshal record: %w", err)
	}
	dataToSave := string(dataToSaveRaw)

	domainRecUpdate.Data = dataToSave
	k.registerDomain(ctx, domainRecUpdate)

	return nil
}

func (k Keeper) payForDomain(ctx sdk.Context, amountUDWS uint64, sender sdk.AccAddress, target sdk.AccAddress) error {
	coinToSend := sdk.Coin{
		Denom:  "udws",
		Amount: sdk.NewIntFromUint64(amountUDWS),
	}

	coins := sdk.Coins{}
	coins = append(coins, coinToSend)
	err := k.bankKeeper.SendCoins(ctx, sender, target, coins)
	if err != nil {
		return err
	}
	return nil
}

// BurnNFT deletes a specified NFT
func (k Keeper) BurnNFT(ctx sdk.Context, tokenID string, owner sdk.AccAddress) error {
	if !k.HasDenomID(ctx, k.dnsDenomName) {
		return sdkerrors.Wrapf(types.ErrInvalidDenom, "denom ID %s not exists", k.dnsDenomName)
	}

	nft, err := k.Authorize(ctx, k.dnsDenomName, tokenID, owner)
	if err != nil {
		return err
	}

	k.deleteNFT(ctx, k.dnsDenomName, nft)
	k.deleteOwner(ctx, k.dnsDenomName, tokenID, owner)
	k.decreaseSupply(ctx, k.dnsDenomName)

	return nil
}
