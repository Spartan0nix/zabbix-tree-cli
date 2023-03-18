package app

import (
	"testing"
)

const (
	URL  = "http://localhost:4444/api_jsonrpc.php"
	USER = "Admin"
	PWD  = "zabbix"
)

func TestInitApi(t *testing.T) {
	client, err := InitApi(URL, USER, PWD)
	if err != nil {
		t.Fatalf("Error when executing InitApi function.\nReason : %v", err)
	}

	if client == nil {
		t.Fatalf("A nil pointer was returned instead of *zabbixgosdk.ZabbixService.")
	}
}

func TestInitApiFailConnectivity(t *testing.T) {
	client, err := InitApi("http://localhost:5555/api_jsonrpc.php", USER, PWD)
	if err == nil {
		t.Fatalf("An error should be returned when failing to connect to the Zabbix server.")
	}

	if client != nil {
		t.Fatalf("A nil pointer was not returned.")
	}
}

func TestInitApiFailAuth(t *testing.T) {
	client, err := InitApi(URL, "random-user", "random-password")
	if err == nil {
		t.Fatalf("An error should be returned when failing to authenticate against the Zabbix server.")
	}

	if client != nil {
		t.Fatalf("A nil pointer was not returned.")
	}
}
