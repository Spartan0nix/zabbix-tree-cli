package render

import (
	"encoding/json"

	"github.com/Spartan0nix/zabbix-tree-cli/internal/tree"
)

// OutputAsJson is used to render a TreeNode to a file in json
func OutputAsJson(tree *tree.TreeNode) ([]byte, error) {
	b, err := json.Marshal(tree)
	if err != nil {
		return nil, err
	}

	return b, nil
}
