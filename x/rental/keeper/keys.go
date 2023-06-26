package keeper

import (
	"bytes"
	"reflect"
	"unsafe"

	"github.com/terramirum/mirumd/x/rental/types"
)

var (
	KeyRentDates = []byte{0x01}
	// store class id contract id to get contract owner of class.
	KeyClassIdContract = []byte{0x02}
	KeyRentSessionId   = []byte{0x03}
	KeySessionIdRenter = []byte{0x04}
	// store contract address and class id to prevent minting nft at the same class.
	KeyContractClassId = []byte{0x05}
	KeyRentDatesOwner  = []byte{0x05}

	Delimiter   = []byte{0x00}
	Placeholder = []byte{0x01}
)

// StoreKey is the store key string for nft
const StoreKey = types.ModuleName

// geting store key all keys include delimiter.
func getStoreWithKey(keyValue []byte, params ...string) []byte {
	newParams := make([][]byte, len(params)+1)
	newParams[0] = keyValue
	for i := 1; i < len(newParams); i++ {
		newParams[i] = UnsafeStrToBytes(params[i-1])
	}
	return getStoreKey(newParams...)
}

func getStoreWithKeyWithDelimiter(keyValue []byte, params ...string) []byte {
	newParams := make([][]byte, len(params)+2)
	newParams[0] = keyValue
	for i := 1; i < len(newParams)-1; i++ {
		newParams[i] = UnsafeStrToBytes(params[i-1])
	}
	newParams[len(newParams)-1] = Delimiter
	return getStoreKey(newParams...)
}

func getStoreKey(params ...[]byte) []byte {
	keyLen := 0
	for i := 0; i < len(params); i++ {
		keyLen += len(params[i])
		if i < len(params)-1 {
			keyLen += len(Delimiter)
		}
	}

	key := make([]byte, keyLen)
	positionLen := 0
	for i := 0; i < len(params); i++ {
		copy(key[positionLen:], params[i])
		positionLen += len(params[i])
		if i < len(params)-1 {
			copy(key[positionLen:], Delimiter)
			positionLen += len(Delimiter)
		}
	}
	return key
}

func getParsedStoreKey(key []byte) []string {
	splittedArray := bytes.Split(key, Delimiter)
	parsed := make([]string, len(splittedArray))
	for i := 0; i < len(parsed); i++ {
		parsed[i] = string(splittedArray[i])
	}
	return parsed
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
