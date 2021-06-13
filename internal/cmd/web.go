// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/Cardinal-Platform/binding"
	"github.com/flamego/flamego"
	"github.com/urfave/cli/v2"

	"github.com/wuhan005/Scrooge/internal/context"
	"github.com/wuhan005/Scrooge/internal/form"
	"github.com/wuhan005/Scrooge/internal/paybob"
	"github.com/wuhan005/Scrooge/internal/route"
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

	paymenter := func(ctx context.Context) {
		client := paybob.NewDefaultClient()
		ctx.Map(client)
	}

	f.Get("/")

	f.Group("/api", func() {
		sponsor := route.NewSponsorHandler()
		f.Get("/sponsor_list", sponsor.List)

		pay := route.NewPayHandler()
		f.Group("/pay", func() {
			f.Post("", binding.Bind(form.NewPayment{}), pay.NewInvoice)
			f.Get("/query", pay.Query)
			f.Get("/cashier", pay.Cashier)
			f.Get("/callback", pay.Callback)
		}, paymenter)
	})

	f.Use(context.Contexter())

	f.Run("0.0.0.0", c.Int("port"))
	return nil
}
