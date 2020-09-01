package cosmos

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/leejw51/cosmos/x/cosmos/keeper"
	"github.com/leejw51/cosmos/x/cosmos/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
    // this line is used by starport scaffolding
		case types.MsgCreateAsset:
			return handleMsgCreateAsset(ctx, k, msg)
		case types.MsgCreateComment:
			return handleMsgCreateComment(ctx, k, msg)
		case types.MsgCreatePost:
			return handleMsgCreatePost(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
