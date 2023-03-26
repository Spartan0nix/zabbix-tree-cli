package api

import (
	"testing"

	zabbixgosdk "github.com/Spartan0nix/zabbix-go-sdk/v2"
)

const (
	ZABBIX_URL  = "http://localhost:4444/api_jsonrpc.php"
	ZABBIX_USER = "Admin"
	ZABBIX_PWD  = "zabbix"
)

var c *zabbixgosdk.ZabbixService

func TestInitService(t *testing.T) {
	client, err := InitService(ZABBIX_URL)
	if err != nil {
		t.Fatalf("Error when initializing service.\nReason : %v", err)
	}

	if client == nil {
		t.Fatal("InitService returned an nil value instead of a *zabbixgosdk.ZabbixService")
	}

	c = client
}

func TestInitServiceFailConnectivity(t *testing.T) {
	_, err := InitService("http://localhost:2222/api_jsonrpc.php")

	if err == nil {
		t.Fatal("A nil error was returned.\nExpected a connectivity error.")
	}
}

func TestAuthenticate(t *testing.T) {
	err := Authenticate(c, ZABBIX_USER, ZABBIX_PWD)
	if err != nil {
		t.Fatalf("Error when executing Authenticate function.\nReason : %v", err)
	}

	if c.HostGroup.Client.Token == "" {
		t.Fatal("No API token was found in the HostGroup service")
	}
}

func TestAuthenticateBadCredentials(t *testing.T) {
	err := Authenticate(c, "random-user", "random-password")

	if err == nil {
		t.Fatal("A nil error was returned.\nExpected an authentification error.")
	}
}

func TestAuthenticateBadPassword(t *testing.T) {
	err := Authenticate(c, ZABBIX_USER, "random-password")

	if err == nil {
		t.Fatal("A nil error was returned.\nExpected an authentification error.")
	}
}
