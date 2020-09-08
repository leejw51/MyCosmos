package cosmos

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/leejw51/cosmos/x/cosmos/types"
	"github.com/leejw51/cosmos/x/cosmos/keeper"
)

func handleMsgCreateSell(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateSell) (*sdk.Result, error) {
	var sell = types.Sell{
		Creator: msg.Creator,
		ID:      msg.ID,
    Item: msg.Item,
    Amount: msg.Amount,
    Info: msg.Info,
	}
	k.CreateSell(ctx, sell)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
