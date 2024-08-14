package keeper

import (
	store "cosmossdk.io/core/store"
	storetypes "cosmossdk.io/store/types"
	nftkeeper "cosmossdk.io/x/nft/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		storeService store.KVStoreService
		cdc          codec.BinaryCodec
		paramstore   paramtypes.Subspace
		nftKeeper    *nftkeeper.Keeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	memKey storetypes.StoreKey,
	nftKeeper *nftkeeper.Keeper,
) *Keeper {
	// // set KeyTable if it has not already been set
	// if !ps.HasKeyTable() {
	// 	ps = ps.WithKeyTable(types.ParamKeyTable())
	// }

	return &Keeper{
		storeService: storeService,
		cdc:          cdc,
		paramstore:   paramtypes.Subspace{},
		nftKeeper:    nftKeeper,
	}
}
