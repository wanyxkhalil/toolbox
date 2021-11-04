package mysqltogostruct

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func readFile() string {
	f, err := os.Open("./data/table.sql")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func TestStructure(t *testing.T) {
	sql := readFile()
	fmt.Printf("%s\n\n", sql)

	table := structure(sql)
	fmt.Printf("%+v\n", table)
}
