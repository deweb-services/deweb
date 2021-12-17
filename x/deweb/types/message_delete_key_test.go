package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/deweb-services/deweb/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgDeleteKey_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteKey
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteKey{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteKey{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
