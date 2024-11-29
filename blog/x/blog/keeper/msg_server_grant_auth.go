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

func (k msgServer) GrantPostAuthorization(goCtx context.Context, msg *types.MsgGrantPostAuthorization) (*types.MsgGrantPostAuthorizationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	val, found := k.GetPost(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", val.Id))
	}
	if msg.Granter != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PostAuthorizationKey))
	auth := types.MsgGrantPostAuthorization{
		Granter: msg.Granter,
		Grantee: msg.Grantee,
		Id:      msg.Id,
	}
	
	authBytes := k.cdc.MustMarshal(&auth)
	store.Set(GetAuthorizationKey(msg.Id, msg.Granter, msg.Grantee), authBytes)
	return &types.MsgGrantPostAuthorizationResponse{}, nil
}
