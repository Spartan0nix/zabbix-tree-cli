package tree

import (
	"fmt"
	"strings"

	zabbixgosdk "github.com/Spartan0nix/zabbix-go-sdk/v2"
	"github.com/Spartan0nix/zabbix-tree-cli/internal/logging"
)

// GenerateHostGroupTree is used to generate a full tree of nodes for the given HostGroups.
// The current node will be used as the root node (or TLD node).
func (t *TreeNode) GenerateHostGroupTree(groups []*zabbixgosdk.HostGroup, logger *logging.Logger) error {
	var err error
	logger.Debug(fmt.Sprintf("starting with '%s' as the root node", t.Name))

	for _, group := range groups {
		logger.Debug(fmt.Sprintf("working on group '%s'", group.Name))

		parts := strings.Split(group.Name, "/")
		logger.Debug(fmt.Sprintf("group '%s' was split in the following part : %v", group.Name, parts))

		if len(parts) == 1 {
			logger.Debug("only one part to work on", fmt.Sprintf("adding node '%s' and returning to the root node.", parts[0]), "\n")

			_, err = t.createChildren(parts[0])
			if err != nil {
				return err
			}

			continue
		}

		parentNode := t

		logger.Debug("more than one parts found", fmt.Sprintf("starting to work from node '%s'", t.Name))

		for len(parts) > 0 {
			nodeName := parts[0]
			logger.Debug(fmt.Sprintf("working on part '%s'", nodeName))
			logger.Debug(fmt.Sprintf("using '%s'", parentNode.Name))

			// Check if the Node is already present in the current node childrens
			exist := searchNode(nodeName, parentNode.Childrens)
			if exist == nil {
				logger.Debug(fmt.Sprintf("no existing node was found for '%s' under '%s", nodeName, parentNode.Name), "creating one now")

				parentNode, err = parentNode.createChildren(nodeName)
				if err != nil {
					return err
				}
			} else {
				logger.Debug(fmt.Sprintf("an existing node was found for '%s' under '%s'", nodeName, parentNode.Name))
				parentNode = exist
			}

			parts = parts[1:]
		}

		logger.Debug("no more part to work on", "moving to next group", "\n")
	}

	return nil
}
