package handler

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

// Protected is the Protected handler
func FindAllApplication(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Protected!\n")
}

func CreateApplication(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Protected!\n")
}

func UpdateApplication(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Protected!\n")
}

func DeleteApplication(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Protected!\n")
}
