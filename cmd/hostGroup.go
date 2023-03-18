package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/Spartan0nix/zabbix-tree-cli/internal/app"
	"github.com/Spartan0nix/zabbix-tree-cli/internal/config"
	"github.com/Spartan0nix/zabbix-tree-cli/internal/render"
	"github.com/Spartan0nix/zabbix-tree-cli/internal/tree"
	"github.com/spf13/cobra"
)

// This command allow to interact with Zabbix HostGroups and runder a graphical output
var hostGroupCmd = &cobra.Command{
	Use:       "host-group [dot|json|shell]",
	Short:     "Render a graph for host groups",
	ValidArgs: []string{"dot", "json", "shell"},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		env, err := config.GetEnvironmentVariables()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		runHostGroup(env, args[0], File, Color)
	},
}

// RunHostGroup is the main entrypoint for the 'host-group' command.
// 1. Initialize the Zabbix API requirements
// 2. Retrieve all HostGroups from the Zabbix Server
// 3. Initialize the TreeNode requirements
// 4. Generate a complete TreeNode for each groups
// 5. Render the TreeNode using the given format
func runHostGroup(env *config.Env, format string, file string, color bool) {
	client, err := app.InitApi(env.ZabbixUrl, env.ZabbixUser, env.ZabbixPwd)
	if err != nil {
		log.Fatalf("Error when initializing zabbix client.\nReason : %v", err)
	}

	groups, err := client.HostGroup.List()
	if err != nil {
		log.Fatalf("Error when retrieving the list of host groups.\nReason : %v", err)
	}

	hash, err := tree.GenerateNewHash(30)
	if err != nil {
		log.Fatalf("Error when generating hash for 'root' node.\nReason : %v", err)
	}

	t := tree.TreeNode{
		Name: "root",
		Id:   hash,
	}

	err = t.GenerateHostGroupTree(groups, false)
	if err != nil {
		log.Fatalf("Error when generating the tree.\nReason : %v", err)
	}

	err = render.RenderTree(file, format, t, color)
	if err != nil {
		log.Fatalf("Error when rendering the tree.\nReason : %v", err)
	}
}
