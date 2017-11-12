package main

import (
	"Matts-Macbook-Pro.local.com/climbpro/game"

	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "climbpro"

	app.Action = func(c *cli.Context) {
		game.Solve()
	}

	app.Run(os.Args)
}
