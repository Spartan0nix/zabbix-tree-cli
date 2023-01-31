package tree

import (
	"strings"

	zabbixgosdk "github.com/Spartan0nix/zabbix-go-sdk/v2"
	"github.com/goccy/go-graphviz/cgraph"
)

// TreeNode defined the structure of each node used to construct a tree
type TreeNode struct {
	Name            string       `json:"name"`
	GraphNode       *cgraph.Node `json:"graphNode,omitempty"`
	ParentGraphNode *cgraph.Node `json:"parentGraphNode,omitempty"`
	Childrens       []TreeNode   `json:"childrens,omitempty"`
}

// Flatten is used to returned a flat representation of a treeNode
func (t *TreeNode) Flatten() []TreeNode {
	out := make([]TreeNode, 0)

	// Add the root node as the first element
	out = append(out, *t)

	// Check of existing childrens
	if len(t.Childrens) > 0 {
		// Execute the function for each childrens
		for _, c := range t.Childrens {
			out = append(out, c.Flatten()...)
		}
	}

	return out
}

// SearchNode is use to search for a TreeNode with the given name in a list of TreeNode
func SearchNode(name string, nodes []TreeNode) *TreeNode {
	for len(nodes) > 0 {
		if nodes[0].Name == name {
			return &nodes[0]
		}

		nodes = nodes[1:]
	}

	return nil
}

// CreateChildren is used to creat a new TreeNode, add it to the current node children and returned the representation of the newly created node.
func (t *TreeNode) CreateChildren(name string, graph *cgraph.Graph) (*TreeNode, error) {
	n := TreeNode{Name: name}

	if graph != nil {
		var err error
		n.GraphNode, err = graph.CreateNode(n.Name)
		if err != nil {
			return nil, err
		}

		n.ParentGraphNode = t.GraphNode
	}

	t.Childrens = append(t.Childrens, n)

	return &n, nil
}

// CreateEdge is used to create an edge between the current node and is parent node.
func (t *TreeNode) CreateEdge(graph *cgraph.Graph) (*cgraph.Graph, error) {
	_, err := graph.CreateEdge("", t.GraphNode, t.ParentGraphNode)
	if err != nil {
		return nil, err
	}

	return graph, nil
}

// addNode is used to add a new node the current node.
// If there is no more part to work on after adding the node, the root node pointer will be returned.
func addNode(root *TreeNode, node *TreeNode, graph *cgraph.Graph, parts []string) (*TreeNode, []string, error) {
	// Create a new children node for the current node
	n, err := node.CreateChildren(parts[0], graph)
	if err != nil {
		return nil, nil, err
	}

	// Remove the first part (which has been worked on already) from the list
	parts = parts[1:]

	// If there is no more parts to work on, return to the root node
	if len(parts) == 0 {
		node = root
	} else {
		// Else, go one level deeper by using the newly added children node as the current node
		node = n
	}

	return node, parts, nil
}

// GenerateHostGroupTree is used to generate an complete tree representation for the given Zabbix Host Groups.
func (t *TreeNode) GenerateHostGroupTree(groups []*zabbixgosdk.HostGroup, graph *cgraph.Graph) error {
	// Start from the root node (t)
	node := t
	var err error

	// Loop over each HostGroup
	// Zabbix support nested group by using '/' between each part
	// For exemple : Templates/Databases/MySQL
	// -> Templates
	// 		-> Databases
	// 			-> MySQL
	for _, group := range groups {
		// Explode each parts of the Hostgroup
		parts := strings.Split(group.Name, "/")

		// Only one part present which means, the hostGroup is a direct children of the root node
		if len(parts) == 1 {
			_, err := node.CreateChildren(parts[0], graph)
			if err != nil {
				return err
			}

			// Restart from the root node (t) for the next hostGroup
			node = t
			continue
		}

		// Loop over each part of the HostGroup
		for len(parts) > 0 {
			// Retrieve the current node childrens
			childrens := node.Childrens

			// If the current node has no children we don't have to search for the current part, it will not exist
			if len(childrens) == 0 {
				// Add a new node
				node, parts, err = addNode(t, node, graph, parts)
				if err != nil {
					return err
				}

				continue
			}

			// Search for an existing node with the given name and return is ptr
			node_ptr := SearchNode(parts[0], childrens)

			// If no node was found, create it
			if node_ptr == nil {
				node, parts, err = addNode(t, node, graph, parts)
				if err != nil {
					return err
				}

				continue
			} else {
				// Remove the part from the list since a node already exist for it
				parts = parts[1:]

				// Go one level deeper by using the existing node as the current node
				node = node_ptr
			}
		}
	}

	return nil
}

// GenerateTreeEdges is used to generate edges between node for a list of node.
func GenerateTreeEdges(graph *cgraph.Graph, nodes []TreeNode) (*cgraph.Graph, error) {
	var err error

	for _, n := range nodes {
		if n.ParentGraphNode != nil {
			graph, err = n.CreateEdge(graph)
			if err != nil {
				return nil, err
			}
		}
	}

	return graph, nil
}
