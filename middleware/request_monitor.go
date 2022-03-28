package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"io/ioutil"
	"iris-demo-new/slog"
	"net/http"
	"strings"
	"time"
)

func IrisRequestHandler(ctx iris.Context)  {
	path := ctx.Request().URL.Path
	method := ctx.Request().Method
	start := time.Now().UnixNano() / 1e6
	ip := ctx.Request().RemoteAddr

	ctx.Request().UserAgent()

	params := getParams(ctx, method)

	ctx.Next()

	end := time.Now().UnixNano() / 1e6
	times := end - start

	request := fmt.Sprintf("[[path]--> %s [method]--> %s [IP]--> %s [time]ms--> %d [params]-->%s]", path, method, ip, times, params)
	slog.Info(request)
}

func getParams(ctx iris.Context, method string) string {
	params := ""
	if method == http.MethodPost || method == http.MethodPut {
		body, err := ioutil.ReadAll(ctx.Request().Body)
		if err == nil {
			defer ctx.Request().Body.Close()
			buf := bytes.NewBuffer(body)
			ctx.Request().Body = ioutil.NopCloser(buf)
			params = string(body)
			if strings.Contains(params, "\r\n") {
				params = strings.ReplaceAll(params, "\r\n", "")
			}
			if strings.Contains(params, "\n") {
				params = strings.ReplaceAll(params, "\n", "")
			}
			params = strings.ReplaceAll(params, " ", "")
		}
	} else {
		data := ctx.FormValues()
		fData := map[string]interface{}{}
		for k, v := range data {
			if len(v) == 1 {
				fData[k] = v[0]
			} else {
				fData[k] = v
			}
		}
		dataJson, _ := json.Marshal(fData)
		params = string(dataJson)
	}
	return params
}
