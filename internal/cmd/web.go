// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"io/fs"
	"net/http"

	"github.com/Cardinal-Platform/binding"
	"github.com/flamego/flamego"
	"github.com/urfave/cli/v2"
	log "unknwon.dev/clog/v2"

	"github.com/wuhan005/Scrooge/frontend"
	"github.com/wuhan005/Scrooge/internal/context"
	"github.com/wuhan005/Scrooge/internal/db"
	"github.com/wuhan005/Scrooge/internal/form"
	"github.com/wuhan005/Scrooge/internal/paybob"
	"github.com/wuhan005/Scrooge/internal/route"
	"github.com/wuhan005/Scrooge/internal/static"
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
	err := db.Init()
	if err != nil {
		log.Fatal("Failed to connect to database: %v", err)
	}

	f := flamego.Classic()

	paymenter := func(ctx context.Context) {
		client := paybob.NewDefaultClient()
		ctx.Map(client)
	}

	fe, err := fs.Sub(frontend.FS, "dist")
	if err != nil {
		log.Fatal("Failed to sub filesystem: %v", err)
	}

	f.Use(flamego.Static(flamego.StaticOptions{
		FileSystem: http.FS(fe),
	}))
	f.NotFound(static.NotFound(http.FS(fe), "index.html"))

	f.Group("/api", func() {
		home := route.NewHomeHandler()
		f.Get("/profile", home.Profile)
		f.Get("/tiers", home.Tiers)
		f.Get("/sponsor_list", home.List)

		pay := route.NewPayHandler()
		f.Group("/pay", func() {
			f.Post("", binding.Bind(form.NewPayment{}), pay.NewInvoice)
			f.Get("/query", pay.Query)
			f.Get("/cashier", pay.Cashier)
			f.Post("/callback", pay.Callback)
		}, paymenter)
	})

	f.Use(context.Contexter())

	// TODO remove the temporary CORS header.
	f.Use(func(ctx context.Context) {
		ctx.ResponseWriter().Header().Set("Access-Control-Allow-Origin", "*")
		ctx.ResponseWriter().Header().Set("Access-Control-Allow-Headers", "*")
		if ctx.Request().Method == http.MethodOptions {
			ctx.ResponseWriter().WriteHeader(http.StatusOK)
		}
	})

	f.Run("0.0.0.0", c.Int("port"))
	return nil
}
