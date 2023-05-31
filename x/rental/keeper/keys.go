package keeper

import (
	"reflect"
	"unsafe"

	"github.com/terramirum/mirumd/x/rental/types"
)

var (
	NftKey           = []byte{0x01}
	ClassTotalSupply = []byte{0x02}
	OwnerKey         = []byte{0x03}

	Delimiter   = []byte{0x00}
	Placeholder = []byte{0x01}
)

// StoreKey is the store key string for nft
const StoreKey = types.ModuleName

// ownerStoreKey returns the byte representation of the nft owner
// Items are stored with the following key: values
// 0x04<classID><Delimiter(1 Byte)><nftID>
func ownerStoreKey(classID, nftID string) []byte {
	// key is of format:
	classIDBz := UnsafeStrToBytes(classID)
	nftIDBz := UnsafeStrToBytes(nftID)

	key := make([]byte, len(OwnerKey)+len(classIDBz)+len(Delimiter)+len(nftIDBz))
	copy(key, OwnerKey)
	copy(key[len(OwnerKey):], classIDBz)
	copy(key[len(OwnerKey)+len(classIDBz):], Delimiter)
	copy(key[len(OwnerKey)+len(classIDBz)+len(Delimiter):], nftIDBz)
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
