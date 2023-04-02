package cmd

import (
	"fmt"
	"os"

	"github.com/Spartan0nix/zabbix-tree-cli/internal/logging"
	"github.com/spf13/cobra"
)

var File string
var Color bool
var Debug bool
var GlobalLogger *logging.Logger

var rootCmd = &cobra.Command{
	Use:           "zabbix-tree-cli",
	Short:         "Render graphical output for certains part of Zabbix server.",
	Long:          "This CLI tool is used to help administrator keeps track of their Zabbix structure by rendering a graphical output (dot, json, shell).",
	Args:          cobra.NoArgs,
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&File, "file", "f", "", "output to a file")
	rootCmd.PersistentFlags().BoolVar(&Color, "color", false, "enable colors in graph output")
	rootCmd.PersistentFlags().BoolVar(&Debug, "debug", false, "enable debug output during execution")
	rootCmd.AddCommand(hostGroupCmd)
	rootCmd.AddCommand(serviceCmd)
	// Init the global logger
	GlobalLogger = logging.NewLogger(logging.Warning)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		GlobalLogger.Error("error when executing root command", fmt.Sprintf("reason : %v", err))
		os.Exit(1)
	}
}
