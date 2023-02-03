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

func TestCheckFileFlag(t *testing.T) {
	err := CheckFileFlag("json", "./local_file")
	if err != nil {
		t.Fatalf("Error when executing CheckFileFlag function.\nReason : %v", err)
	}
}

func TestCheckFileFlagUnsupportedFormat(t *testing.T) {
	err := CheckFileFlag("unsupported_format", "./local_file")
	if err == nil {
		t.Fatalf("CheckFileFlag should return an error when an unsupported format is used.")
	}
}

func TestCheckFileFlagMissingFile(t *testing.T) {
	err := CheckFileFlag("json", "")
	if err == nil {
		t.Fatalf("CheckFileFlag should return an error when the file value is empty for certain formats.")
	}
}

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

	if err := os.Unsetenv("ZABBIX_PWD"); err != nil {
		t.Fatalf("Error when removing ZABBIX_PWD environment variable.\nReason : %v", err)
	}
}

func TestGetEnvironmentVariablesFail(t *testing.T) {
	env, err := GetEnvironmentVariables()
	if err == nil {
		t.Fatalf("GetEnvironmentVariables should returned an error when an environment variable is missing.")
	}

	if env != nil {
		t.Fatalf("An nil pointer should have been returned.\nReturned : %v", env)
	}
}
