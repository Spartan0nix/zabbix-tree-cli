package tree

import (
	"testing"
)

func initTree() *TreeNode {
	return &TreeNode{
		Name: "test-root",
		Childrens: []*TreeNode{
			{Name: "test-root-children"},
		},
	}
}

func TestGenerateNewHash(t *testing.T) {
	hash, err := GenerateNewHash(30)
	if err != nil {
		t.Fatalf("Error when executing GenerateNewHash function\nReason : %v", err)
	}

	if hash == "" {
		t.Fatalf("Empty string returned instead of a new hash")
	}
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

	l := make([]*TreeNode, 0)
	l = append(l, tree)
	l = append(l, tree.Childrens[0])

	expectedNodeName := "test-root-children"
	exist := searchNode(expectedNodeName, l)
	if exist == nil {
		t.Fatalf("A nil pointer was returned instead of a *TreeNode")
	}

	if exist.Name != expectedNodeName {
		t.Fatalf("Expected node '%s'.\nName returned '%s'", expectedNodeName, exist.Name)
	}
}

func TestCreateChildren(t *testing.T) {
	tree := initTree()
	expectedNodeName := "test-root-children2"

	n, err := tree.createChildren(expectedNodeName)
	if err != nil {
		t.Fatalf("Error when creating a new children for the root node.\nReason : %v", err)
	}

	if n == nil {
		t.Fatalf("A nil pointer was returned instead of a *TreeNode")
	}

	if n.Name != expectedNodeName {
		t.Fatalf("Expected node name '%s'.\nName returned '%s'", expectedNodeName, n.Name)
	}

	if n.Id == "" {
		t.Fatalf("A new hash was not generated when the new node was created")
	}
}
