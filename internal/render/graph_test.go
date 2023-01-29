package render

import (
	"testing"

	"github.com/Spartan0nix/zabbix-usergroup-tree/internal/tree"
	"github.com/goccy/go-graphviz"
)

func TestCloseGraph(t *testing.T) {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		t.Fatalf("Error when creating a new graph.\nReason : %v", err)
	}

	CloseGraph(g, graph)
}

func TestRenderGraphToFile(t *testing.T) {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		t.Fatalf("Error when creating a new graph.\nReason : %v", err)
	}

	_, err = graph.CreateNode("root")
	if err != nil {
		t.Fatalf("Error when adding 'root' node to the graph.\nReason : %v", err)
	}

	err = RenderGraphToFile("./test_render.png", g, graph, graphviz.PNG)
	if err != nil {
		t.Fatalf("Error when rendering graph to a file.\nReason : %v", err)
	}

	CloseGraph(g, graph)
}

func TestRenderGraph(t *testing.T) {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		t.Fatalf("Error when creating a new graph.\nReason : %v", err)
	}

	root := tree.TreeNode{Name: "root"}
	root.GraphNode, err = graph.CreateNode("root")
	if err != nil {
		t.Fatalf("Error when adding 'root' node to the graph.\nReason : %v", err)
	}

	node1 := tree.TreeNode{Name: "node1"}
	node1.GraphNode, err = graph.CreateNode("node1")
	if err != nil {
		t.Fatalf("Error when adding 'node1' node to the graph.\nReason : %v", err)
	}
	node1.ParentGraphNode = root.GraphNode

	root.Childrens = append(root.Childrens, node1)

	err = RenderGraph("./test_render.png", root, g, graph, graphviz.PNG)
	if err != nil {
		t.Fatalf("Error when rendering graph.\nReason : %v", err)
	}

	CloseGraph(g, graph)
}
