package utils

import (
	"blog/model/blogs"
	"testing"
)

func BenchmarkGenerateTokenByAccountSpeed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateTokenByAccount(&blogs.Account{Account: "5555555adgjahgfjhafgjah"})
	}
}
