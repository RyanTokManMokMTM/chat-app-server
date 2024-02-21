package serialization

import "encoding/json"

func EncodeJSON(data any) ([]byte, error) {
	return json.Marshal(data)
}

func DecodeJSON(data []byte, to any) error {
	return json.Unmarshal(data, &to)
}
