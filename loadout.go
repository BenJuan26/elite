package elite

import (
	"bufio"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/BenJuan26/elite/loadout"
)

// Loadout contains information about the player's ship, its status, and its modules.
type Loadout struct {
	*JournalEntry
	Ship          string           `json:"Ship"`
	ShipID        int64            `json:"ShipID"`
	ShipName      string           `json:"ShipName"`
	ShipIdent     string           `json:"ShipIdent"`
	HullValue     int64            `json:"HullValue"`
	ModulesValue  int64            `json:"ModulesValue"`
	HullHealth    int64            `json:"HullHealth"`
	UnladenMass   float64          `json:"UnladenMass"`
	CargoCapacity int64            `json:"CargoCapacity"`
	MaxJumpRange  float64          `json:"MaxJumpRange"`
	FuelCapacity  loadout.Fuel     `json:"FuelCapacity"`
	Rebuy         int64            `json:"Rebuy"`
	Modules       []loadout.Module `json:"Modules"`
}

// GetLoadoutFromPath reads the current ship loadout from the journal files at the specified path.
func GetLoadoutFromPath(logPath string) (*Loadout, error) {
	files, _ := ioutil.ReadDir(logPath)

	found := false
	var l *Loadout
	for i := len(files) - 1; i >= 0 && !found; i-- {
		if !journalFilePattern.MatchString(files[i].Name()) {
			continue
		}

		journalFile, err := os.Open(filepath.Join(logPath, files[i].Name()))
		if err != nil {
			return l, err
		}
		defer journalFile.Close()

		scanner := bufio.NewScanner(journalFile)
		for scanner.Scan() {
			var tempEvent Loadout
			json.Unmarshal(scanner.Bytes(), &tempEvent)
			if tempEvent.Event == "Loadout" {
				l = &tempEvent
				found = true
			}
		}
	}

	if !found {
		return l, errors.New("No loadout found in all log files")
	}

	return l, nil
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
