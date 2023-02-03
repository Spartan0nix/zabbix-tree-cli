package app

import (
	"log"

	"github.com/Spartan0nix/zabbix-tree-cli/internal/config"
	"github.com/Spartan0nix/zabbix-tree-cli/internal/render"
)

// RunHostGroup is the main entrypoint for the 'host-group' command.
// 1. Initialize the Zabbix API requirements
// 2. Initialize the TreeNode requirements
// 3. Retrieve all HostGroups from the Zabbix Server
// 4. Generate a complete TreeNode for each groups
// 5. Render the TreeNode using the given format
func RunHostGroup(env *config.Env, format string, file string) {
	client, err := InitApi(env.ZabbixUrl, env.ZabbixUser, env.ZabbixPwd)
	if err != nil {
		log.Fatalf("Error when initializing zabbix client.\nReason : %v", err)
	}

	tree, g, graph, err := InitTree(format)
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
