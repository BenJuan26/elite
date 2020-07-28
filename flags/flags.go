package flags

const (
	// Docked indicates that the ship is docked.
	Docked uint32 = 0x00000001
	// Landed indicates that the ship is landed.
	Landed uint32 = 0x00000002
	// LandingGearDown indicates that the landing gear is deployed.
	LandingGearDown uint32 = 0x00000004
	// ShieldsUp indicates that the ship's shields are up.
	ShieldsUp uint32 = 0x00000008
	// Supercruise indicates that the ship is in supercruise.
	Supercruise uint32 = 0x00000010
	// FlightAssistOff indicates that flight assist is disabled.
	FlightAssistOff uint32 = 0x00000020
	// HardpointsDeployed indicates that the ship's hardpoints are deployed.
	HardpointsDeployed uint32 = 0x00000040
	// InWing indicates whether the player is in a wing.
	InWing uint32 = 0x00000080
	// LightsOn indicates that the ship's lights are on.
	LightsOn uint32 = 0x00000100
	// CargoScoopDeployed indicates that the cargo scoop is deployed.
	CargoScoopDeployed uint32 = 0x00000200
	// SilentRunning indicates that silent running is on.
	SilentRunning uint32 = 0x00000400
	// ScoopingFuel indicates that the ship is currently scooping fuel.
	ScoopingFuel uint32 = 0x00000800
	// SRVHandbrake indicates that the SRV's handbrake is enabled.
	SRVHandbrake uint32 = 0x00001000
	// SRVTurret indicates that the SRV's turret is deployed.
	SRVTurret uint32 = 0x00002000
	// SRVUnderShip indicates that the SRV is positioned under the ship.
	SRVUnderShip uint32 = 0x00004000
	// SRVDriveAssist indicates that the SRV's drive assist is on.
	SRVDriveAssist uint32 = 0x00008000
	// FSDMassLocked indicates that the ship is mass locked.
	FSDMassLocked uint32 = 0x00010000
	// FSDCharging indicates that the FSD is charging.
	FSDCharging uint32 = 0x00020000
	// FSDCooldown indicates that the FSD is cooling down.
	FSDCooldown uint32 = 0x00040000
	// LowFuel indicates that the ship is low on fuel.
	LowFuel uint32 = 0x00080000
	// Overheating indicates that the ship is overheating.
	Overheating uint32 = 0x00100000
	// HasLatLong indicates that latitude and longitude data are available.
	HasLatLong uint32 = 0x00200000
	// IsInDanger indicates that the player is in danger.
	IsInDanger uint32 = 0x00400000
	// BeingInterdicted indicates that the ship is being interdicted.
	BeingInterdicted uint32 = 0x00800000
	// InMainShip indicates that the player is in the ship.
	InMainShip uint32 = 0x01000000
	// InFighter indicates that the player is in a fighter.
	InFighter uint32 = 0x02000000
	// InSRV indicates that the player is in an SRV.
	InSRV uint32 = 0x04000000
	// InAnalysisMode indicates that analysis mode is selected.
	InAnalysisMode uint32 = 0x08000000
	// NightVision indicates that night vision is enabled.
	NightVision uint32 = 0x10000000
	// AltitudeFromAverageRadius indicates that the altitude value is based on the planet's average radius
	// (used at higher altitudes). If it is not set, the Altitude value is based on a raycast to the
	// actual surface below the ship/SRV.
	AltitudeFromAverageRadius uint32 = 0x20000000
	// FSDJump indicates that the ship is undergoing an FSD jump.
	FSDJump uint32 = 0x40000000
	// SRVHighBeam indicates that the SRV's high beams are on.
	SRVHighBeam uint32 = 0x80000000
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
