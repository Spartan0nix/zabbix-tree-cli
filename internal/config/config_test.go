package config

import (
	"os"
	"testing"
)

const (
	URL  = "http://localhost:4444/api_jsonrpc.php"
	USER = "Admin"
	PWD  = "zabbix"
)

func TestGetEnvironmentVariables(t *testing.T) {
	if err := os.Setenv("ZABBIX_URL", URL); err != nil {
		t.Fatalf("Error when setting ZABBIX_URL environment variable.\nReason : %v", err)
	}

	if err := os.Setenv("ZABBIX_USER", USER); err != nil {
		t.Fatalf("Error when setting ZABBIX_USER environment variable.\nReason : %v", err)
	}

	if err := os.Setenv("ZABBIX_PWD", PWD); err != nil {
		t.Fatalf("Error when setting ZABBIX_PWD environment variable.\nReason : %v", err)
	}

	env, err := GetEnvironmentVariables()
	if err != nil {
		t.Fatalf("Error when executing GetEnvironmentVariables function.\nReason : %v", err)
	}

	if env.ZabbixUrl != URL {
		t.Fatalf("Wrong variable retrieved.\nExpected : %s\nReturned : %s", URL, env.ZabbixUrl)
	}

	if env.ZabbixUser != USER {
		t.Fatalf("Wrong variable retrieved.\nExpected : %s\nReturned : %s", USER, env.ZabbixUser)
	}

	if env.ZabbixPwd != PWD {
		t.Fatalf("Wrong variable retrieved.\nExpected : %s\nReturned : %s", PWD, env.ZabbixPwd)
	}
}

func TestGetEnvironmentVariablesMissingUrl(t *testing.T) {
	if err := os.Unsetenv("ZABBIX_URL"); err != nil {
		t.Fatalf("Error when removing ZABBIX_URL environment variable.\nReason : %v", err)
	}

	if err := os.Setenv("ZABBIX_USER", USER); err != nil {
		t.Fatalf("Error when setting ZABBIX_USER environment variable.\nReason : %v", err)
	}

	if err := os.Setenv("ZABBIX_PWD", PWD); err != nil {
		t.Fatalf("Error when setting ZABBIX_PWD environment variable.\nReason : %v", err)
	}

	env, err := GetEnvironmentVariables()
	if err == nil {
		t.Fatalf("GetEnvironmentVariables should returned an error when an environment variable is missing.")
	}

	if env != nil {
		t.Fatalf("An nil pointer should have been returned.\nReturned : %v", env)
	}
}

func TestGetEnvironmentVariablesMissingUser(t *testing.T) {
	if err := os.Setenv("ZABBIX_URL", URL); err != nil {
		t.Fatalf("Error when setting ZABBIX_URL environment variable.\nReason : %v", err)
	}

	if err := os.Unsetenv("ZABBIX_USER"); err != nil {
		t.Fatalf("Error when removing ZABBIX_USER environment variable.\nReason : %v", err)
	}

	if err := os.Setenv("ZABBIX_PWD", PWD); err != nil {
		t.Fatalf("Error when setting ZABBIX_PWD environment variable.\nReason : %v", err)
	}

	env, err := GetEnvironmentVariables()
	if err == nil {
		t.Fatalf("GetEnvironmentVariables should returned an error when an environment variable is missing.")
	}

	if env != nil {
		t.Fatalf("An nil pointer should have been returned.\nReturned : %v", env)
	}
}

func TestGetEnvironmentVariablesMissingPassword(t *testing.T) {
	if err := os.Setenv("ZABBIX_URL", URL); err != nil {
		t.Fatalf("Error when setting ZABBIX_URL environment variable.\nReason : %v", err)
	}

	if err := os.Setenv("ZABBIX_USER", USER); err != nil {
		t.Fatalf("Error when setting ZABBIX_USER environment variable.\nReason : %v", err)
	}

	if err := os.Unsetenv("ZABBIX_PWD"); err != nil {
		t.Fatalf("Error when removing ZABBIX_PWD environment variable.\nReason : %v", err)
	}

	env, err := GetEnvironmentVariables()
	if err == nil {
		t.Fatalf("GetEnvironmentVariables should returned an error when an environment variable is missing.")
	}

	if env != nil {
		t.Fatalf("An nil pointer should have been returned.\nReturned : %v", env)
	}
}
