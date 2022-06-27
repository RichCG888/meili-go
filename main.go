package main

import (
	"github.com/valyala/fasthttp"
	"log"
	_ "meili-api/helper"
	"meili-api/router"
)

func main() {
	log.Fatal(fasthttp.ListenAndServe(":8080", router.SetupRouter().Handler))
}
