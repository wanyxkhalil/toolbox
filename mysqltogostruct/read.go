package mysqltogostruct

import (
	"strings"
)

var typeMap = map[string]string{
	"int":                "int64",
	"integer":            "int64",
	"tinyint":            "int64",
	"smallint":           "int64",
	"mediumint":          "int64",
	"bigint":             "int64",
	"int unsigned":       "int64",
	"integer unsigned":   "int64",
	"tinyint unsigned":   "int64",
	"smallint unsigned":  "int64",
	"mediumint unsigned": "int64",
	"bigint unsigned":    "int64",
	"bit":                "int64",
	"bool":               "bool",
	"enum":               "string",
	"set":                "string",
	"varchar":            "string",
	"char":               "string",
	"tinytext":           "string",
	"mediumtext":         "string",
	"text":               "string",
	"longtext":           "string",
	"blob":               "string",
	"tinyblob":           "string",
	"mediumblob":         "string",
	"longblob":           "string",
	"date":               "time.Time",
	"datetime":           "time.Time",
	"timestamp":          "time.Time",
	"time":               "time.Time",
	"float":              "float64",
	"double":             "float64",
	"decimal":            "float64",
	"binary":             "string",
	"varbinary":          "string",
}

type Table struct {
	Name    string
	Columns []Column
	Comment string
}

type Column struct {
	Name    string
	Type    string
	Comment string
}

func structure(sql string) *Table {
	t := new(Table)
	t.Columns = []Column{}

	// clean
	s := strings.ToLower(sql)
	s = strings.TrimSpace(s)

	// table
	t.Name, t.Comment = getTableInfo(s)

	// column
	i1 := strings.Index(s, "(")
	i2 := strings.LastIndex(s, ")")
	split := strings.Split(s[i1+1:i2], ",")
	for _, cs := range split {
		// if not started with ` then break, that means key rows.
		cs = strings.TrimSpace(cs)
		if cs[0] != 96 {
			break
		}

		t.Columns = append(t.Columns, *getColumnInfo(cs))
	}

	return t
}

// getTableInfo get table name and comment
func getTableInfo(s string) (name, comment string) {
	i1 := strings.Index(s, "(")
	name = getName(s[0:i1])

	i2 := strings.LastIndex(s, ")")
	comment = getComment(s[i2+1:])

	return
}

// getColumnInfo get column info
func getColumnInfo(s string) (c *Column) {
	c = new(Column)
	c.Name = getName(s)
	c.Comment = getComment(s)

	// type
	split := strings.Fields(s)
	c.Type = split[1]

	// remove like (255)
	i := strings.Index(c.Type, "(")
	if i > 0 {
		c.Type = c.Type[0:i]
	}

	// append unsigned
	if split[2] == "unsigned" {
		c.Type = c.Type + " unsigned"
	}

	c.Type = typeMap[c.Type]

	return
}

// getName get name between back quote.
func getName(s string) string {
	i1 := strings.Index(s, "`")
	i2 := strings.LastIndex(s, "`")
	return s[i1+1 : i2]
}

// getComment get comment between single quote, last 1 and last 2.
func getComment(s string) string {
	var i1, i2 int
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == 39 {
			if i2 == 0 {
				i2 = i
			} else if i1 == 0 {
				i1 = i
				break
			}
		}
	}

	if i1 < 1 || i2 < 1 {
		return ""
	}
	return s[i1+1 : i2]
}
