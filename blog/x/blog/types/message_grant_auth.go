package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgGrantPostAuthorization{}

func NewMsgGrantPostAuthorization(granter string, grantee string, postID uint64) *MsgGrantPostAuthorization {
	return &MsgGrantPostAuthorization{
		Granter: granter,
		Grantee: grantee,
		Id:      postID,
	}
}

func (msg *MsgGrantPostAuthorization) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Granter)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid granter address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Grantee)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid grantee address (%s)", err)
	}

	if msg.Id == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "post ID cannot be 0")
	}

	return nil
}
