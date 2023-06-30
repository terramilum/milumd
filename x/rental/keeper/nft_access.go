package keeper

import (
	context "context"
	"fmt"

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
	req := &types.QuerySessionRequest{
		ClassId: classId,
		NftId:   nftId,
		Renter:  renter,
	}

	res, err := k.Sessions(ctx, req)
	if err != nil {
		return nil, err
	}

	hasAccess := false
	for _, sessionDetail := range res.SessionDetail {
		nftRent := sessionDetail.NftRent
		if nftRent.StartDate <= currentDate && nftRent.EndDate >= currentDate {
			hasAccess = true
			break
		}
	}

	return &types.MsgAccessNftResponse{
		HasAccess: hasAccess,
		NftRents:  k.toNftRent(res.SessionDetail),
	}, nil
}
