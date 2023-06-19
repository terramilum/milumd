package keeper

import (
	context "context"
	"errors"

	"crypto/rand"
	"encoding/base64"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/terramirum/mirumd/x/rental/types"
)

func (k Keeper) DeployNft(context context.Context, deployNftRequest *types.MsgDeployNftRequest) (*types.MsgDeployNftResponse, error) {

	ctx := sdk.UnwrapSDKContext(context)
	_, err := sdk.AccAddressFromBech32(deployNftRequest.ContractOwner)
	if err != nil {
		return nil, err
	}

	classId, err := GenerateNonce()
	if err != nil {
		return nil, err
	}

	class := nft.Class{
		Id:          classId,
		Name:        deployNftRequest.Name,
		Symbol:      deployNftRequest.Symbol,
		Description: deployNftRequest.Description,
		Uri:         deployNftRequest.Uri,
	}

	err = k.nftKeeper.SaveClass(ctx, class)
	if err != nil {
		return nil, err
	}

	getClass, exists := k.nftKeeper.GetClass(ctx, class.Id)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.New(fmt.Sprintf("Class not exist %s", class.Id))
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(getStoreWithKey(KeyClassContract, class.Id), []byte(deployNftRequest.ContractOwner))
	store.Set(getStoreWithKey(KeyClassIdContract, deployNftRequest.ContractOwner, class.Id), []byte("1"))

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeDeployNft,
		sdk.NewAttribute(types.AttributeKeyClassId, getClass.Id),
	))

	return &types.MsgDeployNftResponse{
		ClassId: getClass.Id,
	}, nil
}

func GenerateNonce() (string, error) {
	nonceBytes := make([]byte, 32)
	_, err := rand.Read(nonceBytes)
	if err != nil {
		return "", fmt.Errorf("could not generate nonce")
	}

	return base64.URLEncoding.EncodeToString(nonceBytes), nil
}
