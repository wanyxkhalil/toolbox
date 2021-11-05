package mysqltogostruct

import (
	"regexp"
	"strings"
)

var typeMap = map[string]string{
	"tinyint":            "int8",
	"smallint":           "int16",
	"mediumint":          "int32",
	"int":                "int",
	"bigint":             "int64",
	"tinyint unsigned":   "uint8",
	"smallint unsigned":  "uint16",
	"mediumint unsigned": "uint32",
	"int unsigned":       "uint",
	"bigint unsigned":    "uint64",

	"float":                "float32",
	"double":               "float64",
	"decimal default null": "decimal.NullDecimal",
	"decimal not null":     "decimal.Decimal",

	"year":      "uint8",
	"time":      "time.Time",
	"date":      "time.Time",
	"datetime":  "time.Time",
	"timestamp": "time.Time",

	"char":       "string",
	"varchar":    "string",
	"tinytext":   "string",
	"text":       "string",
	"mediumtext": "string",
	"longtext":   "string",
	"enum":       "string",
	"set":        "string",

	"bit":        "byte",
	"binary":     "[]byte",
	"varbinary":  "[]byte",
	"tinyblob":   "[]byte",
	"blob":       "[]byte",
	"mediumblob": "[]byte",
	"longblob":   "[]byte",
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
	for i, cs := range split {
		cs = strings.TrimSpace(cs)

		// if decimal(10,2)，2 line combination
		endWithDigit := regexp.MustCompile(`.*\d$`)
		startWithDigit := regexp.MustCompile(`^\d.*`)

		if endWithDigit.MatchString(cs) {
			continue
		}
		if startWithDigit.MatchString(cs) {
			cs = strings.TrimSpace(split[i-1]) + cs
		}

		// if not started with ` then break, that means key rows.
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

	// decimal
	if c.Type == "decimal" {
		if strings.Contains(s, "default null") {
			c.Type = c.Type + " default null"
		} else if strings.Contains(s, "not null") {
			c.Type = c.Type + " not null"
		}
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
