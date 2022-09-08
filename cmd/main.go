package cmd

import "github.com/spf13/cobra"

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "bftsmart-testing",
	}
	cmd.CompletionOptions.DisableDefaultCmd = true
	cmd.AddCommand(unittestCmd)
	return cmd
}
