package main

import (
	"os"

	"github.com/ForwardGlimpses/OJ/server/pkg/bootstrap"
	"github.com/ForwardGlimpses/OJ/server/pkg/config"

	//"github.com/gin-contrib/cors"
	//"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

var VERSION = "v0.0.1"

func main() {
	app := cli.NewApp()
	app.Name = "console"
	app.Version = VERSION
	app.Commands = []*cli.Command{
		StartCmd,
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

var (
	StartCmd = &cli.Command{
		Name:  "start",
		Usage: "Start server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Runtime configuration files",
				Value:   "configs/config.json",
			},
		},
		Action: func(c *cli.Context) error {
			err := config.Load(c.String("config"))
			if err != nil {
				return err
			}

			err = bootstrap.Run()
			if err != nil {
				return err
			}
			return nil
		},
	}
)
