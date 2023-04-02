package cmd

import (
	"fmt"
	"os"

	"github.com/Spartan0nix/zabbix-tree-cli/internal/app"
	"github.com/Spartan0nix/zabbix-tree-cli/internal/config"
	"github.com/Spartan0nix/zabbix-tree-cli/internal/logging"
	"github.com/Spartan0nix/zabbix-tree-cli/internal/render"
	"github.com/Spartan0nix/zabbix-tree-cli/internal/tree"
	"github.com/spf13/cobra"
)

// This command allow to interact with Zabbix Services and render a graphical output
var serviceCmd = &cobra.Command{
	Use:       "service [dot|json|shell]",
	Short:     "Render a graph for services",
	ValidArgs: []string{"dot", "json", "shell"},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		if Debug {
			GlobalLogger.Level = logging.Debug
		}

		env, err := config.GetEnvironmentVariables()
		if err != nil {
			GlobalLogger.Error("error when reading the required environment variables", fmt.Sprintf("reason : %s", err))
			os.Exit(1)
		}

		runService(env, args[0], File, Color)
	},
}

// runService is the main entrypoint for the 'host-group' command.
// 1. Initialize the Zabbix API requirements
// 3. Initialize the TreeNode requirements
// 4. Generate a complete TreeNode for each services
// 5. Render the TreeNode using the given format
func runService(env *config.Env, format string, file string, color bool) {
	client, err := app.InitApi(env.ZabbixUrl, env.ZabbixUser, env.ZabbixPwd)
	if err != nil {
		GlobalLogger.Error("error when initializing zabbix client", fmt.Sprintf("reason : %v", err))
		os.Exit(1)
	}

	hash, err := tree.GenerateNewHash(30)
	if err != nil {
		GlobalLogger.Error("error when generating hash for 'root' node", fmt.Sprintf("reason : %v", err))
		os.Exit(1)
	}

	t := tree.TreeNode{
		Name: "root",
		Id:   hash,
	}

	err = t.GenerateServiceTree(client, GlobalLogger, "0")
	if err != nil {
		GlobalLogger.Error("error when generating the tree", fmt.Sprintf("reason : %v", err))
		os.Exit(1)
	}

	err = render.RenderTree(file, format, t, color)
	if err != nil {
		GlobalLogger.Error("error when rendering the tree", fmt.Sprintf("reason : %v", err))
		os.Exit(1)
	}
}
