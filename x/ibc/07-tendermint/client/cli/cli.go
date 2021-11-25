package cli

import (
	"github.com/spf13/cobra"

	"github.com/DFWallet/anatha/client/flags"
	"github.com/DFWallet/anatha/codec"
)

// GetTxCmd returns the transaction commands for IBC clients
func GetTxCmd(cdc *codec.Codec, storeKey string) *cobra.Command {
	ics07TendermintTxCmd := &cobra.Command{
		Use:                        "tendermint",
		Short:                      "Tendermint transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
	}

	ics07TendermintTxCmd.AddCommand(flags.PostCommands(
		GetCmdCreateClient(cdc),
		GetCmdUpdateClient(cdc),
		GetCmdSubmitMisbehaviour(cdc),
	)...)

	return ics07TendermintTxCmd
}
