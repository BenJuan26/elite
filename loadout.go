package elite

import (
	"bufio"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// LoadoutFuel contains the main and reserve fuel capacities.
type LoadoutFuel struct {
	Main    float64 `json:"Main"`
	Reserve float64 `json:"Reserve"`
}

// Module contains information about a ship module.
type Module struct {
	Slot        string      `json:"Slot"`
	Item        string      `json:"Item"`
	On          bool        `json:"On"`
	Priority    int64       `json:"Priority"`
	Health      int64       `json:"Health"`
	Engineering Engineering `json:"Engineering"`
}

// Modifier describes an engineering modifier.
type Modifier struct {
	Label         string  `json:"Label"`
	Value         float64 `json:"Value"`
	OriginalValue float64 `json:"OriginalValue"`
	LessIsGood    int64   `json:"LessIsGood"`
}

// Engineering represents the engineering modifications performed on a module.
type Engineering struct {
	Engineer           string     `json:"Engineer"`
	EngineerID         int64      `json:"EngineerID"`
	BlueprintID        int64      `json:"BlueprintID"`
	BlueprintName      string     `json:"BlueprintName"`
	Level              int64      `json:"Level"`
	Quality            float64    `json:"Quality"`
	ExperimentalEffect string     `json:"ExperimentalEffect"`
	Modifiers          []Modifier `json:"Modifiers"`
}

// Loadout contains information about the player's ship, its status, and its modules.
type Loadout struct {
	*JournalEntry
	Ship          string      `json:"Ship"`
	ShipID        int64       `json:"ShipID"`
	ShipName      string      `json:"ShipName"`
	ShipIdent     string      `json:"ShipIdent"`
	HullValue     int64       `json:"HullValue"`
	ModulesValue  int64       `json:"ModulesValue"`
	HullHealth    int64       `json:"HullHealth"`
	UnladenMass   float64     `json:"UnladenMass"`
	CargoCapacity int64       `json:"CargoCapacity"`
	MaxJumpRange  float64     `json:"MaxJumpRange"`
	FuelCapacity  LoadoutFuel `json:"FuelCapacity"`
	Rebuy         int64       `json:"Rebuy"`
	Modules       []Module    `json:"Modules"`
}

// GetLoadoutFromPath reads the current ship loadout from the journal files at the specified path.
func GetLoadoutFromPath(logPath string) (*Loadout, error) {
	files, _ := ioutil.ReadDir(logPath)

	found := false
	var loadout *Loadout
	for i := len(files) - 1; i >= 0 && !found; i-- {
		if !journalFilePattern.MatchString(files[i].Name()) {
			continue
		}

		journalFile, err := os.Open(filepath.Join(logPath, files[i].Name()))
		if err != nil {
			return loadout, err
		}
		defer journalFile.Close()

		scanner := bufio.NewScanner(journalFile)
		for scanner.Scan() {
			var tempEvent Loadout
			json.Unmarshal(scanner.Bytes(), &tempEvent)
			if tempEvent.Event == "Loadout" {
				loadout = &tempEvent
				found = true
			}
		}
	}

	if !found {
		return loadout, errors.New("No loadout found in all log files")
	}

	return loadout, nil
}

// GetLoadout reads the current ship loadout from the journal files.
// It will read them from the default log path, which is the Saved Games
// folder. The full path is:
//
//     C:/Users/<Username>/Saved Games/Frontier Developments/Elite Dangerous
//
// If that path is not suitable, use GetLoadoutFromPath.
func GetLoadout() (*Loadout, error) {
	return GetLoadoutFromPath(defaultLogPath)
}
