package utils

import "github.com/satishcg12/echomers/internal/types"

func NewResponse(status int, message string) types.Response {
	return types.Response{
		Status:  status,
		Message: message,
	}
}

func NewResponseWithData(status int, message string, data interface{}) types.Response {
	return types.Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
