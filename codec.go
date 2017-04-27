package goconf

import (
	"bytes"
	"encoding/json"
	"github.com/BurntSushi/toml"
)

type DecodeFunc func([]byte, interface{}) error

var DecodeFuncMap = map[string]DecodeFunc{
	".toml": func(bytes []byte, data interface{}) error {
		_, err := toml.Decode(string(bytes), data)
		return err
	},
	".json": func(bytes []byte, data interface{}) error {
		return json.Unmarshal(bytes, data)
	},
}

type EncodeFunc func(opts interface{}) (string, error)

var EncodeFuncMap = map[string]EncodeFunc{
	".toml": func(opts interface{}) (string, error) {
		var by bytes.Buffer
		encoder := toml.NewEncoder(&by)
		err := encoder.Encode(opts)
		return by.String(), err
	},
	".json": func(opts interface{}) (string, error) {
		by, err := json.Marshal(opts)
		return string(by), err
	},
}
