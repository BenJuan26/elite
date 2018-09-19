package elite

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"os/user"
)

// Status represents the current state of the player and ship
type Status struct {
	Timestamp string  `json:"timestamp"`
	Event     string  `json:"event"`
	Flags     int32   `json:"Flags"`
	Pips      []int32 `json:"Pips"`
	FireGroup int32   `json:"FireGroup"`
	GuiFocus  int32   `json:"GuiFocus"`
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
	Heading   int32   `json:"Heading"`
	Altitude  int32   `json:"Altitude"`

	// Flags
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

// ExpandFlags parses the flags value and expands it into the corresponding flag fields
func (status *Status) ExpandFlags() {
	status.Docked = status.Flags&(1<<0) > 0
	status.Landed = status.Flags&(1<<1) > 0
	status.LandingGearDown = status.Flags&(1<<2) > 0
	status.ShieldsUp = status.Flags&(1<<3) > 0
	status.Supercruise = status.Flags&(1<<4) > 0
	status.FlightAssistOff = status.Flags&(1<<5) > 0
	status.HardpointsDeployed = status.Flags&(1<<6) > 0
	status.InWing = status.Flags&(1<<7) > 0
	status.LightsOn = status.Flags&(1<<8) > 0
	status.CargoScoopDeployed = status.Flags&(1<<9) > 0
	status.SilentRunning = status.Flags&(1<<10) > 0
	status.ScoopingFuel = status.Flags&(1<<11) > 0
	status.SRVHandbrake = status.Flags&(1<<12) > 0
	status.SRVTurret = status.Flags&(1<<13) > 0
	status.SRVUnderShip = status.Flags&(1<<14) > 0
	status.SRVDriveAssist = status.Flags&(1<<15) > 0
	status.FSDMassLocked = status.Flags&(1<<16) > 0
	status.FSDCharging = status.Flags&(1<<17) > 0
	status.FSDCooldown = status.Flags&(1<<18) > 0
	status.LowFuel = status.Flags&(1<<19) > 0
	status.Overheating = status.Flags&(1<<20) > 0
	status.HasLatLong = status.Flags&(1<<21) > 0
	status.IsInDanger = status.Flags&(1<<22) > 0
	status.BeingInterdicted = status.Flags&(1<<23) > 0
	status.InMainShip = status.Flags&(1<<24) > 0
	status.InFighter = status.Flags&(1<<25) > 0
	status.InSRV = status.Flags&(1<<26) > 0
}

// GetStatus reads the current player and ship status from Status.json
func GetStatus() (*Status, error) {
	currUser, err := user.Current()
	if err != nil {
		return nil, errors.New("Couldn't get current user: " + err.Error())
	}

	statusFilePath := currUser.HomeDir + "\\Saved Games\\Frontier Developments\\Elite Dangerous\\Status.json"
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
