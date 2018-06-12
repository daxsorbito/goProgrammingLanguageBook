// Test program for octal and hex
package main

import "fmt"

func main() {
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // [1] is index of args for Printf, # tells Printf to emmit a 0, 0x or 0X
	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
}
