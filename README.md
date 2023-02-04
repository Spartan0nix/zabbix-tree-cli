# zabbix-tree-cli

## Table of contents

- [Description](#Description)
- [Zabbix-nested-groups](#Zabbix-nested-groups)
- [Usage](#Usage)

## Description

This CLI tool is used to help administrator keeps track of their Zabbix Host Groups structure by rendering a graphical output (PNG, JPG, SVG, json, shell).

## Zabbix-nested-groups

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

An json export is available in the *fixtures* folder.
This export will create a host (fixture-host) and the associate host groups to build an example structure.

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
