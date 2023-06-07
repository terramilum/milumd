package keeper

import (
	"strconv"
	"time"
)

const (
	YYYYMMDDHHMM = "200601021504"
)

func (*Keeper) getNowUtc() int64 {
	now := time.Now().UTC()
	formatted := now.Format(YYYYMMDDHHMM)
	d, _ := strconv.ParseInt(formatted, 10, 16)
	return d
}

func (*Keeper) getNowUtcAddMin(addMin int32) int64 {
	now := time.Now().Add(time.Minute * time.Duration(addMin)).UTC()
	formatted := now.Format(YYYYMMDDHHMM)
	d, _ := strconv.ParseInt(formatted, 10, 16)
	return d
}
