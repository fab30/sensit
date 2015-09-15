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
	defaultHTTPPort       = 8080
	defaultTimeseriesDB   = "sensit"
	defaultTimeseriesHost = "localhost"
	defaultTimeseriesPort = 8086 // default IndexDB port

	defaultConfigFileMode   = os.FileMode(0644)
	defaultConfigFolderMode = os.FileMode(0755)
)

// ApplicationConfig defines the format of an application config file
type ApplicationConfig struct {
	HTTP server.HTTPConfig   `json:"http"`
	DB   timeseries.DBConfig `json:"db"`
}

// InitConfig command initiailize the config file
func InitConfig(c *cli.Context) {
	// Create the default config
	defaultConfig := defaultConfig()

	jsonConfig, err := json.MarshalIndent(defaultConfig, "", "    ")

	if err != nil {
		log.Fatalf("Could not generate json : %s", err)
		return
	}
	// Parse the path
	configPath := c.GlobalString("config")
	if absolutePath, pathErr := filepath.Abs(configPath); pathErr == nil {

		dir := filepath.Dir(absolutePath)
		err = os.MkdirAll(dir, defaultConfigFolderMode)
		if err != nil {
			log.Fatalf("Could not create directory for %s", absolutePath)
			return
		}

		err = ioutil.WriteFile(absolutePath, jsonConfig, defaultConfigFileMode)
		if err != nil {
			log.Fatalf("Could not write config to file %s : %s", absolutePath, err)
		}
	} else {
		log.Fatalf("Could not get config path %s", pathErr)
	}
}

// DefaultConfig returns an application config filled with default values
func defaultConfig() *ApplicationConfig {
	return &ApplicationConfig{
		HTTP: server.HTTPConfig{
			ListenPort: defaultHTTPPort,
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
	if err == nil {
		decoder := json.NewDecoder(file)
		configuration := defaultConfig()

		err = decoder.Decode(configuration)
		if err != nil {
			return nil, err
		}
		return configuration, nil
	}
	return nil, err
}
