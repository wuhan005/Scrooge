// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package db

import (
	"os"
	"time"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/wuhan005/Scrooge/internal/dbutil"
)

var allTables = []interface{}{
	Invoice{},
}

// Init initializes the database.
func Init() error {
	dsn := os.ExpandEnv("postgres://$PGUSER:$PGPASSWORD@$PGHOST:$PGPORT/$PGDATABASE?sslmode=$PGSSLMODE")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return dbutil.Now()
		},
	})
	if err != nil {
		return errors.Wrap(err, "open connection")
	}

	// Migrate databases.
	if db.AutoMigrate(allTables...) != nil {
		return errors.Wrap(err, "auto migrate")
	}

	Invoices = NewInvoiceStore(db)

	return nil
}
