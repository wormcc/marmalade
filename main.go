package main

import (
	"github.com/urfave/cli"
	"github.com/wormcc/marmalade/cmd"
	"os"
)
const appVersion = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Name = "marmalade"
	app.Version = appVersion
	app.Usage = "A api admin platform"
	app.Commands = []cli.Command{
		cmd.Web,
		cmd.GenSecret,
	}
	app.Run(os.Args)
}
