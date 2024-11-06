package utils

import "encoding/json"

/**
*/
// UnmarshalJSON is a reusable function that unmarshals JSON into any given struct.
func StringToStruct(data string, v interface{}) error {
	return json.Unmarshal([]byte(data), v)
}
