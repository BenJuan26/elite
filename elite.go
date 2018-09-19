package elite

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

const (
	FlagDocked             int32 = 0x00000001
	FlagLanded             int32 = 0x00000002
	FlagLandingGearDown    int32 = 0x00000004
	FlagShieldsUp          int32 = 0x00000008
	FlagSupercruise        int32 = 0x00000010
	FlagFlightAssistOff    int32 = 0x00000020
	FlagHardpointsDeployed int32 = 0x00000040
	FlagInWing             int32 = 0x00000080
	FlagLightsOn           int32 = 0x00000100
	FlagCargoScoopDeployed int32 = 0x00000200
	FlagSilentRunning      int32 = 0x00000400
	FlagScoopingFuel       int32 = 0x00000800
	FlagSRVHandbrake       int32 = 0x00001000
	FlagSRVTurret          int32 = 0x00002000
	FlagSRVUnderShip       int32 = 0x00004000
	FlagSRVDriveAssist     int32 = 0x00008000
	FlagFSDMassLocked      int32 = 0x00010000
	FlagFSDCharging        int32 = 0x00020000
	FlagFSDCooldown        int32 = 0x00040000
	FlagLowFuel            int32 = 0x00080000
	FlagOverheating        int32 = 0x00100000
	FlagHasLatLong         int32 = 0x00200000
	FlagIsInDanger         int32 = 0x00400000
	FlagBeingInterdicted   int32 = 0x00800000
	FlagInMainShip         int32 = 0x01000000
	FlagInFighter          int32 = 0x02000000
	FlagInSRV              int32 = 0x04000000
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
	Flags     int32   `json:"Flags"`
	Pips      []int32 `json:"Pips"`
	FireGroup int32   `json:"FireGroup"`
	GuiFocus  int32   `json:"GuiFocus"`
	Latitude  float64 `json:"Latitude,omitempty"`
	Longitude float64 `json:"Longitude,omitempty"`
	Heading   int32   `json:"Heading,omitempty"`
	Altitude  int32   `json:"Altitude,omitempty"`
}

// ExpandFlags parses the flags value and expands it into the corresponding flag fields
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

// GetStatus reads the current player and ship status from Status.json
func GetStatus() (*Status, error) {
	currUser, err := user.Current()
	if err != nil {
		return nil, errors.New("Couldn't get current user: " + err.Error())
	}

	statusFilePath := filepath.FromSlash(currUser.HomeDir + "/Saved Games/Frontier Developments/Elite Dangerous/Status.json")
	statusFile, err := os.Open(statusFilePath)
	if err != nil {
		return nil, errors.New("Couldn't open Status.json file: " + err.Error())
	}

	statusBytes, err := ioutil.ReadAll(statusFile)
	if err != nil {
		return nil, errors.New("Couldn't read Status.json file: " + err.Error())
	}

	status := &Status{}
	if err := json.Unmarshal(statusBytes, status); err != nil {
		return nil, errors.New("Couldn't unmarshal Status.json file: " + err.Error())
	}

	return status, nil
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
