package main

import (
	"net/http"

	"github.com/d5/tengo/v2"
	"github.com/d5/tengo/v2/stdlib"
	tengohttp "github.com/damiva/TengoHTTP"
)

const code = `
fmt := import("fmt")
http := import("httpclient")

resp := http.request("https://httpbin.org/anything", "GET")
fmt.print(is_error(resp) ? resp : resp.body)
`

var r = &http.Request{}
var w = new(http.ResponseWriter)

func main() {
	script := tengo.NewScript([]byte(code))
	modules := stdlib.GetModuleMap("fmt")
	modules.AddBuiltinModule("httpclient",
		tengohttp.GetModuleMAP(*w, r, nil, map[string]tengo.Object{}))
	script.SetImports(modules)
	if _, err := script.Run(); err != nil {
		panic(err)
	}
}
