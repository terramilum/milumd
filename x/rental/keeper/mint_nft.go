package keeper

import (
	context "context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/terramirum/mirumd/x/rental/types"
)

// MintNft implements types.MsgServer
func (k Keeper) MintNft(context context.Context, mintRequest *types.MsgMintNftRequest) (*types.MsgMintNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)
	err := k.nftKeeper.Burn(ctx, mintRequest.ClassId, mintRequest.NftId)
	if err != nil {
		return nil, err
	}
	return &types.MsgMintNftResponse{
		NftId: mintRequest.NftId,
	}, nil
}
