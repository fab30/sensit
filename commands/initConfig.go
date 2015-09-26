package commands

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
	"github.com/joelvim/sensit/server"
	"github.com/joelvim/sensit/timeseries"
)

const (
	defaultTimeseriesDB   = "sensit"
	defaultTimeseriesHost = "localhost"
	defaultTimeseriesPort = 8086 // default IndexDB port

	defaultConfigFileMode   = os.FileMode(0644)
	defaultConfigFolderMode = os.FileMode(0755)

	defaultAuthType = "Token"
)

// ApplicationConfig defines the format of an application config file
type ApplicationConfig struct {
	HTTP server.HTTPConfig   `json:"http"`
	DB   timeseries.DBConfig `json:"db"`
}

// InitConfig command initiailize the config file
func InitConfig(c *cli.Context) {
	// Parse the path
	configPath := c.GlobalString("config")

	log.Printf("Initializing config at %s", configPath)

	absolutePath, pathErr := filepath.Abs(configPath)
	if pathErr != nil {
		log.Fatalf("Could not get asbolute path for config path %s : %s", configPath, pathErr)
		return
	}

	dir := filepath.Dir(absolutePath)
	errMkDir := os.MkdirAll(dir, defaultConfigFolderMode)
	if errMkDir != nil {
		log.Fatalf("Could not create directory for %s : %s", absolutePath, errMkDir)
		return
	}
	// Create the default config
	defaultConfig := defaultConfig()
	jsonConfig, errJSON := json.MarshalIndent(defaultConfig, "", "    ")
	if errJSON != nil {
		log.Fatalf("Could not generate json : %s", errJSON)
		return
	}
	errFile := ioutil.WriteFile(absolutePath, jsonConfig, defaultConfigFileMode)
	if errFile != nil {
		log.Fatalf("Could not write config to file %s : %s", absolutePath, errFile)
		return
	}
	log.Printf("Successfuly wrote config file at %s", absolutePath)
}

// DefaultConfig returns an application config filled with default values
func defaultConfig() *ApplicationConfig {
	return &ApplicationConfig{
		HTTP: server.HTTPConfig{
			AuthType: defaultAuthType,
		},
		DB: timeseries.DBConfig{
			Host:     defaultTimeseriesHost,
			Port:     defaultTimeseriesPort,
			Database: defaultTimeseriesDB,
		},
	}
}

func parseConfig(configPath string) (*ApplicationConfig, error) {
	file, err := os.Open(configPath)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(file)
	configuration := defaultConfig()

	err = decoder.Decode(configuration)
	if err != nil {
		return nil, err
	}
	return configuration, nil
}
