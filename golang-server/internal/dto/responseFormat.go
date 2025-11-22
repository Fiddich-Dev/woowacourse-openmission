package dto

import "net/http"

type ResponseFormat struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Content interface{} `json:"content,omitempty"`
}

func Success(content interface{}) ResponseFormat {
	return ResponseFormat{
		http.StatusOK,
		"OK",
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
