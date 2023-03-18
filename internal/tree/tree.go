package tree

import (
	"crypto/rand"
	"encoding/hex"
)

// TreeNode defined the structure of each node used to construct a tree
type TreeNode struct {
	Name      string      `json:"name"`
	Id        string      `json:"id,omitempty"`
	ParentId  string      `json:"parentId,omitempty"`
	Childrens []*TreeNode `json:"childrens,omitempty"`
}

func GenerateNewHash(length int8) (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

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

func searchNode(name string, nodes []*TreeNode) *TreeNode {
	for len(nodes) > 0 {
		if nodes[0].Name == name {
			return nodes[0]
		}

		nodes = nodes[1:]
	}

	return nil
}

func (t *TreeNode) createChildren(name string) (*TreeNode, error) {
	hash, err := GenerateNewHash(20)
	if err != nil {
		return nil, err
	}

	n := TreeNode{
		Name:     name,
		Id:       hash,
		ParentId: t.Id,
	}

	t.Childrens = append(t.Childrens, &n)

	return &n, nil
}
