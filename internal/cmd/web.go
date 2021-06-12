// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/flamego/flamego"
	"github.com/urfave/cli/v2"
)

var Web = &cli.Command{
	Name:  "web",
	Usage: "Start web server",
	Description: `Scrooge web server is the only thing you need to run,
and it takes care of all the other things for you`,
	Action: runWeb,
	Flags: []cli.Flag{
		intFlag("port, p", 19999, "Temporary port number to prevent conflict"),
	},
}

func runWeb(c *cli.Context) error {
	f := flamego.Classic()

	f.Run("0.0.0.0", c.Int("port"))
	return nil
}
