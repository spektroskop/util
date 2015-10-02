package util

import (
	"encoding/json"
	"io/ioutil"
)

func ReadFileObject(path string, m func([]byte, interface{}) error, object interface{}) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	if err := m(data, object); err != nil {
		return err
	}

	return nil
}

func ReadFileJSON(path string, object interface{}) error {
	return ReadFileObject(path, json.Unmarshal, object)
}
