package cmd

import (
	"fmt"
	"log"
	"os"

	zabbixgosdk "github.com/Spartan0nix/zabbix-go-sdk"
	"github.com/Spartan0nix/zabbix-usergroup-tree/internal/api"
	"github.com/Spartan0nix/zabbix-usergroup-tree/internal/render"
	"github.com/Spartan0nix/zabbix-usergroup-tree/internal/tree"
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
	"github.com/spf13/cobra"
)

// This command allow to interact with Zabbix HostGroups
var hostGroupCmd = &cobra.Command{
	Use:   "host-group",
	Short: "TODO",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		err := checkFileFlag(Format, File)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		RunHostGroup(Format, File)
	},
}

func init() {
	hostGroupCmd.Flags().StringVarP(&Format, "format", "o", "", "output format")
	hostGroupCmd.Flags().StringVarP(&File, "file", "f", "", "output format")
	hostGroupCmd.MarkFlagRequired("format")
}

// checkFileFlag is used to check if the given format requires the 'file' flag to be set.
// If the flag is required and the given file variable is empty, return an error.
func checkFileFlag(format string, file string) error {
	var requiredFile bool

	switch format {
	case "png":
		requiredFile = true
	case "jpg":
		requiredFile = true
	case "svg":
		requiredFile = true
	case "json":
		requiredFile = true
	case "shell":
		requiredFile = false
	default:
		return fmt.Errorf("format '%s' is not supported", format)
	}

	// Since file is empty and the flag is required, return an error
	if requiredFile && file == "" {
		return fmt.Errorf("flag 'file' is required for the format '%s'", format)
	}

	return nil
}

// requireGraph is used to check if the given format requires a graphic output.
func requireGraph(format string) (bool, error) {
	switch format {
	case "png":
		return true, nil
	case "jpg":
		return true, nil
	case "svg":
		return true, nil
	case "json":
		return false, nil
	case "shell":
		return false, nil
	default:
		return false, fmt.Errorf("format '%s' is not supported", format)
	}
}

// initApi is used to initialize the default Zabbix service to interact with the API.
// A connectivity test is also run during this step.
func initApi(url string, user string, password string) (*zabbixgosdk.ZabbixService, error) {
	client, err := api.InitService(url)
	if err != nil {
		return nil, err
	}

	err = api.Authenticate(client, user, password)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// initGraph is used to initialize the default graph elements.
func initGraph(root *tree.TreeNode) (*graphviz.Graphviz, *cgraph.Graph, error) {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		return nil, nil, err
	}

	root.GraphNode, err = graph.CreateNode("root")
	if err != nil {
		return nil, nil, err
	}

	return g, graph, nil
}

// initTree is used to initialize the default tree node and if needed, associate the default graphic elements from initGraph.
func initTree(format string) (*tree.TreeNode, *graphviz.Graphviz, *cgraph.Graph, error) {
	tree := tree.TreeNode{Name: "root"}

	requireGraph, err := requireGraph(format)
	if err != nil {
		return nil, nil, nil, err
	}

	var g *graphviz.Graphviz
	var graph *cgraph.Graph

	if requireGraph {
		g, graph, err = initGraph(&tree)

		if err != nil {
			return nil, nil, nil, err
		}
	}

	return &tree, g, graph, nil
}

// RunHostGroup is the main entrypoint for the 'host-group' command.
// 1. Initialize the Zabbix API requirements
// 2. Initialize the TreeNode requirements
// 3. Retrieve all HostGroups from the Zabbix Server
// 4. Generate a complete TreeNode for each groups
// 5. Render the TreeNode using the given format
func RunHostGroup(format string, file string) {
	client, err := initApi(ZABBIX_URL, ZABBIX_USER, ZABBIX_PWD)
	if err != nil {
		log.Fatalf("Error when initializing zabbix client.\nReason : %v", err)
	}

	tree, g, graph, err := initTree(format)
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
