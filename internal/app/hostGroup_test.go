package app

import (
	"testing"

	"github.com/Spartan0nix/zabbix-tree-cli/internal/config"
)

func TestRunHostGroupWithGraph(t *testing.T) {
	RunHostGroup(&config.Env{
		ZabbixUrl:  URL,
		ZabbixUser: USER,
		ZabbixPwd:  PWD,
	}, "png", "test_render.png")
}

func TestRunHostGroupWithoutGraph(t *testing.T) {
	RunHostGroup(&config.Env{
		ZabbixUrl:  URL,
		ZabbixUser: USER,
		ZabbixPwd:  PWD,
	}, "json", "test_render.json")
}
