package utils

type response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type respErrValidator struct {
	Message string            `json:"message"`
	ErrData map[string]string `json:"error_data"`
	Data    interface{}       `json:"data"`
}

func NewResponse(data interface{}, message string) response {
	return response{message, data}
}

func NewErrValidation(errData map[string]string, message string) respErrValidator {
	return respErrValidator{message, errData, nil}
}
