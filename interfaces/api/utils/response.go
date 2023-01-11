package utils

type response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(data interface{}, message string) response {
	return response{message, data}
}
