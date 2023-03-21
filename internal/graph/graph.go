package graph

import (
	"bytes"
	"fmt"
)

// indent is used to add the required indentation level to the buffer.
// Example : indentation = 2 ; '\t\t' will be added to the buffer.
func indent(indentation *int, b *bytes.Buffer) error {
	for i := 0; i < *indentation; i++ {
		_, err := b.WriteString("\t")
		if err != nil {
			return err
		}
	}

	return nil
}

// writeString is used to write a string to the given buffer while respecting the given indentation level.
func writeString(indentation *int, b *bytes.Buffer, str string) error {
	// Add the required indentation before adding the string
	err := indent(indentation, b)
	if err != nil {
		return err
	}

	// Write the given string
	_, err = b.WriteString(fmt.Sprintf("%s\n", str))
	if err != nil {
		return nil
	}

	return nil
}
