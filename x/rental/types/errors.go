package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/rental module sentinel errors
var (
	ErrFieldIsRequired             = sdkerrors.Register(ModuleName, 1100, "Field is required.")
	ErrNftOwnerCanRent             = sdkerrors.Register(ModuleName, 1101, "Only Nft owner can mint rent")
	ErrNftRentError                = sdkerrors.Register(ModuleName, 1102, "Renting error")
	ErrNftRentMinStartDate         = sdkerrors.Register(ModuleName, 1103, "Minimum start date")
	ErrNftRentNotAvaliableDate     = sdkerrors.Register(ModuleName, 1104, "This period is not avaliable for rent.")
	ErrNftRentAccessGive           = sdkerrors.Register(ModuleName, 1105, "Renter has not right for this session")
	ErrNftClassOwnerTheSame        = sdkerrors.Register(ModuleName, 1106, "Class owner should be the same to mint nft.")
	ErrStartDateBiggerEndDate      = sdkerrors.Register(ModuleName, 1107, "Start date cannot be bigger then end date.")
	ErrQuerySessions               = sdkerrors.Register(ModuleName, 1108, "Renter or Class id must be filled.")
	ErrQuerySessionsNotFound       = sdkerrors.Register(ModuleName, 1109, "Session not found for this renter.")
	ErrQueryOldSessionsNotTransfer = sdkerrors.Register(ModuleName, 1110, "Old session cannot be transfered.")
	ErrSessionOwnerCanTransfer     = sdkerrors.Register(ModuleName, 1111, "Only session owner can transfer.")
	ErrHasNoAccessCurrently        = sdkerrors.Register(ModuleName, 1112, "There has not granted for this nft currently.")
)
