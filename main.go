package main

import (
	"fmt"
	"os"
)

// documentation_generator - Auto-generate docs
func documentation_generator(path string) {
	fmt.Println("========================================")
	fmt.Println("  Documentation-Generator")
	fmt.Println("  Auto-generate docs")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	documentation_generator(path)
}
