package elite

import (
	"fmt"
	"testing"
)

func TestGetStatusFromString(t *testing.T) {
	status, err := GetStatusFromString(`{"timestamp":"2017-12-07T10:31:37Z", "event":"Status", "Flags":16842765, "Pips":[2,8,2], "FireGroup":0, "GuiFocus":5}`)
	if err != nil {
		fmt.Println("Couldn't get status: " + err.Error())
		t.FailNow()
	}

	flags := status.ExpandFlags()
	if !flags.Docked || !flags.ShieldsUp || !flags.InMainShip {
		fmt.Println("Parsed flags were incorrect")
		t.FailNow()
	}
}

func TestGetStatus(t *testing.T) {
	status, err := GetStatus()
	if err != nil {
		fmt.Println("Couldn't get status: " + err.Error())
		t.FailNow()
	}

	flags := status.ExpandFlags()
	if !flags.Docked || !flags.ShieldsUp || !flags.InMainShip {
		fmt.Println("Parsed flags were incorrect")
		t.FailNow()
	}
}
