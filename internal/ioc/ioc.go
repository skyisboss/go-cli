package ioc

import (
	"context"
	"fmt"
	"go-cli/internal/config"
	"sync"

	"github.com/go-redis/redis"
	"github.com/rs/zerolog"
)

type Container struct {
	ctx    context.Context
	once   map[string]*sync.Once
	config *config.Config
	logger *zerolog.Logger

	// Database
	redisService *redis.Client

	// Services
	// walletService     *wallet.Service
}

func (c *Container) init(key string, f func()) {
	if c.once[key] == nil {
		c.once[key] = &sync.Once{}
	}

	c.once[key].Do(f)
}

func NewIoc(ctx context.Context, cfg *config.Config, logger *zerolog.Logger) *Container {
	return &Container{
		config: cfg,
		ctx:    ctx,
		logger: logger,
		once:   make(map[string]*sync.Once, 128),
	}
}

func (c *Container) Context() context.Context {
	return c.ctx
}

func (c *Container) Logger() *zerolog.Logger {
	return c.logger
}

func (c *Container) Config() *config.Config {
	return c.config
}

func (c *Container) RedisService() *redis.Client {
	c.init("service.redis", func() {
		cfg := c.config.Providers.Redis
		// Redis连接格式拼接
		addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
		rds := redis.NewClient(&redis.Options{
			Addr:        addr,
			Password:    cfg.Auth, // 密码
			DB:          cfg.DB,   // 数据库
			PoolSize:    cfg.Pool, // 连接池大小
			IdleTimeout: 300,      // 默认Idle超时时间
		})

		// 验证是否连接到redis服务端
		_, err := rds.Ping().Result()
		if err != nil {
			panic(err)
		}
		c.redisService = rds
	})
	return c.redisService
}
