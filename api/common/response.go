package common

type Error struct {
	Code        ErrorCode `json:"code"`
	Description string    `json:"description"`
}

type Response struct {
	Message any    `json:"message"`
	Error   *Error `json:"error"`
}

func ResponseMessage(message any) Response {
	return Response{
		Message: message,
	}
}
