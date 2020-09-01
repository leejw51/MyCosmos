package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/leejw51/cosmos/x/cosmos/types"
  "github.com/cosmos/cosmos-sdk/codec"
)

func (k Keeper) CreateAsset(ctx sdk.Context, asset types.Asset) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.AssetPrefix + asset.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(asset)
	store.Set(key, value)
}

func listAsset(ctx sdk.Context, k Keeper) ([]byte, error) {
  var assetList []types.Asset
  store := ctx.KVStore(k.storeKey)
  iterator := sdk.KVStorePrefixIterator(store, []byte(types.AssetPrefix))
  for ; iterator.Valid(); iterator.Next() {
    var asset types.Asset
    k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &asset)
    assetList = append(assetList, asset)
  }
  res := codec.MustMarshalJSONIndent(k.cdc, assetList)
  return res, nil
}