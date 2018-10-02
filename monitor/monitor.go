// +build windows

package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/BenJuan26/elite"
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
	Description string
	PNPDeviceID string
}

type errorNoSerialConnection struct {
	message string
}

func (e *errorNoSerialConnection) Error() string {
	return e.message
}

func getSerialPort(pnp string) (*serial.Port, error) {
	var dst []serialPort
	escaped := strings.Replace(pnp, "\\", "\\\\", -1)
	query := "SELECT DeviceID, MaxBaudRate FROM Win32_SerialPort WHERE PNPDeviceID='" + escaped + "'"
	client := &wmi.Client{AllowMissingFields: true}
	err := client.Query(query, &dst)
	if err != nil {
		return nil, fmt.Errorf("Couldn't connect to serial port: %s", err.Error())
	} else if len(dst) < 1 {
		return nil, fmt.Errorf("Couldn't find a PNP Device with ID '%s'", pnp)
	}

	conf := &serial.Config{Name: dst[0].DeviceID, Baud: getBaudRate()}
	s, err := serial.OpenPort(conf)
	if err != nil {
		return nil, fmt.Errorf("Couldn't open serial port: %s", err.Error())
	}

	elog.Info(1, fmt.Sprintf("Connected to serial port %s at baud rate %d\n", dst[0].DeviceID, getBaudRate()))
	return s, nil
}

var errorCount = 0
var lastStatus = &elite.Status{}
var lastSystem = ""
var s *serial.Port

func checkStatusAndUpdate() error {
	elog.Info(1, "In loop")
	if errorCount > 20 {
		return fmt.Errorf("Too many consecutive errors")
	}

	var err error
	if s == nil {
		s, err = getSerialPort(getPNPDeviceID())
		if err != nil {
			elog.Error(1, "No serial connection: "+err.Error())
			return &errorNoSerialConnection{err.Error()}
		}

	}

	status, err := elite.GetStatus()
	if err != nil {
		errorCount = errorCount + 1
		elog.Error(1, "Couldn't get status: "+err.Error())
		if errorCount > 1 {
			elog.Error(1, fmt.Sprintf("Now at %d consecutive errors", errorCount))
		}
		return nil
	}

	system, err := elite.GetStarSystem()
	if err != nil {
		errorCount = errorCount + 1
		elog.Error(1, "Couldn't get star system: "+err.Error())
		if errorCount > 1 {
			elog.Error(1, fmt.Sprintf("Now at %d consecutive errors", errorCount))
		}
		return nil
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
			elog.Error(1, "Couldn't marshal JSON to send to serial: "+err.Error())
			if errorCount > 1 {
				elog.Error(1, fmt.Sprintf("Now at %d consecutive errors", errorCount))
			}
			return nil
		}

		_, err = s.Write(infoBytes)
		if err != nil {
			errorCount = errorCount + 1
			s = nil
			elog.Error(1, "Couldn't write to serial port: "+err.Error())
			return nil
		}
	}

	errorCount = 0
	return nil
}
