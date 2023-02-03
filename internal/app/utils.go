package app

import (
	"fmt"

	zabbixgosdk "github.com/Spartan0nix/zabbix-go-sdk/v2"
	"github.com/Spartan0nix/zabbix-tree-cli/internal/api"
	"github.com/Spartan0nix/zabbix-tree-cli/internal/tree"
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

// RequireGraph is used to check if the given format requires a graphic output.
func RequireGraph(format string) (bool, error) {
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

// InitApi is used to initialize the default Zabbix service to interact with the API.
// A connectivity test is also run during this step.
func InitApi(url string, user string, password string) (*zabbixgosdk.ZabbixService, error) {
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

// InitGraph is used to initialize the default graph elements.
func InitGraph(root *tree.TreeNode) (*graphviz.Graphviz, *cgraph.Graph, error) {
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

// InitTree is used to initialize the default tree node and if needed, associate the default graphic elements from initGraph.
func InitTree(format string) (*tree.TreeNode, *graphviz.Graphviz, *cgraph.Graph, error) {
	tree := tree.TreeNode{Name: "root"}

	requireGraph, err := RequireGraph(format)
	if err != nil {
		return nil, nil, nil, err
	}

	var g *graphviz.Graphviz
	var graph *cgraph.Graph

	if requireGraph {
		g, graph, err = InitGraph(&tree)

		if err != nil {
			return nil, nil, nil, err
		}
	}

	return &tree, g, graph, nil
}
