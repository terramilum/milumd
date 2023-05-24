package keeper

import (
	"milumd/x/rental/types"
)

var _ types.QueryServer = Keeper{}
