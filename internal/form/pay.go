// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package form

type NewPayment struct {
	PriceCents  int `binding:"Required;Range(1,5000000)"`
	SponsorName string
	Comment     string
}
