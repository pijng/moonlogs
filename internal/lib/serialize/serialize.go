package serialize

import (
	"io"

	"encoding/json"

	gojson "github.com/goccy/go-json"
)

func NewJSONEncoder(w io.Writer) *gojson.Encoder {
	return gojson.NewEncoder(w)
}

func NewJSONDecoder(r io.Reader) *json.Decoder {
	return json.NewDecoder(r)
}

func JSONUnmarshal(data []byte, v any) error {
	return gojson.Unmarshal(data, v)
}

func JSONMarshal(v any) ([]byte, error) {
	return gojson.Marshal(v)
}
