package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Sell struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
  Item string `json:"item" yaml:"item"`
  Amount string `json:"amount" yaml:"amount"`
  Info string `json:"info" yaml:"info"`
}