// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package conf

var (
	// App is the application settings.
	App struct {
		Version string `ini:"-"` // Version should only be set by the main package.
	}

	// Profile is the user profile information.
	Profile struct {
		AvatarURL   string `ini:"avatar_url"`
		Name        string `ini:"name"`
		Slogan      string `ini:"slogan"`
		Description string `ini:"description"`
	}

	// Tier is the tier list.
	Tier struct {
		Items []struct {
			Amount  int    `ini:"amount"`
			Comment string `ini:"comment"`
		} `ini:"Tier.Items,,,nonunique"`
	}
)
