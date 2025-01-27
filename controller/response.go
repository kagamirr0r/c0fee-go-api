package controller

type Response struct {
	Code      string       `json:"code"`
	Message   string       `json:"message"`
	Errors    []FieldError `json:"errors"`
	Content   interface{}  `json:"content"`
	Timestamp string       `json:"timestamp"`
}
