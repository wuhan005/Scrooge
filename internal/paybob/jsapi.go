// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package paybob

import (
	"net/http"
	"net/url"
	"strconv"

	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

type JSAPIPayOptions struct {
	UID       string
	TotalFee  int
	Title     string
	Attach    interface{}
	OpenID    string
	NotifyURL string
}

type JSAPIResponse struct {
	OrderID string
	Query   JSAPIQuery
}

type JSAPIQuery struct {
	AppID     string `json:"app_id"`
	TimeStamp string `json:"time_stamp"`
	NonceStr  string `json:"nonce_str"`
	Package   string `json:"package"`
	SignType  string `json:"sign_type"`
	PaySign   string `json:"pay_sign"`
}

// MakeJSAPIPay makes a new JSAPI pay.
// It returns the order UID.
func (c *Client) MakeJSAPIPay(opts JSAPIPayOptions) (*JSAPIResponse, error) {
	u, err := url.Parse("https://paybob.cn/api/jsapi")
	if err != nil {
		return nil, errors.Wrap(err, "parse url")
	}

	attachBody, err := jsoniter.Marshal(opts.Attach)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	q := u.Query()
	q.Set("mchid", c.mchid)
	q.Set("total_fee", strconv.Itoa(opts.TotalFee))
	q.Set("out_trade_no", opts.UID)
	q.Set("title", opts.Title)
	q.Set("attach", string(attachBody))
	q.Set("notify_url", opts.NotifyURL)
	q.Set("openid", opts.OpenID)
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

	var respBody struct {
		ReturnCode    int    `json:"return_code"`
		ReturnMessage string `json:"return_msg"`
		PayJSOrderID  string `json:"payjs_order_id"`
		JSAPI         struct {
			AppID     string `json:"appId"`
			TimeStamp string `json:"timeStamp"`
			NonceStr  string `json:"nonceStr"`
			Package   string `json:"package"`
			SignType  string `json:"signType"`
			PaySign   string `json:"paySign"`
		} `json:"jsapi"`
		Sign string `json:"sign"`
	}
	err = jsoniter.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return nil, errors.Wrap(err, "decode")
	}

	if respBody.ReturnCode != 1 {
		return nil, errors.Errorf("paybob: %q", respBody.ReturnMessage)
	}

	return &JSAPIResponse{
		OrderID: respBody.PayJSOrderID,
		Query: JSAPIQuery{
			AppID:     respBody.JSAPI.AppID,
			TimeStamp: respBody.JSAPI.TimeStamp,
			NonceStr:  respBody.JSAPI.NonceStr,
			Package:   respBody.JSAPI.Package,
			SignType:  respBody.JSAPI.SignType,
			PaySign:   respBody.JSAPI.PaySign,
		},
	}, nil
}
