package serialize

import (
	"io"

	"github.com/goccy/go-json"
)

func NewJSONEncoder(w io.Writer) *json.Encoder {
	return json.NewEncoder(w)
}

func NewJSONDecoder(r io.Reader) *json.Decoder {
	return json.NewDecoder(r)
}

func JSONUnmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func JSONMarshal(v any) ([]byte, error) {
	return json.Marshal(v)
}
