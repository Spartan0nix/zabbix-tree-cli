package app

import (
	zabbixgosdk "github.com/Spartan0nix/zabbix-go-sdk/v2"
	"github.com/Spartan0nix/zabbix-tree-cli/internal/api"
)

// InitApi is used to initialize the default Zabbix service to interact with the API.
// A connectivity test is also run during this step.
func InitApi(url string, user string, password string) (*zabbixgosdk.ZabbixService, error) {
	client, err := api.InitService(url)
	if err != nil {
		return nil, err
	}

	err = api.Authenticate(client, user, password)
	if err != nil {
		return nil, err
	}

	return client, nil
}
