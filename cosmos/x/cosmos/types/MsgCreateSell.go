package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateSell{}

type MsgCreateSell struct {
  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Item string `json:"item" yaml:"item"`
  Amount string `json:"amount" yaml:"amount"`
  Info string `json:"info" yaml:"info"`
}

func NewMsgCreateSell(creator sdk.AccAddress, item string, amount string, info string) MsgCreateSell {
  return MsgCreateSell{
    ID: uuid.New().String(),
		Creator: creator,
    Item: item,
    Amount: amount,
    Info: info,
	}
}

func (msg MsgCreateSell) Route() string {
  return RouterKey
}

func (msg MsgCreateSell) Type() string {
  return "CreateSell"
}

func (msg MsgCreateSell) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateSell) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateSell) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}