// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package conf

import (
	"github.com/pkg/errors"
	"gopkg.in/ini.v1"
	log "unknwon.dev/clog/v2"
)

func init() {
	err := log.NewConsole()
	if err != nil {
		panic("init console logger: " + err.Error())
	}
}

// File is the configuration object.
var File *ini.File

func Init(customConf string) error {
	if customConf == "" {
		customConf = "./conf/scrooge.ini"
	}

	var err error
	File, err = ini.LoadSources(ini.LoadOptions{
		Insensitive:            true,
		AllowNonUniqueSections: true,
	}, customConf)
	if err != nil {
		return errors.Wrap(err, "load ini")
	}

	if err := File.Section("App").MapTo(&App); err != nil {
		return errors.Wrap(err, "mapping [App] section")
	}

	if err := File.Section("Profile").MapTo(&Profile); err != nil {
		return errors.Wrap(err, "mapping [Profile] section")
	}

	if err = File.Section("Tier").StrictMapTo(&Tier); err != nil {
		return errors.Wrap(err, "mapping [Tier] section")
	}

	return nil
}
