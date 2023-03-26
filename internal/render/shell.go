package render

import (
	"fmt"

	"github.com/Spartan0nix/zabbix-tree-cli/internal/tree"
)

// OutputInShell is used to print the content of a treeNode in the shell.
func OutputInShell(t *tree.TreeNode) []byte {
	b := make([]byte, 0)
	out := t.Flatten()

	for _, el := range out {
		b = append(b, []byte("Node:\n")...)
		b = append(b, []byte(fmt.Sprintf("- name      : %s\n", el.Name))...)
		b = append(b, []byte(fmt.Sprintf("- parent    : %s\n\n", el.ParentId))...)
	}

	return b
}
