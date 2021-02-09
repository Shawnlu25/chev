package keep

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

// ParseExportedNoteInJSON ...
func ParseExportedNoteInJSON(filepath string) (ExportedNote, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return ExportedNote{}, err
	}

	result := ExportedNote{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return ExportedNote{}, err
	}

	return result, nil
}

// ParseExportedNotesFromDir ...
func ParseExportedNotesFromDir(dir string) ([]ExportedNote, error) {
	matchedFilepaths, err := filepath.Glob(dir + "*.json")
	if err != nil {
		return nil, err
	}

	notes := []ExportedNote{}
	for i := 0; i < len(matchedFilepaths); i++ {
		note, err := ParseExportedNoteInJSON(matchedFilepaths[i])
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}
