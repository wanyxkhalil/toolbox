package mkpasswd

import (
	"fmt"
	"testing"
)

func BenchmarkGenerateString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateString(4,4,2,1)
	}
}

func GenerateSpecialChars() {
	for i := 33; i <= 47; i++ {
		fmt.Printf("%c",i)
	}
	for i := 58; i <= 64; i++ {
		fmt.Printf("%c", i)
	}
	fmt.Println()
}
