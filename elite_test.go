package elite_test

import (
	"fmt"
	"testing"

	"github.com/BenJuan26/elite"
)

var testLogPath = "./test"

func TestGetStatusFromString(t *testing.T) {
	status, err := elite.GetStatusFromBytes([]byte(`{"timestamp":"2017-12-07T10:31:37Z", "event":"Status", "Flags":16842765, "Pips":[2,8,2], "FireGroup":0, "GuiFocus":5}`))
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

func TestGetStatusFromPath(t *testing.T) {
	status, err := elite.GetStatusFromPath(testLogPath)
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
	_, err := elite.GetStatus()
	if err != nil {
		fmt.Println("Couldn't get status: " + err.Error())
		t.FailNow()
	}
}

func TestGetStarSystemFromPath(t *testing.T) {
	sys, err := elite.GetStarSystemFromPath(testLogPath)
	if err != nil {
		fmt.Println("An error occurred while getting the star system: " + err.Error())
		t.FailNow()
	}

	if sys != "Sol" {
		fmt.Println("Incorrect star system: Expecting Sol, got " + sys)
		t.FailNow()
	}
}

func TestGetLoadoutFromPath(t *testing.T) {
	loadout, err := elite.GetLoadoutFromPath(testLogPath)
	if err != nil {
		fmt.Println("Couldn't get loadout: " + err.Error())
		t.FailNow()
	}

	expectedShipName := "dora winifred"
	if loadout.ShipName != expectedShipName {
		fmt.Printf("Incorrect ship name: Expecting %s, got %s\n", expectedShipName, loadout.ShipName)
		t.FailNow()
	}
}

func Example() {
	// Errors not handled here
	system, _ := elite.GetStarSystem()
	fmt.Println("Current star system is " + system)

	status, _ := elite.GetStatus()
	if status.Flags.Docked {
		fmt.Println("Ship is docked")
	} else {
		fmt.Println("Ship is not docked")
	}
}
