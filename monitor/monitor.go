// +build windows

package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/BenJuan26/elite"
	"github.com/BenJuan26/elite/monitor/config"
	"github.com/StackExchange/wmi"
	"github.com/tarm/serial"
)

type controllerInfo struct {
	Timestamp  string   `json:"timestamp"`
	Flags      uint32   `json:"Flags"`
	Pips       [3]int32 `json:"Pips"`
	FireGroup  int32    `json:"FireGroup"`
	StarSystem string   `json:"StarSystem"`
}

type serialPort struct {
	MaxBaudRate int
	DeviceID    string
}

func getSerialPort(deviceDescription string) *serial.Port {
	var dst []serialPort
	err := wmi.Query("SELECT DeviceID, MaxBaudRate FROM Win32_SerialPort WHERE Description='"+deviceDescription+"'", &dst)
	if err != nil {
		panic(err)
	} else if len(dst) < 1 {
		panic("Couldn't find a serial device with a description matching '" + deviceDescription + "'; is the config.json correct?")
	}

	conf := &serial.Config{Name: dst[0].DeviceID, Baud: config.GetBaudRate()}
	s, err := serial.OpenPort(conf)
	if err != nil {
		panic("Couldn't open serial port: " + err.Error())
	}

	fmt.Printf("Connected to serial port %s at baud rate %d\n", dst[0].DeviceID, config.GetBaudRate())

	return s
}

func main() {
	s := getSerialPort(config.GetDeviceDescription())

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

			n, err := s.Write(infoBytes)
			if err != nil {
				fmt.Println(err)
				errorCount = errorCount + 1
				time.Sleep(5 * time.Millisecond)
				continue
			} else {
				fmt.Printf("Wrote %d bytes\n", n)
			}
		}

		errorCount = 0
		time.Sleep(time.Duration(config.GetPollInterval()) * time.Millisecond)
	}
}
