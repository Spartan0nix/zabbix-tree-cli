package cmd

import (
	"os"
	"os/exec"
	"testing"

	"github.com/Spartan0nix/zabbix-tree-cli/internal/config"
)

func TestRunService(t *testing.T) {
	runService(&config.Env{
		ZabbixUrl:  URL,
		ZabbixUser: USER,
		ZabbixPwd:  PWD,
	}, "dot", "", false)
}

// https://go.dev/talks/2014/testing.slide#23
func TestRunServiceFailAuth(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		runService(&config.Env{
			ZabbixUrl:  URL,
			ZabbixUser: "random-user",
			ZabbixPwd:  PWD,
		}, "dot", "", false)

		return
	}

	// Execute test in a subprocess
	cmd := exec.Command(os.Args[0], "-test.run=TestRunServiceFailAuth")
	// Run the desired command when running test in suprocess
	cmd.Env = append(cmd.Env, "BE_CRASHER=1")
	err := cmd.Run()

	//
	exit, _ := err.(*exec.ExitError)
	if exit.ExitCode() != 1 {
		t.Fatalf("expected exit code 1.\nCode returned : %d\nError returned : %s", exit.ExitCode(), err)
	}
}
