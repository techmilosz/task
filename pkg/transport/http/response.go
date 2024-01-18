package api

import "encoding/json"

type ErrorResponse struct {
	Message string `json:"message,omitempty"`
}

func mustMarshalError(message string) []byte {
	v := ErrorResponse{Message: message}
	payload, err := json.Marshal(v)
	if err != nil {
		return []byte("{}")
	}

	return payload
}
