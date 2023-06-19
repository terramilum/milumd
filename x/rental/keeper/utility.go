package keeper

import (
	"strconv"
	"time"
)

const (
	YYYYMMDDHHMM = "200601021504"
)

func getNowUtc() int64 {
	now := time.Now().UTC()
	formatted := now.Format(YYYYMMDDHHMM)
	d, _ := strconv.ParseInt(formatted, 10, 64)
	return d
}

func getNowUtcAddMin(addMin int32) int64 {
	now := time.Now().Add(time.Minute * time.Duration(addMin)).UTC()
	formatted := now.Format(YYYYMMDDHHMM)
	d, _ := strconv.ParseInt(formatted, 10, 64)
	return d
}
