package counter

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

var (
	strContentType     = []byte("Content-Type")
	strApplicationJSON = []byte("application/json")
)

func doJSONWrite(ctx *fasthttp.RequestCtx, code int, obj interface{}) {
	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)
	ctx.Response.SetStatusCode(code)
	if err := json.NewEncoder(ctx).Encode(obj); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}

type CounterHandler struct {
	userCounterService *UserCounter
}

func NewCounterHandler() *CounterHandler {
	return &CounterHandler{userCounterService: newUserCounter()}
}

func (h *CounterHandler) HandleCommon(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/":
		h.handleIndex(ctx)
	case "/count":
		h.handleCounter(ctx)
	default:
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}
}

func (h *CounterHandler) handleIndex(ctx *fasthttp.RequestCtx) {
	userId := string(ctx.QueryArgs().Peek("user_id"))
	ctx.Response.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody([]byte{})
	go h.userCounterService.incrUser(userId)
}

func (h *CounterHandler) handleCounter(ctx *fasthttp.RequestCtx) {
	doJSONWrite(ctx, fasthttp.StatusOK, CounterResponse{Count: h.userCounterService.getRobotCount()})
}
