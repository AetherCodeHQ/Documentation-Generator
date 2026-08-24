package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var reFunc = regexp.MustCompile("^func[[:space:]]+([A-Za-z0-9_]+)")
var reType = regexp.MustCompile("^type[[:space:]]+([A-Za-z0-9_]+)")

func main() {
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}
	fmt.Println("# API Documentation\n")
	filepath.Walk(dir, func(p string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		if filepath.Ext(p) != ".go" || strings.HasSuffix(p, "_test.go") {
			return nil
		}
		f, err := os.Open(p)
		if err != nil {
			return nil
		}
		defer f.Close()
		fmt.Printf("## %s\n\n", filepath.ToSlash(p))
		sc := bufio.NewScanner(f)
		ln := 0
		for sc.Scan() {
			ln++
			line := strings.TrimSpace(sc.Text())
			if m := reFunc.FindStringSubmatch(line); m != nil {
				fmt.Printf("- **%s()** (func, line %d)\n", m[1], ln)
			} else if m := reType.FindStringSubmatch(line); m != nil {
				fmt.Printf("- **%s** (type, line %d)\n", m[1], ln)
			}
		}
		fmt.Println()
		return nil
	})
}
