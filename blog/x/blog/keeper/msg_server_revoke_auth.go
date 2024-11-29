package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"blog/x/blog/types"
)

func (k msgServer) RevokePostAuthorization(goCtx context.Context, msg *types.MsgRevokePostAuthorization) (*types.MsgRevokePostAuthorizationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	val, found := k.GetPost(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("post %d doesn't exist", msg.Id))
	}
	if msg.Granter != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "only post creator can revoke authorization")
	}
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PostAuthorizationKey))
	store.Delete(GetAuthorizationKey(msg.Id, msg.Granter, msg.Grantee))
	return &types.MsgRevokePostAuthorizationResponse{}, nil
}
