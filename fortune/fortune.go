package fortune

import (
	"bufio"
	"fmt"
	"github.com/wanyxkhalil/toolbox/util"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

var rander = rand.New(rand.NewSource(time.Now().UnixNano()))

func Run(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 读取有多少条
	r := bufio.NewReader(f)
	lines := []int{0}
	var i int
	for {
		str, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}

		if strings.TrimSpace(str) == "%" {
			lines = append(lines, i)
		}
		i++
	}

	if len(lines) == 0 {
		panic("database is empty.")
	}

	// 随机选取第 x 条
	x := rander.Intn(len(lines) - 1)
	strs := util.ReadLines(path, lines[x], lines[x+1])

	// 打印区间中的文字
	for _, s := range strs {
		trimStr := strings.TrimSpace(s)
		if len(trimStr) == 0 || trimStr == "%" {
			continue
		}
		fmt.Printf("%s\n", s)
	}
}
