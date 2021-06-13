// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package db

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
)

var _ InvoiceStore = (*invoices)(nil)

// Invoices is the default instance of the InvoiceStore.
var Invoices InvoiceStore

// InvoiceStore is the persistent interface for invoices.
type InvoiceStore interface {
	// Create creates a new invoice with the given options and returns the invoice's uid.
	Create(ctx context.Context, opts CreateInvoiceOptions) (string, error)
	// Update updates the invoice with the given uid.
	Update(ctx context.Context, uid string, opts UpdateInvoiceOptions) error
	// Get returns the invoices list with the given options.
	// The zero value in the options will be ignored.
	Get(ctx context.Context, opts GetInvoiceOptions) ([]*Invoice, error)
	// List returns the summary invoice data which contains the sponsor's information, subtotal, count.
	// It sorted by sponsor's subtotal.
	List(ctx context.Context) ([]*InvoiceSummary, error)
	// GetByID gets the invoice with the given id.
	// It returns ErrInvoiceNotExists error when the invoice dose not exist.
	GetByID(ctx context.Context, id uint) (*Invoice, error)
	// GetByUID gets the invoice with the given uid.
	// It returns ErrInvoiceNotExists error when the invoice dose not exist.
	GetByUID(ctx context.Context, uid string) (*Invoice, error)
}

// NewInvoiceStore returns a InvoiceStore instance with the given database connection.
func NewInvoiceStore(db *gorm.DB) InvoiceStore {
	return &invoices{DB: db}
}

// Invoice represents the invoice.
type Invoice struct {
	gorm.Model

	UID        string
	OrderID    string
	Paid       bool
	PriceCents int

	SponsorName    string
	SponsorWebSite string
	SponsorOpenID  string

	Comment string
}

type invoices struct {
	*gorm.DB
}

type CreateInvoiceOptions struct {
	OrderID        string
	PriceCents     int
	SponsorName    string
	SponsorWebSite string
	SponsorOpenID  string
	Comment        string
}

func (db *invoices) Create(ctx context.Context, opts CreateInvoiceOptions) (string, error) {
	uid := strings.ReplaceAll(uuid.NewV4().String(), "-", "")

	return uid, db.WithContext(ctx).Create(&Invoice{
		UID:            uid,
		OrderID:        opts.OrderID,
		PriceCents:     opts.PriceCents,
		SponsorName:    opts.SponsorName,
		SponsorWebSite: opts.SponsorWebSite,
		SponsorOpenID:  opts.SponsorOpenID,
		Comment:        opts.Comment,
	}).Error
}

type UpdateInvoiceOptions struct {
	OrderID       string
	SponsorOpenID string
	Paid          bool
}

func (db *invoices) Update(ctx context.Context, uid string, opts UpdateInvoiceOptions) error {
	return db.WithContext(ctx).Model(&Invoice{}).Where("uid = ?", uid).
		Update("paid", opts.Paid).
		Updates(&Invoice{
			OrderID:       opts.OrderID,
			SponsorOpenID: opts.SponsorOpenID,
		}).Error
}

type GetInvoiceOptions struct {
	UID           string
	OrderID       string
	SponsorName   string
	SponsorOpenID string
}

func (db *invoices) Get(ctx context.Context, opts GetInvoiceOptions) ([]*Invoice, error) {
	query := db.WithContext(ctx).Model(&Invoice{})

	if opts.UID != "" {
		query = query.Where("uid = ?", opts.UID)
	}
	if opts.OrderID != "" {
		query = query.Where("order_id = ?", opts.OrderID)
	}
	if opts.SponsorName != "" {
		query = query.Where("sponsor_name = ?", opts.SponsorName)
	}
	if opts.SponsorOpenID != "" {
		query = query.Where("sponsor_open_id = ?", opts.SponsorOpenID)
	}

	var invoices []*Invoice
	return invoices, query.Order("id DESC").Find(&invoices).Error
}

type InvoiceSummary struct {
	SubtotalCents  int
	Count          int
	SponsorName    string
	SponsorWebSite string
}

func (db *invoices) List(ctx context.Context) ([]*InvoiceSummary, error) {
	panic("implement me")
}

var ErrInvoiceNotExists = errors.New("invoice dose not exist")

func (db *invoices) GetByID(ctx context.Context, id uint) (*Invoice, error) {
	var invoice Invoice
	err := db.WithContext(ctx).Model(&Invoice{}).Where("id = ?", id).First(&invoice).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrInvoiceNotExists
		}
		return nil, err
	}
	return &invoice, nil
}

func (db *invoices) GetByUID(ctx context.Context, uid string) (*Invoice, error) {
	var invoice Invoice
	err := db.WithContext(ctx).Model(&Invoice{}).Where("uid = ?", uid).First(&invoice).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrInvoiceNotExists
		}
		return nil, err
	}
	return &invoice, nil
}
