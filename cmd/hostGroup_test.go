package cmd

import (
	"testing"

	"github.com/Spartan0nix/zabbix-tree-cli/internal/config"
)

const (
	URL  = "http://localhost:4444/api_jsonrpc.php"
	USER = "Admin"
	PWD  = "zabbix"
)

func TestRunHostGroup(t *testing.T) {
	runHostGroup(&config.Env{
		ZabbixUrl:  URL,
		ZabbixUser: USER,
		ZabbixPwd:  PWD,
	}, "dot", "", false)
}

// func TestRunHostGroup(t *testing.T) {
// 	runHostGroup(&config.Env{
// 		ZabbixUrl:  URL,
// 		ZabbixUser: USER,
// 		ZabbixPwd:  PWD,
// 	}, "dot", "", false)
// }
