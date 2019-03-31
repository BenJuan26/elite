package elite

import (
	"fmt"
	"testing"
)

func TestGetStatusFromString(t *testing.T) {
	status, err := GetStatusFromBytes([]byte(`{"timestamp":"2017-12-07T10:31:37Z", "event":"Status", "Flags":16842765, "Pips":[2,8,2], "FireGroup":0, "GuiFocus":5}`))
	if err != nil {
		fmt.Println("Couldn't get status: " + err.Error())
		t.FailNow()
	}

	if !status.Flags.Docked || !status.Flags.ShieldsUp || !status.Flags.InMainShip || !status.Flags.LandingGearDown || !status.Flags.FSDMassLocked {
		fmt.Println("Parsed flags were incorrect")
		t.FailNow()
	}

	if status.Flags.Landed ||
		status.Flags.Supercruise ||
		status.Flags.FlightAssistOff ||
		status.Flags.HardpointsDeployed ||
		status.Flags.InWing ||
		status.Flags.LightsOn ||
		status.Flags.CargoScoopDeployed ||
		status.Flags.SilentRunning ||
		status.Flags.ScoopingFuel ||
		status.Flags.SRVHandbrake ||
		status.Flags.SRVTurret ||
		status.Flags.SRVUnderShip ||
		status.Flags.SRVDriveAssist ||
		status.Flags.FSDCharging ||
		status.Flags.FSDCooldown ||
		status.Flags.LowFuel ||
		status.Flags.Overheating ||
		status.Flags.HasLatLong ||
		status.Flags.IsInDanger ||
		status.Flags.BeingInterdicted ||
		status.Flags.InFighter ||
		status.Flags.InSRV ||
		status.Flags.InAnalysisMode ||
		status.Flags.NightVision {
		fmt.Println("Parsed flags were incorrect")
		t.FailNow()
	}
}

func TestGetStatus(t *testing.T) {
	_, err := GetStatus()
	if err != nil {
		fmt.Println("Couldn't get status: " + err.Error())
		t.FailNow()
	}
}
