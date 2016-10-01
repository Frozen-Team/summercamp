package actions

import (
	"bytes"
	"encoding/json"
)

// Response represents the response from any action of our application.
// Actions tests unmarshal data into this struct
type Response struct {
	Meta struct {
		HasError bool     `json:"has-error"`
		Errors   []string `json:"errors"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func ReadResponse(body *bytes.Buffer) (*Response, error) {
	var response Response
	err := json.Unmarshal(body.Bytes(), &response)
	return &response, err
}
