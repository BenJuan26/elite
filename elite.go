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

const (
	FlagDocked             uint32 = 0x00000001
	FlagLanded             uint32 = 0x00000002
	FlagLandingGearDown    uint32 = 0x00000004
	FlagShieldsUp          uint32 = 0x00000008
	FlagSupercruise        uint32 = 0x00000010
	FlagFlightAssistOff    uint32 = 0x00000020
	FlagHardpointsDeployed uint32 = 0x00000040
	FlagInWing             uint32 = 0x00000080
	FlagLightsOn           uint32 = 0x00000100
	FlagCargoScoopDeployed uint32 = 0x00000200
	FlagSilentRunning      uint32 = 0x00000400
	FlagScoopingFuel       uint32 = 0x00000800
	FlagSRVHandbrake       uint32 = 0x00001000
	FlagSRVTurret          uint32 = 0x00002000
	FlagSRVUnderShip       uint32 = 0x00004000
	FlagSRVDriveAssist     uint32 = 0x00008000
	FlagFSDMassLocked      uint32 = 0x00010000
	FlagFSDCharging        uint32 = 0x00020000
	FlagFSDCooldown        uint32 = 0x00040000
	FlagLowFuel            uint32 = 0x00080000
	FlagOverheating        uint32 = 0x00100000
	FlagHasLatLong         uint32 = 0x00200000
	FlagIsInDanger         uint32 = 0x00400000
	FlagBeingInterdicted   uint32 = 0x00800000
	FlagInMainShip         uint32 = 0x01000000
	FlagInFighter          uint32 = 0x02000000
	FlagInSRV              uint32 = 0x04000000
)

// StatusFlags contains boolean flags describing the player and ship
type StatusFlags struct {
	Docked             bool
	Landed             bool
	LandingGearDown    bool
	ShieldsUp          bool
	Supercruise        bool
	FlightAssistOff    bool
	HardpointsDeployed bool
	InWing             bool
	LightsOn           bool
	CargoScoopDeployed bool
	SilentRunning      bool
	ScoopingFuel       bool
	SRVHandbrake       bool
	SRVTurret          bool
	SRVUnderShip       bool
	SRVDriveAssist     bool
	FSDMassLocked      bool
	FSDCharging        bool
	FSDCooldown        bool
	LowFuel            bool
	Overheating        bool
	HasLatLong         bool
	IsInDanger         bool
	BeingInterdicted   bool
	InMainShip         bool
	InFighter          bool
	InSRV              bool
}

// Status represents the current state of the player and ship
type Status struct {
	Timestamp string  `json:"timestamp"`
	Event     string  `json:"event"`
	Flags     uint32  `json:"Flags"`
	Pips      []int32 `json:"Pips"`
	FireGroup int32   `json:"FireGroup"`
	GuiFocus  int32   `json:"GuiFocus"`
	Latitude  float64 `json:"Latitude,omitempty"`
	Longitude float64 `json:"Longitude,omitempty"`
	Heading   int32   `json:"Heading,omitempty"`
	Altitude  int32   `json:"Altitude,omitempty"`
}

type starSystemEvent struct {
	Timestamp  string `json:"timestamp"`
	Event      string `json:"event"`
	StarSystem string `json:"StarSystem,omitempty"`
}

var statusFilePath string
var logPath string

func init() {
	currUser, _ := user.Current()
	logPath = filepath.FromSlash(currUser.HomeDir + "/Saved Games/Frontier Developments/Elite Dangerous")
	statusFilePath = filepath.Join(logPath, "Status.json")
}

// ExpandFlags parses the flags value and returns the flags in a StatusFlags struct
func (status *Status) ExpandFlags() StatusFlags {
	flags := StatusFlags{}

	flags.Docked = status.Flags&FlagDocked != 0
	flags.Landed = status.Flags&FlagLanded != 0
	flags.LandingGearDown = status.Flags&FlagLandingGearDown != 0
	flags.ShieldsUp = status.Flags&FlagShieldsUp != 0
	flags.Supercruise = status.Flags&FlagSupercruise != 0
	flags.FlightAssistOff = status.Flags&FlagFlightAssistOff != 0
	flags.HardpointsDeployed = status.Flags&FlagHardpointsDeployed != 0
	flags.InWing = status.Flags&FlagInWing != 0
	flags.LightsOn = status.Flags&FlagLightsOn != 0
	flags.CargoScoopDeployed = status.Flags&FlagCargoScoopDeployed != 0
	flags.SilentRunning = status.Flags&FlagSilentRunning != 0
	flags.ScoopingFuel = status.Flags&FlagScoopingFuel != 0
	flags.SRVHandbrake = status.Flags&FlagSRVHandbrake != 0
	flags.SRVTurret = status.Flags&FlagSRVTurret != 0
	flags.SRVUnderShip = status.Flags&FlagSRVUnderShip != 0
	flags.SRVDriveAssist = status.Flags&FlagSRVDriveAssist != 0
	flags.FSDMassLocked = status.Flags&FlagFSDMassLocked != 0
	flags.FSDCharging = status.Flags&FlagFSDCharging != 0
	flags.FSDCooldown = status.Flags&FlagFSDCooldown != 0
	flags.LowFuel = status.Flags&FlagLowFuel != 0
	flags.Overheating = status.Flags&FlagOverheating != 0
	flags.HasLatLong = status.Flags&FlagHasLatLong != 0
	flags.IsInDanger = status.Flags&FlagIsInDanger != 0
	flags.BeingInterdicted = status.Flags&FlagBeingInterdicted != 0
	flags.InMainShip = status.Flags&FlagInMainShip != 0
	flags.InFighter = status.Flags&FlagInFighter != 0
	flags.InSRV = status.Flags&FlagInSRV != 0

	return flags
}

// GetStarSystem returns the current star system
func GetStarSystem() (string, error) {
	files, _ := ioutil.ReadDir(logPath)
	journalFilePattern, err := regexp.Compile(`^Journal\.\d{12}\.\d{2}\.log$`)
	if err != nil {
		return "", err
	}

	found := false
	var event starSystemEvent
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
			var tempEvent starSystemEvent
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

// GetStatus reads the current player and ship status from Status.json
func GetStatus() (*Status, error) {
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

		status := &Status{}
		if err := json.Unmarshal(statusBytes, status); err != nil {
			retries = retries - 1
			time.Sleep(3 * time.Millisecond)
			continue
		}

		return status, nil
	}

	return nil, errors.New("Couldn't get status after 5 attempts")
}

// GetStatusFromString reads the current player and ship status from Status.json
func GetStatusFromString(content string) (*Status, error) {
	statusBytes := []byte(content)

	status := &Status{}
	if err := json.Unmarshal(statusBytes, status); err != nil {
		return nil, errors.New("Couldn't unmarshal Status.json file: " + err.Error())
	}

	return status, nil
}

// GetStatusLastModified gets the time the status.json file was last modified
func GetStatusLastModified() (time.Time, error) {
	info, err := os.Stat(statusFilePath)
	if err != nil {
		return time.Time{}, err
	}

	return info.ModTime(), nil
}
