package utils

import (
	"bytes"
	"encoding/json"
)

func PrettyJSON(byteBuffer []byte) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, byteBuffer, "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}
