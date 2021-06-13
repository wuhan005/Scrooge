// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package paybob

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/wuhan005/gadget"
)

func (c *Client) Sign(query url.Values) string {
	return strings.ToUpper(gadget.Md5(fmt.Sprintf("%s&key=%s", query.Encode(), c.key)))
}
