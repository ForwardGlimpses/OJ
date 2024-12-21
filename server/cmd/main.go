package main

import (
	"os"

	"github.com/ForwardGlimpses/OJ/server/pkg/bootstrap"
	"github.com/ForwardGlimpses/OJ/server/pkg/config"
	"github.com/ForwardGlimpses/OJ/server/pkg/logs"
	"github.com/urfave/cli/v2"
)

var VERSION = "v0.0.1"

func main() {
	logs.Init() //初始化日志配置

	app := cli.NewApp()
	app.Name = "console"
	app.Version = VERSION
	app.Commands = []*cli.Command{
		StartCmd,
	}
	err := app.Run(os.Args)
	if err != nil {
		logs.Error("Failed to start the application:", err)
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
			},
		},
		Action: func(c *cli.Context) error {
			err := config.Load(c.String("config"))
			if err != nil {
				logs.Error("Failed to load config:", err)
				return err
			}

			err = bootstrap.Run()
			if err != nil {
				logs.Error("Failed to run bootstrap:", err)
				panic(err)
			}
			return nil
		},
	}
)
