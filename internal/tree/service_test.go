package tree

import (
	"log"
	"testing"

	zabbixgosdk "github.com/Spartan0nix/zabbix-go-sdk/v2"
)

const (
	ZABBIX_URL  = "http://localhost:4444/api_jsonrpc.php"
	ZABBIX_USER = "Admin"
	ZABBIX_PWD  = "zabbix"
)

var c *zabbixgosdk.ZabbixService

func init() {
	client := zabbixgosdk.NewZabbixService()

	client.Auth.Client.Url = ZABBIX_URL
	client.Service.Client.Url = ZABBIX_URL
	client.Auth.User = &zabbixgosdk.ApiUser{
		User: ZABBIX_USER,
		Pwd:  ZABBIX_PWD,
	}

	err := client.Authenticate()
	if err != nil {
		log.Fatalf("Error when creating testing client.\nReason : %v", err)
	}

	c = client
}

func TestGenerateServiceTree(t *testing.T) {
	tree := TreeNode{Name: "test-root"}

	err := tree.GenerateServiceTree(c, nil, "0")
	if err != nil {
		t.Fatalf("Error executing GenerateServiceTree function.\nReason : %v", err)
	}
}
