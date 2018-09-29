package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

type config struct {
	DeviceDescription string `json:"device_description"`
	PollInterval      int    `json:"poll_interval"`
}

var configData *config

// TODO: Allow for a cmdline flag specifying the config file path
func init() {
	loadConfig("config.json")
}

func loadConfig(path string) {
	_, err := os.Stat(path)
	if err != nil {
		fmt.Println("WARNING: config.json not found! Using sample values for now, see config-sample.json for a sample")
		configData = &config{"Arduino Uno", 5}
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

// GetDeviceDescription gets the configured serial device description
func GetDeviceDescription() string {
	return configData.DeviceDescription
}

// GetPollInterval gets the configured polling interval in milliseconds
func GetPollInterval() int {
	return configData.PollInterval
}
