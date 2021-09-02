package utils

import "encoding/json"

func Marshal(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func Unmarshal(data []byte, dest interface{}) error {
	if err := json.Unmarshal(data, dest); err != nil {
		return err
	}

	return nil
}
