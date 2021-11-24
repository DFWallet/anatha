package client

import (
	govclient "github.com/DFWallet/anatha/x/gov/client"
	"github.com/DFWallet/anatha/x/upgrade/client/cli"
	"github.com/DFWallet/anatha/x/upgrade/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitUpgradeProposal, rest.ProposalRESTHandler)
