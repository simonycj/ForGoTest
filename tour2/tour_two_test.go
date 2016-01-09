package tour2

import (
	//	"fmt"
	"testing"
)

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Demo_basic_loop()
	}
}
