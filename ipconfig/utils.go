package ipconfig

import "context"

type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

func GetIpInfoList(c context.Context, ctx *app.RequestContext) {
	defer func() {
		if err := recover(); err!=nil {
			ctx.Json(consts.StatusBadRequest, utils.H{"err":err})

		}
	}
	//step1: build context from clint request
	ipConfCtx := domain.BuildIpConfContext(&c, ctx);
	//step2: Ask for ip list to return
	eds := domain.Dispatch (ipConfCtx)
	//step3: set the result(top 5 ip recommend) in context
	ipConfCtx.AppCtx.Json(consts.StatusOK, packRes(top5Endports(eds)))

	//?? no return??
}