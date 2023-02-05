package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/Spartan0nix/zabbix-tree-cli/internal/app"
	"github.com/Spartan0nix/zabbix-tree-cli/internal/config"
	"github.com/Spartan0nix/zabbix-tree-cli/internal/render"
	"github.com/spf13/cobra"
)

// This command allow to interact with Zabbix HostGroups and runder a graphical output
var hostGroupCmd = &cobra.Command{
	Use:       "host-group [png|jpg|svg|json|shell]",
	Short:     "Render a graph for host groups",
	ValidArgs: []string{"png", "jpg", "svg", "json", "shell"},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),

	PreRun: func(cmd *cobra.Command, args []string) {
		err := config.CheckFileFlag(args[0], File)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},

	Run: func(cmd *cobra.Command, args []string) {
		env, err := config.GetEnvironmentVariables()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		runHostGroup(env, args[0], File)
	},
}

func init() {
	hostGroupCmd.Flags().StringVarP(&File, "file", "f", "", "output format")
}

// RunHostGroup is the main entrypoint for the 'host-group' command.
// 1. Initialize the Zabbix API requirements
// 2. Initialize the TreeNode requirements
// 3. Retrieve all HostGroups from the Zabbix Server
// 4. Generate a complete TreeNode for each groups
// 5. Render the TreeNode using the given format
func runHostGroup(env *config.Env, format string, file string) {
	client, err := app.InitApi(env.ZabbixUrl, env.ZabbixUser, env.ZabbixPwd)
	if err != nil {
		log.Fatalf("Error when initializing zabbix client.\nReason : %v", err)
	}

	tree, g, graph, err := app.InitTree(format)
	if err != nil {
		log.Fatalf("Error when initializing the tree node.\nReason : %v", err)
	}

	if graph != nil {
		defer render.CloseGraph(g, graph)
	}

	groups, err := client.HostGroup.List()
	if err != nil {
		log.Fatalf("Error when retrieving the list of host groups.\nReason : %v", err)
	}

	err = tree.GenerateHostGroupTree(groups, graph)
	if err != nil {
		log.Fatalf("Error when generating the tree.\nReason : %v", err)
	}

	err = render.RenderOutput(file, format, *tree, g, graph)
	if err != nil {
		log.Fatalf("Error when rendering the tree.\nReason : %v", err)
	}
}
