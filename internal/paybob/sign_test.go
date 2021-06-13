// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package paybob

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_Sign(t *testing.T) {
	client := NewClient("12345", "579a57f933397e0f441ba37f239d3721")

	query := url.Values{}
	query.Set("mchid", "12345")
	query.Set("total_fee", "1")
	query.Set("out_trade_no", "123123123123")

	got := client.Sign(query)
	want := "E3114A647C5A4CDD2CFC35019556DCD3"
	assert.Equal(t, want, got)

	// Empty key.
	client = NewClient("12345", "")
	got = client.Sign(query)
	want = "077AA74308BAC9FD5474B97C01D6FEA0"
	assert.Equal(t, want, got)
}
