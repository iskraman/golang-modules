package jsonlib

import (
	"encoding/json"
	"io"
)

// Memory -> Json
func Encoding(v interface{}) ([]byte, error) {
	b, err := json.Marshal(v)
	return b, err
}

func EncodingIndent(v interface{}) ([]byte, error) {
	b, err := json.MarshalIndent(v, "", "  ")
	return b, err
}

// Memory -> IO Stream
func EncodingStream(w io.Writer, v interface{}) error {
	enc := json.NewEncoder(w)
	err := enc.Encode(v)
	return err
}

func EncodingIndentStream(w io.Writer, v interface{}) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	err := enc.Encode(v)
	return err
}

// Json -> Memory
func Decoding(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	return err
}

func DecodingStream(r io.Reader, v interface{}) error {
	dec := json.NewDecoder(r)
	err := dec.Decode(v)
	return err
}
