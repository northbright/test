package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	f := filepath.Join(os.TempDir(), "MY_APP", "README.md")
	fmt.Printf("f: %v\n", f)

	// Get parent dir of f via path.Dir.
	dir := path.Dir(f)
	fmt.Printf("get dir using path.Dir(), dir: %v\n", dir)

	dir = filepath.Dir(f)
	fmt.Printf("get dir using filepath.Dir(), dir: %v\n", dir)

}
