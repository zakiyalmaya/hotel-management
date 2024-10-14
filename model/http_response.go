package model

type HttpResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewHttpResponse(code int, message string, data interface{}) HttpResponse {
	return HttpResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}