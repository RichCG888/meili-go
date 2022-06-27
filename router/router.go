package router

import (
	"github.com/fasthttp/router"
	"meili-api/controller"
)

func SetupRouter() *router.Router {

	route := router.New()

	index := new(controller.IndexController)

	route.GET("/indexes", index.GetAllIndexes) //查询所有索引

	route.GET("/index/{index}", index.GetIndex) //查询指定索引

	route.POST("/index", index.CreateIndex) //创建索引

	route.PUT("/index/{index}", index.UpdateIndex) //更新索引

	route.DELETE("/index/{index}", index.DeleteIndex) //删除索引

	document := new(controller.DocumentController)

	route.GET("/document/{index}/documents", document.GetAllDocument) //查询所有文档

	route.GET("/document/{index}/documents/{id}", document.GetDocument) //查询指定文档

	route.POST("/document/{index}/documents", document.CreateOrUpdateDocument) //创建或更新文档

	route.DELETE("/document/{index}/document", document.DeleteDocument) //删除索引下的指定文档

	route.DELETE("/document/{index}/documents", document.DeleteAllDocument) //删除索引下的所有文档

	return route
}
