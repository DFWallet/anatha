package cli

import (
	"github.com/spf13/cobra"

	"github.com/DFWallet/anatha/client/flags"
	"github.com/DFWallet/anatha/codec"
	"github.com/DFWallet/anatha/x/ibc/09-localhost/types"
)

// GetTxCmd returns the transaction commands for IBC
func GetTxCmd(cdc *codec.Codec, storeKey string) *cobra.Command {
	ics09LocalhostTxCmd := &cobra.Command{
		Use:                        types.SubModuleName,
		Short:                      "Localhost transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
	}

	ics09LocalhostTxCmd.AddCommand(flags.PostCommands(
		GetCmdCreateClient(cdc),
	)...)

	return ics09LocalhostTxCmd
}
