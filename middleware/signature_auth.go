package middleware

import "github.com/kataras/iris/v12"

func SignatureAuth(Ctx iris.Context)  {
	//do some check logic
	Ctx.Next()
}