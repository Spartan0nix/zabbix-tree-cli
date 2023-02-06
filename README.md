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

A json export is available in the *fixtures* folder.
This export will create a host (fixture-host) and the associates host groups to build an example structure.

Using the *docker-compose.test.yml* stack combine with the export file can give you a good preview of the possibility available with this CLI tool.

### Required environment variables

To use this tool, you will need to set up the following variables :
- ZABBIX_URL
- ZABBIX_USER
- ZABBIX_PWD

You can simply export the variable in your current shell :
```bash
export ZABBIX_URL="http://<zabbix-server-IP-or-DNS>:<port>/zabbix/api_jsonrpc.php"
export ZABBIX_USER="some-zabbix-user"
export ZABBIX_PWD="some-zabbix-user-password"
```

Adding this configuration to your ~/.bashrc or ~/.zshrc will make the configuration persistent between shell.


### Install

Each time a new release is created, the cli is compiled and the resultant binaries are pushed as assets.
Currently, only the linux/amd64 arch is supported (due to the 'go-graphviz' module build restrictions).

1. Script (available in the *scripts* folder):

    ```bash
    bash scripts/install.sh
    ```

2. Manually :

    ```bash
    # Retrieve the archive for release v1.0.0
    wget https://github.com/Spartan0nix/zabbix-tree-cli/releases/download/v1.0.0/zabbix-tree-cli_1.0.0_linux_amd64.tar.gz
    # Remove previous install
    sudo rm -r /usr/local/zabbix-tree-cli
    # Create a folder to store the binary
    sudo mkdir /usr/local/zabbix-tree-cli
    # Extract the archive
    sudo tar -C /usr/local/zabbix-tree-cli -xzf zabbix-tree-cli_1.0.0_linux_amd64.tar.gz
    # Update your PATH
    export PATH=$PATH:/usr/local/zabbix-tree-cli
    ```

### Uninstall

1. Script (available in the *scripts* folder):

    ```bash
    bash scripts/uninstall.sh
    ```

2. Manually :

    ```bash
    sudo rm -r /usr/local/zabbix-tree-cli
    ```

### Run
```
zabbix-tree-cli [command]

Usage:
   [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  host-group  Render a Zabbix host groups graph.

Flags:
  -h, --help   help for this command

Use " [command] --help" for more information about a command.
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

    - Completion for fish haven't been tested.
    - Binaries for windows can't be compiled.