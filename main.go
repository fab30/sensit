package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/joelvim/sensit/commands"
)

func main() {

	app := cli.NewApp()
	app.Name = "sensit"
	app.Usage = "Manage sensit data"
	app.EnableBashCompletion = true
	app.Version = "0.1.0"

	// List the global flags of the application
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "config, c",
			Value:  "config.json",
			Usage:  "path to the configuration file",
			EnvVar: "SENSIT_CONFIG",
		},
	}

	// Application subcommands
	app.Commands = []cli.Command{
		{
			Name:   "daemon",
			Usage:  "launch the http server that receives data from the sensit server",
			Action: commands.HTTPServer,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "http-port",
					Value: 8080,
					Usage: "path to the configuration file",
				},
			},
		},
		{
			Name:   "init",
			Usage:  "init the config at the given path",
			Action: commands.InitConfig,
		},
	}

	app.Run(os.Args)
}
