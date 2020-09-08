package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateBuy{}

type MsgCreateBuy struct {
  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Item string `json:"item" yaml:"item"`
  Amount string `json:"amount" yaml:"amount"`
  Info string `json:"info" yaml:"info"`
}

func NewMsgCreateBuy(creator sdk.AccAddress, item string, amount string, info string) MsgCreateBuy {
  return MsgCreateBuy{
    ID: uuid.New().String(),
		Creator: creator,
    Item: item,
    Amount: amount,
    Info: info,
	}
}

func (msg MsgCreateBuy) Route() string {
  return RouterKey
}

func (msg MsgCreateBuy) Type() string {
  return "CreateBuy"
}

func (msg MsgCreateBuy) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateBuy) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateBuy) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}