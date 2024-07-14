package route

import (
	"net/http"

	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/auth"
	// "goyave.dev/goyave/v5/cors"
	// "goyave.dev/goyave/v5/middleware/parse"
	// _ "goyave.dev/template/http/controller/user"
)

func Register(_ *goyave.Server, router *goyave.Router) {
	router.GlobalMiddleware(auth.ConfigBasicAuth()).SetMeta(auth.MetaAuth, true)
	router.Get("/hello", func(response *goyave.Response, request *goyave.Request) {
		response.String(http.StatusOK, "Hello world")
	})
}