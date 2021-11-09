package mdtoc

import (
	"bufio"
	"fmt"
	"github.com/wanyxkhalil/toolbox/util"
	"io"
	"os"
	"strings"
)

type header struct {
	level int
	name  string
	line  int
}

func Run(p string) {
	headers := getHeaders(p)
	printHeaders(headers)
}

func printHeaders(headers []header) {
	fmt.Printf("\nTable of Contents\n=================\n\n")
	for _, h := range headers {
		prefix := strings.Repeat("    ", h.level-1)
		fmt.Printf("%s* [%s](#%s)\n", prefix, h.name, convertName(h.name))
	}
	fmt.Println()
}

func getHeaders(p string) []header {
	f, err := os.Open(p)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	i := 1
	var headers []header
	var lines []int

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		if strings.HasPrefix(str, "#") {
			ss := strings.SplitN(str, " ", 2)
			if len(ss) == 2 {
				i := header{
					level: len(ss[0]),
					name:  strings.TrimSpace(ss[1]),
					line:  i,
				}
				headers = append(headers, i)
			}
		}

		if strings.HasPrefix(str, "```") {
			lines = append(lines, i)
		}

		i++
	}

	// 清除代码块中的 header
	j := 0
	for _, a := range headers {
		firstGreater := util.BinarySearchFirstGreater(lines, a.line)
		if firstGreater < 0 || firstGreater%2 == 0 {
			headers[j] = a
			j++
		}
	}

	return headers[:j]
}

func convertName(s string) string {
	s = strings.ToLower(s)
	return strings.ReplaceAll(s, " ", "-")
}
