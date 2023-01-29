package tree

import (
	"log"
	"testing"

	zabbixgosdk "github.com/Spartan0nix/zabbix-go-sdk"
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

func initTree() *TreeNode {
	return &TreeNode{
		Name: "root",
		Childrens: []TreeNode{
			{Name: "node1"},
		},
	}
}

func initGraph() (*TreeNode, *graphviz.Graphviz, *cgraph.Graph, error) {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		return nil, nil, nil, err
	}

	root := TreeNode{Name: "root"}
	root.GraphNode, err = graph.CreateNode("root")
	if err != nil {
		return nil, nil, nil, err
	}

	node1 := TreeNode{Name: "node1"}
	node1.GraphNode, err = graph.CreateNode("node1")
	if err != nil {
		return nil, nil, nil, err
	}
	node1.ParentGraphNode = root.GraphNode

	root.Childrens = append(root.Childrens, node1)

	return &root, g, graph, nil
}

func closeGraph(g *graphviz.Graphviz, graph *cgraph.Graph) {
	if err := graph.Close(); err != nil {
		log.Fatalf("Error when closing the graph.\nReason : %v", err)
	}
	g.Close()
}

func TestFlatten(t *testing.T) {
	tree := initTree()

	n := tree.Flatten()

	if len(n) == 0 {
		t.Fatal("An empty list of TreeNode was returned")
	}

	if len(n) != 2 {
		t.Fatalf("Expected a list of 2 TreeNode.\nA list of '%d' was returned.", len(n))
	}
}

func TestSearchNode(t *testing.T) {
	tree := initTree()

	l := make([]TreeNode, 0)
	l = append(l, *tree)
	l = append(l, tree.Childrens[0])

	existingNode := SearchNode("node1", l)
	if existingNode == nil {
		t.Fatalf("A nil pointer was returned instead of a *TreeNode")
	}

	if existingNode.Name != "node1" {
		t.Fatalf("Expected node 'node1'.\nName returned '%s'", existingNode.Name)
	}
}

func TestCreateChildren(t *testing.T) {
	tree, g, graph, err := initGraph()
	if err != nil {
		t.Fatalf("Error when initializing tree and graph.\nReason : %v", err)
	}

	newNode, err := tree.CreateChildren("node2", graph)
	if err != nil {
		t.Fatalf("Error when creating a new children for the current node.\nReason : %v", err)
	}

	if newNode == nil {
		t.Fatalf("A nil pointer was returned instead of a *TreeNode")
	}

	if newNode.Name != "node2" {
		t.Fatalf("Expected node name 'node2'.\nName returned '%s'", newNode.Name)
	}

	closeGraph(g, graph)
}

func TestCreateEdge(t *testing.T) {
	tree, g, graph, err := initGraph()
	if err != nil {
		t.Fatalf("Error when initializing tree and graph.\nReason : %v", err)
	}

	_, err = tree.Childrens[0].CreateEdge(graph)
	if err != nil {
		t.Fatalf("Error when creating edge between 'node1' and 'root'.\nReason : %v", err)
	}

	closeGraph(g, graph)
}

func TestAddNode(t *testing.T) {
	tree, g, graph, err := initGraph()
	if err != nil {
		t.Fatalf("Error when initializing tree and graph.\nReason : %v", err)
	}

	parts := make([]string, 0)
	parts = append(parts, "node2")
	node := &tree.Childrens[0]

	node, parts, err = addNode(tree, node, graph, parts)
	if err != nil {
		t.Fatalf("Erro when adding a node to 'node1'.\nReason : %v", err)
	}

	if len(parts) > 0 {
		t.Fatalf("Expected length of parts list to be 0.\nLength of parts list is '%d'.", len(parts))
	}

	if node == nil || node.Name != "root" {
		t.Fatalf("addNode method did not returned the pointer of the root node (expected since the length of the parts list is 0).")
	}

	closeGraph(g, graph)
}

func TestGenerateHostGroupTree(t *testing.T) {
	_, g, graph, err := initGraph()
	if err != nil {
		t.Fatalf("Error when initializing tree and graph.\nReason : %v", err)
	}

	tree := TreeNode{Name: "root"}
	groups := make([]*zabbixgosdk.HostGroup, 0)
	groups = append(groups, &zabbixgosdk.HostGroup{
		Name: "Templates",
	})
	groups = append(groups, &zabbixgosdk.HostGroup{
		Name: "Templates/Modules",
	})

	err = tree.GenerateHostGroupTree(groups, graph)
	if err != nil {
		t.Fatalf("Error when generating the tree.\nReason : %v", err)
	}

	closeGraph(g, graph)
}

func TestGenerateTreeEdges(t *testing.T) {
	tree, g, graph, err := initGraph()
	if err != nil {
		t.Fatalf("Error when initializing tree and graph.\nReason : %v", err)
	}

	l := make([]TreeNode, 0)
	l = append(l, *tree)
	l = append(l, tree.Childrens[0])

	graph, err = GenerateTreeEdges(graph, l)
	if err != nil {
		t.Fatalf("Error when generating tree edges.\nReason : %v", err)
	}

	closeGraph(g, graph)
}
