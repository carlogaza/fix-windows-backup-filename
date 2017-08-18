package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var flagPath = flag.String("path", "", "path of file want to rename")
var currentDir = ""

func match(s string) string {
	i := strings.Index(s, " (2017")
	if i >= 0 {
		j := strings.Index(s, "UTC)")
		if j >= 0 {
			return s[i : j+4]
		}
	}
	return ""
}

func partialRename(path string, f os.FileInfo, err error) (e error) {
	if currentDir != filepath.Dir(path) {
		currentDir = filepath.Dir(path)
		fmt.Println("\n" + "--------------------------------------")
		fmt.Println("Go to directory " + currentDir + " ...")
		fmt.Println("--------------------------------------" + "\n")
	}

	if strings.Contains(f.Name(), " (2017") && strings.Contains(f.Name(), "UTC)") {
		base := filepath.Base(path) // file name
		dir := filepath.Dir(path)   // directory name

		r := match(base)
		renameTo := filepath.Join(dir, strings.Replace(base, r, "", 1))
		fmt.Println("Renaming on file " + path + " ...")
		os.Rename(path, renameTo)
		fmt.Println("File has been renamed to : " + renameTo)
	}
	return
}

func init() {
	flag.Parse()
}

func main() {
	if *flagPath == "" {
		flag.Usage()
		os.Exit(0)
	}

	filepath.Walk(*flagPath, partialRename)
}
