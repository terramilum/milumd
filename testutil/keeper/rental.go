package keeper

import (
	"testing"

	"github.com/terramirum/mirumd/x/rental/keeper"
	"github.com/terramirum/mirumd/x/rental/types"

	"cosmossdk.io/store"
	storetypes "cosmossdk.io/store/types"
	dbm "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func RentalKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	//nftKeeper := nftkeeper.NewKeeper(storeKey, cdc)
	// paramsSubspace := typesparams.NewSubspace(cdc,
	// 	types.Amino,
	// 	storeKey,
	// 	memStoreKey,
	// 	"RentalParams",
	// )

	k := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		nil,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}
