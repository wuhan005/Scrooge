// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by Apache-2.0
// license that can be found in the LICENSE file.

package static

import (
	"net/http"

	"github.com/flamego/flamego"
)

func NotFound(fs http.FileSystem, file string) flamego.Handler {
	return flamego.Handler(func(c flamego.Context) {
		if c.Request().Method != http.MethodGet && c.Request().Method != http.MethodHead {
			return
		}
		
		f, err := fs.Open(file)
		if err != nil {
			return
		}
		defer func() { _ = f.Close() }()

		fi, err := f.Stat()
		if err != nil {
			return // File exists but failed to open.
		}

		http.ServeContent(c.ResponseWriter(), c.Request().Request, file, fi.ModTime(), f)
	})
}
