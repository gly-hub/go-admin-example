package api

import (
	"fmt"
	"github.com/gly-hub/toolbox/stringx"
	"github.com/spf13/cobra"
	"github.com/team-dandelion/go-dandelion/application"
	"github.com/team-dandelion/go-dandelion/config"
	"github.com/team-dandelion/go-dandelion/logger"
	"go-admin-example/authorize/boot"
	"go-admin-example/authorize/internal/service"
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
		application.RpcServer(service.NewRpcApi())
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
