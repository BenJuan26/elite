package elite

import (
	"bufio"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// StarSystemEvent is an event that contains the current star system.
// It may be a Location, FSDJump, or SupercruiseExit event.
type StarSystemEvent struct {
	*JournalEntry
	StarSystem string `json:"StarSystem,omitempty"`
}

// GetStarSystem returns the current star system.
func GetStarSystem() (string, error) {
	return GetStarSystemFromPath(defaultLogPath)
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
			json.Unmarshal([]byte(scanner.Text()), &tempEvent)
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
