package order

import (
	"context"
	"net/http"
	"strconv"

	"github.com/q4Zar/go-rest-api/dto"
	"github.com/q4Zar/go-rest-api/http/middleware"
	"github.com/q4Zar/go-rest-api/service"
	"goyave.dev/filter"
	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/auth"
	"goyave.dev/goyave/v5/database"
	"goyave.dev/goyave/v5/util/typeutil"
)

type Service interface {
	Index(ctx context.Context, request *filter.Request) (*database.PaginatorDTO[*dto.Order], error)
	Create(ctx context.Context, createDTO *dto.CreateOrder) error
	Delete(ctx context.Context, id uint) error
	IsOwner(ctx context.Context, resourceID, ownerID uint) (bool, error)
}

type Controller struct {
	goyave.Component
	OrderService Service
}

func NewController() *Controller {
	return &Controller{}
}

func (ctrl *Controller) Init(server *goyave.Server) {
	ctrl.Component.Init(server)
	ctrl.OrderService = server.Service(service.Order).(Service)
}

func (ctrl *Controller) RegisterRoutes(router *goyave.Router) {
	subrouter := router.Subrouter("/orders")

	authRouter := subrouter.Group().SetMeta(auth.MetaAuth, true)
	authRouter.Post("/", ctrl.Create).ValidateBody(ctrl.CreateRequest)
	authRouter.Get("/", ctrl.Index).ValidateQuery(filter.Validation)

	ownedRouter := authRouter.Group()
	ownerMiddleware := middleware.NewOwner("orderD", ctrl.OrderService)
	ownedRouter.Middleware(ownerMiddleware)
	ownedRouter.Delete("/{orderID:[0-9]+}", ctrl.Delete)
}

func (ctrl *Controller) Index(response *goyave.Response, request *goyave.Request) {
	paginator, err := ctrl.OrderService.Index(request.Context(), filter.NewRequest(request.Query))
	if response.WriteDBError(err) {
		return
	}
	response.JSON(http.StatusOK, paginator)
}

func (ctrl *Controller) Create(response *goyave.Response, request *goyave.Request) {
	createDTO := typeutil.MustConvert[*dto.CreateOrder](request.Data)
	createDTO.UserID = request.User.(*dto.InternalUser).ID
	err := ctrl.OrderService.Create(request.Context(), createDTO)
	if err != nil {
		response.Error(err)
		return
	}
	response.Status(http.StatusCreated)
}

func (ctrl *Controller) Delete(response *goyave.Response, request *goyave.Request) {
	id, err := strconv.ParseUint(request.RouteParams["orderID"], 10, 64)
	if err != nil {
		response.Status(http.StatusNotFound)
		return
	}

	err = ctrl.OrderService.Delete(request.Context(), uint(id))
	response.WriteDBError(err)
}
