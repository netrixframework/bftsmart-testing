package cmd

import (
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/netrixframework/bftsmart-testing/client"
	"github.com/netrixframework/bftsmart-testing/tests"
	"github.com/netrixframework/bftsmart-testing/util"
	"github.com/netrixframework/netrix/config"
	"github.com/netrixframework/netrix/strategies"
	"github.com/netrixframework/netrix/strategies/pct"
	"github.com/spf13/cobra"
)

var pctTestStrategy = &cobra.Command{
	Use: "pct-test",
	RunE: func(cmd *cobra.Command, args []string) error {
		termCh := make(chan os.Signal, 1)
		signal.Notify(termCh, os.Interrupt, syscall.SIGTERM)

		var strategy strategies.Strategy = pct.NewPCTStrategyWithTestCase(
			&pct.PCTStrategyConfig{
				RandSrc:        rand.NewSource(time.Now().UnixMilli()),
				MaxEvents:      1000,
				Depth:          6,
				RecordFilePath: "/home/nagendra/data/testing/bftsmart/t",
			},
			tests.DelayProposeForP(),
		)

		strategy = strategies.NewStrategyWithProperty(strategy, tests.DelayProposeForPProperty())

		bftSmartClient := client.NewBFTSmartClient(&client.BFTSmartClientConfig{
			CodePath: "/home/nagendra/code/bftsmart",
		})

		driver := strategies.NewStrategyDriver(
			&config.Config{
				APIServerAddr: "127.0.0.1:7074",
				NumReplicas:   4,
				LogConfig: config.LogConfig{
					Format: "json",
					Level:  "info",
					Path:   "/home/nagendra/data/testing/bftsmart/t/checker.log",
				},
			},
			&util.BFTSmartParser{},
			strategy,
			&strategies.StrategyConfig{
				Iterations:       100,
				IterationTimeout: 30 * time.Second,
				SetupFunc: func(ctx *strategies.Context) {
					go bftSmartClient.Set("name", "srinidhi")
				},
			},
		)

		go func() {
			<-termCh
			driver.Stop()
		}()

		if err := driver.Start(); err != nil {
			panic(err)
		}
		return nil
	},
}
