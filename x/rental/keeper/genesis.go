package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/terramirum/mirumd/x/rental/types"
	rental "github.com/terramirum/mirumd/x/rental/types"
)

// InitGenesis initializes the nft module's genesis state from a given
// genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, data *rental.GenesisState) {
	store := ctx.KVStore(k.storeKey)
	for _, class := range data.ClassOwners {
		err := k.saveContractOwner(ctx, class.ClassId, class.ContractOwner)
		if err != nil {
			panic(err)
		}
	}

	for _, rent := range data.RentedNfts {
		rentRequest := &types.MsgMintRentRequest{
			ClassId:   rent.ClassId,
			NftId:     rent.NftId,
			Renter:    rent.Renter,
			StartDate: rent.StartDate,
			EndDate:   rent.EndDate,
		}
		err := k.saveSessionOfNft(store, rentRequest)
		if err != nil {
			panic(err)
		}
	}
}

// ExportGenesis returns a GenesisState for a given context.
func (k Keeper) ExportGenesis(ctx sdk.Context) *rental.GenesisState {
	var classOwners []*types.ClassOwner
	classIdOwners := k.GetAllClassIdsOwners(ctx)
	for key, val := range classIdOwners {
		classOwners = append(classOwners, &types.ClassOwner{
			ClassId:       key,
			ContractOwner: val,
		})
	}

	return &rental.GenesisState{
		ClassOwners: classOwners,
		RentedNfts:  k.GetAllSessionOfNft(ctx),
	}
}
