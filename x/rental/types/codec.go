package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgDeployNftRequest{}, "rental.MsgDeployNftRequest", nil)
	cdc.RegisterConcrete(&MsgMintNftRequest{}, "rental.MsgMintNftRequest", nil)
	cdc.RegisterConcrete(&MsgBurnNftRequest{}, "rental.MsgBurnNftRequest", nil)
	cdc.RegisterConcrete(&MsgMintRentRequest{}, "rental.MsgMintRentRequest", nil)
	cdc.RegisterConcrete(&MsgBurnRentRequest{}, "rental.MsgBurnRentRequest", nil)
	cdc.RegisterConcrete(&MsgAccessNftRequest{}, "rental.MsgAccessNftRequest", nil)
	cdc.RegisterConcrete(&MsgSendSessionRequest{}, "rental.MsgSendSessionRequest", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeployNftRequest{},
		&MsgMintNftRequest{},
		&MsgBurnNftRequest{},
		&MsgMintRentRequest{},
		&MsgBurnRentRequest{},
		&MsgAccessNftRequest{},
		&MsgSendSessionRequest{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
