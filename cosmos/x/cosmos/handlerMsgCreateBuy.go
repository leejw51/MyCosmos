package cosmos

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/leejw51/cosmos/x/cosmos/types"
	"github.com/leejw51/cosmos/x/cosmos/keeper"
)

func handleMsgCreateBuy(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateBuy) (*sdk.Result, error) {
	var buy = types.Buy{
		Creator: msg.Creator,
		ID:      msg.ID,
    Item: msg.Item,
    Amount: msg.Amount,
    Info: msg.Info,
	}
	k.CreateBuy(ctx, buy)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
