package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var File string

var rootCmd = &cobra.Command{
	Use:           "zabbix-tree-cli",
	Short:         "Render graphical output for certains part of Zabbix server.",
	Long:          `This CLI tool is used to help administrator keeps track of their Zabbix Host Groups structure by rendering a graphical output (PNG, JPG, SVG, json, shell).`,
	Args:          cobra.NoArgs,
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	rootCmd.AddCommand(hostGroupCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
