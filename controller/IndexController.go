package controller

import (
	"github.com/valyala/fasthttp"
	"meili-api/api"
	"meili-api/common"
)

type IndexController struct{}

func (that *IndexController) GetAllIndexes(ctx *fasthttp.RequestCtx) {
	indexes, err := api.GetAllIndexes()
	if err != nil {
		common.Error(ctx, false, err, common.ERROR)
		return
	}
	common.Success(ctx, true, indexes)
}

func (that *IndexController) GetIndex(ctx *fasthttp.RequestCtx) {
	value := ctx.UserValue("index")
	if value == nil {
		common.Error(ctx, false, "非法参数", common.ERROR)
		return
	}
	indexes, err := api.GetIndex(value.(string))
	if err != nil {
		common.Error(ctx, false, err, common.ERROR)
		return
	}
	common.Success(ctx, true, indexes)
}

func (that *IndexController) CreateIndex(ctx *fasthttp.RequestCtx) {
	uid := string(ctx.PostArgs().Peek("uid"))
	primaryKey := string(ctx.PostArgs().Peek("primaryKey"))
	if len(uid) == 0 {
		common.Error(ctx, false, "非法参数", common.ERROR)
		return
	}
	task, err := api.CreateIndex(uid, primaryKey)
	if err != nil {
		common.Error(ctx, false, err, common.ERROR)
		return
	}
	common.Success(ctx, true, task.UID)
}

func (that *IndexController) UpdateIndex(ctx *fasthttp.RequestCtx) {
	uid := ctx.UserValue("index")
	primaryKey := string(ctx.PostArgs().Peek("primaryKey"))
	if uid == nil || len(primaryKey) == 0 {
		common.Error(ctx, false, "非法参数", common.ERROR)
		return
	}
	task, err := api.UpdateIndex(uid.(string), primaryKey)
	if err != nil {
		common.Error(ctx, false, err, common.ERROR)
		return
	}
	common.Success(ctx, true, task.UID)
}

func (that *IndexController) DeleteIndex(ctx *fasthttp.RequestCtx) {
	uid := ctx.UserValue("index")
	if uid == nil {
		common.Error(ctx, false, "非法参数", common.ERROR)
		return
	}
	task, err := api.DeleteIndex(uid.(string))
	if err != nil {
		common.Error(ctx, false, err, common.ERROR)
		return
	}
	common.Success(ctx, true, task.UID)
}
