package cmd

import (
	"fmt"
	"go-cli/internal/config"
	"os"

	"github.com/spf13/cobra"
)

var (
	Commit     = "none"
	Version    = "1.0.0"
	configPath = "config.yml"
)

// 根命令，开发使用：go run main.go [命令]
// 例如启动api服务：go run main.go api-server
var rootCmd = &cobra.Command{
	Use:   "go-cli",
	Short: "go-cli 程序",
	Long:  "go-cli 程序",
}

func Execute() {

	InitConfig()
	rootCmd.AddCommand(mqConsumerCommand)
	rootCmd.AddCommand(mqProducerCommand)
	rootCmd.AddCommand(apiCommand)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// 初始化配置
func InitConfig() *config.Config {
	cfg, err := config.New(Commit, Version, configPath)
	if err != nil {
		fmt.Printf("unable to initialize config: %s\n", err.Error())
		os.Exit(1)
	}

	return cfg
}
