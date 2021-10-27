package przelewy24api

import (
	"bytes"
	"encoding/json"
)

func HTMLUnescapedJSON(object interface{}) []byte {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(object)
	if err != nil {
		return nil
	}
	return buffer.Bytes()
}
