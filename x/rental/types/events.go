package types

const (
	// WasmModuleEventType is stored with any contract TX that returns non empty EventAttributes
	ModuleEventType = "rental"

	EventTypeDeployNft     = "rental_deploy"
	EventTypeMintNft       = "rental_mint"
	EventTypeBurnNft       = "rental_burn"
	EventTypeRentNft       = "rental_mint"
	EventTypeRentSend      = "rental_send"
	EventTypeBurnRentNft   = "rental_burn"
	EventTypeAccessNft     = "rental_access"
	EventTypeAccessGiveNft = "rental_access_give"
)

const (
	AttributeReservedPrefix = "_"

	AttributeKeyClassId         = "class_id"
	AttributeKeySessionId       = "session_id"
	AttributeKeyNftId           = "nft_id"
	AttributeKeyNftUri          = "nft_uri"
	AttributeKeyNftReciever     = "nft_reciever"
	AttributeKeyNftRentReciever = "rent_reciever"
	AttributeKeyNftRentStart    = "rent_start"
	AttributeKeyNftRentEnd      = "rent_end"
	AttributeKeyNftRentAccess   = "has_access"
	AttributeKeyNftCurrentDate  = "current_date"
	AttributeKeyNftRenter       = "renter"
	AttributeKeyNftNewRenter    = "new_renter"
	AttributeKeyNftFromRenter   = "from_renter"
	AttributeKeyNftToRenter     = "to_renter"
)
