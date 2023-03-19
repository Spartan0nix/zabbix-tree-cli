package graph

import (
	"bytes"
	"fmt"
	"math/rand"

	"github.com/Spartan0nix/zabbix-tree-cli/internal/tree"
	"github.com/lucasb-eyer/go-colorful"
)

// writeDotNode is used to write the representation of a node
func writeDotNode(indentation *int, b *bytes.Buffer, t *tree.TreeNode) error {
	// Build the default node config
	nodeConfig := fmt.Sprintf("\"%s\" [label=\"%s\",shape=box];", t.Id, t.Name)

	err := writeString(indentation, b, nodeConfig)
	if err != nil {
		return err
	}

	return nil
}

// writeConnections is used to write connections between the given node and it's children
func writeDotConnections(indentation *int, b *bytes.Buffer, t *tree.TreeNode) error {
	nodes := t.Childrens

	for len(nodes) > 0 {
		err := writeString(indentation, b, fmt.Sprintf("\"%s\" -> \"%s\";", t.Id, nodes[0].Id))
		if err != nil {
			return err
		}

		nodes = nodes[1:]
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
		// That color will be inherited by direct node in the subgraph
		bgColor := colorful.Hcl(rand.Float64()*360.0, rand.Float64(), 0.6+rand.Float64()*0.4)
		err = writeString(indentation, b, fmt.Sprintf("bgcolor=\"%s\"", bgColor.Hex()))
		if err != nil {
			return err
		}
	}

	// Write a node representation in the subgraph instead of simply adding a label to the subgraph
	err = writeDotNode(indentation, b, t)
	if err != nil {
		return err
	}

	// Write the connections between the node representing the current subgraph and the (future) nodes representing the children
	err = writeDotConnections(indentation, b, t)
	if err != nil {
		return err
	}

	// Execute the same process for every children of the current node
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

// writeGraphConfiguration is used to write the given config to a graph block
func writeDotGraphConfiguration(indentation *int, b *bytes.Buffer, configs []string) error {
	for _, c := range configs {
		err := writeString(indentation, b, c)
		if err != nil {
			return err
		}
	}

	return nil
}

// RenderDotGraph is used to render the given Tree to a dot representation
func RenderDotGraph(t tree.TreeNode, color bool) (*bytes.Buffer, error) {
	var buffer bytes.Buffer
	indentation := 0

	// Open the graph
	err := writeString(&indentation, &buffer, "digraph {")
	if err != nil {
		return nil, err
	}

	graphConfig := []string{
		// Define the layout
		"layout=\"dot\";",
		// Increase the separator from 0.5 to 1
		"ranksep=\"1.0 equally\";",
	}

	if !color {
		// If colors are not needed, hide the default graph and subgraph shape line
		graphConfig = append(graphConfig, "graph[style=filled,color=white];")
	}

	// Write the default graph configuration
	err = writeDotGraphConfiguration(&indentation, &buffer, graphConfig)
	if err != nil {
		return nil, err
	}

	// We are now writting in the graph block body
	indentation++

	// Start building the subgraph hierarchy with the 'root' node
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
