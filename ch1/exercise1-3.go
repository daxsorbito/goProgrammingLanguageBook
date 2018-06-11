// Expriment to measure the difference in running time between our potentially inefficient versions
// and the one that uses strings.Join (Section 1.6 illustrates part of the time package, and Secion 11.4
// shows how to write benchmark test for systematic perfomance evaluation.)
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Printf("Result for Range%v\n", usingRange(os.Args[1:]))
	fmt.Printf("%d elapsed for Range\n", time.Since(start).Nanoseconds())
	start = time.Now()
	fmt.Printf("Result for Join%v\n", usingJoin(os.Args[1:]))
	fmt.Printf("%d elapsed for Join\n", time.Since(start).Nanoseconds())
}

func usingRange(args []string) string {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	return s
}

func usingJoin(args []string) string {
	return strings.Join(args, " ")
}
