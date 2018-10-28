package main

import (
	"encoding/csv"
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

func walkWithExtraParams(extension string, matches *[]string) filepath.WalkFunc {
	return func(path string, f os.FileInfo, err error) error {
		parts := strings.Split(path, ".")
		n := len(parts)
		if parts[n-1] == extension {
			fmt.Printf("File match found: %s\n", path)
			*matches = append(*matches, path)
		}
		return nil
	}
}

func main() {
	fmt.Println(banner())
	flag.Parse()
	root := flag.Arg(0)
	extension := flag.Arg(1)
	fmt.Printf("Searching for .%s files in %s...\n", extension, root)
	matches := make([]string, 0)
	filepath.Walk(root, walkWithExtraParams(extension, &matches))
	fmt.Println("Search complete.")
	fmt.Println("Creating report")
	file, err := os.Create("creport.csv")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, value := range matches {
		s := make([]string, 0)
		s = append([]string{value})
		err := writer.Write(s)
		if err != nil {
			fmt.Println(err)
		}
	}
}
