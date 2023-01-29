package render

import (
	"testing"

	"github.com/Spartan0nix/zabbix-usergroup-tree/internal/tree"
)

func TestOutputAsJson(t *testing.T) {
	root := tree.TreeNode{
		Name: "root",
		Childrens: []tree.TreeNode{
			{Name: "node1"},
		},
	}

	err := OutputAsJson("./test_render.json", &root)
	if err != nil {
		t.Fatalf("Error when rendering tree node to json.\nReason : %v", err)
	}
}
