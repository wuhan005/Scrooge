// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package route

import (
	"fmt"
	"net/url"

	log "unknwon.dev/clog/v2"

	"github.com/wuhan005/Scrooge/internal/context"
	"github.com/wuhan005/Scrooge/internal/db"
	"github.com/wuhan005/Scrooge/internal/form"
	"github.com/wuhan005/Scrooge/internal/paybob"
)

type Pay struct{}

// NewPayHandler creates a new Pay.
func NewPayHandler() *Pay {
	return &Pay{}
}

func (*Pay) NewInvoice(ctx context.Context, client *paybob.Client, f form.NewPayment) error {
	u, err := url.Parse(ctx.Host)
	if err != nil {
		log.Error("Failed to parse host: %v", err)
		return ctx.Error(50000, "internal server error")
	}

	// To frontend `/cashier`.
	u.Path = "/cashier"

	uid, err := db.Invoices.Create(ctx.Request().Context(), db.CreateInvoiceOptions{
		PriceCents:  f.PriceCents,
		SponsorName: f.SponsorName,
		Comment:     f.Comment,
	})
	if err != nil {
		log.Error("Failed to create new invoice: %v", err)
		return ctx.ServerError()
	}

	q := u.Query()
	q.Set("uid", uid)
	u.RawQuery = q.Encode()

	// Redirect to get user OpenID.
	redirectURL := client.GetOpenIDRedirectURL(u.String())
	return ctx.Success(map[string]interface{}{
		"uid":          uid,
		"redirect_url": redirectURL,
	})
}

func (*Pay) Query(ctx context.Context) error {
	uid := ctx.Query("uid")
	invoice, err := db.Invoices.GetByUID(ctx.Request().Context(), uid)
	if err != nil {
		if err == db.ErrInvoiceNotExists {
			return ctx.Error(40400, "账单不存在")
		}
		log.Error("Failed to get invoice: %v", err)
		return ctx.ServerError()
	}

	return ctx.Success(map[string]interface{}{
		"Paid": invoice.Paid,
	})
}

func (*Pay) Cashier(ctx context.Context, client *paybob.Client) error {
	openID := ctx.Query("openID")
	uid := ctx.Query("uid")
	invoice, err := db.Invoices.GetByUID(ctx.Request().Context(), uid)
	if err != nil {
		if err == db.ErrInvoiceNotExists {
			return ctx.Error(40400, "账单不存在")
		}
		log.Error("Failed to get invoice: %v", err)
		return ctx.ServerError()
	}

	if invoice.Paid {
		return ctx.Error(40300, "订单已支付")
	}

	notifyURL := ctx.Host + "api/pay/callback"

	resp, err := client.MakeJSAPIPay(paybob.JSAPIPayOptions{
		UID:       uid,
		TotalFee:  invoice.PriceCents,
		Title:     fmt.Sprintf("赞助大茄子 %.02f 元", float64(invoice.PriceCents)/100),
		Attach:    nil,
		OpenID:    openID,
		NotifyURL: notifyURL,
	})
	if err != nil {
		log.Error("Failed to make JSAPI pay: %v", err)
		return ctx.ServerError()
	}

	if invoice.SponsorOpenID == "" {
		// Update invoice orderID and sponsor WeChat OpenID.
		err = db.Invoices.Update(ctx.Request().Context(), invoice.UID, db.UpdateInvoiceOptions{
			OrderID:       resp.OrderID,
			SponsorOpenID: openID,
		})
		if err != nil {
			log.Error("Failed to update invoice: %v", err)
			return ctx.ServerError()
		}
	}

	return ctx.Success(map[string]interface{}{
		"uid":         invoice.UID,
		"price_cents": invoice.PriceCents,
		"order_id":    resp.OrderID,
		"query":       resp.Query,
	})
}

func (*Pay) Callback(ctx context.Context, client *paybob.Client) error {
	err := ctx.Request().ParseForm()
	if err != nil {
		log.Error("Failed to parse form: %v", err)
		return ctx.Error(40000, "请求体错误")
	}

	f := ctx.Request().Form
	sign := f.Get("sign")
	f.Del("sign")
	if client.Sign(f) != sign {
		return ctx.Error(40300, "签名错误")
	}

	uid := f.Get("out_trade_no")
	invoice, err := db.Invoices.GetByUID(ctx.Request().Context(), uid)
	if err != nil {
		if err == db.ErrInvoiceNotExists {
			return ctx.Error(40400, "账单不存在")
		}
		log.Error("Failed to get invoice: %v", err)
		return ctx.ServerError()
	}
	if invoice.Paid {
		return ctx.Success("")
	}

	resp, err := client.Check(invoice.OrderID)
	if err != nil {
		log.Error("Failed to check invoice status: %v", err)
		return ctx.ServerError()
	}
	if resp.Status == 0 {
		return ctx.ServerError()
	}

	err = db.Invoices.Update(ctx.Request().Context(), uid, db.UpdateInvoiceOptions{
		Paid: true,
	})
	if err != nil {
		log.Error("Failed to update invoice: %v", err)
		return ctx.ServerError()
	}

	return ctx.Success("")
}
