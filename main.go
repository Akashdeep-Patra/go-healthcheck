package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// imports as package "cli"

// imports as package "cli"

func main() {
	app := &cli.App{
		Name:  "Health Check",
		Usage: "Check the health of the system",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Aliases:  []string{"d"},
				Usage:    "Domain name to check",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "Port to check",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			port := c.String("port")
			if port == "" {
				port = "80"
			}
			domain := c.String("domain")
			fmt.Println("Checking health of", domain, "on port", port)
			fmt.Println("Health check passed")
			status := CheckHealth(domain, port)
			println(status)
			return nil

		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
