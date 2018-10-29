package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
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
			fmt.Printf("  File match found: %s\n", path)
			*matches = append(*matches, path)
		}
		return nil
	}
}

func writeToCSV(filename string, data []string) {
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()
	// header
	writer.Write([]string{"", "File"})
	// rows
	for row, value := range data {
		s := []string{strconv.Itoa(row + 1), value}
		err := writer.Write(s)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	fmt.Println(banner())
	root := flag.String("path", "/", "The directory to search")
	extension := flag.String("ext", "crypto", "The extension of the files to find")
	flag.Parse()
	fmt.Printf("Searching for .%s files in %s...\n", *extension, *root)
	matches := make([]string, 0)
	filepath.Walk(*root, walkWithExtraParams(*extension, &matches))
	fmt.Println("Search complete.")
	fmt.Println("Creating report...")
	t := time.Now()
	dateStr := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	filename := fmt.Sprintf("%s-c-report.csv", dateStr)
	writeToCSV(filename, matches)
	fmt.Println("Report created.")
}
