package main

import (
	"fmt"
	"math"
	"os"
	"path"
	"path/filepath"
)

func convertBytes(bytes int64) string {
	var result float64
	var unit string

	floated := float64(bytes)
	units := []string{"Bytes", "KB", "MB", "GB"}

	if floated != 0 {
		base := math.Log(floated) / math.Log(1024)
		floor := math.Floor(base)
		unit = units[int(floor)]
		result = math.Pow(1024, base-floor)
	} else {
		result = 0
		unit = units[0]
	}

	humanReadable := fmt.Sprintf("%.2f %s", result, unit)
	return humanReadable
}

func recursiveWalk(rootpath string) ([]string, error) {
	var files []string
	err := filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func getHelp() {
	_, name := path.Split(os.Args[0])
	fmt.Printf("%s (total size): util for count total size of files inside directory\n", name)
	fmt.Printf("Usage: %s <path>\n", name)
}

func main() {
	var totalSize int64
	var path string

	if len(os.Args) == 1 {
		fmt.Println("No path given, working in current directory")
		path = "./"
	} else if help := os.Args[1]; help == "-h" || help == "--help" {
		getHelp()
		os.Exit(0)
	} else if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
		fmt.Println("Unknown argument or given path does not exists")
		os.Exit(1)
	} else {
		path = os.Args[1]
	}

	files, _ := recursiveWalk(path)

	for _, file := range files {
		fi, _ := os.Stat(file)
		// fmt.Printf("%d\t%s\n", fi.Size(), file)
		totalSize += fi.Size()
	}

	fmt.Printf("%s\t%s\n", convertBytes(totalSize), path)
}
