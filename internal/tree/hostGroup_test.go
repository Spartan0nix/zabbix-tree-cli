package tree

import (
	"testing"

	zabbixgosdk "github.com/Spartan0nix/zabbix-go-sdk/v2"
)

func TestGenerateHostGroupTree(t *testing.T) {
	tree := TreeNode{Name: "test-root"}

	groups := make([]*zabbixgosdk.HostGroup, 0)
	groups = append(groups, &zabbixgosdk.HostGroup{
		Name: "Templates",
	})
	groups = append(groups, &zabbixgosdk.HostGroup{
		Name: "Templates/Modules",
	})

	err := tree.GenerateHostGroupTree(groups, false)
	if err != nil {
		t.Fatalf("Error when generating the tree.\nReason : %v", err)
	}

	if len(tree.Childrens) != 1 {
		t.Fatalf("Expected number of children for '%s' is 1.\nReturned : %d", tree.Name, len(tree.Childrens))
	}

	if len(tree.Childrens[0].Childrens) != 1 {
		t.Fatalf("Expected number of children for '%s' is 1.\nReturned : %d", tree.Childrens[0].Name, len(tree.Childrens[0].Childrens))
	}
}

func TestGenerateHostGroupTreeSinglePart(t *testing.T) {
	tree := TreeNode{Name: "test-root"}

	groups := make([]*zabbixgosdk.HostGroup, 0)
	groups = append(groups, &zabbixgosdk.HostGroup{
		Name: "Templates",
	})

	err := tree.GenerateHostGroupTree(groups, false)
	if err != nil {
		t.Fatalf("Error when generating the tree.\nReason : %v", err)
	}

	if len(tree.Childrens) != 1 {
		t.Fatalf("Expected number of children for '%s' is 1.\nReturned : %d", tree.Name, len(tree.Childrens))
	}
}
