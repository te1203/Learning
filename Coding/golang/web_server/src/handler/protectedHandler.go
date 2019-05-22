package handler

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

// Protected is the Protected handler
func Protected(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Protected!\n")
}
