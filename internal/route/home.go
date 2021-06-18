// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package route

import (
	"github.com/wuhan005/Scrooge/internal/context"
)

type Home struct{}

// NewHomeHandler creates a new Home.
func NewHomeHandler() *Home {
	return &Home{}
}

func (*Home) Profile(ctx context.Context) error {
	// TODO read from database.
	return ctx.Success(map[string]interface{}{
		"avatar_url":  "https://avatars.githubusercontent.com/u/12731778",
		"name":        "E99p1ant",
		"slogan":      "Be cool, but also be warm.",
		"description": "说点什么？",
	})
}

// Tiers returns the tiers.
func (*Home) Tiers(ctx context.Context) error {
	type tier struct {
		Amount  int    `json:"amount"`
		Comment string `json:"comment"`
	}

	// TODO read from database.
	tiers := []tier{
		{Amount: 5, Comment: "谢谢老板"},
		{Amount: 10, Comment: "谢谢老板"},
		{Amount: 50, Comment: "谢谢老板"},
		{Amount: 100, Comment: "谢谢老板"},
		{Amount: 0, Comment: "谢谢老板"},
	}
	return ctx.Success(tiers)
}

// List returns the sponsor list.
func (*Home) List(ctx context.Context) error {
	panic("implement me")
}
