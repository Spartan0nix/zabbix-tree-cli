package render

import (
	"os"
	"testing"

	"github.com/Spartan0nix/zabbix-tree-cli/internal/tree"
)

func TestWriteToFile(t *testing.T) {
	err := WriteToFile("test.json", []byte{})
	if err != nil {
		t.Fatalf("Error when executing WriteToFile function\nReason : %v", err)
	}

	err = os.Remove("test.json")
	if err != nil {
		t.Fatalf("Error when trying to remove the file created during this test ('test.json')")
	}
}

func TestOutputTree(t *testing.T) {
	// Hide shell output
	// Stderr while still be display
	os.Stdout = nil

	err := OutputTree("", []byte{})
	if err != nil {
		t.Fatalf("Error when executing OutputTree function\nReason : %v", err)
	}
}

func TestOutputTreeToFile(t *testing.T) {
	err := OutputTree("test.json", []byte{})
	if err != nil {
		t.Fatalf("Error when executing OutputTree function\nReason : %v", err)
	}

	err = os.Remove("test.json")
	if err != nil {
		t.Fatalf("Error when trying to remove the file created during this test ('test.json')")
	}
}

func TestRenderTreeDot(t *testing.T) {
	// Hide shell output
	// Stderr while still be display
	os.Stdout = nil

	err := RenderTree("", "dot", tree.TreeNode{
		Name: "test-root",
		Id:   "123456789",
	}, false)

	if err != nil {
		t.Fatalf("Error when executing RenderTree function\nReason : %v", err)
	}
}

func TestRenderTreeDotColor(t *testing.T) {
	// Hide shell output
	// Stderr while still be display
	os.Stdout = nil

	err := RenderTree("", "dot", tree.TreeNode{
		Name: "test-root",
		Id:   "123456789",
	}, true)

	if err != nil {
		t.Fatalf("Error when executing RenderTree function\nReason : %v", err)
	}
}

func TestRenderTreeDotToFile(t *testing.T) {
	err := RenderTree("test.dot", "dot", tree.TreeNode{
		Name: "test-root",
		Id:   "123456789",
	}, false)

	if err != nil {
		t.Fatalf("Error when executing RenderTree function\nReason : %v", err)
	}

	err = os.Remove("test.dot")
	if err != nil {
		t.Fatalf("Error when trying to remove the file created during this test ('test.dot')")
	}
}

func TestRenderTreeDotToFileColor(t *testing.T) {
	err := RenderTree("test.dot", "dot", tree.TreeNode{
		Name: "test-root",
		Id:   "123456789",
	}, true)

	if err != nil {
		t.Fatalf("Error when executing RenderTree function\nReason : %v", err)
	}

	err = os.Remove("test.dot")
	if err != nil {
		t.Fatalf("Error when trying to remove the file created during this test ('test.dot')")
	}
}

func TestRenderTreeShell(t *testing.T) {
	// Hide shell output
	// Stderr while still be display
	os.Stdout = nil

	err := RenderTree("", "shell", tree.TreeNode{
		Name: "test-root",
		Id:   "123456789",
	}, false)

	if err != nil {
		t.Fatalf("Error when executing RenderTree function\nReason : %v", err)
	}
}

func TestRenderTreeShellColor(t *testing.T) {
	// Hide shell output
	// Stderr while still be display
	os.Stdout = nil

	err := RenderTree("", "shell", tree.TreeNode{
		Name: "test-root",
		Id:   "123456789",
	}, true)

	if err != nil {
		t.Fatalf("Error when executing RenderTree function\nReason : %v", err)
	}
}

func TestRenderTreeShellToFile(t *testing.T) {
	err := RenderTree("test.txt", "shell", tree.TreeNode{
		Name: "test-root",
		Id:   "123456789",
	}, false)

	if err != nil {
		t.Fatalf("Error when executing RenderTree function\nReason : %v", err)
	}

	err = os.Remove("test.txt")
	if err != nil {
		t.Fatalf("Error when trying to remove the file created during this test ('test.txt')")
	}
}

func TestRenderTreeShellToFileColor(t *testing.T) {
	err := RenderTree("test.txt", "shell", tree.TreeNode{
		Name: "test-root",
		Id:   "123456789",
	}, true)

	if err != nil {
		t.Fatalf("Error when executing RenderTree function\nReason : %v", err)
	}

	err = os.Remove("test.txt")
	if err != nil {
		t.Fatalf("Error when trying to remove the file created during this test ('test.txt')")
	}
}

func TestRenderTreeJson(t *testing.T) {
	// Hide shell output
	// Stderr while still be display
	os.Stdout = nil

	err := RenderTree("", "json", tree.TreeNode{
		Name: "test-root",
		Id:   "123456789",
	}, false)

	if err != nil {
		t.Fatalf("Error when executing RenderTree function\nReason : %v", err)
	}
}

func TestRenderTreeJsonColor(t *testing.T) {
	// Hide shell output
	// Stderr while still be display
	os.Stdout = nil

	err := RenderTree("", "json", tree.TreeNode{
		Name: "test-root",
		Id:   "123456789",
	}, true)

	if err != nil {
		t.Fatalf("Error when executing RenderTree function\nReason : %v", err)
	}
}

func TestRenderTreeJsonToFile(t *testing.T) {
	err := RenderTree("test.json", "json", tree.TreeNode{
		Name: "test-root",
		Id:   "123456789",
	}, false)

	if err != nil {
		t.Fatalf("Error when executing RenderTree function\nReason : %v", err)
	}

	err = os.Remove("test.json")
	if err != nil {
		t.Fatalf("Error when trying to remove the file created during this test ('test.json')")
	}
}

func TestRenderTreeJsonToFileColor(t *testing.T) {
	err := RenderTree("test.json", "json", tree.TreeNode{
		Name: "test-root",
		Id:   "123456789",
	}, true)

	if err != nil {
		t.Fatalf("Error when executing RenderTree function\nReason : %v", err)
	}

	err = os.Remove("test.json")
	if err != nil {
		t.Fatalf("Error when trying to remove the file created during this test ('test.json')")
	}
}
