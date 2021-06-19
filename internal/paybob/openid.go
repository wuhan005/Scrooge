// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed a MIT-style
// license that can be found in the LICENSE file.

package paybob

import (
	"fmt"
	"net/http"
	"net/url"
)

// GetOpenIDRedirectURL returns the URL which should be redirect to get the user open ID.
func (c *Client) GetOpenIDRedirectURL(redirectURL string) string {
	// The `mchid` must be the first parameter.
	return fmt.Sprintf("https://paybob.cn/api/openid?mchid=%s&callback_url=%s", c.mchid, url.QueryEscape(redirectURL))
}

// ParseOpenID gets the WeChat OpenID from a HTTP request.
// It returns a empty string when not find.
func ParseOpenID(req *http.Request) string {
	return req.URL.Query().Get("openid")
}
