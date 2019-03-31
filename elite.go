// +build windows

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
	// FlagDocked indicates that the ship is docked.
	FlagDocked uint32 = 0x00000001
	// FlagLanded indicates that the ship is landed.
	FlagLanded uint32 = 0x00000002
	// FlagLandingGearDown indicates that the landing gear is deployed.
	FlagLandingGearDown uint32 = 0x00000004
	// FlagShieldsUp indicates that the ship's shields are up.
	FlagShieldsUp uint32 = 0x00000008
	// FlagSupercruise indicates that the ship is in supercruise.
	FlagSupercruise uint32 = 0x00000010
	// FlagFlightAssistOff indicates that flight assist is disabled.
	FlagFlightAssistOff uint32 = 0x00000020
	// FlagHardpointsDeployed indicates that the ship's hardpoints are deployed.
	FlagHardpointsDeployed uint32 = 0x00000040
	// FlagInWing indicates whether the player is in a wing.
	FlagInWing uint32 = 0x00000080
	// FlagLightsOn indicates that the ship's lights are on.
	FlagLightsOn uint32 = 0x00000100
	// FlagCargoScoopDeployed indicates that the cargo scoop is deployed.
	FlagCargoScoopDeployed uint32 = 0x00000200
	// FlagSilentRunning indicates that silent running is on.
	FlagSilentRunning uint32 = 0x00000400
	// FlagScoopingFuel indicates that the ship is currently scooping fuel.
	FlagScoopingFuel uint32 = 0x00000800
	// FlagSRVHandbrake indicates that the SRV's handbrake is enabled.
	FlagSRVHandbrake uint32 = 0x00001000
	// FlagSRVTurret indicates that the SRV's turret is deployed.
	FlagSRVTurret uint32 = 0x00002000
	// FlagSRVUnderShip indicates that the SRV is positioned under the ship.
	FlagSRVUnderShip uint32 = 0x00004000
	// FlagSRVDriveAssist indicates that the SRV's drive assist is on.
	FlagSRVDriveAssist uint32 = 0x00008000
	// FlagFSDMassLocked indicates that the ship is mass locked.
	FlagFSDMassLocked uint32 = 0x00010000
	// FlagFSDCharging indicates that the FSD is charging.
	FlagFSDCharging uint32 = 0x00020000
	// FlagFSDCooldown indicates that the FSD is cooling down.
	FlagFSDCooldown uint32 = 0x00040000
	// FlagLowFuel indicates that the ship is low on fuel.
	FlagLowFuel uint32 = 0x00080000
	// FlagOverheating indicates that the ship is overheating.
	FlagOverheating uint32 = 0x00100000
	// FlagHasLatLong indicates that latitude and longitude data are available.
	FlagHasLatLong uint32 = 0x00200000
	// FlagIsInDanger indicates that the player is in danger.
	FlagIsInDanger uint32 = 0x00400000
	// FlagBeingInterdicted indicates that the ship is being interdicted.
	FlagBeingInterdicted uint32 = 0x00800000
	// FlagInMainShip indicates that the player is in the ship.
	FlagInMainShip uint32 = 0x01000000
	// FlagInFighter indicates that the player is in a fighter.
	FlagInFighter uint32 = 0x02000000
	// FlagInSRV indicates that the player is in an SRV.
	FlagInSRV uint32 = 0x04000000
	// FlagInAnalysisMode indicates that analysis mode is selected.
	FlagInAnalysisMode uint32 = 0x08000000
	// FlagNightVision indicates that night vision is enabled.
	FlagNightVision uint32 = 0x10000000

	// GuiFocusNone indicates that there is no menu panel focused.
	GuiFocusNone uint32 = 0
	// GuiFocusInternalPanel indicates that the internal menu panel is focused.
	GuiFocusInternalPanel uint32 = 1
	// GuiFocusExternalPanel indicates that the external menu panel is focused.
	GuiFocusExternalPanel uint32 = 2
	// GuiFocusCommsPanel indicates that the comms menu panel is focused.
	GuiFocusCommsPanel uint32 = 3
	// GuiFocusRolePanel indicates that the role menu panel is focused.
	GuiFocusRolePanel uint32 = 4
	// GuiFocusStationServices indicates that the station services menu is focused.
	GuiFocusStationServices uint32 = 5
	// GuiFocusGalaxyMap indicates that the galaxy map is open.
	GuiFocusGalaxyMap uint32 = 6
	// GuiFocusSystemMap indicates that the system map is open.
	GuiFocusSystemMap uint32 = 7
	// GuiFocusOrrery indicates that the orrery is open.
	GuiFocusOrrery uint32 = 8
	// GuiFocusFSSMode indicates that the FSS is open.
	GuiFocusFSSMode uint32 = 9
	// GuiFocusSAAMode indicates that the SAA is focused.
	GuiFocusSAAMode uint32 = 10
	// GuiFocusCodex indicates that the codex is focused.
	GuiFocusCodex uint32 = 11

	// GuiFocusLeft is a helper alias for GuiFocusExternalPanel.
	GuiFocusLeft uint32 = GuiFocusExternalPanel
	// GuiFocusRight is a helper alias for GuiFocusInternalPanel.
	GuiFocusRight uint32 = GuiFocusInternalPanel
	// GuiFocusTop is a helper alias for GuiFocusCommsPanel.
	GuiFocusTop uint32 = GuiFocusCommsPanel
	// GuiFocusBottom is a helper alias for GuiFocusRolePanel.
	GuiFocusBottom uint32 = GuiFocusRolePanel
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
	InAnalysisMode     bool
	NightVision        bool
}

// Status represents the current state of the player and ship
type Status struct {
	Timestamp string      `json:"timestamp"`
	Event     string      `json:"event"`
	Flags     StatusFlags `json:"-"`
	RawFlags  uint32      `json:"Flags"`
	Pips      [3]int32    `json:"Pips"`
	FireGroup int32       `json:"FireGroup"`
	GuiFocus  int32       `json:"GuiFocus"`
	Latitude  float64     `json:"Latitude,omitempty"`
	Longitude float64     `json:"Longitude,omitempty"`
	Heading   int32       `json:"Heading,omitempty"`
	Altitude  int32       `json:"Altitude,omitempty"`
}

type starSystemEvent struct {
	Timestamp  string `json:"timestamp"`
	Event      string `json:"event"`
	StarSystem string `json:"StarSystem,omitempty"`
}

var defaultLogPath string

func init() {
	currUser, _ := user.Current()
	homeDir := currUser.HomeDir
	defaultLogPath = filepath.FromSlash(homeDir + "/Saved Games/Frontier Developments/Elite Dangerous")
}

// ExpandFlags parses the flags value and returns the flags in a StatusFlags struct
func (status *Status) ExpandFlags() {
	status.Flags.Docked = status.RawFlags&FlagDocked != 0
	status.Flags.Landed = status.RawFlags&FlagLanded != 0
	status.Flags.LandingGearDown = status.RawFlags&FlagLandingGearDown != 0
	status.Flags.ShieldsUp = status.RawFlags&FlagShieldsUp != 0
	status.Flags.Supercruise = status.RawFlags&FlagSupercruise != 0
	status.Flags.FlightAssistOff = status.RawFlags&FlagFlightAssistOff != 0
	status.Flags.HardpointsDeployed = status.RawFlags&FlagHardpointsDeployed != 0
	status.Flags.InWing = status.RawFlags&FlagInWing != 0
	status.Flags.LightsOn = status.RawFlags&FlagLightsOn != 0
	status.Flags.CargoScoopDeployed = status.RawFlags&FlagCargoScoopDeployed != 0
	status.Flags.SilentRunning = status.RawFlags&FlagSilentRunning != 0
	status.Flags.ScoopingFuel = status.RawFlags&FlagScoopingFuel != 0
	status.Flags.SRVHandbrake = status.RawFlags&FlagSRVHandbrake != 0
	status.Flags.SRVTurret = status.RawFlags&FlagSRVTurret != 0
	status.Flags.SRVUnderShip = status.RawFlags&FlagSRVUnderShip != 0
	status.Flags.SRVDriveAssist = status.RawFlags&FlagSRVDriveAssist != 0
	status.Flags.FSDMassLocked = status.RawFlags&FlagFSDMassLocked != 0
	status.Flags.FSDCharging = status.RawFlags&FlagFSDCharging != 0
	status.Flags.FSDCooldown = status.RawFlags&FlagFSDCooldown != 0
	status.Flags.LowFuel = status.RawFlags&FlagLowFuel != 0
	status.Flags.Overheating = status.RawFlags&FlagOverheating != 0
	status.Flags.HasLatLong = status.RawFlags&FlagHasLatLong != 0
	status.Flags.IsInDanger = status.RawFlags&FlagIsInDanger != 0
	status.Flags.BeingInterdicted = status.RawFlags&FlagBeingInterdicted != 0
	status.Flags.InMainShip = status.RawFlags&FlagInMainShip != 0
	status.Flags.InFighter = status.RawFlags&FlagInFighter != 0
	status.Flags.InSRV = status.RawFlags&FlagInSRV != 0
	status.Flags.InAnalysisMode = status.RawFlags&FlagInAnalysisMode != 0
	status.Flags.NightVision = status.RawFlags&FlagNightVision != 0
}

// GetStarSystem returns the current star system
func GetStarSystem() (string, error) {
	return GetStarSystemFromPath(defaultLogPath)
}

// GetStarSystemFromPath returns the current star system using the specified log path
func GetStarSystemFromPath(logPath string) (string, error) {
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
	return GetStatusFromPath(defaultLogPath)
}

// GetStatusFromPath reads the current player and ship status from Status.json at the specified log path
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

// GetStatusFromBytes reads the current player and ship status from the string contained in the byte array
func GetStatusFromBytes(content []byte) (*Status, error) {
	status := &Status{}
	if err := json.Unmarshal(content, status); err != nil {
		return nil, errors.New("Couldn't unmarshal Status.json file: " + err.Error())
	}

	status.ExpandFlags()
	return status, nil
}
