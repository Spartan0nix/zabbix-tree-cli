package render

import (
	"testing"

	"github.com/Spartan0nix/zabbix-tree-cli/internal/tree"
)

func TestOutputInShell(t *testing.T) {
	root := &tree.TreeNode{
		Name: "test-root",
		Id:   "123456789",
		Childrens: []*tree.TreeNode{
			{
				Name: "test-root-children",
				Id:   "987654321",
			},
		},
	}

	b := OutputInShell(root)
	if len(b) == 0 {
		t.Fatalf("An empty slice of byte was returned")
	}
}
