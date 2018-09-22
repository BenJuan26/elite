package main

import (
	"encoding/binary"
	"time"

	"github.com/BenJuan26/elite"
	"github.com/tarm/serial"
)

var buff = make([]byte, 4)

func main() {
	conf := &serial.Config{Name: "COM6", Baud: 9600}
	s, err := serial.OpenPort(conf)
	if err != nil {
		panic("Couldn't open port: " + err.Error())
	}

	errorCount := 0
	lastTime := ""
	for {
		status, err := elite.GetStatus()
		if err != nil {
			errorCount = errorCount + 1
			if errorCount > 20 {
				panic("Can't read status file: " + err.Error())
			}
		} else if status.Timestamp != lastTime {
			errorCount = 0

			binary.BigEndian.PutUint32(buff, status.Flags)
			_, err = s.Write(buff)
			if err != nil {
				panic("Couldn't write to device: " + err.Error())
			}
		}

		time.Sleep(5 * time.Millisecond)
	}
}
