// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package paybob

import (
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/wuhan005/gadget"
)

func (c *Client) Sign(query url.Values) string {
	if query == nil {
		return ""
	}
	// Remove the empty value.
	for k, v := range query {
		if len(v) == 0 || v[0] == "" {
			query.Del(k)
		}
	}

	var buf strings.Builder
	keys := make([]string, 0, len(query))
	for k := range query {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := query[k]
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(v)
		}
	}

	return strings.ToUpper(gadget.Md5(fmt.Sprintf("%s&key=%s", buf.String(), c.key)))
}
