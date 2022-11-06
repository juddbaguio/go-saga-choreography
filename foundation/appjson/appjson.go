package appjson

import "encoding/json"

func EncodeJSONByte(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func DecodeJSONByte(data []byte, dest interface{}) error {
	return json.Unmarshal(data, dest)
}
