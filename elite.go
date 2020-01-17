// Package elite provides real-time data from Elite Dangerous through
// files written to disk by the game.
package elite

import (
	"bufio"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"time"
)

// Fuel contains fuel readouts for the ship.
type Fuel struct {
	Main      float64 `json:"FuelMain"`
	Reservoir float64 `json:"FuelReservoir"`
}

// Status represents the current state of the player and ship.
type Status struct {
	Timestamp string      `json:"timestamp"`
	Event     string      `json:"event"`
	Flags     StatusFlags `json:"-"`
	RawFlags  uint32      `json:"Flags"`
	Pips      [3]int32    `json:"Pips"`
	FireGroup int32       `json:"FireGroup"`
	GuiFocus  int32       `json:"GuiFocus"`
	Fuel      Fuel        `json:"Fuel"`
	Cargo     int64       `json:"Cargo"`
	Latitude  float64     `json:"Latitude,omitempty"`
	Longitude float64     `json:"Longitude,omitempty"`
	Heading   int32       `json:"Heading,omitempty"`
	Altitude  int32       `json:"Altitude,omitempty"`
}

// JournalEntry is a minimal entry in the Journal file.
// It is primarily intended for embedding within event types,
// such as StarSystemEvent.
type JournalEntry struct {
	Timestamp string `json:"timestamp"`
	Event     string `json:"event"`
}

// StarSystemEvent is an event that contains the current star system.
// It may be a Location, FSDJump, or SupercruiseExit event.
type StarSystemEvent struct {
	*JournalEntry
	StarSystem string `json:"StarSystem,omitempty"`
}

var defaultLogPath string
var journalFilePattern *regexp.Regexp

func init() {
	currUser, _ := user.Current()
	homeDir := currUser.HomeDir
	defaultLogPath = filepath.FromSlash(homeDir + "/Saved Games/Frontier Developments/Elite Dangerous")
	journalFilePattern = regexp.MustCompile(`^Journal\.\d{12}\.\d{2}\.log$`)
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

// GetStatus reads the current player and ship status from Status.json.
// It will read them from the default log path, which is the Saved Games
// folder. The full path is:
//
//     C:/Users/<Username>/Saved Games/Frontier Developments/Elite Dangerous
//
// If that path is not suitable, use GetStatusFromPath.
func GetStatus() (*Status, error) {
	return GetStatusFromPath(defaultLogPath)
}

// GetStatusFromPath reads the current player and ship status from Status.json at the specified log path.
func GetStatusFromPath(logPath string) (*Status, error) {
	statusFilePath := filepath.FromSlash(logPath + "/Status.json")
	retries := 5
	for retries > 0 {
		statusFile, err := os.Open(statusFilePath)
		if err != nil {
			retries = retries - 1
			time.Sleep(3 * time.Millisecond)
			continue
		}
		defer statusFile.Close()

		statusBytes, err := ioutil.ReadAll(statusFile)
		if err != nil {
			retries = retries - 1
			time.Sleep(3 * time.Millisecond)
			continue
		}

		return GetStatusFromBytes(statusBytes)
	}

	return nil, errors.New("Couldn't get status after 5 attempts")
}

// GetStatusFromBytes reads the current player and ship status from the string contained in the byte array.
func GetStatusFromBytes(content []byte) (*Status, error) {
	status := &Status{}
	if err := json.Unmarshal(content, status); err != nil {
		return nil, errors.New("Couldn't unmarshal Status.json file: " + err.Error())
	}

	status.ExpandFlags()
	return status, nil
}
