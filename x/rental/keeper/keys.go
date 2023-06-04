package keeper

import (
	"reflect"
	"unsafe"

	"github.com/terramirum/mirumd/x/rental/types"
)

var (
	RenterDates = []byte{0x01}

	Delimiter   = []byte{0x00}
	Placeholder = []byte{0x01}
)

// StoreKey is the store key string for nft
const StoreKey = types.ModuleName

func renterDatesStoreKey(classID, nftID, address string) []byte {
	// key is of format:
	classIDBz := UnsafeStrToBytes(classID)
	nftIDBz := UnsafeStrToBytes(nftID)

	key := make([]byte, len(RenterDates)+len(classIDBz)+len(nftIDBz)+len(Delimiter)+len(address))
	copy(key, RenterDates)
	copy(key[len(RenterDates):], classIDBz)
	copy(key[len(RenterDates)+len(classIDBz):], nftIDBz)
	copy(key[len(RenterDates)+len(classIDBz)+len(nftIDBz):], Delimiter)
	copy(key[len(RenterDates)+len(classIDBz)+len(nftIDBz)+len(Delimiter):], address)
	return key
}

// UnsafeStrToBytes uses unsafe to convert string into byte array. Returned bytes
// must not be altered after this function is called as it will cause a segmentation fault.
func UnsafeStrToBytes(s string) []byte {
	var buf []byte
	sHdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bufHdr := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	bufHdr.Data = sHdr.Data
	bufHdr.Cap = sHdr.Len
	bufHdr.Len = sHdr.Len
	return buf
}

// UnsafeBytesToStr is meant to make a zero allocation conversion
// from []byte -> string to speed up operations, it is not meant
// to be used generally, but for a specific pattern to delete keys
// from a map.
func UnsafeBytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
