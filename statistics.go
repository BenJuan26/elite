package elite

import (
	"bufio"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/BenJuan26/elite/stats"
)

// Statistics is the main statistics object.
// The statistics are divided into different categories contains by several fields.
type Statistics struct {
	*JournalEntry
	BankAccount     stats.BankAccount     `json:"Bank_Account"`
	Combat          stats.Combat          `json:"Combat"`
	Crime           stats.Crime           `json:"Crime"`
	Smuggling       stats.Smuggling       `json:"Smuggling"`
	Trading         stats.Trading         `json:"Trading"`
	Mining          stats.Mining          `json:"Mining"`
	Exploration     stats.Exploration     `json:"Exploration"`
	Passengers      stats.Passengers      `json:"Passengers"`
	SearchAndRescue stats.SearchAndRescue `json:"Search_And_Rescue"`
	Crafting        stats.Crafting        `json:"Crafting"`
	// Crew            stats.Crew            `json:"Crew"`
	Multicrew      stats.Multicrew      `json:"Multicrew"`
	MaterialTrader stats.MaterialTrader `json:"Material_Trader_Stats"`
}

// GetStatisticsFromPath returns game statistics using the specified log path.
func GetStatisticsFromPath(logPath string) (*Statistics, error) {
	files, _ := ioutil.ReadDir(logPath)

	found := false
	var stats *Statistics
	for i := len(files) - 1; i >= 0 && !found; i-- {
		if !journalFilePattern.MatchString(files[i].Name()) {
			continue
		}

		journalFile, err := os.Open(filepath.Join(logPath, files[i].Name()))
		if err != nil {
			return stats, err
		}
		defer journalFile.Close()

		scanner := bufio.NewScanner(journalFile)
		for scanner.Scan() {
			var tempEvent Statistics
			json.Unmarshal([]byte(scanner.Text()), &tempEvent)
			if tempEvent.Event == "Statistics" {
				stats = &tempEvent
				found = true
			}
		}
	}

	if !found {
		return stats, errors.New("No location found in all log files")
	}

	return stats, nil
}

// GetStatistics reads the game statistics from the journal files.
// It will read them from the default log path, which is the Saved Games
// folder. The full path is:
//
//     C:/Users/<Username>/Saved Games/Frontier Developments/Elite Dangerous
//
// If that path is not suitable, use GetStatisticsFromPath.
func GetStatistics() (*Statistics, error) {
	return GetStatisticsFromPath(defaultLogPath)
}
