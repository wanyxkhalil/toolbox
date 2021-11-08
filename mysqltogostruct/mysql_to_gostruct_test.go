package mysqltogostruct

import (
	"fmt"
	"testing"
)

func TestToTable(t *testing.T) {
	sql := readSource("./data/table.sql")
	fmt.Printf("%s\n\n", sql)

	table := toTable(sql)
	fmt.Printf("%+v\n", table)
}

func TestRun(t *testing.T) {
	//Run("/Users/khalil/Projects/github/toolbox/mysqltogostruct/data/table.sql", "/Users/khalil/Downloads")
}
