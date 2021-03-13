package models

// Response : type response struct
type ResponseModel struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Response(code int, message string) ResponseModel {
	return ResponseModel{
		Code:    code,
		Message: message,
	}
}
