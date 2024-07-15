package route

import (
	"net/http"

	"github.com/q4Zar/go-rest-api/http/controller/currency"
	"github.com/q4Zar/go-rest-api/http/controller/user"
	// "github.com/q4Zar/go-rest-api/service"
	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/auth"
	"goyave.dev/goyave/v5/log"
)

func Register(server *goyave.Server, router *goyave.Router) {
	router.GlobalMiddleware(log.CombinedLogMiddleware())
	router.GlobalMiddleware(auth.ConfigBasicAuth()).SetMeta(auth.MetaAuth, true)

	router.Get("/hello", func(response *goyave.Response, request *goyave.Request) {
		response.String(http.StatusOK, "Hello world")
	})

	router.Controller(user.NewController())
	router.Controller(currency.NewController())
}