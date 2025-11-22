package dto

import "net/http"

const (
	messageOK = "OK"
)

type ResponseFormat struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Content interface{} `json:"content,omitempty"`
}

func Success(content interface{}) ResponseFormat {
	return ResponseFormat{
		http.StatusOK,
		messageOK,
		content,
	}
}

func Fail(status int, message string) ResponseFormat {
	return ResponseFormat{
		status,
		message,
		nil,
	}
}
