// Copyright 2018 Gin Core Team.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

type uriBinding struct{}

func (uriBinding) Name() string {
	return "uri"
}

func (uriBinding) Bind(m map[string][]string, obj interface{}) error {
	if err := mapURI(obj, m); err != nil {
		return err
	}
	return validate(obj)
}

func (uriBinding) BindOnly(m map[string][]string, obj interface{}) error {
	return mapURI(obj, m)
}
