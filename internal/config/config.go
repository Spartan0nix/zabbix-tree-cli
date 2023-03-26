package config

import (
	"fmt"
	"os"
)

type Env struct {
	ZabbixUrl  string
	ZabbixUser string
	ZabbixPwd  string
}

// GetEnvironmentVariables is used to retrive the required environment variables for the Zabbix API.
func GetEnvironmentVariables() (*Env, error) {
	vars := Env{}
	if vars.ZabbixUrl = os.Getenv("ZABBIX_URL"); vars.ZabbixUrl == "" {
		return nil, fmt.Errorf("required environment variable 'ZABBIX_URL' is not set")
	}

	if vars.ZabbixUser = os.Getenv("ZABBIX_USER"); vars.ZabbixUser == "" {
		return nil, fmt.Errorf("required environment variable 'ZABBIX_USER' is not set")
	}

	if vars.ZabbixPwd = os.Getenv("ZABBIX_PWD"); vars.ZabbixPwd == "" {
		return nil, fmt.Errorf("required environment variable 'ZABBIX_PWD' is not set")
	}

	return &vars, nil
}
