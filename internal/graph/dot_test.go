package graph

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/Spartan0nix/zabbix-tree-cli/internal/tree"
)

func TestWriteDotNode(t *testing.T) {
	var b bytes.Buffer
	indentation := 0

	err := writeDotNode(&indentation, &b, &tree.TreeNode{
		Name: "test-root",
		Id:   "123456789",
	})

	if err != nil {
		t.Fatalf("Error when executing writeDotNote function.\nReason : %v", err)
	}

	expectedOutput := fmt.Sprintf("\"%s\" [label=\"%s\",shape=box];\n", "123456789", "test-root")
	if b.String() != expectedOutput {
		t.Fatalf("Wrong output returned\nExpected : %s\nReturned : %s", expectedOutput, b.String())
	}
}

func TestWriteDotConnections(t *testing.T) {
	var b bytes.Buffer
	indentation := 0

	err := writeDotConnections(&indentation, &b, &tree.TreeNode{
		Name: "test-root",
		Id:   "123456789",
		Childrens: []*tree.TreeNode{
			{
				Name: "test-root-children",
				Id:   "987654321",
			},
		},
	})

	if err != nil {
		t.Fatalf("Error when executing writeDotConnections.\nReason : %v", err)
	}

	expectedOutput := fmt.Sprintf("\"%s\" -> \"%s\";\n", "123456789", "987654321")
	if b.String() != expectedOutput {
		t.Fatalf("Wrong output returned\nExpected : %s\nReturned : %s", expectedOutput, b.String())
	}
}

func TestWriteDotConnectionsNoChild(t *testing.T) {
	var b bytes.Buffer
	indentation := 0

	err := writeDotConnections(&indentation, &b, &tree.TreeNode{
		Name: "test-root",
		Id:   "123456789",
	})

	if err != nil {
		t.Fatalf("Error when executing writeDotConnections.\nReason : %v", err)
	}

	if b.String() != "" {
		t.Fatalf("Wrong output returned\nExpected nothing to be added to the buffer\nReturned : %s", b.String())
	}
}

func TestWriteDotSubGraph(t *testing.T) {
	var b bytes.Buffer
	indentation := 0

	err := writeDotSubGraph(&indentation, &b, &tree.TreeNode{
		Name: "test-root",
		Id:   "123456789",
		Childrens: []*tree.TreeNode{
			{
				Name: "test-root-children",
				Id:   "987654321",
			},
		},
	}, false)

	if err != nil {
		t.Fatalf("Error when executing writeDotSubGraph.\nReason : %v", err)
	}

	expectedOutput := fmt.Sprintf("subgraph cluster%s {\n", "123456789")
	expectedOutput += fmt.Sprintf("\t\"%s\" [label=\"%s\",shape=box];\n", "123456789", "test-root")
	expectedOutput += fmt.Sprintf("\t\"%s\" -> \"%s\";\n", "123456789", "987654321")
	expectedOutput += fmt.Sprintf("\t\"%s\" [label=\"%s\",shape=box];\n", "987654321", "test-root-children")
	expectedOutput += "}\n"

	if b.String() != expectedOutput {
		t.Fatalf("Wrong output returned\nExpected : %s\nReturned : %s", expectedOutput, b.String())
	}
}

func TestWriteDotSubGraphWithColor(t *testing.T) {
	var b bytes.Buffer
	indentation := 0

	err := writeDotSubGraph(&indentation, &b, &tree.TreeNode{
		Name: "test-root",
		Id:   "123456789",
		Childrens: []*tree.TreeNode{
			{
				Name: "test-root-children",
				Id:   "987654321",
			},
		},
	}, true)

	if err != nil {
		t.Fatalf("Error when executing writeDotSubGraph.\nReason : %v", err)
	}

}

func TestWriteDotSubGraphNoChild(t *testing.T) {
	var b bytes.Buffer
	indentation := 0

	err := writeDotSubGraph(&indentation, &b, &tree.TreeNode{
		Name: "test-root",
		Id:   "123456789",
	}, false)

	if err != nil {
		t.Fatalf("Error when executing writeDotSubGraph.\nReason : %v", err)
	}

	expectedOutput := fmt.Sprintf("\"%s\" [label=\"%s\",shape=box];\n", "123456789", "test-root")

	if b.String() != expectedOutput {
		t.Fatalf("Wrong output returned\nExpected : %s\nReturned : %s", expectedOutput, b.String())
	}
}

func TestWriteDotGraphConfiguration(t *testing.T) {
	var b bytes.Buffer
	indentation := 0

	err := writeDotGraphConfiguration(&indentation, &b, []string{
		"random-config1",
		"random-config2",
	})

	if err != nil {
		t.Fatalf("Error when executing writeDotGraphConfiguration\nReason : %v", err)
	}

	expectedOutput := "random-config1\n"
	expectedOutput += "random-config2\n"

	if b.String() != expectedOutput {
		t.Fatalf("Wrong output returned\nExpected : %s\nReturned : %s", expectedOutput, b.String())
	}
}

func TestRenderDotGraph(t *testing.T) {
	b, err := RenderDotGraph(tree.TreeNode{
		Name: "test-root",
		Id:   "123456789",
		Childrens: []*tree.TreeNode{
			{
				Name: "test-root-children",
				Id:   "987654321",
			},
		},
	}, false)

	if err != nil {
		t.Fatalf("Error when executing RenderDotGraph\nReason : %v", err)
	}

	expectedOutput := "digraph {\n"
	expectedOutput += "layout=\"dot\";\n"
	expectedOutput += "ranksep=\"1.0 equally\";\n"
	expectedOutput += "graph[style=filled,color=white];\n"
	expectedOutput += fmt.Sprintf("\tsubgraph cluster%s {\n", "123456789")
	expectedOutput += fmt.Sprintf("\t\t\"%s\" [label=\"%s\",shape=box];\n", "123456789", "test-root")
	expectedOutput += fmt.Sprintf("\t\t\"%s\" -> \"%s\";\n", "123456789", "987654321")
	expectedOutput += fmt.Sprintf("\t\t\"%s\" [label=\"%s\",shape=box];\n", "987654321", "test-root-children")
	expectedOutput += "\t}\n"
	expectedOutput += "}\n"

	if b.String() != expectedOutput {
		t.Fatalf("Wrong output returned\nExpected : %s\nReturned : %s", expectedOutput, b.String())
	}
}

func TestRenderDotGraphWithColor(t *testing.T) {
	_, err := RenderDotGraph(tree.TreeNode{
		Name: "test-root",
		Id:   "123456789",
		Childrens: []*tree.TreeNode{
			{
				Name: "test-root-children",
				Id:   "987654321",
			},
		},
	}, false)

	if err != nil {
		t.Fatalf("Error when executing RenderDotGraph\nReason : %v", err)
	}
}
