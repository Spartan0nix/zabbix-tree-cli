package graph

import (
	"bytes"
	"fmt"
	"math/rand"

	"github.com/Spartan0nix/zabbix-tree-cli/internal/tree"
	"github.com/lucasb-eyer/go-colorful"
)

// writeDotNode is used to write the representation of a node
// If the given TreeNode as a parentId set, add the connection between the two
func writeDotNode(indentation *int, b *bytes.Buffer, t *tree.TreeNode) error {
	// Write a label for the current node
	err := writeString(indentation, b, fmt.Sprintf("\"%s\" [label=\"%s\"];", t.Id, t.Name))
	if err != nil {
		return err
	}

	// Write the connection to the parent
	if t.ParentId != "" {
		err = writeString(indentation, b, fmt.Sprintf("\"%s\" -> \"%s\";", t.ParentId, t.Id))
		if err != nil {
			return err
		}
	}

	return nil
}

// writeSubGraph is used to write a subgraph in the buffer representing the given TreeNode
// If the TreeNode has child nodes, the function will be executed on them
func writeDotSubGraph(indentation *int, b *bytes.Buffer, t *tree.TreeNode, color bool) error {

	// If the node as no child, don't add a subgraph, simply a node representation
	if len(t.Childrens) == 0 {
		err := writeDotNode(indentation, b, t)
		if err != nil {
			return err
		}

		return nil
	}

	// Open the subgraph
	err := writeString(indentation, b, fmt.Sprintf("subgraph cluster%s {", t.Id))
	if err != nil {
		return err
	}

	// Increase the indentation level since we are writting inside the subgraph body
	*indentation++

	if color {
		// Add a random color to the graph
		// That color will be inherited by node in the subgraph
		bgColor := colorful.Hcl(rand.Float64()*360.0, rand.Float64(), 0.6+rand.Float64()*0.4)
		err = writeString(indentation, b, fmt.Sprintf("bgcolor=\"%s\"", bgColor.Hex()))
		if err != nil {
			return err
		}
	}

	err = writeDotNode(indentation, b, t)
	if err != nil {
		return err
	}

	for _, c := range t.Childrens {
		err := writeDotSubGraph(indentation, b, c, color)
		if err != nil {
			return err
		}
	}

	// Decrease the indentation level since we are going to close the subgraph body
	*indentation--

	// Close the subgraph
	err = writeString(indentation, b, "}")
	if err != nil {
		return err
	}

	return nil
}

// // writeGraphConfiguration is used to write the given config to a graph block
// func writeDotGraphConfiguration(indentation *int, b *bytes.Buffer, config string) error {
// 	// Remove shape from the node.
// 	// The shape of each subgraph will be used instead
// 	err := writeString(indentation, b, fmt.Sprintf("graph %s;", config))
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// RenderDotGraph is used to render the given Tree to a dot representation
func RenderDotGraph(t tree.TreeNode, color bool) (*bytes.Buffer, error) {
	var buffer bytes.Buffer
	indentation := 0

	// Open the graph
	err := writeString(&indentation, &buffer, "digraph {")
	if err != nil {
		return nil, err
	}

	// err = writeDotGraphConfiguration(&indentation, &buffer, "[style=filled,color=white]")
	// if err != nil {
	// 	return nil
	// }

	indentation++

	err = writeDotSubGraph(&indentation, &buffer, &t, color)
	if err != nil {
		return nil, err
	}

	// Close the graph
	err = writeString(&indentation, &buffer, "}")
	if err != nil {
		return nil, err
	}

	return &buffer, nil
}
