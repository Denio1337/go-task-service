package response

// Constructs a successful response with data
func SuccessResponse(data any) *Response {
	return &Response{
		Success: true,
		Error:   "",
		Data:    data,
	}
}

// Constructs an error response with a message
func ErrorResponse(message string) *Response {
	return &Response{
		Success: false,
		Error:   message,
		Data:    nil,
	}
}
