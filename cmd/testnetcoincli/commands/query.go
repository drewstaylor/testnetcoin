package commands

import (
	"github.com/spf13/cobra"

	proofcmd "github.com/tendermint/light-client/commands/proofs"

	"github.com/tendermint/basecoin/docs/guide/counter/plugins/counter"
)

//CounterQueryCmd CLI command to query the counter state
var CounterQueryCmd = &cobra.Command{
	Use:   "testnetcoin",
	Short: "Query counter state, with proof",
	RunE:  counterQueryCmd,
}

func counterQueryCmd(cmd *cobra.Command, args []string) error {
	key := testnetcoin.New().StateKey()

	var cp testnetcoin.CounterPluginState
	proof, err := proofcmd.GetAndParseAppProof(key, &cp)
	if err != nil {
		return err
	}

	return proofcmd.OutputProof(cp, proof.BlockHeight())
}
