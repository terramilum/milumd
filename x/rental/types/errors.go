package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/rental module sentinel errors
var (
	ErrSample                  = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrNftOwnerCanRent         = sdkerrors.Register(ModuleName, 1101, "Only Nft owner can mint rent")
	ErrNftRentError            = sdkerrors.Register(ModuleName, 1102, "Renting error")
	ErrNftRentMinStartDate     = sdkerrors.Register(ModuleName, 1103, "Minimum start date")
	ErrNftRentNotAvaliableDate = sdkerrors.Register(ModuleName, 1104, "This period is not avaliable for rent.")
	ErrNftRentAccessGive       = sdkerrors.Register(ModuleName, 1105, "Renter has not right for this session")
	ErrNftClassOwnerTheSame    = sdkerrors.Register(ModuleName, 1106, "Class owner should be the same to mint nft.")
)
