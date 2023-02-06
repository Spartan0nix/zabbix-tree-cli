#!/bin/bash

set -e

INSTALL_DIR="/usr/local/zabbix-tree-cli"

if [[ -d $INSTALL_DIR ]]
then
    echo "Removing directory '$INSTALL_DIR'"
    sudo rm -r $INSTALL_DIR
else
    echo "No '$INSTALL_DIR' directory found"
fi

echo ""
echo "--------------------------------------------------------------------------------------------------"
echo "'zabbix-tree-cli' was succesfuly uninstalled."
echo ""
echo "Don't forget to remove '$INSTALL_DIR' from your path if you added configuration in your ~/.bashrc or ~/.zshrc."
echo ""
echo "Thanks you for using this cli and don't forget to check out the documentation at 'https://github.com/Spartan0nix/zabbix-tree-cli#zabbix-tree-cli'"
echo "--------------------------------------------------------------------------------------------------"

exit 0