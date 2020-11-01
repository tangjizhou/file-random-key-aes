package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	fmt.Println(filepath.Match("/a/b/c/c*", "/a/b/c/csds"))        // true
	fmt.Println(filepath.Match("/a/b/c/c？", "/a/b/c/csds"))        // false
	fmt.Println(filepath.Match("/a/b/c/[a-z]s*", "/a/b/c/csds"))   // true
	fmt.Println(filepath.Match("/a/b/c/[a-z]", "/a/b/c/cc"))       // false
	fmt.Println(filepath.Match("/a/b/c/[a-z][a-z]", "/a/b/c/cc"))  // true
	fmt.Println(filepath.Match("/a/b/c/[a-z][a-z]", "/a/b/c/ccc")) // false

}
