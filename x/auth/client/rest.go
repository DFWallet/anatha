package client

import (
	"log"
	"net/http"

	"github.com/DFWallet/anatha/client/context"
	"github.com/DFWallet/anatha/client/flags"
	sdk "github.com/DFWallet/anatha/types"
	"github.com/DFWallet/anatha/types/errors"
	"github.com/DFWallet/anatha/types/rest"
	"github.com/DFWallet/anatha/x/auth/types"
)

// WriteGenerateStdTxResponse writes response for the generate only mode.
func WriteGenerateStdTxResponse(w http.ResponseWriter, cliCtx context.CLIContext, br rest.BaseReq, msgs []sdk.Msg) {
	gasAdj, ok := rest.ParseFloat64OrReturnBadRequest(w, br.GasAdjustment, flags.DefaultGasAdjustment)
	if !ok {
		return
	}

	simAndExec, gas, err := flags.ParseGas(br.Gas)
	if rest.CheckBadRequestError(w, err) {
		return
	}

	txBldr := types.NewTxBuilder(
		GetTxEncoder(cliCtx.Codec), br.AccountNumber, br.Sequence, gas, gasAdj,
		br.Simulate, br.ChainID, br.Memo, br.Fees, br.GasPrices,
	)

	if br.Simulate || simAndExec {
		if gasAdj < 0 {
			rest.WriteErrorResponse(w, http.StatusBadRequest, errors.ErrorInvalidGasAdjustment.Error())
			return
		}

		txBldr, err = EnrichWithGas(txBldr, cliCtx, msgs)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		if br.Simulate {
			rest.WriteSimulationResponse(w, cliCtx.Codec, txBldr.Gas())
			return
		}
	}

	stdMsg, err := txBldr.BuildSignMsg(msgs)
	if rest.CheckBadRequestError(w, err) {
		return
	}

	output, err := cliCtx.Codec.MarshalJSON(types.NewStdTx(stdMsg.Msgs, stdMsg.Fee, nil, stdMsg.Memo))
	if rest.CheckInternalServerError(w, err) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(output); err != nil {
		log.Printf("could not write response: %v", err)
	}

}
