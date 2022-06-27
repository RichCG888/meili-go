package common

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

type Response struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

func Success(ctx *fasthttp.RequestCtx, state bool, data interface{}) {
	ctx.SetStatusCode(SUCCESS)
	ctx.SetContentType(ContentType)
	r(ctx, data, state)
}

func r(ctx *fasthttp.RequestCtx, data interface{}, state bool) {
	res := Response{
		Status: state,
		Data:   data,
	}
	bytes, err := json.Marshal(res)
	if err != nil {
		ctx.SetBody([]byte(err.Error()))
		return
	}
	ctx.SetBody(bytes)
}

func Error(ctx *fasthttp.RequestCtx, state bool, data interface{}, code int) {
	ctx.SetStatusCode(code)
	ctx.SetContentType(ContentType)
	r(ctx, data, state)
}
