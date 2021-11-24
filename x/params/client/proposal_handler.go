package client

import (
	govclient "github.com/DFWallet/anatha/x/gov/client"
	"github.com/DFWallet/anatha/x/params/client/cli"
	"github.com/DFWallet/anatha/x/params/client/rest"
)

// ProposalHandler handles param change proposals
var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
