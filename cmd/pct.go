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

var pctStrategy = &cobra.Command{
	Use: "pct",
	RunE: func(cmd *cobra.Command, args []string) error {
		termCh := make(chan os.Signal, 1)
		signal.Notify(termCh, os.Interrupt, syscall.SIGTERM)

		var strategy strategies.Strategy = pct.NewPCTStrategy(&pct.PCTStrategyConfig{
			RandSrc:        rand.NewSource(time.Now().UnixMilli()),
			MaxEvents:      1000,
			Depth:          10,
			RecordFilePath: "/home/nagendra/data/testing/bftsmart/t",
		})

		bftSmartClient := client.NewBFTSmartClient(&client.BFTSmartClientConfig{
			CodePath: "/Users/srinidhin/Local/github/bft-smart",
		})

		strategy = strategies.NewStrategyWithProperty(strategy, tests.DelayProposeForPProperty())

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
				Iterations:       1000,
				IterationTimeout: 40 * time.Second,
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
