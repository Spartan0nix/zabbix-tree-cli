package app

import (
	"testing"

	"github.com/Spartan0nix/zabbix-tree-cli/internal/tree"
)

const (
	URL  = "http://localhost:4444/api_jsonrpc.php"
	USER = "Admin"
	PWD  = "zabbix"
)

func TestRequireGraphPNG(t *testing.T) {
	b, err := RequireGraph("png")
	if err != nil {
		t.Fatalf("Error when executing RequireGraph function.\nReason : %v", err)
	}

	if !b {
		t.Fatalf("png format should require a graph (%t was returned).", b)
	}
}

func TestRequireGraphJPG(t *testing.T) {
	b, err := RequireGraph("jpg")
	if err != nil {
		t.Fatalf("Error when executing RequireGraph function.\nReason : %v", err)
	}

	if !b {
		t.Fatalf("jpg format should require a graph (%t was returned).", b)
	}
}

func TestRequireGraphSVG(t *testing.T) {
	b, err := RequireGraph("svg")
	if err != nil {
		t.Fatalf("Error when executing RequireGraph function.\nReason : %v", err)
	}

	if !b {
		t.Fatalf("svg format should require a graph (%t was returned).", b)
	}
}

func TestRequireGraphJson(t *testing.T) {
	b, err := RequireGraph("json")
	if err != nil {
		t.Fatalf("Error when executing RequireGraph function.\nReason : %v", err)
	}

	if b {
		t.Fatalf("json format should not require a graph (%t was returned).", b)
	}
}

func TestRequireGraphShell(t *testing.T) {
	b, err := RequireGraph("shell")
	if err != nil {
		t.Fatalf("Error when executing RequireGraph function.\nReason : %v", err)
	}

	if b {
		t.Fatalf("shell format should not require a graph (%t was returned).", b)
	}
}

func TestRequireGraphUnsupportedFormat(t *testing.T) {
	_, err := RequireGraph("unsupported_format")
	if err == nil {
		t.Fatalf("RequireGraph should return an error when using an unsupported format.")
	}
}

func TestInitApi(t *testing.T) {
	client, err := InitApi(URL, USER, PWD)
	if err != nil {
		t.Fatalf("Error when executing InitApi function.\nReason : %v", err)
	}

	if client == nil {
		t.Fatalf("A nil pointer was returned instead of *zabbixgosdk.ZabbixService.")
	}
}

func TestInitApiFailConnectivity(t *testing.T) {
	client, err := InitApi("http://localhost:5555/api_jsonrpc.php", USER, PWD)
	if err == nil {
		t.Fatalf("An error should be returned when failing to connect to the Zabbix server.")
	}

	if client != nil {
		t.Fatalf("A nil pointer was not returned.")
	}
}

func TestInitApiFailAuth(t *testing.T) {
	client, err := InitApi(URL, "random-user", "random-password")
	if err == nil {
		t.Fatalf("An error should be returned when failing to authenticate against the Zabbix server.")
	}

	if client != nil {
		t.Fatalf("A nil pointer was not returned.")
	}
}

func TestInitGraph(t *testing.T) {
	root := tree.TreeNode{
		Name: "root",
	}

	g, graph, err := InitGraph(&root)
	if err != nil {
		t.Fatalf("Error when executing InitGraph function;\nReason : %v", err)
	}

	if g == nil {
		t.Fatalf("A nil pointer was returned instead of *graphviz.Graphviz.")
	}

	if graph == nil {
		t.Fatalf("A nil pointer was returned instead of *cgraph.Graph.")
	}

	if err := graph.Close(); err != nil {
		t.Fatalf("Error when closing the graph.\nReason : %v", err)
	}
	g.Close()
}

func TestInitTreeWithGraph(t *testing.T) {
	root, g, graph, err := InitTree("png")
	if err != nil {
		t.Fatalf("Error when executing InitTree function.\nReason : %v", err)
	}

	if root == nil {
		t.Fatalf("A nil pointer was returned instead of *tree.TreeNode.")
	}

	if g == nil {
		t.Fatalf("A nil pointer was returned instead of *graphviz.Graphviz.")
	}

	if graph == nil {
		t.Fatalf("A nil pointer was returned instead of *cgraph.Graph.")
	}

	if root.Name != "root" {
		t.Fatalf("Root tree node should be named 'root' not '%s'.", root.Name)
	}
}

func TestInitTreeWithoutGraph(t *testing.T) {
	root, g, graph, err := InitTree("json")
	if err != nil {
		t.Fatalf("Error when executing InitTree function.\nReason : %v", err)
	}

	if root == nil {
		t.Fatalf("A nil pointer was returned instead of *tree.TreeNode.")
	}

	if g != nil {
		t.Fatalf("A graphviz.Graphviz pointer was returned instead of a nil pointer.")
	}

	if graph != nil {
		t.Fatalf("A cgraph.Graph pointer was returned instead of a nil pointer.")
	}

	if root.Name != "root" {
		t.Fatalf("Root tree node should be named 'root' not '%s'.", root.Name)
	}
}
