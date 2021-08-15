package binding

import (
	"net/http"
)

type requestBinding struct{}

func (requestBinding) Name() string {
	return "request"
}

func (b requestBinding) Bind(obj interface{}, req *http.Request, form map[string][]string) error {
	if err := b.BindOnly(obj, req, form); err != nil {
		return err
	}

	return validate(obj)
}

func (b requestBinding) BindOnly(obj interface{}, req *http.Request, uriMap map[string][]string) error {

	if err := Uri.BindOnly(uriMap, obj); err != nil {
		return err
	}

	if err := b.bindingQuery(req, obj); err != nil {
		return err
	}

	binders := []Binding{Header, Cookie}
	for _, binder := range binders {
		if err := binder.BindOnly(req, obj); err != nil {
			return err
		}
	}

	// default json
	if req.Method == http.MethodPut || req.Method == http.MethodPost {
		contentType := req.Header.Get("Content-Type")
		if contentType == "" {
			contentType = MIMEJSON
		}
		bb := Default(req.Method, contentType)
		return bb.BindOnly(req, obj)
	}
	return nil
}

func (b requestBinding) bindingQuery(req *http.Request, obj interface{}) error {
	values := req.URL.Query()
	return mapFormByTag(obj, values, "query")
}
