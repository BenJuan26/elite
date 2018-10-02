// +build windows

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"text/tabwriter"

	"github.com/StackExchange/wmi"
	"github.com/pkg/errors"
)

type config struct {
	PNPDeviceID string `json:"pnp_device_id"`
	BaudRate    int    `json:"baud_rate"`
}

var configData *config

func loadConfig(path string) {
	_, err := os.Stat(path)
	if err != nil {
		panic("config.json not found")
	}

	if configData == nil {
		configData = new(config)
		buff, err := ioutil.ReadFile(path)
		if err != nil {
			panic(errors.Wrap(err, "Problem reading config file"))
		}
		err = json.Unmarshal(buff, configData)
		if err != nil {
			panic(errors.Wrap(err, "Problem unmarshaling config structure"))
		}
	}
}

func configPath() string {
	exepath, err := exePath()
	if err != nil {
		panic("Couldn't find executable path: " + err.Error())
	}

	return filepath.FromSlash(filepath.Dir(exepath) + "/config.json")
}

// InteractiveConfig prompts the user for the necessary config parameters
// and writes them to config.json
func interactiveConfig() error {
	var dst []serialPort
	client := &wmi.Client{AllowMissingFields: true}
	err := client.Query("SELECT DeviceID, MaxBaudRate, PNPDeviceID, Description FROM Win32_SerialPort", &dst)
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)
	fmt.Fprintln(w, "#\tDevice ID\tMax Baud Rate\tDescription")
	for i, device := range dst {
		fmt.Fprintf(w, "%d)\t%s\t%d\t%s\n", i+1, device.DeviceID, device.MaxBaudRate, device.Description)
	}
	w.Flush()

	fmt.Printf("\nEnter selection or c to cancel: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	deviceSelection := scanner.Text()

	deviceIndex, err := strconv.ParseInt(deviceSelection, 10, 64)
	if err != nil {
		fmt.Println("Cancelled")
		return nil
	}
	if int(deviceIndex) > len(dst) || deviceIndex < 1 {
		return fmt.Errorf("Invalid selection: %s", deviceSelection)
	}

	pnp := dst[deviceIndex-1].PNPDeviceID

	fmt.Printf("Enter baud rate: ")
	scanner.Scan()
	baudRateSelection := scanner.Text()

	baudRate, err := strconv.ParseInt(baudRateSelection, 10, 32)
	if err != nil {
		return fmt.Errorf("Invalid selection: %s", baudRateSelection)
	}

	conf := config{pnp, int(baudRate)}
	confBytes, err := json.Marshal(conf)
	if err != nil {
		return fmt.Errorf("Couldn't marshal config into JSON: %s", err.Error())
	}

	err = ioutil.WriteFile(configPath(), confBytes, 0777)
	if err != nil {
		return fmt.Errorf("Couldn't write config.json file: %s", err.Error())
	}

	return nil
}

// GetPNPDeviceID gets the configured serial device PNP ID
func getPNPDeviceID() string {
	if configData == nil {
		loadConfig(configPath())
	}
	return configData.PNPDeviceID
}

// GetBaudRate gets the configured serial baud rate
func getBaudRate() int {
	if configData == nil {
		loadConfig(configPath())
	}
	return configData.BaudRate
}
