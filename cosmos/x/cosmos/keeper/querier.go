package keeper

import (
  // this line is used by starport scaffolding
	"github.com/leejw51/cosmos/x/cosmos/types"
		
	
		
	
		
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for cosmos clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
    // this line is used by starport scaffolding # 2
		case types.QueryListAsset:
			return listAsset(ctx, k)
		case types.QueryListComment:
			return listComment(ctx, k)
		case types.QueryListPost:
			return listPost(ctx, k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown cosmos query endpoint")
		}
	}
}