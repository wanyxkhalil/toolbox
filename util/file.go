package util

import (
	"bufio"
	"io"
	"os"
)

func ReadLines(name string, start, end int) (lines []string) {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	var i int
	for {
		str, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if i < start {
		} else if i <= end {
			lines = append(lines, str)
		}

		i++
	}

	return
}
