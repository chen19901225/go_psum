package main

import (
	"go_psum/pkg/runner"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	var nameRaw, excludeName string
	var showDetail = 0
	var verbose = 1
	app := &cli.App{
		Name:    "go_psum",
		Usage:   "go_psum ",
		Version: "0.0.1",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "name",
				Usage:       "search process by name",
				Required:    true,
				Destination: &nameRaw,
			},
			&cli.StringFlag{
				Name:        "exclude",
				Usage:       "exclude process name",
				Required:    false,
				Destination: &excludeName,
			},
			&cli.IntFlag{
				Name:        "show",
				Usage:       "show detail?",
				Required:    false,
				DefaultText: "1",
				Destination: &showDetail,
			},
			&cli.IntFlag{
				Name:        "verbose",
				Usage:       "show log detail",
				Required:    false,
				Value:       1,
				DefaultText: "1",
				Destination: &verbose,
			},
		},
		Action: func(c *cli.Context) error {
			runner.Run(nameRaw, excludeName, showDetail, verbose)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}

}
