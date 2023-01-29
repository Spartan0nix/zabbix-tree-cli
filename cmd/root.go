package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	ZABBIX_URL  = "http://localhost:8080/api_jsonrpc.php"
	ZABBIX_USER = "Admin"
	ZABBIX_PWD  = "zabbix"
)

func init() {
	rootCmd.AddCommand(hostGroupCmd)
}

var File string
var Format string

var rootCmd = &cobra.Command{Use: ""}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
