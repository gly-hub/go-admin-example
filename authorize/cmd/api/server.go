package api

import (
	"fmt"
	"go-admin-example/authorize/boot"
	"go-admin-example/authorize/internal/service"
	"github.com/gly-hub/go-dandelion/application"
	"github.com/gly-hub/go-dandelion/config"
	"github.com/gly-hub/go-dandelion/logger"
	"github.com/gly-hub/toolbox/stringx"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/signal"
)

var (
	env      string
	StartCmd = &cobra.Command{
		Use:          "server",
		Short:        "Start RPC server",
		Example:      "authorize server -e local",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&env, "env", "e", "local", "Env")
}

func setup() {
	// 配置初始化
	config.InitConfig(env)
	// 应用初始化
	application.Init()
	// 初始化服务方法
	boot.Init()
}

func run() error {
	// 初始化rpc model
	go func() {
		application.RpcServer(new(service.RpcApi))
	}()
	content, _ := ioutil.ReadFile("./static/authorize.txt")
	fmt.Println(logger.Green(string(content)))
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Printf("%s Shutdown Server ... \r\n", stringx.GetCurrentTimeStr())
    logger.Info("Server exiting")
    return nil
}
