#!/bin/bash

set -e

TMP_INSTALL_DIR="/tmp"
TMP_INSTALL_ARCHIVE_PATH="$TMP_INSTALL_DIR/zabbix-tree-cli.tar.gz"
INSTALL_DIR="/usr/local/zabbix-tree-cli"

# Retrieve the list of release objects from github API
payload=$(curl --silent https://api.github.com/repos/Spartan0nix/zabbix-tree-cli/releases)

echo "Listing available version"
# Retrieve all tag_name from the available releases
releases=$(echo $payload | jq -r '.[].tag_name')
# Replace space by newline to render a better ouput in the shell
echo $releases | tr '[[:space:]]' '\n'

release_tag=""
re=v[0-9]\.[0-9]\.[0-9]
# Loop while user as not provided a proper release format (vX.Y.Z)
while [[ ! "$release_tag" =~ $re ]]
do
    echo "Select one version for the list above :"
    read release_tag
done

# Extract the selected release
selected_release=$(echo $payload | jq --arg v $release_tag '.[] | select(.tag_name == $v)')
if [[ $selected_release == "" ]]
then
    echo "No release '$release_tag' found"
    exit 1
fi

# Retrive the asset URL
asset_url=$(echo $selected_release | jq -r '.assets[] | select(.name | test("tar.gz")) | .browser_download_url')
if [[ $asset_url == "" ]]
then
    echo "No assets found for release '$release_tag'"
    exit 1
fi

# Check for required directory
echo ""
echo "Checking if '$INSTALL_DIR' exist"
if [[ ! -d $INSTALL_DIR ]]
then
    echo "Creating directory now"
    sudo mkdir $INSTALL_DIR
else
    echo "Directory already exist"
fi

echo ""
echo "Downloading asset to '$TMP_INSTALL_ARCHIVE_PATH'"
wget $asset_url -O $TMP_INSTALL_ARCHIVE_PATH -q --show-progress

echo ""
echo "Extracting archive in '$INSTALL_DIR'"
sudo tar -C $INSTALL_DIR -xzf $TMP_INSTALL_ARCHIVE_PATH

echo ""
echo "Removing archive '$TMP_INSTALL_ARCHIVE_PATH'"
rm $TMP_INSTALL_ARCHIVE_PATH

echo ""
echo "--------------------------------------------------------------------------------------------------"
echo "Release '$release_tag' of the 'zabbix-tree-cli' was succesfuly installed to '$INSTALL_DIR'."
echo ""
echo "Don't forget to add '$INSTALL_DIR' to your path :"
echo "$ export PATH=\$PATH:$INSTALL_DIR"
echo ""
echo "Thanks you for using this cli and don't forget to check out the documentation at 'https://github.com/Spartan0nix/zabbix-tree-cli#zabbix-tree-cli'"
echo "--------------------------------------------------------------------------------------------------"

exit 0
