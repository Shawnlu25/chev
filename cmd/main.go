package main

import (
	"flag"
	"fmt"

	"github.com/Shawnlu25/chev/internal/keep"
)

var path string

func init() {
	const (
		defaultPath = "./Keep"
		pathUsage   = "Path to Google Keep export"
	)
	flag.StringVar(&path, "p", defaultPath, pathUsage)
}

func main() {
	flag.Parse()
	_, err := keep.ParseExportedNotesFromDir(path)

	if err != nil {
		panic(fmt.Errorf("Error parsing keep export: %v", err))
	}

	//for _, note := range notes {
	//fmt.Println(note)
	//}
}
