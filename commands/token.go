package commands

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
	"github.com/joelvim/sensit/authtoken"
)

// Token prints the token generated for the config credentials
func Token(c *cli.Context) {
	configPath := c.GlobalString("config")

	config, err := parseConfig(configPath)

	if err != nil {
		log.Fatalf("Error parsing the config : %s", err)
		return
	}

	fmt.Print(string(authtoken.Token(config.HTTP.Login, config.HTTP.Password, config.HTTP.Salt)))
}
