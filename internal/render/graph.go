package render

import (
	"log"

	"github.com/Spartan0nix/zabbix-tree-cli/internal/tree"
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

// CloseGraph is used to close the given graph objects.
func CloseGraph(g *graphviz.Graphviz, graph *cgraph.Graph) {
	if err := graph.Close(); err != nil {
		log.Fatalf("Error when closing the graph.\nReason : %v", err)
	}
	g.Close()
}

// RenderGraphToFile is used to render a graphic to a file.
func RenderGraphToFile(file string, g *graphviz.Graphviz, graph *cgraph.Graph, format graphviz.Format) error {
	err := g.RenderFilename(graph, format, file)
	if err != nil {
		return err
	}

	return nil
}

// RenderGraph is used to render a graphic to a file after generating the tree edges.
// Each node of the given TreeNode will be linked to it's parent node.
func RenderGraph(file string, t tree.TreeNode, g *graphviz.Graphviz, graph *cgraph.Graph, format graphviz.Format) error {
	var err error

	// Generate each nodes links.
	graph, err = tree.GenerateTreeEdges(graph, t.Flatten())
	if err != nil {
		return err
	}

	// Render the tree to the given file
	err = RenderGraphToFile(file, g, graph, format)

	return err
}
