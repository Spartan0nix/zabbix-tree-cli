package render

import (
	"fmt"
	"os"

	"github.com/Spartan0nix/zabbix-tree-cli/internal/graph"
	"github.com/Spartan0nix/zabbix-tree-cli/internal/tree"
)

// WriteToFile is used to write the given slice of byte to a file.
func WriteToFile(file string, b []byte) error {
	err := os.WriteFile(file, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

// OutputTree is used to write the given slice of byte to the appropriate format.
// If the file argument is empty, print the output to the shell, otherwise write to the given file.
func OutputTree(file string, b []byte) error {
	var err error

	if file == "" {
		fmt.Println(string(b))
	} else {
		err = WriteToFile(file, b)
	}

	return err
}

func RenderTree(file string, format string, t tree.TreeNode, color bool) error {
	var err error
	var b []byte

	switch format {
	case "dot":
		buffer, err := graph.RenderDotGraph(t, color)
		if err != nil {
			return err
		}

		b = buffer.Bytes()

	case "shell":
		b = OutputInShell(&t)

	case "json":
		b, err = OutputAsJson(&t)
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("format '%s' is not supported", format)
	}

	err = OutputTree(file, b)
	if err != nil {
		return err
	}

	return nil
}
