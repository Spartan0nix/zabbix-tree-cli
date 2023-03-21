package tree

import (
	"fmt"
	"strings"

	zabbixgosdk "github.com/Spartan0nix/zabbix-go-sdk/v2"
)

// GenerateHostGroupTree is used to generate a full tree of nodes for the given HostGroups.
// The current node will be used as the root node (or TLD node).
func (t *TreeNode) GenerateHostGroupTree(groups []*zabbixgosdk.HostGroup, debug bool) error {
	var err error
	if debug {
		fmt.Printf("Starting with '%s' as the root node", t.Name)
	}

	for _, group := range groups {
		if debug {
			fmt.Printf("Working on group '%s'\n", group.Name)
		}

		parts := strings.Split(group.Name, "/")
		if debug {
			fmt.Printf("Group '%s' was split in the following part : %v\n", group.Name, parts)
		}

		if len(parts) == 1 {
			if debug {
				fmt.Printf("Only one part to work on.\nAdding node '%s' and returning to the root node.\n\n", parts[0])
			}
			_, err = t.createChildren(parts[0])
			if err != nil {
				return err
			}

			continue
		}

		parentNode := t

		if debug {
			fmt.Printf("More than one parts found.\nStarting to work from node '%s'\n", t.Name)
		}
		for len(parts) > 0 {
			nodeName := parts[0]
			if debug {
				fmt.Printf("Working on part '%s'\n", nodeName)
				fmt.Printf("Using '%s'\n", parentNode.Name)
			}

			// Check if the Node is already present in the current node childrens
			exist := searchNode(nodeName, parentNode.Childrens)
			if exist == nil {
				if debug {
					fmt.Printf("No existing node was found for '%s' under '%s\nCreating one now\n", nodeName, parentNode.Name)
				}

				parentNode, err = parentNode.createChildren(nodeName)
				if err != nil {
					return err
				}
			} else {
				if debug {
					fmt.Printf("An existing node was found for '%s' under '%s'\n", nodeName, parentNode.Name)
				}
				parentNode = exist
			}

			parts = parts[1:]
		}

		if debug {
			fmt.Printf("No more part to work on.\nMoving to next group\n\n")
		}
	}

	return nil
}
