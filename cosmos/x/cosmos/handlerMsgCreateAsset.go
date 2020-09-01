package cosmos

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/leejw51/cosmos/x/cosmos/types"
	"github.com/leejw51/cosmos/x/cosmos/keeper"
)

func handleMsgCreateAsset(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateAsset) (*sdk.Result, error) {
	var asset = types.Asset{
		Creator: msg.Creator,
		ID:      msg.ID,
    Title: msg.Title,
    Body: msg.Body,
	}
	k.CreateAsset(ctx, asset)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
