package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/terramirum/mirumd/x/rental/types"
)

// RentNftBurn implements types.MsgServer
func (k Keeper) RentNftBurn(context context.Context, burnRentRequest *types.MsgBurnRentRequest) (*types.MsgBurnRentResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)
	store := ctx.KVStore(k.storeKey)

	sessionIdKey := getStoreWithKey(KeyRentSessionId, burnRentRequest.ClassId, burnRentRequest.NftId, burnRentRequest.SessionId)
	store.Delete(sessionIdKey)

	//k.clearSessionIdRenters(ctx, burnRentRequest.ClassId, burnRentRequest.NftId, burnRentRequest.SessionId)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeBurnRentNft,
		sdk.NewAttribute(types.AttributeKeySessionId, string(sessionIdKey)),
	))

	return &types.MsgBurnRentResponse{}, nil
}
