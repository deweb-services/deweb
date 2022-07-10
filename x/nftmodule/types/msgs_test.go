package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/deweb-services/deweb/x/nftmodule/types"
)

// ---------------------------------------- Msgs --------------------------------------------------

func TestMsgTransferNFTValidateBasicMethod(t *testing.T) {
	newMsgTransferNFT := types.NewMsgTransferDomain(denomID, "", id, tokenURI, uriHash, tokenData, address.String(), address2.String())
	err := newMsgTransferNFT.ValidateBasic()
	require.Error(t, err)

	newMsgTransferNFT = types.NewMsgTransferDomain(denomID, denom, "", tokenURI, uriHash, tokenData, "", address2.String())
	err = newMsgTransferNFT.ValidateBasic()
	require.Error(t, err)

	newMsgTransferNFT = types.NewMsgTransferDomain(denomID, denom, "", tokenURI, uriHash, tokenData, address.String(), "")
	err = newMsgTransferNFT.ValidateBasic()
	require.Error(t, err)

	newMsgTransferNFT = types.NewMsgTransferDomain(denomID, denom, id, tokenURI, uriHash, tokenData, address.String(), address2.String())
	err = newMsgTransferNFT.ValidateBasic()
	require.NoError(t, err)
}

func TestMsgTransferNFTGetSignBytesMethod(t *testing.T) {
	newMsgTransferNFT := types.NewMsgTransferDomain(denomID, denom, id, tokenURI, uriHash, tokenData, address.String(), address2.String())
	sortedBytes := newMsgTransferNFT.GetSignBytes()
	expected := `{"type":"irismod/nft/MsgTransferNFT","value":{"data":"https://google.com/token-1.json","denom_id":"denom","id":"denom","name":"id1","recipient":"cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgp0ctjdj","sender":"cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqjwl8sq","uri":"https://google.com/token-1.json","uri_hash":"uriHash"}}`
	require.Equal(t, expected, string(sortedBytes))
}

func TestMsgTransferNFTGetSignersMethod(t *testing.T) {
	newMsgTransferNFT := types.NewMsgTransferDomain(denomID, denom, id, tokenURI, uriHash, tokenData, address.String(), address2.String())
	signers := newMsgTransferNFT.GetSigners()
	require.Equal(t, 1, len(signers))
	require.Equal(t, address.String(), signers[0].String())
}

func TestMsgEditNFTValidateBasicMethod(t *testing.T) {
	newMsgEditNFT := types.NewMsgEditDomain(id, tokenData, "")

	err := newMsgEditNFT.ValidateBasic()
	require.Error(t, err)

	newMsgEditNFT = types.NewMsgEditDomain("", tokenData, address.String())
	err = newMsgEditNFT.ValidateBasic()
	require.Error(t, err)

	newMsgEditNFT = types.NewMsgEditDomain(id, tokenData, address.String())
	err = newMsgEditNFT.ValidateBasic()
	require.Error(t, err)

	newMsgEditNFT = types.NewMsgEditDomain(id, tokenData, address.String())
	err = newMsgEditNFT.ValidateBasic()
	require.NoError(t, err)
}

func TestMsgEditNFTGetSignBytesMethod(t *testing.T) {
	newMsgEditNFT := types.NewMsgEditDomain(id, tokenData, address.String())
	sortedBytes := newMsgEditNFT.GetSignBytes()
	expected := `{"type":"irismod/nft/MsgEditNFT","value":{"data":"https://google.com/token-1.json","denom_id":"denom","id":"id1","name":"report","sender":"cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqjwl8sq","uri":"https://google.com/token-1.json","uri_hash":"uriHash"}}`
	require.Equal(t, expected, string(sortedBytes))
}

func TestMsgEditNFTGetSignersMethod(t *testing.T) {
	newMsgEditNFT := types.NewMsgEditDomain(id, tokenData, address.String())
	signers := newMsgEditNFT.GetSigners()
	require.Equal(t, 1, len(signers))
	require.Equal(t, address.String(), signers[0].String())
}

func TestMsgMsgMintNFTValidateBasicMethod(t *testing.T) {
	newMsgMintNFT := types.NewMsgRegisterDomain(id, tokenData, "", address2.String())
	err := newMsgMintNFT.ValidateBasic()
	require.Error(t, err)

	newMsgMintNFT = types.NewMsgRegisterDomain("", tokenData, address.String(), address2.String())
	err = newMsgMintNFT.ValidateBasic()
	require.Error(t, err)

	newMsgMintNFT = types.NewMsgRegisterDomain(id, tokenData, address.String(), address2.String())
	err = newMsgMintNFT.ValidateBasic()
	require.Error(t, err)

	newMsgMintNFT = types.NewMsgRegisterDomain(id, tokenData, address.String(), address2.String())
	err = newMsgMintNFT.ValidateBasic()
	require.NoError(t, err)
}

func TestMsgMintNFTGetSignBytesMethod(t *testing.T) {
	newMsgMintNFT := types.NewMsgRegisterDomain(id, tokenData, address.String(), address2.String())
	sortedBytes := newMsgMintNFT.GetSignBytes()
	expected := `{"type":"irismod/nft/MsgMintNFT","value":{"data":"https://google.com/token-1.json","denom_id":"denom","id":"id1","name":"report","recipient":"cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgp0ctjdj","sender":"cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqjwl8sq","uri":"https://google.com/token-1.json","uri_hash":"uriHash"}}`
	require.Equal(t, expected, string(sortedBytes))
}

func TestMsgMsgBurnNFTValidateBasicMethod(t *testing.T) {
	newMsgBurnNFT := types.NewMsgRemoveDomain("", id)
	err := newMsgBurnNFT.ValidateBasic()
	require.Error(t, err)

	newMsgBurnNFT = types.NewMsgRemoveDomain(address.String(), "")
	err = newMsgBurnNFT.ValidateBasic()
	require.Error(t, err)

	newMsgBurnNFT = types.NewMsgRemoveDomain(address.String(), id)
	err = newMsgBurnNFT.ValidateBasic()
	require.Error(t, err)

	newMsgBurnNFT = types.NewMsgRemoveDomain(address.String(), id)
	err = newMsgBurnNFT.ValidateBasic()
	require.NoError(t, err)
}

func TestMsgBurnNFTGetSignBytesMethod(t *testing.T) {
	newMsgBurnNFT := types.NewMsgRemoveDomain(address.String(), id)
	sortedBytes := newMsgBurnNFT.GetSignBytes()
	expected := `{"type":"irismod/nft/MsgBurnNFT","value":{"denom_id":"denom","id":"id1","sender":"cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqjwl8sq"}}`
	require.Equal(t, expected, string(sortedBytes))
}

func TestMsgBurnNFTGetSignersMethod(t *testing.T) {
	newMsgBurnNFT := types.NewMsgRemoveDomain(address.String(), id)
	signers := newMsgBurnNFT.GetSigners()
	require.Equal(t, 1, len(signers))
	require.Equal(t, address.String(), signers[0].String())
}
