// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package route

import (
	"time"

	log "unknwon.dev/clog/v2"

	"github.com/wuhan005/Scrooge/internal/context"
	"github.com/wuhan005/Scrooge/internal/db"
)

type Home struct{}

// NewHomeHandler creates a new Home.
func NewHomeHandler() *Home {
	return &Home{}
}

func (*Home) Profile(ctx context.Context) error {
	return ctx.Success(map[string]interface{}{
		"avatar_url":  "https://avatars.githubusercontent.com/u/12731778",
		"name":        "E99p1ant",
		"slogan":      "Be cool, but also be warm.",
		"description": "说点什么？",
	})
}

// Tiers returns the tiers.
func (*Home) Tiers(ctx context.Context) error {
	type tier struct {
		Amount  int    `json:"amount"`
		Comment string `json:"comment"`
	}

	// TODO read from database.
	tiers := []tier{
		{Amount: 5, Comment: "谢谢老板"},
		{Amount: 10, Comment: "谢谢老板"},
		{Amount: 50, Comment: "谢谢老板"},
		{Amount: 100, Comment: "谢谢老板"},
		{Amount: 0, Comment: "谢谢老板"},
	}
	return ctx.Success(tiers)
}

// List returns the sponsor list.
func (*Home) List(ctx context.Context) error {
	// Get all the paid invoices.
	paidInvoices, err := db.Invoices.Get(ctx.Request().Context(), db.GetInvoiceOptions{
		Paid: true,
	})
	if err != nil {
		log.Error("Failed to get invoices: %v", err)
		return ctx.ServerError()
	}

	// Sort by sponsor name.
	sponsor := make(map[string][]*db.Invoice)
	for _, invoice := range paidInvoices {
		invoice := invoice
		if sponsor[invoice.SponsorName] == nil {
			sponsor[invoice.SponsorName] = make([]*db.Invoice, 0)
		}
		sponsor[invoice.SponsorName] = append(sponsor[invoice.SponsorName], invoice)
	}

	type sponsorInvoice struct {
		PriceCents int       `json:"price_cents"`
		Comment    string    `json:"comment"`
		CreatedAt  time.Time `json:"created_at"`
	}

	type sponsorItem struct {
		Index    int               `json:"index"`
		Name     string            `json:"name"`
		Subtotal int               `json:"subtotal"`
		Count    int               `json:"count"`
		Invoices []*sponsorInvoice `json:"invoices"`
	}

	var resp []*sponsorItem

	list, err := db.Invoices.List(ctx.Request().Context())
	if err != nil {
		log.Error("Failed to get sponsor list: %v", err)
		return ctx.ServerError()
	}

	for i, item := range list {
		invoices := make([]*sponsorInvoice, 0, len(sponsor[item.SponsorName]))

		for _, invoice := range sponsor[item.SponsorName] {
			invoices = append(invoices, &sponsorInvoice{
				PriceCents: invoice.PriceCents,
				Comment:    invoice.Comment,
				CreatedAt:  invoice.CreatedAt,
			})
		}

		resp = append(resp, &sponsorItem{
			Index:    i + 1,
			Name:     item.SponsorName,
			Subtotal: item.Subtotal,
			Count:    len(invoices),
			Invoices: invoices,
		})
	}

	return ctx.Success(resp)
}
