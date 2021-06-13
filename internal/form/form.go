// Copyright 2021 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package form

import (
	"net/http"

	"github.com/Cardinal-Platform/binding"
	"github.com/flamego/flamego"
	jsoniter "github.com/json-iterator/go"
	log "unknwon.dev/clog/v2"
)

func init() {
	binding.CustomErrorHandler = func(ctx flamego.Context, errors binding.Errors) {
		ctx.ResponseWriter().Header().Set("Content-Type", "application/json")
		ctx.ResponseWriter().WriteHeader(http.StatusBadRequest)

		log.Error("%+v", errors)
		var msg string
		if errors.Len() != 0 {
			msg = errors[0].Message
		}

		err := jsoniter.NewEncoder(ctx.ResponseWriter()).Encode(
			map[string]interface{}{
				"error": 40000,
				"msg":   msg, // FIXME error message should be more human-friendly.
			},
		)
		if err != nil {
			log.Error("Failed to encode: %v", err)
		}
	}
}
