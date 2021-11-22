package util

import (
	"fmt"
	"testing"
)

func TestReadLine(t *testing.T) {
	lines := ReadLines("/Users/khalil/Projects/github/toolbox/fortune/data/data.txt", 5, 9)
	fmt.Printf("%v", lines)
}
