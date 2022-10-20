package cmd

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/netrixframework/bftsmart-testing/tests"
	"github.com/netrixframework/bftsmart-testing/util"
	"github.com/netrixframework/netrix/config"
	"github.com/netrixframework/netrix/log"
	"github.com/netrixframework/netrix/strategies"
	"github.com/netrixframework/netrix/strategies/unittest"
	"github.com/netrixframework/netrix/types"
	"github.com/spf13/cobra"
)

func logStepFunc(e *types.Event, ctx *strategies.Context) {
	if !e.IsMessageReceive() {
		return
	}
	message, ok := ctx.GetMessage(e)
	if !ok || message.ParsedMessage == nil {
		return
	}
	bftMessage, ok := message.ParsedMessage.(*util.BFTSmartMessage)
	if !ok {
		return
	}
	ctx.Logger.With(log.LogParams{
		"message": bftMessage.String(),
		"from":    message.From,
		"to":      message.To,
	}).Info("Message received")

}

var stratTestCmd = &cobra.Command{
	Use: "test",
	RunE: func(cmd *cobra.Command, args []string) error {
		termCh := make(chan os.Signal, 1)
		signal.Notify(termCh, os.Interrupt, syscall.SIGTERM)

		var strategy strategies.Strategy = unittest.NewTestCaseStrategy(tests.DummyTest())

		server := strategies.NewStrategyDriver(
			&config.Config{
				APIServerAddr: "127.0.0.1:7074",
				NumReplicas:   4,
				LogConfig: config.LogConfig{
					Format: "json",
					Path:   "/Users/srinidhin/Local/data/testing/bftsmart/t/checker.log",
				},
			},
			&util.BFTSmartParser{},
			strategy,
			&strategies.StrategyConfig{
				Iterations:       10,
				IterationTimeout: 30 * time.Second,
				StepFunc:         logStepFunc,
			},
		)

		go func() {
			<-termCh
			server.Stop()
		}()
		server.Start()
		return nil
	},
}
