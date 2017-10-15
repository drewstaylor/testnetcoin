package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/tendermint/tmlibs/cli"

	"github.com/tendermint/basecoin/cmd/basecoin/commands"
	"github.com/tendermint/basecoin/docs/guide/counter/plugins/counter"
	"github.com/tendermint/basecoin/types"
)

func main() {
	var RootCmd = &cobra.Command{
		Use:   "testnetcoin",
		Short: "Funnel coins between Ethereum testnets, and never stop the dev Q_Q",
	}

	RootCmd.AddCommand(
		commands.InitCmd,
		commands.StartCmd,
		commands.UnsafeResetAllCmd,
		commands.VersionCmd,
	)

	commands.RegisterStartPlugin("testnetcoin", func() types.Plugin { return testnetcoin.New() })
	cmd := cli.PrepareMainCmd(RootCmd, "CT", os.ExpandEnv("$HOME/.testnetcoin"))
	cmd.Execute()
}
