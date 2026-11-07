package humanduration_test

import (
	"fmt"
	"time"

	"github.com/inkochetkov/humanduration"
)

func ExampleFormat() {
	fmt.Println(humanduration.Format(25 * time.Hour))
	fmt.Println(humanduration.Format(90 * time.Second))
	fmt.Println(humanduration.Format(500 * time.Millisecond))
	// Output:
	// 1 day 1 hour
	// 1 minute 30 seconds
	// 500 milliseconds
}

func ExampleFormatShort() {
	fmt.Println(humanduration.FormatShort(90 * time.Second))
	fmt.Println(humanduration.FormatShort(9 * 24 * time.Hour))
	// Output:
	// 1m 30s
	// 1w 2d
}