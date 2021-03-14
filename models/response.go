package models

// Response : type response struct
type ResponseModel struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(Success bool, message string, data interface{}) ResponseModel {
	return ResponseModel{
		Success: Success,
		Message: message,
		Data:    data,
	}
}
