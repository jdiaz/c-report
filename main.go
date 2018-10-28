package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func banner() string {
	return `
   ____   ____                       _    
  / ___| |  _ \ ___ _ __   ___  _ __| |_  
 | |     | |_) / _ \ '_ \ / _ \| '__| __| 
 | |___  |  _ <  __/ |_) | (_) | |  | |_  
  \____| |_| \_\___| .__/ \___/|_|   \__| 
                   |_|                    
 `
}

func visit(path string, f os.FileInfo, err error) error {
	parts := strings.Split(path, ".")
	n := len(parts)
	if parts[n-1] == "crypto" {
		fmt.Printf("File match in: %s\n", path)
	}
	return nil
}

func main() {
	fmt.Println(banner())
	fmt.Printf("Searching for .crypto files...\n")
	flag.Parse()
	root := flag.Arg(0)
	filepath.Walk(root, visit)
	fmt.Println("Search complete.")
}
