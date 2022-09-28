package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/rest"

	"websocket_demo/internal/config"
	"websocket_demo/internal/handler"
	"websocket_demo/internal/svc"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "websocket",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.api.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(apiCmd)
	apiCmd.Flags().StringVarP(&configFile, "conf", "f", "etc/websocket.yaml", "config file")
}

var configFile string

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := config.LoadCfg(configFile)

		server := rest.MustNewServer(c.RestConf)
		defer server.Stop()

		ctx := svc.NewServiceContext(*c)
		handler.RegisterHandlers(server, ctx)

		fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
		server.Start()
	},
}
