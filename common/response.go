package common

import "time"

var Jst, _ = time.LoadLocation("Asia/Tokyo")

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type Response struct {
	Code      string       `json:"code"`
	Message   string       `json:"message"`
	Errors    []FieldError `json:"errors"`
	Content   interface{}  `json:"content"`
	TimeStamp string       `json:"timestamp"`
}

func GenerateErrorResponse(code, message string, fieldErrors []FieldError) Response {
	if len(fieldErrors) == 0 {
		fieldErrors = []FieldError{}
	}
	return Response{
		Code:      code,
		Message:   message,
		Errors:    fieldErrors,
		Content:   nil,
		TimeStamp: time.Now().In(Jst).Format(time.RFC3339),
	}
}
