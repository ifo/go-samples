package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	filetype string
	verbose  bool
)

func init() {
	flag.StringVar(&filetype, "filetype", "go", "The filetype to search for")
	flag.BoolVar(&verbose, "v", false, "Show lots of information")
	flag.Parse()
}

func main() {
	// Receive input for which file type to find.
	// init does this

	// Search current directory and subdirectories for all files of that type.
	numberOfFiles := Search(filetype)

	// Print the number of files found.
	fmt.Printf("%d files found with the extension %s.\n", numberOfFiles, filetype)
}

func Search(ftype string) int {
	fileCount := 0

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			if verbose {
				fmt.Printf("There was an issue accessing %s: %v\n", path, err)
			}
			return err
		}

		if strings.HasSuffix(info.Name(), ftype) {
			if verbose {
				fmt.Printf("Found file with filetype %s: %s\n", ftype, path)
			}
			fileCount++
		}

		// No problems here.
		return nil
	})

	return fileCount
}
