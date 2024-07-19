package route

import (
	"github.com/q4Zar/go-rest-api/http/controller/asset"
	"github.com/q4Zar/go-rest-api/http/controller/user"

	"github.com/q4Zar/go-rest-api/http/controller/order"
	"github.com/q4Zar/go-rest-api/service"
	userservice "github.com/q4Zar/go-rest-api/service/user"
	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/auth"
	"goyave.dev/goyave/v5/cors"
	"goyave.dev/goyave/v5/log"
	"goyave.dev/goyave/v5/middleware/parse"
)

func Register(server *goyave.Server, router *goyave.Router) {
	router.CORS(cors.Default())
	router.GlobalMiddleware(log.CombinedLogMiddleware())

	userService := server.Service(service.User).(*userservice.Service)

	authenticator := auth.NewJWTAuthenticator(userService)
	authMiddleware := auth.Middleware(authenticator)
	router.GlobalMiddleware(authMiddleware)
	router.GlobalMiddleware(&parse.Middleware{})

	loginController := auth.NewJWTController(userService, "Password")

	router.Controller(loginController)
	router.Controller(asset.NewController())
	router.Controller(order.NewController())
	router.Controller(user.NewController())

	// router.Post("/orders", controllers.CreateOrder)
	// router.Get("/orders", controllers.GetOrders)
	// router.Get("/assets", controllers.GetAssets)
}
