package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/terramirum/mirumd/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) context.Context {
	_, ctx := keepertest.RentalKeeper(t)
	return sdk.WrapSDKContext(ctx)
}
