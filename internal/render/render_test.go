package render

import (
	"os"
	"testing"

	"github.com/Spartan0nix/zabbix-tree-cli/internal/tree"
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

func initTree() *tree.TreeNode {
	return &tree.TreeNode{
		Name: "root",
		Childrens: []tree.TreeNode{
			{Name: "node1"},
		},
	}
}

func initGraph() (*tree.TreeNode, *graphviz.Graphviz, *cgraph.Graph, error) {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		return nil, nil, nil, err
	}

	root := tree.TreeNode{Name: "root"}
	root.GraphNode, err = graph.CreateNode("root")
	if err != nil {
		return nil, nil, nil, err
	}

	node1 := tree.TreeNode{Name: "node1"}
	node1.GraphNode, err = graph.CreateNode("node1")
	if err != nil {
		return nil, nil, nil, err
	}
	node1.ParentGraphNode = root.GraphNode

	root.Childrens = append(root.Childrens, node1)

	return &root, g, graph, nil
}

func TestRenderOutputJPG(t *testing.T) {
	tree, g, graph, err := initGraph()
	if err != nil {
		t.Fatalf("Error when initializing tree and graph.\nReason : %v", err)
	}

	err = RenderOutput("./test_render.jpg", "jpg", *tree, g, graph)
	if err != nil {
		t.Fatalf("Error when rendering tree to jpg.\nReason : %v", err)
	}

	CloseGraph(g, graph)
}

func TestRenderOutputPNG(t *testing.T) {
	tree, g, graph, err := initGraph()
	if err != nil {
		t.Fatalf("Error when initializing tree and graph.\nReason : %v", err)
	}

	err = RenderOutput("./test_render.png", "png", *tree, g, graph)
	if err != nil {
		t.Fatalf("Error when rendering tree to png.\nReason : %v", err)
	}

	CloseGraph(g, graph)
}

func TestRenderOutputSVG(t *testing.T) {
	tree, g, graph, err := initGraph()
	if err != nil {
		t.Fatalf("Error when initializing tree and graph.\nReason : %v", err)
	}

	err = RenderOutput("./test_render.svg", "svg", *tree, g, graph)
	if err != nil {
		t.Fatalf("Error when rendering tree to svg.\nReason : %v", err)
	}

	CloseGraph(g, graph)
}

func TestRenderOutputShell(t *testing.T) {
	tree := initTree()

	// Prevent shell output
	// We only when error to be returned
	os.Stdout = nil

	err := RenderOutput("", "shell", *tree, nil, nil)
	if err != nil {
		t.Fatalf("Error when rendering tree to shell.\nReason : %v", err)
	}
}

func TestRenderOutputJson(t *testing.T) {
	tree := initTree()

	err := RenderOutput("./test_render.json", "json", *tree, nil, nil)
	if err != nil {
		t.Fatalf("Error when rendering tree to json.\nReason : %v", err)
	}
}
