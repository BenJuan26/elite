package main

import (
	"encoding/json"
	"time"

	"github.com/BenJuan26/elite"
	"github.com/tarm/serial"
)

type controllerInfo struct {
	Timestamp  string   `json:"timestamp"`
	Flags      uint32   `json:"Flags"`
	Pips       [3]int32 `json:"Pips"`
	FireGroup  int32    `json:"FireGroup"`
	StarSystem string   `json:"StarSystem"`
}

func main() {
	conf := &serial.Config{Name: "COM6", Baud: 9600}
	s, err := serial.OpenPort(conf)
	if err != nil {
		panic("Couldn't open port: " + err.Error())
	}

	errorCount := 0
	lastStatus := &elite.Status{}
	lastSystem := ""
	for {
		if errorCount > 20 {
			panic("Too many errors")
		}

		status, err := elite.GetStatus()
		if err != nil {
			errorCount = errorCount + 1
			time.Sleep(5 * time.Millisecond)
			continue
		}

		system, err := elite.GetStarSystem()
		if err != nil {
			errorCount = errorCount + 1
			time.Sleep(5 * time.Millisecond)
			continue
		}

		if status.Timestamp != lastStatus.Timestamp || lastSystem != system {
			lastStatus = status
			lastSystem = system

			info := controllerInfo{
				Timestamp:  status.Timestamp,
				Flags:      status.Flags,
				Pips:       status.Pips,
				FireGroup:  status.FireGroup,
				StarSystem: system,
			}

			infoBytes, err := json.Marshal(info)
			if err != nil {
				errorCount = errorCount + 1
				time.Sleep(5 * time.Millisecond)
				continue
			}

			_, err = s.Write(infoBytes)
			if err != nil {
				errorCount = errorCount + 1
				time.Sleep(5 * time.Millisecond)
				continue
			}
		}

		errorCount = 0
		time.Sleep(5 * time.Millisecond)
	}
}
