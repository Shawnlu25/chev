package event

import (
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/Shawnlu25/chev/internal/keep"
)

var timeRegexPattern *regexp.Regexp

func init() {
	timeRegexPattern = regexp.MustCompile(`([01][0-9]|2[0-3])[0-5][0-9]`)
}

// ParseEventsFromKeepExportedNotes ...
func ParseEventsFromKeepExportedNotes(notes []keep.ExportedNote) ([]Event, error) {
	events := []Event{}
	for _, note := range notes {
		eventsFromNote, err := parseEventsFromKeepExportedNote(note)
		if err != nil {
			return nil, err
		}
		events = append(events, eventsFromNote...)
	}
	return events, nil
}

func parseEventsFromKeepExportedNote(note keep.ExportedNote) ([]Event, error) {
	// Check if title can be parsed to date
	_, err := time.Parse("20060102", note.Title)
	if err != nil {
		return nil, nil
	}

	events := []Event{}

	for _, line := range strings.Split(note.TextContent, "\n") {
		event := Event{}

		// Retrieve time
		if len(line) < 9 {
			continue
		}
		startTimeStr := timeRegexPattern.FindString(line)
		endTimeStr := timeRegexPattern.FindString(line[4:])

		// Parse start time
		if startTimeStr != "" {
			startTime, err := time.Parse("200601021504", note.Title+startTimeStr)
			if err != nil {
				log.Println(err)
				continue
			}
			event.StartTime = startTime
		} else {
			continue
		}

		if endTimeStr != "" {
			endTime, err := time.Parse("200601021504", note.Title+endTimeStr)
			if err != nil {
				log.Println(err)
				continue
			}
			event.Duration = endTime.Sub(event.StartTime)
			event.Description = line[9:]
		} else {
			event.Description = line[4:]
		}

		events = append(events, event)
	}
	return events, nil
}
