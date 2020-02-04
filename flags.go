package elite

import (
	"github.com/BenJuan26/elite/flags"
)

// StatusFlags contains boolean flags describing the player and ship.
type StatusFlags struct {
	Docked                    bool
	Landed                    bool
	LandingGearDown           bool
	ShieldsUp                 bool
	Supercruise               bool
	FlightAssistOff           bool
	HardpointsDeployed        bool
	InWing                    bool
	LightsOn                  bool
	CargoScoopDeployed        bool
	SilentRunning             bool
	ScoopingFuel              bool
	SRVHandbrake              bool
	SRVTurret                 bool
	SRVUnderShip              bool
	SRVDriveAssist            bool
	FSDMassLocked             bool
	FSDCharging               bool
	FSDCooldown               bool
	LowFuel                   bool
	Overheating               bool
	HasLatLong                bool
	IsInDanger                bool
	BeingInterdicted          bool
	InMainShip                bool
	InFighter                 bool
	InSRV                     bool
	InAnalysisMode            bool
	NightVision               bool
	AltitudeFromAverageRadius bool
	FSDJump                   bool
	SRVHighBeam               bool
}

// ExpandFlags parses the RawFlags and sets the Flags values accordingly.
func (status *Status) ExpandFlags() {
	status.Flags.Docked = status.RawFlags&flags.Docked != 0
	status.Flags.Landed = status.RawFlags&flags.Landed != 0
	status.Flags.LandingGearDown = status.RawFlags&flags.LandingGearDown != 0
	status.Flags.ShieldsUp = status.RawFlags&flags.ShieldsUp != 0
	status.Flags.Supercruise = status.RawFlags&flags.Supercruise != 0
	status.Flags.FlightAssistOff = status.RawFlags&flags.FlightAssistOff != 0
	status.Flags.HardpointsDeployed = status.RawFlags&flags.HardpointsDeployed != 0
	status.Flags.InWing = status.RawFlags&flags.InWing != 0
	status.Flags.LightsOn = status.RawFlags&flags.LightsOn != 0
	status.Flags.CargoScoopDeployed = status.RawFlags&flags.CargoScoopDeployed != 0
	status.Flags.SilentRunning = status.RawFlags&flags.SilentRunning != 0
	status.Flags.ScoopingFuel = status.RawFlags&flags.ScoopingFuel != 0
	status.Flags.SRVHandbrake = status.RawFlags&flags.SRVHandbrake != 0
	status.Flags.SRVTurret = status.RawFlags&flags.SRVTurret != 0
	status.Flags.SRVUnderShip = status.RawFlags&flags.SRVUnderShip != 0
	status.Flags.SRVDriveAssist = status.RawFlags&flags.SRVDriveAssist != 0
	status.Flags.FSDMassLocked = status.RawFlags&flags.FSDMassLocked != 0
	status.Flags.FSDCharging = status.RawFlags&flags.FSDCharging != 0
	status.Flags.FSDCooldown = status.RawFlags&flags.FSDCooldown != 0
	status.Flags.LowFuel = status.RawFlags&flags.LowFuel != 0
	status.Flags.Overheating = status.RawFlags&flags.Overheating != 0
	status.Flags.HasLatLong = status.RawFlags&flags.HasLatLong != 0
	status.Flags.IsInDanger = status.RawFlags&flags.IsInDanger != 0
	status.Flags.BeingInterdicted = status.RawFlags&flags.BeingInterdicted != 0
	status.Flags.InMainShip = status.RawFlags&flags.InMainShip != 0
	status.Flags.InFighter = status.RawFlags&flags.InFighter != 0
	status.Flags.InSRV = status.RawFlags&flags.InSRV != 0
	status.Flags.InAnalysisMode = status.RawFlags&flags.InAnalysisMode != 0
	status.Flags.NightVision = status.RawFlags&flags.NightVision != 0
	status.Flags.AltitudeFromAverageRadius = status.RawFlags&flags.AltitudeFromAverageRadius != 0
	status.Flags.FSDJump = status.RawFlags&flags.FSDJump != 0
	status.Flags.SRVHighBeam = status.RawFlags&flags.SRVHighBeam != 0
}
