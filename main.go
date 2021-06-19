package main

import (
	"os"

	"github.com/urfave/cli/v2"
	log "unknwon.dev/clog/v2"

	"github.com/wuhan005/Scrooge/internal/cmd"
)

var (
	CommitSHA string
)

func main() {
	app := cli.NewApp()
	app.Name = "Scrooge"
	app.Usage = ""
	app.Version = ""
	app.Commands = []*cli.Command{
		cmd.Web,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal("Failed to start application: %v", err)
	}
}
