package keeper

import (
	"reflect"
	"unsafe"

	"github.com/terramirum/mirumd/x/rental/types"
)

var (
	RentDates       = []byte{0x01}
	ClassContract   = []byte{0x02}
	RentSessionId   = []byte{0x03}
	SessionIdRenter = []byte{0x04}
	ClassIdContract = []byte{0x05}

	Delimiter   = []byte{0x00}
	Placeholder = []byte{0x01}
)

// StoreKey is the store key string for nft
const StoreKey = types.ModuleName

func renterDatesStoreKey(classID, nftID, address string) []byte {
	// key is of format:
	classIDBz := UnsafeStrToBytes(classID)
	nftIDBz := UnsafeStrToBytes(nftID)

	key := make([]byte, len(RentDates)+len(classIDBz)+len(nftIDBz)+len(Delimiter)+len(address))
	copy(key, RentDates)
	copy(key[len(RentDates):], classIDBz)
	copy(key[len(RentDates)+len(classIDBz):], nftIDBz)
	copy(key[len(RentDates)+len(classIDBz)+len(nftIDBz):], Delimiter)
	copy(key[len(RentDates)+len(classIDBz)+len(nftIDBz)+len(Delimiter):], address)
	return key
}

func nftRentDatesStoreKey(classID, nftID string) []byte {
	// key is of format:
	classIDBz := UnsafeStrToBytes(classID)
	nftIDBz := UnsafeStrToBytes(nftID)

	key := make([]byte, len(RentSessionId)+len(classIDBz)+len(nftIDBz))
	copy(key, RentSessionId)
	copy(key[len(RentSessionId):], classIDBz)
	copy(key[len(RentSessionId)+len(classIDBz):], nftIDBz)
	return key
}

func nftRentDatesSessionIdStoreKey(classID, nftID, sessionId string) []byte {
	// key is of format:
	nftId := UnsafeStrToBytes(classID + nftID)
	sessionIdz := UnsafeStrToBytes(sessionId)

	key := make([]byte, len(RentSessionId)+len(nftId)+len(Delimiter)+len(sessionIdz)+len(Delimiter))
	copy(key, RentSessionId)
	copy(key[len(RentSessionId):], nftId)
	copy(key[len(RentSessionId)+len(nftId):], Delimiter)
	copy(key[len(RentSessionId)+len(nftId)+len(Delimiter):], sessionIdz)
	copy(key[len(RentSessionId)+len(nftId)+len(Delimiter)+len(sessionIdz):], Delimiter)
	return key
}

func nftSessionIdRentersStoreKey(classID, nftID, sessionId string) []byte {
	// key is of format:
	nftId := UnsafeStrToBytes(classID + nftID)
	sessionIdz := UnsafeStrToBytes(sessionId)

	key := make([]byte, len(SessionIdRenter)+len(nftId)+len(Delimiter)+len(sessionIdz)+len(Delimiter))
	copy(key, SessionIdRenter)
	copy(key[len(SessionIdRenter):], nftId)
	copy(key[len(SessionIdRenter)+len(nftId):], Delimiter)
	copy(key[len(SessionIdRenter)+len(nftId)+len(Delimiter):], sessionIdz)
	copy(key[len(SessionIdRenter)+len(nftId)+len(Delimiter)+len(sessionIdz):], Delimiter)
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

// classStoreKey returns the byte representation of the nft class key
func classContractAddressKey(classID string) []byte {
	key := make([]byte, len(ClassContract)+len(classID))
	copy(key, ClassContract)
	copy(key[len(ClassContract):], classID)
	return key
}

// classStoreKey returns the byte representation of the nft class key
func contractAddressClassIdKey(contractAddress string) []byte {
	key := make([]byte, len(ClassIdContract)+len(contractAddress))
	copy(key, ClassIdContract)
	copy(key[len(ClassIdContract):], contractAddress)
	return key
}
