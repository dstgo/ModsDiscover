package resp

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Response struct {
	body struct {
		Code int    `json:"code,omitempty"`
		Data any    `json:"data,omitempty"`
		Msg  string `json:"msg,omitempty"`
	}

	status int
	err    error

	ctx *app.RequestContext
}

func (resp *Response) Code(code int) *Response {
	resp.body.Code = code
	return resp
}

func (resp *Response) Data(data any) *Response {
	resp.body.Data = data
	return resp
}

func (resp *Response) Msg(msg string) *Response {
	resp.body.Msg = msg
	return resp
}

func (resp *Response) Error(err error) *Response {
	resp.err = err
	return resp
}

func (resp *Response) Status(status int) *Response {
	resp.status = status
	return resp
}

func (resp *Response) Do() {
	ctx := resp.ctx
	if ctx == nil {
		return
	}

	if resp.body.Code == 0 {
		resp.body.Code = resp.status
	}

	if resp.body.Msg == "" && resp.err != nil {
		resp.body.Msg = resp.err.Error()
	}

	if resp.err != nil {
		resp.ctx.Error(resp.err)
	}

	ctx.JSON(resp.status, resp.body)
}

func New(ctx *app.RequestContext) *Response {
	return &Response{ctx: ctx}
}

func Ok(ctx *app.RequestContext) *Response {
	return &Response{ctx: ctx, status: consts.StatusOK}
}

func Failed(ctx *app.RequestContext) *Response {
	return &Response{ctx: ctx, status: consts.StatusBadRequest}
}
