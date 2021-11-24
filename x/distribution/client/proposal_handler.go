package client

import (
	"github.com/DFWallet/anatha/x/distribution/client/cli"
	"github.com/DFWallet/anatha/x/distribution/client/rest"
	govclient "github.com/DFWallet/anatha/x/gov/client"
)

// param change proposal handler
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)
