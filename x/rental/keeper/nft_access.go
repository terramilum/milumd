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
	currentDate := getNowUtc()

	response, err := k.getNftAccesses(ctx, currentDate, accessNftRequest.ClassId, accessNftRequest.NftId, accessNftRequest.Renter)
	if err != nil {
		return nil, err
	}

	hasAccess := "0"
	if response.HasAccess {
		hasAccess = "1"
	}
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeAccessNft,
		sdk.NewAttribute(types.AttributeKeyNftCurrentDate, fmt.Sprintf("%d", currentDate)),
		sdk.NewAttribute(types.AttributeKeyClassId, accessNftRequest.ClassId),
		sdk.NewAttribute(types.AttributeKeyNftId, accessNftRequest.NftId),
		sdk.NewAttribute(types.AttributeKeyNftRentAccess, hasAccess),
	))

	return response, nil
}

func (k Keeper) getNftAccesses(ctx sdk.Context, currentDate int64, classId, nftId, renter string) (*types.MsgAccessNftResponse, error) {
	store := ctx.KVStore(k.storeKey)
	response := &types.MsgAccessNftResponse{
		HasAccess: false,
		NftRents:  []*types.NftRent{},
	}

	renterKey := getStoreWithKey(KeyRentSessionId, classId, nftId, renter)
	allSessionStore := prefix.NewStore(store, renterKey)
	iterator := allSessionStore.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var nftRent types.NftRent
		k.cdc.MustUnmarshal(iterator.Value(), &nftRent)
		if nftRent.StartDate <= currentDate && nftRent.EndDate >= currentDate {
			response.HasAccess = true
			break
		}
		response.NftRents = append(response.NftRents, &nftRent)
	}
	return response, nil
}
