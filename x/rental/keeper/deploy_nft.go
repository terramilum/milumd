package keeper

import (
	context "context"
	"errors"

	"crypto/rand"
	"encoding/base64"
	"fmt"

	"cosmossdk.io/store/prefix"
	codec "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"cosmossdk.io/x/nft"
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

	nftDetail, err := codec.NewAnyWithValue(deployNftRequest.Detail)
	if err != nil {
		return nil, err
	}

	class := nft.Class{
		Id:          classId,
		Name:        deployNftRequest.Name,
		Symbol:      deployNftRequest.Symbol,
		Description: deployNftRequest.Description,
		Uri:         deployNftRequest.Uri,
		Data:        nftDetail,
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

	err = k.saveContractOwner(ctx, class.Id, deployNftRequest.ContractOwner)
	if err != nil {
		return nil, err
	}

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

func (k Keeper) saveContractOwner(ctx sdk.Context, classId, contractOwner string) error {
	store := ctx.KVStore(k.storeKey)
	store.Set(getStoreWithKey(KeyClassIdContract, classId), []byte(contractOwner))
	store.Set(getStoreWithKey(KeyContractClassId, contractOwner, classId), []byte("1"))
	return nil
}

func (k Keeper) GetAllClassIdsOwners(ctx sdk.Context) map[string]string {
	classIdContractOwners := make(map[string]string)
	store := ctx.KVStore(k.storeKey)
	classIdContractKey := getStoreWithKey(KeyClassIdContract)
	allClassIds := prefix.NewStore(store, classIdContractKey)
	iterator := allClassIds.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		keys := getParsedStoreKey(iterator.Key())
		classIdContractOwners[keys[0]] = string(iterator.Value())
	}
	return classIdContractOwners
}
