package cosmos

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/leejw51/cosmos/x/cosmos/types"
	"github.com/leejw51/cosmos/x/cosmos/keeper"
)

func handleMsgCreateComment(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateComment) (*sdk.Result, error) {
	var comment = types.Comment{
		Creator: msg.Creator,
		ID:      msg.ID,
    Title: msg.Title,
    Body: msg.Body,
	}
	k.CreateComment(ctx, comment)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
