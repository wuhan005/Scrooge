// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestInvoices(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	t.Parallel()

	db, cleanup := newTestDB(t)
	store := NewInvoiceStore(db)

	for _, tc := range []struct {
		name string
		test func(t *testing.T, ctx context.Context, db *invoices)
	}{
		{"Create", testInvoiceCreate},
		{"Update", testInvoiceUpdate},
		{"Get", testInvoiceGet},
		{"List", testInvoiceList},
		{"GetByID", testInvoiceGetByID},
		{"GetByUID", testInvoiceGetByUID},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Cleanup(func() {
				err := cleanup("invoices")
				if err != nil {
					t.Fatal(err)
				}
			})
			tc.test(t, context.Background(), store.(*invoices))
		})
	}
}

func testInvoiceCreate(t *testing.T, ctx context.Context, db *invoices) {
	uid, err := db.Create(ctx, CreateInvoiceOptions{
		OrderID:       "579a57f933397e0f441ba37f239d3721",
		PriceCents:    2000, // ￥20.00
		SponsorName:   "Scrooge",
		SponsorOpenID: "876742a2b7950be1491959b76713606a",
		Comment:       "Well Done!",
	})
	assert.Nil(t, err)
	assert.NotZero(t, uid)

	got, err := db.GetByUID(ctx, uid)
	assert.Nil(t, err)

	got.CreatedAt = time.Time{}
	got.UpdatedAt = time.Time{}

	want := &Invoice{
		Model: gorm.Model{
			ID: 1,
		},
		UID:           uid,
		OrderID:       "579a57f933397e0f441ba37f239d3721",
		PriceCents:    2000,
		SponsorName:   "Scrooge",
		SponsorOpenID: "876742a2b7950be1491959b76713606a",
		Comment:       "Well Done!",
	}
	assert.Equal(t, want, got)
}

func testInvoiceUpdate(t *testing.T, ctx context.Context, db *invoices) {
	uid, err := db.Create(ctx, CreateInvoiceOptions{
		PriceCents:  2000, // ￥20.00
		SponsorName: "Scrooge",
		Comment:     "Well Done!",
	})
	assert.Nil(t, err)
	assert.NotZero(t, uid)

	err = db.Update(ctx, uid, UpdateInvoiceOptions{
		OrderID:       "579a57f933397e0f441ba37f239d3721",
		Paid:          true,
		SponsorOpenID: "876742a2b7950be1491959b76713606a",
	})
	assert.Nil(t, err)

	got, err := db.GetByUID(ctx, uid)
	assert.Nil(t, err)

	got.CreatedAt = time.Time{}
	got.UpdatedAt = time.Time{}

	want := &Invoice{
		Model: gorm.Model{
			ID: 1,
		},
		UID:           uid,
		OrderID:       "579a57f933397e0f441ba37f239d3721",
		Paid:          true,
		PriceCents:    2000,
		SponsorName:   "Scrooge",
		SponsorOpenID: "876742a2b7950be1491959b76713606a",
		Comment:       "Well Done!",
	}
	assert.Equal(t, want, got)
}

func testInvoiceGet(t *testing.T, ctx context.Context, db *invoices) {
	uid1, err := db.Create(ctx, CreateInvoiceOptions{
		OrderID:       "579a57f933397e0f441ba37f239d3721",
		PriceCents:    2000, // ￥20.00
		SponsorName:   "Scrooge",
		SponsorOpenID: "876742a2b7950be1491959b76713606a",
		Comment:       "Well Done!",
	})
	assert.Nil(t, err)
	assert.NotZero(t, uid1)

	uid2, err := db.Create(ctx, CreateInvoiceOptions{
		OrderID:       "9e66623ec3649dd2eabdb2b711ad18bf",
		PriceCents:    5000, // ￥50.00
		SponsorName:   "Scrooge Mcduck",
		SponsorOpenID: "8b3348fe50baa6bb487fd931203a3d73",
	})
	assert.Nil(t, err)
	assert.NotZero(t, uid2)

	uid3, err := db.Create(ctx, CreateInvoiceOptions{
		OrderID:       "a0270de1b2e9279410829d2f6fb831bc",
		PriceCents:    8000, // ￥80.00
		SponsorName:   "Scrooge",
		SponsorOpenID: "876742a2b7950be1491959b76713606a",
		Comment:       "Excellent!",
	})
	assert.Nil(t, err)
	assert.NotZero(t, uid3)

	// Get all the invoices.
	got, err := db.Get(ctx, GetInvoiceOptions{})
	assert.Nil(t, err)

	for _, invoice := range got {
		invoice.UID = ""
		invoice.CreatedAt = time.Time{}
		invoice.UpdatedAt = time.Time{}
	}

	want := []*Invoice{
		{
			Model: gorm.Model{
				ID: 3,
			},
			OrderID:       "a0270de1b2e9279410829d2f6fb831bc",
			PriceCents:    8000,
			SponsorName:   "Scrooge",
			SponsorOpenID: "876742a2b7950be1491959b76713606a",
			Comment:       "Excellent!",
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			OrderID:       "9e66623ec3649dd2eabdb2b711ad18bf",
			PriceCents:    5000,
			SponsorName:   "Scrooge Mcduck",
			SponsorOpenID: "8b3348fe50baa6bb487fd931203a3d73",
			Comment:       "",
		},
		{
			Model: gorm.Model{
				ID: 1,
			},
			OrderID:       "579a57f933397e0f441ba37f239d3721",
			PriceCents:    2000,
			SponsorName:   "Scrooge",
			SponsorOpenID: "876742a2b7950be1491959b76713606a",
			Comment:       "Well Done!",
		},
	}
	assert.Equal(t, want, got)

	// Filter by sponsor name
	got, err = db.Get(ctx, GetInvoiceOptions{
		SponsorName: "Scrooge",
	})
	assert.Nil(t, err)

	for _, invoice := range got {
		invoice.UID = ""
		invoice.CreatedAt = time.Time{}
		invoice.UpdatedAt = time.Time{}
	}

	want = []*Invoice{
		{
			Model: gorm.Model{
				ID: 3,
			},
			OrderID:       "a0270de1b2e9279410829d2f6fb831bc",
			PriceCents:    8000,
			SponsorName:   "Scrooge",
			SponsorOpenID: "876742a2b7950be1491959b76713606a",
			Comment:       "Excellent!",
		},
		{
			Model: gorm.Model{
				ID: 1,
			},
			OrderID:       "579a57f933397e0f441ba37f239d3721",
			PriceCents:    2000,
			SponsorName:   "Scrooge",
			SponsorOpenID: "876742a2b7950be1491959b76713606a",
			Comment:       "Well Done!",
		},
	}
	assert.Equal(t, want, got)

	// Filer paid invoice

	// Invoice 2
	err = db.Update(ctx, uid2, UpdateInvoiceOptions{
		Paid: true,
	})
	assert.Nil(t, err)

	// Invoice 3
	err = db.Update(ctx, uid3, UpdateInvoiceOptions{
		Paid: true,
	})
	assert.Nil(t, err)

	got, err = db.Get(ctx, GetInvoiceOptions{
		Paid: true,
	})
	assert.Nil(t, err)

	for _, invoice := range got {
		invoice.UID = ""
		invoice.CreatedAt = time.Time{}
		invoice.UpdatedAt = time.Time{}
	}

	want = []*Invoice{
		{
			Model: gorm.Model{
				ID: 3,
			},
			OrderID:       "a0270de1b2e9279410829d2f6fb831bc",
			PriceCents:    8000,
			SponsorName:   "Scrooge",
			SponsorOpenID: "876742a2b7950be1491959b76713606a",
			Comment:       "Excellent!",
			Paid:          true,
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			OrderID:       "9e66623ec3649dd2eabdb2b711ad18bf",
			PriceCents:    5000,
			SponsorName:   "Scrooge Mcduck",
			SponsorOpenID: "8b3348fe50baa6bb487fd931203a3d73",
			Comment:       "",
			Paid:          true,
		},
	}
	assert.Equal(t, want, got)
}

func testInvoiceList(t *testing.T, ctx context.Context, db *invoices) {
	uid1, err := db.Create(ctx, CreateInvoiceOptions{
		OrderID:       "579a57f933397e0f441ba37f239d3721",
		PriceCents:    2000, // ￥20.00
		SponsorName:   "Scrooge",
		SponsorOpenID: "876742a2b7950be1491959b76713606a",
		Comment:       "Well Done!",
	})
	assert.Nil(t, err)
	assert.NotZero(t, uid1)
	err = db.Update(ctx, uid1, UpdateInvoiceOptions{Paid: true})
	assert.Nil(t, err)

	uid2, err := db.Create(ctx, CreateInvoiceOptions{
		OrderID:       "9e66623ec3649dd2eabdb2b711ad18bf",
		PriceCents:    5000, // ￥50.00
		SponsorName:   "E99p1ant",
		SponsorOpenID: "8b3348fe50baa6bb487fd931203a3d73",
	})
	assert.Nil(t, err)
	assert.NotZero(t, uid2)
	err = db.Update(ctx, uid2, UpdateInvoiceOptions{Paid: true})
	assert.Nil(t, err)

	uid3, err := db.Create(ctx, CreateInvoiceOptions{
		OrderID:       "51e92d5b304cccbe18b32ac108421657",
		PriceCents:    8000, // ￥80.00
		SponsorName:   "E99p1ant",
		SponsorOpenID: "8b3348fe50baa6bb487fd931203a3d73",
	})
	assert.Nil(t, err)
	assert.NotZero(t, uid3)
	err = db.Update(ctx, uid3, UpdateInvoiceOptions{Paid: true})
	assert.Nil(t, err)

	uid4, err := db.Create(ctx, CreateInvoiceOptions{
		OrderID:       "a0270de1b2e9279410829d2f6fb831bc",
		PriceCents:    12000, // ￥120.00
		SponsorName:   "Scrooge",
		SponsorOpenID: "876742a2b7950be1491959b76713606a",
		Comment:       "Excellent!",
	})
	assert.Nil(t, err)
	assert.NotZero(t, uid4)

	got, err := db.List(ctx)
	assert.Nil(t, err)

	want := []*SponsorList{
		{
			SponsorName: "E99p1ant",
			Subtotal:    13000,
		},
		{
			SponsorName: "Scrooge",
			Subtotal:    2000,
		},
	}
	assert.Equal(t, want, got)
}

func testInvoiceGetByID(t *testing.T, ctx context.Context, db *invoices) {
	uid, err := db.Create(ctx, CreateInvoiceOptions{
		OrderID:       "579a57f933397e0f441ba37f239d3721",
		PriceCents:    2000, // ￥20.00
		SponsorName:   "Scrooge",
		SponsorOpenID: "876742a2b7950be1491959b76713606a",
		Comment:       "Well Done!",
	})
	assert.Nil(t, err)
	assert.NotZero(t, uid)

	uid, err = db.Create(ctx, CreateInvoiceOptions{
		OrderID:       "9e66623ec3649dd2eabdb2b711ad18bf",
		PriceCents:    5000, // ￥50.00
		SponsorName:   "Scrooge Mcduck",
		SponsorOpenID: "8b3348fe50baa6bb487fd931203a3d73",
	})
	assert.Nil(t, err)
	assert.NotZero(t, uid)

	got, err := db.GetByID(ctx, 2)
	assert.Nil(t, err)

	got.UID = ""
	got.CreatedAt = time.Time{}
	got.UpdatedAt = time.Time{}

	want := &Invoice{
		Model: gorm.Model{
			ID: 2,
		},
		OrderID:       "9e66623ec3649dd2eabdb2b711ad18bf",
		PriceCents:    5000,
		SponsorName:   "Scrooge Mcduck",
		SponsorOpenID: "8b3348fe50baa6bb487fd931203a3d73",
		Comment:       "",
	}
	assert.Equal(t, want, got)
}

func testInvoiceGetByUID(t *testing.T, ctx context.Context, db *invoices) {
	uid, err := db.Create(ctx, CreateInvoiceOptions{
		OrderID:       "579a57f933397e0f441ba37f239d3721",
		PriceCents:    2000, // ￥20.00
		SponsorName:   "Scrooge",
		SponsorOpenID: "876742a2b7950be1491959b76713606a",
		Comment:       "Well Done!",
	})
	assert.Nil(t, err)
	assert.NotZero(t, uid)

	uid, err = db.Create(ctx, CreateInvoiceOptions{
		OrderID:       "9e66623ec3649dd2eabdb2b711ad18bf",
		PriceCents:    5000, // ￥50.00
		SponsorName:   "Scrooge Mcduck",
		SponsorOpenID: "8b3348fe50baa6bb487fd931203a3d73",
	})
	assert.Nil(t, err)
	assert.NotZero(t, uid)

	got, err := db.GetByUID(ctx, uid)
	assert.Nil(t, err)

	got.UID = ""
	got.CreatedAt = time.Time{}
	got.UpdatedAt = time.Time{}

	want := &Invoice{
		Model: gorm.Model{
			ID: 2,
		},
		OrderID:       "9e66623ec3649dd2eabdb2b711ad18bf",
		PriceCents:    5000,
		SponsorName:   "Scrooge Mcduck",
		SponsorOpenID: "8b3348fe50baa6bb487fd931203a3d73",
		Comment:       "",
	}
	assert.Equal(t, want, got)
}
