package ibc_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	abci "github.com/tendermint/tendermint/abci/types"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/DFWallet/anatha/codec"
	"github.com/DFWallet/anatha/simapp"
	sdk "github.com/DFWallet/anatha/types"
	channeltypes "github.com/DFWallet/anatha/x/ibc/04-channel/types"
	ibctmtypes "github.com/DFWallet/anatha/x/ibc/07-tendermint/types"
)

const (
	connectionID  = "connectionidone"
	clientID      = "clientidone"
	connectionID2 = "connectionidtwo"
	clientID2     = "clientidtwo"

	port1 = "firstport"
	port2 = "secondport"

	channel1 = "firstchannel"
	channel2 = "secondchannel"

	channelOrder   = channeltypes.ORDERED
	channelVersion = "1.0"

	trustingPeriod time.Duration = time.Hour * 24 * 7 * 2
	ubdPeriod      time.Duration = time.Hour * 24 * 7 * 3
	maxClockDrift  time.Duration = time.Second * 10
)

type IBCTestSuite struct {
	suite.Suite

	cdc    *codec.Codec
	ctx    sdk.Context
	app    *simapp.SimApp
	header ibctmtypes.Header
}

func (suite *IBCTestSuite) SetupTest() {
	isCheckTx := false
	suite.app = simapp.Setup(isCheckTx)

	privVal := tmtypes.NewMockPV()
	pubKey, err := privVal.GetPubKey()
	suite.Require().NoError(err)

	now := time.Now().UTC()

	val := tmtypes.NewValidator(pubKey, 10)
	valSet := tmtypes.NewValidatorSet([]*tmtypes.Validator{val})

	suite.header = ibctmtypes.CreateTestHeader("chainID", 10, now, valSet, []tmtypes.PrivValidator{privVal})

	suite.cdc = suite.app.Codec()
	suite.ctx = suite.app.BaseApp.NewContext(isCheckTx, abci.Header{})
}

func TestIBCTestSuite(t *testing.T) {
	suite.Run(t, new(IBCTestSuite))
}