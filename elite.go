// Package elite provides real-time data from Elite Dangerous through
// files written to disk by the game.
package elite

import (
	"os/user"
	"path/filepath"
	"regexp"
)

// JournalEntry is a minimal entry in the Journal file.
// It is primarily intended for embedding within event types,
// such as StarSystemEvent.
type JournalEntry struct {
	Timestamp string `json:"timestamp"`
	Event     string `json:"event"`
}

var defaultLogPath string
var journalFilePattern *regexp.Regexp

func init() {
	currUser, _ := user.Current()
	homeDir := currUser.HomeDir
	defaultLogPath = filepath.FromSlash(homeDir + "/Saved Games/Frontier Developments/Elite Dangerous")
	journalFilePattern = regexp.MustCompile(`^Journal\.\d{12}\.\d{2}\.log$`)
}
