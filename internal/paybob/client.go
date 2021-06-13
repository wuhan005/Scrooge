// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package paybob

import (
	"os"
)

// Client is a paybob client.
type Client struct {
	mchid string
	key   string
}

// NewClient creates a new paybob client with the given arguments.
func NewClient(mchid, key string) *Client {
	return &Client{
		mchid: mchid,
		key:   key,
	}
}

// NewDefaultClient creates a new paybob client with the environment variables.
func NewDefaultClient() *Client {
	return &Client{
		mchid: os.Getenv("PAYBOB_MCHID"),
		key:   os.Getenv("PAYBOB_KEY"),
	}
}
