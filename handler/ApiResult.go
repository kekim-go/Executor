package handler

import ctx "github.com/kekim-go/Executor/ctx"

// ApiResultHandler : ApiResult 핸들러
type ApiResultHandler struct {
	Ctx *ctx.Context
}

// NewApiResultHandler : ApiResultHandler 생성자
func NewApiResultHandler(ctx *ctx.Context) *ApiResultHandler {
	return &ApiResultHandler{
		Ctx: ctx,
	}
}
