package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	// "github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/leejw51/cosmos/x/cosmos/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group cosmos queries under a subcommand
	cosmosQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cosmosQueryCmd.AddCommand(
		flags.GetCommands(
      // this line is used by starport scaffolding
			GetCmdListSell(queryRoute, cdc),
			GetCmdListBuy(queryRoute, cdc),
			GetCmdListAsset(queryRoute, cdc),
			GetCmdListComment(queryRoute, cdc),
			GetCmdListPost(queryRoute, cdc),
		)...,
	)

	return cosmosQueryCmd
}
