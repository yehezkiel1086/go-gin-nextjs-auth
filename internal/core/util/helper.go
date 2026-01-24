package util

import "encoding/json"

func Deserialize(src []byte, dst any) error {
	return json.Unmarshal(src, dst)
}

func Serialize(data any) ([]byte, error) {
	return json.Marshal(data)
}
