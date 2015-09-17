package commands

import (
	"log"

	"github.com/codegangsta/cli"
	"github.com/joelvim/sensit/server"
)

// HTTPServer command launches the HTTP server
func HTTPServer(c *cli.Context) {

	configPath := c.GlobalString("config")

	config, err := parseConfig(configPath)

	if err != nil {
		log.Fatalf("Error parsing the config : %s", err)
		return
	}

	config.HTTP.ListenPort = c.Int("http-port")

	// Serve the API
	server.API(config.DB, config.HTTP)
}
