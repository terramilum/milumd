package keeper

import (
	context "context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/terramirum/mirumd/x/rental/types"
)

// BurnNft implements types.MsgServer
func (k Keeper) BurnNft(context context.Context, burnRequest *types.MsgBurnNftRequest) (*types.MsgBurnNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)
	err := k.nftKeeper.Burn(ctx, burnRequest.ClassId, burnRequest.NftId)
	if err != nil {
		return nil, err
	}
	return &types.MsgBurnNftResponse{
		NftId: burnRequest.NftId,
	}, nil
}
