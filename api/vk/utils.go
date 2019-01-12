package vk

import (
	"net/url"

	"github.com/google/go-querystring/query"
)

// MergeURLValues merges mergeWith into base
//
// Can be useful when implementing API.Request
func MergeURLValues(base, mergeWith url.Values) {
	for k, v := range mergeWith {
		if old, ok := base[k]; ok {
			base[k] = append(old, v...)
		} else {
			base[k] = v
		}
	}
}

// BuildRequestParams is a helper function which can be used when implementing API.Request
//
// params can be:
// 	- nil
//  - url.Values
//  - url tagged struct (https://godoc.org/github.com/google/go-querystring/query)
func BuildRequestParams(params interface{}) (url.Values, error) {
	switch v := params.(type) {
	case nil:
		return make(url.Values), nil
	case url.Values:
		return v, nil
	default:
		return query.Values(params)
	}
}
