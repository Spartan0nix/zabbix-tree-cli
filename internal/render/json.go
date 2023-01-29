package render

import (
	"encoding/json"
	"os"

	"github.com/Spartan0nix/zabbix-usergroup-tree/internal/tree"
)

// OutputAsJson is used to render a TreeNode to a file in json
func OutputAsJson(path string, tree *tree.TreeNode) error {
	b, err := json.Marshal(tree)
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write(b)
	if err != nil {
		return err
	}

	return nil
}
