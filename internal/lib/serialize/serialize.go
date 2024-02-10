package serialize

import (
	"io"

	stdjson "encoding/json"

	"github.com/goccy/go-json"
)

func NewJSONEncoder(w io.Writer) *json.Encoder {
	return json.NewEncoder(w)
}

func NewJSONDecoder(r io.Reader) *stdjson.Decoder {
	return stdjson.NewDecoder(r)
}

func JSONUnmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func JSONMarshal(v any) ([]byte, error) {
	return json.Marshal(v)
}
