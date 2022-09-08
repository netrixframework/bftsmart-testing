package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/netrixframework/bftsmart-testing/util"
	"github.com/netrixframework/netrix/config"
	"github.com/netrixframework/netrix/testlib"
	"github.com/spf13/cobra"
)

var unittestCmd = &cobra.Command{
	Use: "unit",
	RunE: func(cmd *cobra.Command, args []string) error {
		termCh := make(chan os.Signal, 1)
		signal.Notify(termCh, os.Interrupt, syscall.SIGTERM)

		server, err := testlib.NewTestingServer(
			&config.Config{
				APIServerAddr: "127.0.0.1:7074",
				NumReplicas:   4,
				LogConfig: config.LogConfig{
					Format: "json",
					Path:   "/tmp/bftsmart/log/checker.log",
				},
			},
			&util.BFTSmartParser{},
			[]*testlib.TestCase{},
		)
		if err != nil {
			return err
		}
		go func() {
			<-termCh
			server.Stop()
		}()
		server.Start()
		return nil
	},
}
