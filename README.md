# zabbix-tree-cli

[![Go Package](https://pkg.go.dev/badge/github.com/Spartan0nix/zabbix-tree-cli?status.svg)](https://pkg.go.dev/github.com/Spartan0nix/zabbix-tree-cli)
[![Go report](https://goreportcard.com/badge/github.com/Spartan0nix/zabbix-tree-cli)](https://goreportcard.com/report/github.com/Spartan0nix/zabbix-tree-cli)
![Test workflow](https://github.com/Spartan0nix/zabbix-tree-cli/actions/workflows/test.yml/badge.svg)
![Build workflow](https://github.com/Spartan0nix/zabbix-tree-cli/actions/workflows/build.yml/badge.svg)

## Table of contents

- [Description](#description)
- [Zabbix-nested-groups](#zabbix-nested-groups)
- [Usage](#usage)
  - [Fixtures (optional)](#fixtures-(optional))
  - [Required environment variables](#required-environment-variables)
  - [Install](#install)
  - [Run](#run)
  - [Completion](#completion)

## Description

This CLI tool is used to help administrator keeps track of their Zabbix Host Groups structure by rendering a graphical output (PNG, JPG, SVG, json, shell).

## Zabbix nested groups

Zabbix supports nested host groups by separating each part in the host groups with '/'.
For example :
```
"Templates/modules/network devices"
-> "Templates"
    -> "modules"
        -> "network devices"
```

## Usage

### Fixtures (optional)

Export files are available in the *fixtures* folder.

- Host Groups
The *'zbx_export_hosts.json'* export will create a host (fixture-host) and the associate host groups to build an example structure since there is no direct export function for the host groups.

- Service
The *'export_services.json'* contains service object definitions used by the *'service.go'* script to create sample services.

<u>Script usage :</u>
```bash
go run fixtures/service.go --url <ZABBIX-URL> --user <ZABBIX_USER> --password <ZABBIX_PWD> --action import
```

Or if the environment variables *"ZABBIX_URL"*, *"ZABBIX_USER"*, *"ZABBIX_PWD"* are defined :
```bash
make import-services
```

> Using the *docker-compose.test.yml* stack combine with the export files can give you a good preview of the possibilities available with this CLI tool.

### Required environment variables

To use this tool, you will need to set up the following variables :
- ZABBIX_URL
- ZABBIX_USER
- ZABBIX_PWD

You can simply export the variable in your current shell :

<u>Linux :</u>
```bash
export ZABBIX_URL="http://<zabbix-server-IP-or-DNS>:<port>/zabbix/api_jsonrpc.php"
export ZABBIX_USER="some-zabbix-user"
export ZABBIX_PWD="some-zabbix-user-password"
```
Adding this configuration to your ~/.bashrc or ~/.zshrc will make the configuration persistent between shell.

<u>Windows (example for powershell) :</u>
```powershell
$env:ZABBIX_URL="http://<zabbix-server-IP-or-DNS>:<port>/zabbix/api_jsonrpc.php"
$env:ZABBIX_USER="some-zabbix-user"
$env:ZABBIX_PWD="some-zabbix-user-password"
```

### Install

1. With a script (available in the *scripts* folder):

    ```bash
    bash scripts/install.sh
    ```

2. Manually :

    Each time a new release is created, the cli is compiled and the resultant binaries are pushed as assets (https://github.com/Spartan0nix/zabbix-tree-cli/releases).

    ```bash
    # Create a temp installation folder
    mkdir /tmp/zabbix-cli

    # Retrieve the archive for release $RELEASE
    wget -O /tmp/zabbix-cli/zabbix-cli.tar.gz https://github.com/Spartan0nix/zabbix-tree-cli/releases/download/$RELEASE/zabbix-tree-cli_$RELEASE_linux_amd64.tar.gz

    # Remove previous install
    sudo rm /usr/local/bin/zabbix-tree-cli

    # Extract the archive
    tar -C /tmp/zabbix-cli -xzf /tmp/zabbix-cli/zabbix-cli.tar.gz

    # Move the binairy
    sudo mv /tmp/zabbix-cli/zabbix-tree-cli /usr/local/bin

    # Update permissions
    sudo chown $(id -un):$(id -gn) /usr/local/bin/zabbix-tree-cli

    # Remove temp installation folder
    rm -r /tmp/zabbix-cli
    ```

### Uninstall

1. With a script (available in the *scripts* folder):

    ```bash
    bash scripts/uninstall.sh
    ```

2. Manually :

    ```bash
    sudo rm /usr/local/bin/zabbix-tree-cli
    ```

### Run
```
This CLI tool is used to help administrator keeps track of their Zabbix structure by rendering a graphical output (dot, json, shell).

Usage:
  zabbix-tree-cli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  host-group  Render a graph for host groups
  service     Render a graph for services

Flags:
      --color         enable colors in graph output
      --debug         enable debug output during execution
  -f, --file string   output to a file
  -h, --help          help for zabbix-tree-cli

Use "zabbix-tree-cli [command] --help" for more information about a command.
```

### Completion

1. Zsh completion

    If shell completion is not enabled in your current shell (oh-my-zsh not running for example), add the following config to your .zshrc :

    ```bash
    echo "autoload -U compinit; compinit" >> ~/.zshrc
    ```

    - To load completions only in the current shell :
    ```bash
    source <(zabbix-tree-cli completion zsh); compdef _zabbix-tree-cli zabbix-tree-cli
    ```

    - To make the configuration persistent between shells :
    ```bash
    zabbix-tree-cli completion zsh > "${fpath[1]}/_zabbix-tree-cli"
    ```

2. Bash completion

    To use completion scripts with bash, you will need to install the "bash-completion" package following your package manager recommendations.


    - To load completions only in the current shell
    ```bash
    source <(zabbix-tree-cli completion bash)
    ```

    - To make the configuration persistent between shells :
    ```bash
    zabbix-tree-cli completion bash > /etc/bash_completion.d/zabbix-tree-cli
    ```

3. Other completions

    - Completion for fish and powershell are available but haven't been tested.