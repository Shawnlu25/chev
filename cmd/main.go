package main

import (
	"flag"
	"fmt"

	"github.com/Shawnlu25/chev/internal/event"
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
	notes, err := keep.ParseExportedNotesFromDir(path)

	if err != nil {
		panic(fmt.Errorf("Error parsing keep export: %v", err))
	}

	events, err := event.ParseEventsFromKeepExportedNotes(notes)

	if err != nil {
		panic(fmt.Errorf("Error parsing events: %v", err))
	}

	for _, event := range events {
		fmt.Println(event)
	}
}
