package types

import (
	"github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _, _, _, _, _, _ sdk.Msg = &MsgDeployNftRequest{}, &MsgMintNftRequest{}, &MsgBurnNftRequest{}, &MsgMintNftRequest{},
	&MsgBurnNftRequest{}, &MsgAccessNftRequest{}
var _ sdk.Msg = &MsgRentGiveAccessRequest{}

// GetSigners implements types.Msg
func (m *MsgAccessNftRequest) GetSigners() []types.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Renter)
	return []sdk.AccAddress{addr}
}

// ValidateBasic implements types.Msg
func (m *MsgAccessNftRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Renter); err != nil {
		return sdkerrors.Wrap(err, "invalid authority address")
	}
	return nil
}

// GetSigners implements types.Msg
func (m *MsgRentGiveAccessRequest) GetSigners() []types.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Renter)
	return []sdk.AccAddress{addr}
}

// ValidateBasic implements types.Msg
func (m *MsgRentGiveAccessRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Renter); err != nil {
		return sdkerrors.Wrap(err, "invalid authority address")
	}
	return nil
}

// GetSigners implements types.Msg
func (m *MsgMintRentRequest) GetSigners() []types.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.ContractOwner)
	return []sdk.AccAddress{addr}
}

// ValidateBasic implements types.Msg
func (m *MsgMintRentRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Renter); err != nil {
		return sdkerrors.Wrap(err, "invalid authority address")
	}
	return nil
}

// GetSigners implements types.Msg
func (m *MsgBurnRentRequest) GetSigners() []types.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.ContractOwner)
	return []sdk.AccAddress{addr}
}

// ValidateBasic implements types.Msg
func (m *MsgBurnRentRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.ContractOwner); err != nil {
		return sdkerrors.Wrap(err, "invalid authority address")
	}
	return nil
}

// GetSigners implements types.Msg
func (m *MsgDeployNftRequest) GetSigners() []types.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.ContractOwner)
	return []sdk.AccAddress{addr}
}

// ValidateBasic implements types.Msg
func (m *MsgDeployNftRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.ContractOwner); err != nil {
		return sdkerrors.Wrap(err, "invalid authority address")
	}
	return nil
}

// GetSigners implements types.Msg
func (m *MsgMintNftRequest) GetSigners() []types.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.ContractOwner)
	return []sdk.AccAddress{addr}
}

// ValidateBasic implements types.Msg
func (m *MsgMintNftRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.ContractOwner); err != nil {
		return sdkerrors.Wrap(err, "invalid authority address")
	}
	return nil
}

// GetSigners implements types.Msg
func (m *MsgBurnNftRequest) GetSigners() []types.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.ContractOwner)
	return []sdk.AccAddress{addr}
}

// ValidateBasic implements types.Msg
func (m *MsgBurnNftRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.ContractOwner); err != nil {
		return sdkerrors.Wrap(err, "invalid authority address")
	}
	return nil
}
