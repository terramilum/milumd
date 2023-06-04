package keeper

import (
	context "context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/terramirum/mirumd/x/rental/types"
)

// MintNft implements types.MsgServer
func (k Keeper) MintNft(context context.Context, mintRequest *types.MsgMintNftRequest) (*types.MsgMintNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)

	nfts := k.nftKeeper.GetNFTsOfClass(ctx, mintRequest.ClassId)
	nftId := fmt.Sprintf("%d", len(nfts)+1)

	nft := nft.NFT{
		ClassId: mintRequest.ClassId,
		Id:      nftId,
		Uri:     "/" + nftId,
	}

	reciever, err := sdk.AccAddressFromBech32(mintRequest.ContractOwner)
	if err != nil {
		return nil, err
	}

	err = k.nftKeeper.Mint(ctx, nft, reciever)
	if err != nil {
		return nil, err
	}

	return &types.MsgMintNftResponse{
		NftId: mintRequest.NftId,
	}, nil
}
