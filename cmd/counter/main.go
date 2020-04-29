package main

import (
	"github.com/cooli88/counter/internal/counter"
	"github.com/valyala/fasthttp"
)

const address = ":8080"

func main() {
	if err := fasthttp.ListenAndServe(address, counter.NewCounterHandler().HandleCommon); err != nil {
		panic(err)
	}
}
