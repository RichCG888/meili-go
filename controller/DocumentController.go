package controller

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"meili-api/api"
	"meili-api/common"
)

type DocumentController struct{}

func (that *DocumentController) GetAllDocument(ctx *fasthttp.RequestCtx) {
	value := ctx.UserValue("index")
	if value == nil {
		common.Error(ctx, false, "非法参数", common.ERROR)
		return
	}
	documents, err := api.GetAllDocument(value.(string))
	if err != nil {
		common.Error(ctx, false, err, common.ERROR)
		return
	}
	common.Success(ctx, true, documents)
}

func (that *DocumentController) GetDocument(ctx *fasthttp.RequestCtx) {
	uid := ctx.UserValue("index")
	id := ctx.UserValue("id")
	if uid == nil || id == nil {
		common.Error(ctx, false, "非法参数", common.ERROR)
		return
	}
	indexes, err := api.GetDocument(uid.(string), id.(string))
	if err != nil {
		common.Error(ctx, false, err, common.ERROR)
		return
	}
	common.Success(ctx, true, indexes)
}

func (that *DocumentController) CreateOrUpdateDocument(ctx *fasthttp.RequestCtx) {
	uid := ctx.UserValue("index")
	primaryKey := string(ctx.PostArgs().Peek("primaryKey"))
	if uid == nil {
		common.Error(ctx, false, "非法参数", common.ERROR)
		return
	}
	dMap := make([]map[string]interface{}, 0)

	err := json.Unmarshal(ctx.PostBody(), &dMap)
	if err != nil {
		common.Error(ctx, false, "JSON格式错误", common.ERROR)
		return
	}
	task, err := api.CreateOrUpdateDocument(uid.(string), primaryKey, dMap)
	if err != nil {
		common.Error(ctx, false, err, common.ERROR)
		return
	}
	common.Success(ctx, true, task.UID)
}

func (that *DocumentController) DeleteDocument(ctx *fasthttp.RequestCtx) {
	uid := ctx.UserValue("index")
	if uid == nil {
		common.Error(ctx, false, "非法参数", common.ERROR)
		return
	}
	ids := make(map[string][]string, 0)
	err := json.Unmarshal(ctx.PostBody(), &ids)
	if err != nil {
		common.Error(ctx, false, "JSON格式错误", common.ERROR)
		return
	}
	if len(ids["ids"]) == 0 {
		common.Error(ctx, false, "ID不能为空", common.ERROR)
		return
	}
	task, err := api.DeleteDocument(uid.(string), ids["ids"])
	if err != nil {
		common.Error(ctx, false, err, common.ERROR)
		return
	}
	common.Success(ctx, true, task.UID)
}

func (that *DocumentController) DeleteAllDocument(ctx *fasthttp.RequestCtx) {
	uid := ctx.UserValue("index")
	if uid == nil {
		common.Error(ctx, false, "非法参数", common.ERROR)
		return
	}
	task, err := api.DeleteAllDocument(uid.(string))
	if err != nil {
		common.Error(ctx, false, err, common.ERROR)
		return
	}
	common.Success(ctx, true, task.UID)
}
