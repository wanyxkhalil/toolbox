package mysqltogostruct

import (
	"bufio"
	"fmt"
	"github.com/wanyxkhalil/toolbox/util"
	"os"
	"path"
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

var importMap = map[string]string{
	"decimal.NullDecimal": "github.com/shopspring/decimal",
	"decimal.Decimal":     "github.com/shopspring/decimal",
	"time.Time":           "time",
}

type Table struct {
	Name    string
	Columns []Column
	Comment string
	Imports util.Set
}

type Column struct {
	Name    string
	Type    string
	Comment string
}

func (t *Table) toFile(dirPath string) {
	p := path.Join(dirPath, t.Name+".go")
	f, err := os.OpenFile(p, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	writer.WriteString(fmt.Sprintf("package %s\n\n", strings.ToLower(path.Base(dirPath))))

	// import
	if len(t.Imports) > 0 {
		writer.WriteString("import (\n")
		for s := range t.Imports {
			writer.WriteString(fmt.Sprintf("\t\"%s\"\n", s))
		}
		writer.WriteString(")\n\n")
	}

	// struct
	writer.WriteString(fmt.Sprintf("// %s %s\n", util.ToCamel(t.Name), t.Comment))
	writer.WriteString(fmt.Sprintf("type %s struct {\n", util.ToCamel(t.Name)))
	for _, column := range t.Columns {
		if len(column.Comment) > 0 {
			writer.WriteString(fmt.Sprintf("\t// %s\n", column.Comment))
		}
		writer.WriteString(fmt.Sprintf("\t%s\t%s\n", util.ToCamel(column.Name), column.Type))
	}
	writer.WriteString("}\n")

	writer.Flush()
}

func toTable(sql string) *Table {
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

	t.Imports = getImports(t.Columns)
	return t
}

func getImports(columns []Column) (s util.Set) {
	s = make(util.Set)
	for _, column := range columns {
		if v, ok := importMap[column.Type]; ok {
			s.Add(v)
		}
	}
	return
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
