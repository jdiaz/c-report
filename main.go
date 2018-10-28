package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func visit(path string, f os.FileInfo, err error) error {
	parts := strings.Split(path, ".")
	if len(parts) != 2 {
		return nil
	}
	if parts[1] == "crypto" {
		fmt.Printf("File match in: %s\n", path)
	}
	return nil
}

func main() {
	file, err := ioutil.ReadFile("banner.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	banner := string(file)
	fmt.Println(banner)
	fmt.Printf("Searching for .crypto files...\n")
	flag.Parse()
	root := flag.Arg(0)
	filepath.Walk(root, visit)
	fmt.Println("Search complete.")
}
