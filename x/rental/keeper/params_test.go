package keeper_test

import (
	"testing"

	testkeeper "mirumd/testutil/keeper"
	"mirumd/x/rental/types"

	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.RentalKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
