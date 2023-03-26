package graph

import (
	"bytes"
	"testing"
)

func TestIndent(t *testing.T) {
	var b bytes.Buffer
	indentation := 2

	err := indent(&indentation, &b)
	if err != nil {
		t.Fatalf("Error when executing indent function\nReason : %v", err)
	}

	expectedOutput := "\t\t"
	if b.String() != expectedOutput {
		t.Fatalf("Wrong output returned\nExpected : %s\nReturned : %s", expectedOutput, b.String())
	}
}

func TestIndentNoIndentation(t *testing.T) {
	var b bytes.Buffer
	indentation := 0

	err := indent(&indentation, &b)
	if err != nil {
		t.Fatalf("Error when executing indent function\nReason : %v", err)
	}

	if b.String() != "" {
		t.Fatalf("Wrong output returned\nExpected : %s\nReturned : %s", "", b.String())
	}
}

func TestWriteString(t *testing.T) {
	var b bytes.Buffer
	indentation := 1

	err := writeString(&indentation, &b, "random-string")
	if err != nil {
		t.Fatalf("Error when executing writeString function\nReason : %v", err)
	}

	expectedOutput := "\trandom-string\n"
	if b.String() != expectedOutput {
		t.Fatalf("Wrong output returned\nExpected : %s\nReturned : %s", expectedOutput, b.String())
	}
}

func TestWriteStringNoIndentation(t *testing.T) {
	var b bytes.Buffer
	indentation := 0

	err := writeString(&indentation, &b, "random-string")
	if err != nil {
		t.Fatalf("Error when executing writeString function\nReason : %v", err)
	}

	expectedOutput := "random-string\n"
	if b.String() != expectedOutput {
		t.Fatalf("Wrong output returned\nExpected : %s\nReturned : %s", expectedOutput, b.String())
	}
}
