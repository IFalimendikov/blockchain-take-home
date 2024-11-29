package keeper

import (
	"blog/x/blog/types"
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Key format: postID + granter + grantee
func GetAuthorizationKey(postID uint64, granter, grantee string) []byte {
	key := append(GetPostIDBytes(postID), []byte(granter)...)
	return append(key, []byte(grantee)...)
}

func (k Keeper) CheckAuthorization(ctx sdk.Context, postID uint64, granter, grantee string) bool {
	if granter == grantee {
		return true
	}
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PostAuthorizationKey))
	authBytes := store.Get(GetAuthorizationKey(postID, granter, grantee))
	return authBytes != nil
}
