package cmd

import (
	"fmt"
	"os"

	"github.com/Spartan0nix/zabbix-tree-cli/internal/app"
	"github.com/Spartan0nix/zabbix-tree-cli/internal/config"
	"github.com/spf13/cobra"
)

// This command allow to interact with Zabbix HostGroups
var hostGroupCmd = &cobra.Command{
	Use:   "host-group",
	Short: "TODO",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		err := config.CheckFileFlag(Format, File)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		env, err := config.GetEnvironmentVariables()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		app.RunHostGroup(env, Format, File)
	},
}

func init() {
	hostGroupCmd.Flags().StringVarP(&Format, "format", "o", "", "output format")
	hostGroupCmd.Flags().StringVarP(&File, "file", "f", "", "output format")
	hostGroupCmd.MarkFlagRequired("format")
}
