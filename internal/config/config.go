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

// CheckFileFlag is used to check if the given format requires the 'file' flag to be set.
// If the flag is required and the given file variable is empty, return an error.
func CheckFileFlag(format string, file string) error {
	var requiredFile bool

	switch format {
	case "png":
		requiredFile = true
	case "jpg":
		requiredFile = true
	case "svg":
		requiredFile = true
	case "json":
		requiredFile = true
	case "shell":
		requiredFile = false
	default:
		return fmt.Errorf("format '%s' is not supported", format)
	}

	// Since file is empty and the flag is required, return an error
	if requiredFile && file == "" {
		return fmt.Errorf("flag 'file' is required for the format '%s'", format)
	}

	return nil
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
