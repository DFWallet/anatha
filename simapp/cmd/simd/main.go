package main

import (
	"encoding/json"
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/DFWallet/anatha/baseapp"
	"github.com/DFWallet/anatha/client/debug"
	"github.com/DFWallet/anatha/client/flags"
	"github.com/DFWallet/anatha/server"
	"github.com/DFWallet/anatha/simapp"
	"github.com/DFWallet/anatha/store"
	sdk "github.com/DFWallet/anatha/types"
	"github.com/DFWallet/anatha/x/bank"
	genutilcli "github.com/DFWallet/anatha/x/genutil/client/cli"
	"github.com/DFWallet/anatha/x/staking"
)

const flagInvCheckPeriod = "inv-check-period"

var invCheckPeriod uint

func main() {
	appCodec, cdc := simapp.MakeCodecs()

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(sdk.Bech32PrefixAccAddr, sdk.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(sdk.Bech32PrefixValAddr, sdk.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(sdk.Bech32PrefixConsAddr, sdk.Bech32PrefixConsPub)
	config.Seal()

	ctx := server.NewDefaultContext()
	cobra.EnableCommandSorting = false
	rootCmd := &cobra.Command{
		Use:               "simd",
		Short:             "Simulation Daemon (server)",
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}

	rootCmd.AddCommand(
		genutilcli.InitCmd(ctx, cdc, simapp.ModuleBasics, simapp.DefaultNodeHome),
		genutilcli.CollectGenTxsCmd(ctx, cdc, bank.GenesisBalancesIterator{}, simapp.DefaultNodeHome),
		genutilcli.MigrateGenesisCmd(ctx, cdc),
		genutilcli.GenTxCmd(
			ctx, cdc, simapp.ModuleBasics, staking.AppModuleBasic{},
			bank.GenesisBalancesIterator{}, simapp.DefaultNodeHome, simapp.DefaultCLIHome,
		),
		genutilcli.ValidateGenesisCmd(ctx, cdc, simapp.ModuleBasics),
		AddGenesisAccountCmd(ctx, cdc, appCodec, simapp.DefaultNodeHome, simapp.DefaultCLIHome),
		flags.NewCompletionCmd(rootCmd, true),
		debug.Cmd(cdc))

	server.AddCommands(ctx, cdc, rootCmd, newApp, exportAppStateAndTMValidators)

	// prepare and add flags
	executor := cli.PrepareBaseCmd(rootCmd, "GA", simapp.DefaultNodeHome)
	rootCmd.PersistentFlags().UintVar(&invCheckPeriod, flagInvCheckPeriod,
		0, "Assert registered invariants every N blocks")
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func newApp(logger log.Logger, db dbm.DB, traceStore io.Writer) abci.Application {
	var cache sdk.MultiStorePersistentCache

	if viper.GetBool(server.FlagInterBlockCache) {
		cache = store.NewCommitKVStoreCacheManager()
	}

	skipUpgradeHeights := make(map[int64]bool)
	for _, h := range viper.GetIntSlice(server.FlagUnsafeSkipUpgrades) {
		skipUpgradeHeights[int64(h)] = true
	}

	return simapp.NewSimApp(
		logger, db, traceStore, true, skipUpgradeHeights,
		viper.GetString(flags.FlagHome), invCheckPeriod,
		baseapp.SetPruning(store.NewPruningOptionsFromString(viper.GetString("pruning"))),
		baseapp.SetMinGasPrices(viper.GetString(server.FlagMinGasPrices)),
		baseapp.SetHaltHeight(viper.GetUint64(server.FlagHaltHeight)),
		baseapp.SetHaltTime(viper.GetUint64(server.FlagHaltTime)),
		baseapp.SetInterBlockCache(cache),
	)
}

func exportAppStateAndTMValidators(
	logger log.Logger, db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool, jailWhiteList []string,
) (json.RawMessage, []tmtypes.GenesisValidator, *abci.ConsensusParams, error) {

	var simApp *simapp.SimApp
	if height != -1 {
		simApp = simapp.NewSimApp(logger, db, traceStore, false, map[int64]bool{}, "", uint(1))
		err := simApp.LoadHeight(height)
		if err != nil {
			return nil, nil, nil, err
		}
	} else {
		simApp = simapp.NewSimApp(logger, db, traceStore, true, map[int64]bool{}, "", uint(1))
	}
	return simApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
}
