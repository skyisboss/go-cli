package cmd

import (
	"context"
	"go-cli/internal/boot"
	"go-cli/pkg/graceful"
	"log"

	"github.com/spf13/cobra"
)

// api服务
var apiCommand = &cobra.Command{
	Use:     "web-server",
	Short:   "提供api服务接口",
	Example: "web33 api-server",
	Run:     runWeb,
}

func runWeb(_ *cobra.Command, _ []string) {
	var (
		ctx  = context.Background()
		cfg  = InitConfig()
		boot = boot.New(ctx, cfg)
	)
	go func() {
		boot.WebServer()
	}()
	graceful.AddCallback(func() error {
		log.Println("shutting down")
		return nil
	})
	if err := graceful.WaitShutdown(); err != nil {
		boot.Logger().Error().Err(err).Msg("unable to shutdown service gracefully")
		return
	}
}
