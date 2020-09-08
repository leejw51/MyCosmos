package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/leejw51/cosmos/x/cosmos/types"
  "github.com/cosmos/cosmos-sdk/codec"
)

func (k Keeper) CreateSell(ctx sdk.Context, sell types.Sell) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.SellPrefix + sell.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(sell)
	store.Set(key, value)
}

func listSell(ctx sdk.Context, k Keeper) ([]byte, error) {
  var sellList []types.Sell
  store := ctx.KVStore(k.storeKey)
  iterator := sdk.KVStorePrefixIterator(store, []byte(types.SellPrefix))
  for ; iterator.Valid(); iterator.Next() {
    var sell types.Sell
    k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &sell)
    sellList = append(sellList, sell)
  }
  res := codec.MustMarshalJSONIndent(k.cdc, sellList)
  return res, nil
}