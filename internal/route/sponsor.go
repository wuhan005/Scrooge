// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package route

import (
	log "unknwon.dev/clog/v2"

	"github.com/wuhan005/Scrooge/internal/context"
	"github.com/wuhan005/Scrooge/internal/db"
)

type Sponsor struct{}

// NewSponsorHandler creates a new Sponsor.
func NewSponsorHandler() *Sponsor {
	return &Sponsor{}
}

// List returns the sponsor list.
func (*Sponsor) List(ctx context.Context) error {
	invoices, err := db.Invoices.Get(ctx.Request().Context(), db.GetInvoiceOptions{})
	if err != nil {
		log.Error("Failed to get invoices: %v", err)
		return ctx.ServerError()
	}

	sponsorSet := make(map[string][]*db.Invoice)

	for _, invoice := range invoices {
		invoice := invoice
		if sponsorSet[invoice.SponsorName] == nil {
			sponsorSet[invoice.SponsorName] = make([]*db.Invoice, 0)
		}

		sponsorSet[invoice.SponsorName] = append(sponsorSet[invoice.SponsorName], invoice)
	}
}
