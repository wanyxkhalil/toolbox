package mysqltogostruct

import (
	"log"
	"os"
)

func Run(src, dst string) {
	in, err := os.Open(src)
	if err != nil {
		log.Fatalf("Read file error：%v", err.Error())
		return
	}
	defer in.Close()

}
