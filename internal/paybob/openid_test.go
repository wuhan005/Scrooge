// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by Apache-2.0
// license that can be found in the LICENSE file.

package paybob

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetOpenIDRedirectURL(t *testing.T) {
	client := NewClient("123456789", "")
	got := client.GetOpenIDRedirectURL("https://github.red/openid")
	want := "https://paybob.cn/api/openid?callback_url=https%3A%2F%2Fgithub.red%2Fopenid&mchid=123456789"
	assert.Equal(t, want, got)
}

func Test_ParseOpenID(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "https://www.example.com/abc?uid=32&openid=abcdefghijk", nil)
	assert.Nil(t, err)

	got := ParseOpenID(req)
	want := "abcdefghijk"
	assert.Equal(t, want, got)

	// OpenID not  found
	req, err = http.NewRequest(http.MethodGet, "https://www.example.com/abc?uid=47", nil)
	assert.Nil(t, err)

	got = ParseOpenID(req)
	assert.Zero(t, got)
}
