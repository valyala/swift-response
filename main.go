package main

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"math/rand"
)

func main() {
	if err := fasthttp.ListenAndServe(":8090", requestHandler); err != nil {
		log.Fatalf("error in server: %s", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/json":
		jsonHandler(ctx)
	default:
		ctx.Error("page not found", fasthttp.StatusNotFound)
	}
}

func jsonHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json; charset=utf-8")

	a := make(map[string]int, 10)
	for i := 0; i < 10; i++ {
		k := fmt.Sprintf("Test Number %d", i+1)
		a[k] = rand.Intn(1000)
	}
	if err := json.NewEncoder(ctx).Encode(a); err != nil {
		log.Printf("error when encoding json: %s", err)
	}
}
