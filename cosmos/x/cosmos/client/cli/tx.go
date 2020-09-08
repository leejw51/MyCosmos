package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/leejw51/cosmos/x/cosmos/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	cosmosTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cosmosTxCmd.AddCommand(flags.PostCommands(
    // this line is used by starport scaffolding
		GetCmdCreateSell(cdc),
		GetCmdCreateBuy(cdc),
		GetCmdCreateAsset(cdc),
		GetCmdCreateComment(cdc),
		GetCmdCreatePost(cdc),
	)...)

	return cosmosTxCmd
}
