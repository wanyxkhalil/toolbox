package mysqltogostruct

import (
	"io"
	"os"
	"strings"
)

func Run(src, dst string) {
	s := readSource(src)
	arr := splitSQL(s)

	for _, sql := range arr {
		table := toTable(sql)
		table.toFile(dst)
	}
}

func readSource(p string) string {
	f, err := os.Open(p)
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

func splitSQL(s string) (arr []string) {
	arr = strings.Split(s, ";")
	i := 0
	for _, sql := range arr {
		s = strings.TrimSpace(sql)
		if len(s) > 0 {
			arr[i] = s
			i++
		}
	}
	return arr[:i]
}
