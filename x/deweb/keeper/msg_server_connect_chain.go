package keeper

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/deweb-services/deweb/x/deweb/types"
	"strings"
)

func (k msgServer) ConnectChain(goCtx context.Context, msg *types.MsgConnectChain) (*types.MsgConnectChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	externalAddress := msg.GetAddress()
	externalChain := msg.GetChain()
	creator := msg.GetCreator()

	itemToStore := types.ChainAddressMapping{
		ExtAddress: externalAddress,
		Chain:      externalChain,
	}

	idVal := creator + "_" + externalChain + "_" + externalAddress
	mappingWriteError := k.writeChainMappingRecord(ctx, itemToStore, idVal)
	if mappingWriteError != nil {
		return nil, fmt.Errorf("error writing  message to store: %w", mappingWriteError)
	}

	toUserErr := k.appendMappingRecordToUser(ctx, creator, idVal)
	if toUserErr != nil {
		return nil, fmt.Errorf("error adding created record to user")
	}

	return &types.MsgConnectChainResponse{}, nil
}

func (k msgServer) writeChainMappingRecord(ctx sdk.Context, mappingRec types.ChainAddressMapping, idVal string) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ConnectChainRecords))

	recordValue, err := k.cdc.Marshal(&mappingRec)
	if err != nil {
		return err
	}

	store.Set([]byte(idVal), recordValue)
	return nil
}

func (k msgServer) appendMappingRecordToUser(ctx sdk.Context, address string, recordID string) error {
	k.Logger(ctx).Error(fmt.Sprintf("[appendUserRecordID] Process"))

	existRecords, err := k.getChainMappingRecordsIDs(ctx, address)
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
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.UserConnectChainRecords))

	resRecords := types.RecordsToUser{Records: existRecords}
	recordValue, err := k.cdc.Marshal(&resRecords)
	if err != nil {
		return err
	}
	recKey := []byte(address)
	store.Set(recKey, recordValue)
	return nil
}
