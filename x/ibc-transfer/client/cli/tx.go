package cli

import (
	"bufio"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/DFWallet/anatha/client/context"
	"github.com/DFWallet/anatha/client/flags"
	"github.com/DFWallet/anatha/codec"
	sdk "github.com/DFWallet/anatha/types"
	authclient "github.com/DFWallet/anatha/x/auth/client"
	authtypes "github.com/DFWallet/anatha/x/auth/types"
	"github.com/DFWallet/anatha/x/ibc-transfer/types"
)

// IBC transfer flags
var (
	FlagNode1    = "node1"
	FlagNode2    = "node2"
	FlagFrom1    = "from1"
	FlagFrom2    = "from2"
	FlagChainID2 = "chain-id2"
	FlagSequence = "packet-sequence"
	FlagTimeout  = "timeout"
)

// GetTransferTxCmd returns the command to create a NewMsgTransfer transaction
func GetTransferTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer [src-port] [src-channel] [dest-height] [receiver] [amount]",
		Short: "Transfer fungible token through IBC",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := authtypes.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc).WithBroadcastMode(flags.BroadcastBlock)

			sender := cliCtx.GetFromAddress()
			srcPort := args[0]
			srcChannel := args[1]
			destHeight, err := strconv.Atoi(args[2])
			if err != nil {
				return err
			}

			// parse coin trying to be sent
			coins, err := sdk.ParseCoins(args[4])
			if err != nil {
				return err
			}

			msg := types.NewMsgTransfer(srcPort, srcChannel, uint64(destHeight), coins, sender, args[3])
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	return cmd
}
