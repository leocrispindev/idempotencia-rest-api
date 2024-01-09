package model

type ResponseError struct {
	Text string `json:"message"`
}

type ResponseSuccess struct {
	Text string      `json:"message"`
	Data interface{} `json:"data,omitempty"`
}
