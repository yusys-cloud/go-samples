// Author: yangzq80@gmail.com
// Date: 2020-11-26
// https://github.com/urfave/cli/blob/master/docs/v2/manual.md#arguments
package cli

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"testing"
)

func TestBase(t *testing.T) {
	app := &cli.App{
		Name:     "monitor",
		Usage:    "Cluster state monitor",
		HelpName: "cluster-monitor",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "clusterType",
				Aliases: []string{"c"},
				Value:   "redis",
				Usage:   "redis,rabbitMQ",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action: func(c *cli.Context) error {
					fmt.Println("added task: ", c.Args().First())
					return nil
				}, Flags: []cli.Flag{
					&cli.BoolFlag{Name: "serve", Aliases: []string{"s"}},
					&cli.BoolFlag{Name: "option", Aliases: []string{"o"}},
					&cli.StringFlag{Name: "message", Aliases: []string{"m"}},
				},
			},
			{
				Name:    "complete",
				Aliases: []string{"c"},
				Usage:   "complete a task on the list",
				Action: func(c *cli.Context) error {
					fmt.Println(c.String("message"))
					fmt.Println("completed task: ", c.Args().First())
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
