// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package db

import (
	"testing"

	"gorm.io/gorm"

	"github.com/wuhan005/Scrooge/internal/dbutil"
)

func newTestDB(t *testing.T) (*gorm.DB, func(...string) error) {
	return dbutil.NewTestDB(t, allTables...)
}
