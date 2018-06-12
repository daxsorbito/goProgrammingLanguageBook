// Test program for Rune
package main

import (
	"fmt"
)

func main() {
	ascii := 'a'
	unicode := 'œ'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]q\n", newline)
}
