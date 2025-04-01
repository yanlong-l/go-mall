package app

import (
	"github.com/gin-gonic/gin"
	"github.com/yanlong-l/go-mall/common/errcode"
	"github.com/yanlong-l/go-mall/common/logger"
)

type response struct {
	ctx        *gin.Context
	Code       int         `json:"code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data,omitempty"`
	RequestId  string      `json:"request_id,omitempty"`
	Pagination *pagination `json:"pagination,omitempty"`
}

func NewResponse(ctx *gin.Context) *response {
	return &response{ctx: ctx}
}

// SetPagination 设置Response的分页信息
func (r *response) SetPagination(pagination *pagination) *response {
	r.Pagination = pagination
	return r
}

func (r *response) Success(data any) {
	r.Code = errcode.Success.Code
	r.Msg = errcode.Success.Msg
	r.Data = data
	requestId := ""
	if _, exists := r.ctx.Get("traceid"); exists {
		val, _ := r.ctx.Get("traceid")
		requestId = val.(string)
	}
	r.RequestId = requestId
	r.ctx.JSON(errcode.Success.HttpStatusCode(), r)
}

func (r *response) SuccessOk() {
	r.Success("")
}

func (r *response) Error(err *errcode.AppError) {
	r.Code = err.Code
	r.Msg = err.Msg
	requestId := ""
	if _, exists := r.ctx.Get("traceid"); exists {
		val, _ := r.ctx.Get("traceid")
		requestId = val.(string)
	}
	r.RequestId = requestId
	logger.Error(r.ctx, "api_response_error", "err", err)
	r.ctx.JSON(err.HttpStatusCode(), r)
}
