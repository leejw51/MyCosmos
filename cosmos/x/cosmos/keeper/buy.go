package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/leejw51/cosmos/x/cosmos/types"
  "github.com/cosmos/cosmos-sdk/codec"
)

func (k Keeper) CreateBuy(ctx sdk.Context, buy types.Buy) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.BuyPrefix + buy.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(buy)
	store.Set(key, value)
}

func listBuy(ctx sdk.Context, k Keeper) ([]byte, error) {
  var buyList []types.Buy
  store := ctx.KVStore(k.storeKey)
  iterator := sdk.KVStorePrefixIterator(store, []byte(types.BuyPrefix))
  for ; iterator.Valid(); iterator.Next() {
    var buy types.Buy
    k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &buy)
    buyList = append(buyList, buy)
  }
  res := codec.MustMarshalJSONIndent(k.cdc, buyList)
  return res, nil
}