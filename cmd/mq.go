package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// 消息队列消费端
var mqConsumerCommand = &cobra.Command{
	Use:     "mq-consumer",
	Short:   "消息队列 消费者",
	Example: "web33 mq-consumer",
	Run:     runConsumer,
}

// 消息队列生产者
var mqProducerCommand = &cobra.Command{
	Use:     "mq-producer",
	Short:   "消息队列 生产者",
	Example: "web33 mq-producer",
	Run:     runProducer,
}

func runConsumer(_ *cobra.Command, _ []string) {
	fmt.Println("消息队列消费端")
	// var (
	// 	ctx     = context.Background()
	// 	cfg     = RegisterConfig()
	// 	service = boot.New(ctx, cfg)
	// 	// users   = service.Container()
	// 	logger = service.Logger()
	// )

	// logger.Info().Msg("info")
	// service.Ioc()
}

func runProducer(_ *cobra.Command, _ []string) {
	fmt.Println("消息队列生产者")
	// var (
	// 	ctx     = context.Background()
	// 	cfg     = RegisterConfig()
	// 	service = boot.New(ctx, cfg)
	// 	// users   = service.Container()
	// 	logger = service.Logger()
	// )

	// logger.Info().Msg("info")
	// service.Ioc()
}
