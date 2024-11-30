package types

import (
	"testing"

	"blog/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgRevokePostAuthorization_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRevokePostAuthorization
		err  error
	}{
		{
			name: "invalid granter address",
			msg: MsgRevokePostAuthorization{
				Granter: "invalid_address",
				Grantee: sample.AccAddress(),
				Id:      1,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid grantee address",
			msg: MsgRevokePostAuthorization{
				Granter: sample.AccAddress(),
				Grantee: "invalid_address",
				Id:      1,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid post ID",
			msg: MsgRevokePostAuthorization{
				Granter: sample.AccAddress(),
				Grantee: sample.AccAddress(),
				Id:      0,
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "valid message",
			msg: MsgRevokePostAuthorization{
				Granter: sample.AccAddress(),
				Grantee: sample.AccAddress(),
				Id:      1,
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
