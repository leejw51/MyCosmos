package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateAsset{}

type MsgCreateAsset struct {
  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Title string `json:"title" yaml:"title"`
  Body string `json:"body" yaml:"body"`
}

func NewMsgCreateAsset(creator sdk.AccAddress, title string, body string) MsgCreateAsset {
  return MsgCreateAsset{
    ID: uuid.New().String(),
		Creator: creator,
    Title: title,
    Body: body,
	}
}

func (msg MsgCreateAsset) Route() string {
  return RouterKey
}

func (msg MsgCreateAsset) Type() string {
  return "CreateAsset"
}

func (msg MsgCreateAsset) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateAsset) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateAsset) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}