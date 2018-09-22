package main

import (
	"encoding/binary"
	"time"

	"github.com/BenJuan26/elite"
	"github.com/tarm/serial"
)

func main() {
	conf := &serial.Config{Name: "COM6", Baud: 9600}
	s, err := serial.OpenPort(conf)
	if err != nil {
		panic("Couldn't open port: " + err.Error())
	}

	lastMod := time.Date(1980, time.January, 1, 0, 0, 0, 0, time.UTC)
	for {
		mod, err := elite.GetStatusLastModified()
		if err != nil {
			panic(err)
		}

		if mod.After(lastMod) {
			status, err := elite.GetStatus()
			if err != nil {
				panic(err)
			}

			bytes := make([]byte, 4)
			binary.BigEndian.PutUint32(bytes, status.Flags)
			_, err = s.Write(bytes)
			if err != nil {
				panic("Couldn't write to device: " + err.Error())
			}
		}

		time.Sleep(50 * time.Millisecond)
	}
}
