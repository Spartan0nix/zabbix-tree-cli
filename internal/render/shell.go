package render

import (
	"fmt"

	"github.com/Spartan0nix/zabbix-tree-cli/internal/tree"
)

// OutputInShell is used to print the content of a treeNode in the shell
func OutputInShell(t *tree.TreeNode) {
	out := t.Flatten()
	for _, el := range out {
		fmt.Printf("Node:\n")
		fmt.Printf("- name      : %s\n", el.Name)
		if el.ParentGraphNode != nil {
			fmt.Printf("- parent    : %s\n", el.ParentGraphNode.Name())
		}
		fmt.Println("")
	}
}
