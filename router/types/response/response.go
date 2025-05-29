package response

type Response struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Data    any    `json:"data"`
}

func SuccessResponse(data any) *Response {
	return &Response{
		Success: true,
		Error:   "",
		Data:    data,
	}
}

func ErrorResponse(message string) *Response {
	return &Response{
		Success: false,
		Error:   message,
		Data:    nil,
	}
}
