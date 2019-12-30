package to_test

import (
	"fmt"
	"testing"

	"github.com/qianlnk/to"
)

func TestTo(t *testing.T) {
	var testVal = []interface{}{
		"abc",
		123,
		"1321",
		"11123.21",
		123.12,
		-12313,
		-123.32,
		"-123.12",
	}

	for _, v := range testVal {
		fmt.Println("---", v, "---")
		fmt.Println("string ", to.String(v))
		fmt.Println("int64  ", to.Int64(v))
		fmt.Println("float64", to.Float64(v))
	}
}

func BenchmarkTo(b *testing.B) {
	var testVal = []interface{}{
		"abc",
		123,
		"1321",
		"11123.21",
		123.12,
		-12313,
		-123.32,
		"-123.12",
	}

	for i := 0; i < b.N; i++ {
		for _, v := range testVal {
			to.String(v)
			//util.ToInt64(v)
			//util.ToFloat64(v)
		}
	}
}
