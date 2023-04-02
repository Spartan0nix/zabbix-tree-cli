package tree

import (
	zabbixgosdk "github.com/Spartan0nix/zabbix-go-sdk/v2"
	"github.com/Spartan0nix/zabbix-tree-cli/internal/api"
	"github.com/Spartan0nix/zabbix-tree-cli/internal/logging"
)

// GenerateServiceTree is used to generate a full tree of nodes for the available services on the server.
func (t *TreeNode) GenerateServiceTree(client *zabbixgosdk.ZabbixService, logger *logging.Logger, parentId string) error {
	childServices, err := api.ListChildServices(client, parentId)
	if err != nil {
		return err
	}

	for _, service := range childServices {
		node, err := t.createChildren(service.Name)
		if err != nil {
			return err
		}

		if len(service.Children) > 0 {
			err = node.GenerateServiceTree(client, logger, service.ServiceId)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
