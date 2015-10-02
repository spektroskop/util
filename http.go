package util

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func Post(uri string, object interface{}, marshal func(interface{}) ([]byte, error), content string) (io.Reader, error) {
	data, err := marshal(object)
	if err != nil {
		return nil, err
	}

	res, err := http.Post(uri, content, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	return res.Body, nil
}

func PostJSON(uri string, object interface{}) (io.Reader, error) {
	// marshal := func(object interface{}) ([]byte, error) {
	// 	return json.MarshalIndent(object, "", "  ")
	// }

	return Post(uri, object, json.Marshal /*marshal,*/, "application/json")
}
