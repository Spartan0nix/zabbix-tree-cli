package render

import (
	"os"
	"testing"

	"github.com/Spartan0nix/zabbix-tree-cli/internal/tree"
)

func TestOutputInShell(t *testing.T) {
	root := &tree.TreeNode{
		Name: "root",
		Childrens: []tree.TreeNode{
			{Name: "node1"},
		},
	}

	// Prevent shell output
	// We only when error to be returned
	os.Stdout = nil

	OutputInShell(root)
}
