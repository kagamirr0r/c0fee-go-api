package controller

import "time"

// 共通エラーレスポンス生成関数
func GenerateErrorResponse(code, message string, fieldErrors []FieldError) Response {
	if len(fieldErrors) == 0 {
		fieldErrors = []FieldError{}
	}
	return Response{
		Code:      code,
		Message:   message,
		Errors:    fieldErrors,
		Content:   nil,
		Timestamp: time.Now().In(jst).Format(time.RFC3339),
	}
}
