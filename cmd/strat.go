package cmd

import "github.com/spf13/cobra"

func strategyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "strat",
	}
	cmd.AddCommand(stratTestCmd)
	return cmd
}
