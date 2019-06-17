package main

import (
	"fmt"
	"github.com/teirm/learning_go/ch_2/popcount"
	"time"
)

func main() {

	var testValue uint64 = 1231412

	start := time.Now()
	fmt.Printf("PopCount: %d\n", popcount.PopCount(testValue))
	fmt.Printf("%.2ds elapsed\n", time.Since(start).Nanoseconds())

	start = time.Now()
	fmt.Printf("PopCount2: %d\n", popcount.PopCount2(testValue))
	fmt.Printf("%.2ds elapsed\n", time.Since(start).Nanoseconds())

	start = time.Now()
	fmt.Printf("PopCount3: %d\n", popcount.PopCount3(testValue))
	fmt.Printf("%.2ds elapsed\n", time.Since(start).Nanoseconds())

	start = time.Now()
	fmt.Printf("PopCount4: %d\n", popcount.PopCount4(testValue))
	fmt.Printf("%.2ds elapsed\n", time.Since(start).Nanoseconds())
}
