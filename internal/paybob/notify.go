// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package paybob

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

type NotifyResponse struct {
	ReturnCode    int    `json:"return_code"`
	TotalFee      int    `json:"total_fee"`
	OutTradeNo    string `json:"out_trade_no"`
	PayJSOrderID  string `json:"payjs_order_id"`
	TransactionID string `json:"transaction_id"`
	TimeEnd       string `json:"time_end"`
	OpenID        string `json:"open_id"`
	Attach        string `json:"attach"`
	MchID         string `json:"mchid"`
	Sign          string `json:"sign"`
}

// ParseNotify parse the notify request body.
func (c *Client) ParseNotify(req *http.Request) (*NotifyResponse, error) {
	var requestBody NotifyResponse
	return &requestBody, jsoniter.NewDecoder(req.Body).Decode(&requestBody)
}
