package elite

import (
	"bufio"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	stationDocked   = "Docked"
	stationUndocked = "Undocked"
)

// StarSystemEvent is an event that contains the current star system.
// It may be a Location, FSDJump, or SupercruiseExit event.
type StarSystemEvent struct {
	*JournalEntry
	StarSystem string `json:"StarSystem,omitempty"`
}

type StationEvent struct {
	*JournalEntry
	StationName string `json:"StationName,omitempty"`
}

// GetStarSystem returns the current star system.
func GetStarSystem() (string, error) {
	return GetStarSystemFromPath(defaultLogPath)
}

func GetCurrentStation() (string, error) {
	return GetCurrentStationFromPath(defaultLogPath)
}

// GetStarSystemFromPath returns the current star system using the specified log path.
func GetStarSystemFromPath(logPath string) (string, error) {
	files, _ := ioutil.ReadDir(logPath)

	found := false
	var event StarSystemEvent
	for i := len(files) - 1; i >= 0 && !found; i-- {
		if !journalFilePattern.MatchString(files[i].Name()) {
			continue
		}

		journalFile, err := os.Open(filepath.Join(logPath, files[i].Name()))
		if err != nil {
			return "", err
		}
		defer journalFile.Close()

		scanner := bufio.NewScanner(journalFile)
		for scanner.Scan() {
			var tempEvent StarSystemEvent
			err := json.Unmarshal([]byte(scanner.Text()), &tempEvent)
			if err != nil {
				return "", err
			}
			if tempEvent.Event == "FSDJump" || tempEvent.Event == "Location" {
				event = tempEvent
				found = true
			}
		}
	}

	if !found {
		return "", errors.New("No location found in all log files")
	}

	return event.StarSystem, nil
}

func GetCurrentStationFromPath(logPath string) (string, error) {
	files, _ := ioutil.ReadDir(logPath)

	found := false
	var event StationEvent
	for i := len(files) - 1; i >= 0 && !found; i-- {
		if !journalFilePattern.MatchString(files[i].Name()) {
			continue
		}

		journalFile, err := os.Open(filepath.Join(logPath, files[i].Name()))
		if err != nil {
			return "", err
		}
		defer journalFile.Close()

		scanner := bufio.NewScanner(journalFile)
		for scanner.Scan() {
			var tempEvent StationEvent
			err := json.Unmarshal([]byte(scanner.Text()), &tempEvent)
			if err != nil {
				return "", err
			}

			switch tempEvent.Event {
			case stationDocked:
				event = tempEvent
				found = true
			case stationUndocked:
				event = tempEvent
				event.StationName = ""
				found = true
			}
		}
	}

	if !found {
		return "", errors.New("No location found in all log files")
	}

	return event.StationName, nil
}
