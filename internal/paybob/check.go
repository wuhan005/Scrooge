// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package paybob

import (
	"net/http"
	"net/url"

	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

type CheckAPIResponse struct {
	ReturnCode    int                 `json:"return_code"`
	MchID         string              `json:"mch_id"`
	OutTradeNo    string              `json:"out_trade_no"`
	PayJSOrderID  string              `json:"payjs_order_id"`
	TransactionID string              `json:"transaction_id"`
	Status        int                 `json:"status"`
	OpenID        string              `json:"openid"`
	TotalFee      int                 `json:"total_fee"`
	PaidTime      string              `json:"paid_time"`
	Attach        jsoniter.RawMessage `json:"attach"`
	Sign          string              `json:"sign"`
}

// Check checks the paybob invoice payment status.
func (c *Client) Check(payJSOrderID string) (*CheckAPIResponse, error) {
	u, err := url.Parse("https://paybob.cn/api/check")
	if err != nil {
		return nil, errors.Wrap(err, "parse url")
	}

	q := u.Query()
	q.Set("payjs_order_id", payJSOrderID)
	sign := c.Sign(q)
	q.Set("sign", sign)

	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, errors.Wrap(err, "http GET")
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var respBody CheckAPIResponse
	err = jsoniter.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return nil, errors.Wrap(err, "decode")
	}

	return &respBody, nil
}
