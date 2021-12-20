package keeper

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/deweb-services/deweb/x/deweb/types"
)

const MaxMessageLength = 1000

func (k msgServer) SaveWallet(goCtx context.Context, msg *types.MsgSaveWallet) (*types.MsgSaveWalletResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	encryptedKey := msg.GetEncryptedKey()
	if len(encryptedKey) > MaxMessageLength {
		return nil, fmt.Errorf("received mesage greater then maximum lengtx %d", MaxMessageLength)
	}
	creator := msg.GetCreator()
	if msg.Chain == "" {
		// Maybe create a list of allowed chains
		return nil, fmt.Errorf("parameter chain required")
	}
	userRec := types.UserWalletRec{
		Address:      msg.Address,
		EncryptedKey: encryptedKey,
		Chain:        msg.Chain,
	}
	recordID := creator + "_" + userRec.Address
	err := k.writeUserKeyRecord(ctx, userRec, recordID)
	if err != nil {
		return nil, fmt.Errorf("error writing  message to store: %w", err)
	}
	err = k.appendUserRecordID(ctx, creator, recordID)
	if err != nil {
		return nil, fmt.Errorf("error appending created record to user records list: %w", err)
	}
	k.Logger(ctx).Error(fmt.Sprintf("[SaveWallet] Saved encrypted key with id: %s: %s", recordID, encryptedKey))
	// Update the post count
	return &types.MsgSaveWalletResponse{}, nil
}

func (k msgServer) writeUserKeyRecord(ctx sdk.Context, userRec types.UserWalletRec, idVal string) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.RecordsKey))

	recordValue, err := k.cdc.Marshal(&userRec)
	if err != nil {
		return err
	}

	store.Set([]byte(idVal), recordValue)
	return nil
}

func (k msgServer) appendUserRecordID(ctx sdk.Context, address string, recordID string) error {
	k.Logger(ctx).Error(fmt.Sprintf("[appendUserRecordID] Process"))

	existRecords, err := k.getUserRecordsIDs(ctx, address)
	if err != nil {
		return err
	}
	if existRecords == nil {
		existRecords = make([]string, 0, 1)
	}
	for _, recID := range existRecords {
		if recID == recordID {
			return nil
		}
	}
	existRecords = append(existRecords, recordID)
	k.Logger(ctx).Error(fmt.Sprintf("[appendUserRecordID] user records: %s", strings.Join(existRecords, ",")))
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.UsersRecords))

	resRecords := types.RecordsToUser{Records: existRecords}
	recordValue, err := k.cdc.Marshal(&resRecords)
	if err != nil {
		return err
	}
	recKey := []byte(address)
	store.Set(recKey, recordValue)
	return nil
}
