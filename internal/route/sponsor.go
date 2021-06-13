// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by Apache-2.0
// license that can be found in the LICENSE file.

package route

type Sponsor struct{}

// NewSponsorHandler creates a new Sponsor.
func NewSponsorHandler() *Sponsor {
	return &Sponsor{}
}

// List returns the sponsor list.
func (*Sponsor) List() {

}
