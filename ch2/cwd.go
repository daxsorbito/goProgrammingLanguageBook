// Test program for package declaration variables
package main

import (
	"fmt"
	"log"
	"os"
)

var cwd string

func init() {
	// var err error
	// cwd, err = os.Getwd()

	cwd, err := os.Getwd()

	if err != nil {
		log.Fatalf("os.Getwd failed: %v,", err)
	}
	// log.Printf("Working directory = %s", cwd)
}

func main() {
	fmt.Println("The value of cwd from Main is: %v", cwd)
}
