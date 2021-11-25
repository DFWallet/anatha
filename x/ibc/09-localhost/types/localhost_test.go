package types_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/DFWallet/anatha/codec"
	"github.com/DFWallet/anatha/simapp"
	sdk "github.com/DFWallet/anatha/types"
	clientexported "github.com/DFWallet/anatha/x/ibc/02-client/exported"
)

const (
	height = 4
)

type LocalhostTestSuite struct {
	suite.Suite

	aminoCdc *codec.Codec
	cdc      codec.Marshaler
	store    sdk.KVStore
}

func (suite *LocalhostTestSuite) SetupTest() {
	isCheckTx := false
	app := simapp.Setup(isCheckTx)

	suite.aminoCdc = app.Codec()
	suite.cdc = app.AppCodec()
	ctx := app.BaseApp.NewContext(isCheckTx, abci.Header{Height: 1})
	suite.store = app.IBCKeeper.ClientKeeper.ClientStore(ctx, clientexported.ClientTypeLocalHost)
}

func TestLocalhostTestSuite(t *testing.T) {
	suite.Run(t, new(LocalhostTestSuite))
}