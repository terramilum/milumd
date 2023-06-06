package keeper

import (
	context "context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/terramirum/mirumd/x/rental/types"
)

// NftAccess implements types.MsgServer
func (k Keeper) NftAccess(context context.Context, accessNftRequest *types.MsgAccessNftRequest) (*types.MsgAccessNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)
	store := ctx.KVStore(k.storeKey)

	hasAccess := "0"
	currentDate := k.getNowUtc()
	renterSessions := k.getRenterSessions(ctx, accessNftRequest.ClassId, accessNftRequest.NftId, accessNftRequest.Renter)
	for _, v := range renterSessions {
		sessionKey := store.Get(v)
		bz := store.Get(sessionKey)
		var nftRent types.NftRent
		k.cdc.MustUnmarshal(bz, &nftRent)
		if nftRent.StartDate <= currentDate && nftRent.EndDate >= currentDate {
			hasAccess = "1"
			break
		}
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeAccessNft,
		sdk.NewAttribute(types.AttributeKeyNftCurrentDate, fmt.Sprintf("%d", currentDate)),
		sdk.NewAttribute(types.AttributeKeyClassId, accessNftRequest.ClassId),
		sdk.NewAttribute(types.AttributeKeyNftId, accessNftRequest.NftId),
		sdk.NewAttribute(types.AttributeKeyNftRentAccess, hasAccess),
	))

	return &types.MsgAccessNftResponse{
		HasAccess: hasAccess == "1",
	}, nil
}

func (k Keeper) getRenterSessions(ctx sdk.Context, classId, nftId, renter string) (sessionKeys [][]byte) {
	store := ctx.KVStore(k.storeKey)
	key := renterDatesStoreKey(classId, nftId, renter)
	allSessionStore := prefix.NewStore(store, key)
	iterator := allSessionStore.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		sessionKeys = append(sessionKeys, iterator.Value())
	}
	return sessionKeys
}
