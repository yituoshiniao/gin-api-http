package util

import (
	"net/url"
)

func ValuesDelKey(values url.Values, key ...string) {
	for _, k := range key {
		if _, ok := values[k]; ok {
			delete(values, k)
		}
	}

	//return true
}
