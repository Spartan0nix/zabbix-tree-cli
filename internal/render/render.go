package render

import (
	"fmt"

	"github.com/Spartan0nix/zabbix-tree-cli/internal/tree"
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

// RenderOutput is used to render a tree node depending of the format passed
func RenderOutput(file string, format string, t tree.TreeNode, g *graphviz.Graphviz, graph *cgraph.Graph) error {
	var err error

	switch format {

	case "jpg":
		if err = RenderGraph(file, t, g, graph, graphviz.JPG); err != nil {
			return err
		}

	case "png":
		if err = RenderGraph(file, t, g, graph, graphviz.PNG); err != nil {
			return err
		}

	case "svg":
		if err = RenderGraph(file, t, g, graph, graphviz.SVG); err != nil {
			return err
		}

	case "shell":
		OutputInShell(&t)

	case "json":
		err = OutputAsJson(file, &t)

	default:
		err = fmt.Errorf("format '%s' is not supported", format)
	}

	return err
}
