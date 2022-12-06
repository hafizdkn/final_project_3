package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IResponse interface {
	WriteJsonRespnse(ctx *gin.Context, resp *Response)
	AbortJsonRespnse(ctx *gin.Context, resp *Response)
	SuccessCreateResponse(payload interface{}, message string) *Response
	SuccessLoginResponse(payload interface{}, token, message string) *Response
	SuccessResponse(payload interface{}, message string) *Response
	InternalServerError(err error) *Response
	BadRequestResponse(err error) *Response
	UnprocessAbleEntityResponse(err interface{}) *Response
	UnauthorizedResponse(err error, message string) *Response
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Token   string      `json:"token,omitempty"`
}

type response struct{}

var NewResponse IResponse = &response{}

func (r *response) WriteJsonRespnse(ctx *gin.Context, resp *Response) {
	ctx.JSON(resp.Status, resp)
}

func (r *response) AbortJsonRespnse(ctx *gin.Context, resp *Response) {
	ctx.AbortWithStatusJSON(resp.Status, resp)
}

func (r *response) SuccessCreateResponse(payload interface{}, message string) *Response {
	return &Response{
		Status:  http.StatusCreated,
		Message: message,
		Payload: payload,
	}
}

func (r *response) SuccessLoginResponse(payload interface{}, token, message string) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: message,
		Payload: payload,
		Token:   token,
	}
}

func (r *response) SuccessResponse(payload interface{}, message string) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: message,
		Payload: payload,
	}
}

func (r *response) InternalServerError(err error) *Response {
	return &Response{
		Status:  http.StatusInternalServerError,
		Message: "Internal server error",
		Error:   err.Error(),
	}
}

func (r *response) BadRequestResponse(err error) *Response {
	return &Response{
		Status:  http.StatusBadRequest,
		Message: "Bad request",
		Error:   err.Error(),
	}
}

func (r *response) UnprocessAbleEntityResponse(err interface{}) *Response {
	return &Response{
		Status:  http.StatusUnprocessableEntity,
		Message: "Unprocess able entity",
		Error:   err,
	}
}

func (r *response) UnauthorizedResponse(err error, message string) *Response {
	return &Response{
		Status:  http.StatusUnauthorized,
		Message: message,
		Error:   err.Error(),
	}
}
