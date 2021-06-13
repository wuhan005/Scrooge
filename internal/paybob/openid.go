// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed a MIT-style
// license that can be found in the LICENSE file.

package paybob

import (
	"net/http"
	"net/url"
)

// GetOpenIDRedirectURL returns the URL which should be redirect to get the user open ID.
func (c *Client) GetOpenIDRedirectURL(redirectURL string) string {
	u, err := url.Parse("https://paybob.cn/api/openid")
	if err != nil {
		return ""
	}

	q := u.Query()
	q.Set("mchid", c.mchid)
	q.Set("callback_url", redirectURL)
	u.RawQuery = q.Encode()
	return u.String()
}

// ParseOpenID gets the WeChat OpenID from a HTTP request.
// It returns a empty string when not find.
func ParseOpenID(req *http.Request) string {
	return req.URL.Query().Get("openid")
}
