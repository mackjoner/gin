// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"compress/gzip"
	"encoding/json"
	"net/http"
)

type jsonBinding struct{}

func (jsonBinding) Name() string {
	return "json"
}

func (jsonBinding) Bind(req *http.Request, obj interface{}) error {
	var decoder *json.Decoder
	switch req.Header.Get("Content-Encoding") {
	case "gzip":
		gz, err := gzip.NewReader(req.Body)
		if err != nil {
			return err
		}
		defer gz.Close()
		decoder = json.NewDecoder(gz)
	default:
		decoder = json.NewDecoder(req.Body)
	}
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return validate(obj)
}
