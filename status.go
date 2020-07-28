package elite

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
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
	Cargo     float64     `json:"Cargo"`
	Latitude  float64     `json:"Latitude,omitempty"`
	Longitude float64     `json:"Longitude,omitempty"`
	Heading   int32       `json:"Heading,omitempty"`
	Altitude  int32       `json:"Altitude,omitempty"`
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
