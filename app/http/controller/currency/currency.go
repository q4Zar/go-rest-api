package currency

import (
	"context"
	"net/http"
	"strconv"
	"log"

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
	Index(ctx context.Context, request *filter.Request) (*database.PaginatorDTO[*dto.Currency], error)
	GetByID(ctx context.Context, id uint) (*dto.ShowCurrency, error)
	Create(ctx context.Context, createDTO *dto.CreateCurrency) error
	Update(ctx context.Context, id uint, updateDTO *dto.UpdateCurrency) error
	Delete(ctx context.Context, id uint) error
	IsOwner(ctx context.Context, resourceID, ownerID uint) (bool, error)
}

type Controller struct {
	goyave.Component
	CurrencyService Service
}

func NewController() *Controller {
	return &Controller{}
}

func (ctrl *Controller) Init(server *goyave.Server) {
	ctrl.Component.Init(server)
	ctrl.CurrencyService = server.Service(service.Currency).(Service)
}

func (ctrl *Controller) RegisterRoutes(router *goyave.Router) {
	subrouter := router.Subrouter("/currencies")
	subrouter.Get("/", ctrl.Index).ValidateQuery(filter.Validation)
	
	authRouter := subrouter.Group().SetMeta(auth.MetaAuth, true)
	authRouter.Post("/", ctrl.Create).ValidateBody(ctrl.CreateRequest)
	
	ownedRouter := authRouter.Group()
	ownerMiddleware := middleware.NewOwner("currencyID", ctrl.CurrencyService)
	ownedRouter.Middleware(ownerMiddleware)
	ownedRouter.Get("/{currencyID:[0-9]+}", ctrl.Show)
	ownedRouter.Patch("/{currencyID:[0-9]+}", ctrl.Update).ValidateBody(ctrl.UpdateRequest)
	ownedRouter.Delete("/{currencyID:[0-9]+}", ctrl.Delete)
}

func (ctrl *Controller) Index(response *goyave.Response, request *goyave.Request) {
	paginator, err := ctrl.CurrencyService.Index(request.Context(), filter.NewRequest(request.Query))
	if response.WriteDBError(err) {
		return
	}
	response.JSON(http.StatusOK, paginator)
}

func (ctrl *Controller) Show(response *goyave.Response, request *goyave.Request) {
	id, err := strconv.ParseUint(request.RouteParams["currencyID"], 10, 64)
	log.Println(id)
	if err != nil {
		response.Status(http.StatusNotFound)
		return
	}

	currency, err := ctrl.CurrencyService.GetByID(request.Context(), uint(id))
	if response.WriteDBError(err) {
		return
	}
	response.JSON(http.StatusOK, currency)
}

func (ctrl *Controller) Create(response *goyave.Response, request *goyave.Request) {
	createDTO := typeutil.MustConvert[*dto.CreateCurrency](request.Data)
	createDTO.OwnerID = request.User.(*dto.InternalUser).ID

	err := ctrl.CurrencyService.Create(request.Context(), createDTO)
	if err != nil {
		response.Error(err)
		return
	}
	response.Status(http.StatusCreated)
}

func (ctrl *Controller) Update(response *goyave.Response, request *goyave.Request) {
	id, err := strconv.ParseUint(request.RouteParams["currencyID"], 10, 64)
	if err != nil {
		response.Status(http.StatusNotFound)
		return
	}

	updateDTO := typeutil.MustConvert[*dto.UpdateCurrency](request.Data)

	err = ctrl.CurrencyService.Update(request.Context(), uint(id), updateDTO)
	response.WriteDBError(err)
}

func (ctrl *Controller) Delete(response *goyave.Response, request *goyave.Request) {
	id, err := strconv.ParseUint(request.RouteParams["currencyID"], 10, 64)
	if err != nil {
		response.Status(http.StatusNotFound)
		return
	}

	err = ctrl.CurrencyService.Delete(request.Context(), uint(id))
	response.WriteDBError(err)
}