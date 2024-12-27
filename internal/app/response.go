package app

import (
	"github.com/gin-gonic/gin"
	errcode2 "github.com/ljinf/template_project/pkg/errcode"
	"github.com/ljinf/template_project/pkg/logger"
)

type response struct {
	ctx        *gin.Context
	Code       int         `json:"code"`
	Msg        string      `json:"msg"`
	RequestId  string      `json:"request_id"`
	Data       interface{} `json:"data,omitempty"`
	Pagination *pagination `json:"pagination,omitempty"`
}

func Success(ctx *gin.Context, data interface{}) {
	resp := &response{ctx: ctx}
	resp.success(data)
}

func SuccessOk(ctx *gin.Context) {
	resp := &response{ctx: ctx}
	resp.success("")
}

func SuccessWithPagination(ctx *gin.Context, data interface{}, pagination *pagination) {
	resp := &response{ctx: ctx}
	resp.Pagination = pagination
	resp.success(data)
}

func Error(ctx *gin.Context, err *errcode2.AppError) {
	resp := &response{ctx: ctx}
	resp.error(err)
}

// SetPagination 设置Response的分页信息
func (r *response) setPagination(pagination *pagination) *response {
	r.Pagination = pagination
	return r
}

func (r *response) success(data interface{}) {
	r.Code = errcode2.Success.Code()
	r.Msg = errcode2.Success.Msg()
	requestId := ""
	if _, exists := r.ctx.Get("traceid"); exists {
		val, _ := r.ctx.Get("traceid")
		requestId = val.(string)
	}
	r.RequestId = requestId
	r.Data = data

	r.ctx.JSON(errcode2.Success.HttpStatusCode(), r)
}

func (r *response) successOk() {
	r.success("")
}

func (r *response) error(err *errcode2.AppError) {
	r.Code = err.Code()
	r.Msg = err.Msg()
	if _, exists := r.ctx.Get("traceid"); exists {
		val, _ := r.ctx.Get("traceid")
		r.RequestId = val.(string)
	}
	// 兜底记一条响应错误, 项目自定义的AppError中有错误链条, 方便出错后排查问题
	logger.Error(r.ctx, "api_response_error", "err", err)
	r.ctx.JSON(err.HttpStatusCode(), r)
}
