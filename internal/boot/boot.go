package boot

import (
	"context"
	"go-cli/internal/apps/web/router"
	"go-cli/internal/config"
	"go-cli/internal/ioc"
	"go-cli/internal/logs"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Boot struct {
	config *config.Config
	ctx    context.Context
	logger *zerolog.Logger
	ioc    *ioc.Container
	// beforeRun []BeforeRun
}

func New(ctx context.Context, cfg *config.Config) *Boot {
	hostname, _ := os.Hostname()
	logger := logs.New(cfg.Logger, "go-cli", cfg.GitVersion, cfg.Env, hostname)
	ioc := ioc.NewIoc(ctx, cfg, &logger)

	return &Boot{
		config: cfg,
		ctx:    ctx,
		ioc:    ioc,
		logger: &logger,
	}
}

func (b *Boot) Ioc() *ioc.Container {
	return b.ioc
}

func (b *Boot) Logger() *zerolog.Logger {
	return b.logger
}

// 后台服务
func (b *Boot) WebServer() {
	app := gin.New()
	app.Use(gin.Logger(), gin.Recovery())
	router.New(app, b.ioc)

	app.Run(":52033")
}
