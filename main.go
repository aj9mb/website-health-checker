package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:      "website-health-checker",
		Usage:     "Tool to check whether a website is down",
		UsageText: "website-health-checker [-d | --domain <domain>] [-p | --port <port>] [-t | --timeout <seconds>]",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Aliases:  []string{"d"},
				Usage:    "pass the `DOMAIN` of site which needs to be checked",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "specifies which `PORT` to use",
				Required: false,
				Value:    "80",
			},
			&cli.IntFlag{
				Name:     "timeout",
				Aliases:  []string{"t"},
				Usage:    "can be used to specify max timeout in seconds `TIMEOUT`",
				Required: false,
				Value:    5,
			},
		},
		Action: func(c *cli.Context) error {
			status := Check(c.String("domain"), c.String("port"), c.Int("timeout"))
			fmt.Println(status)
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
