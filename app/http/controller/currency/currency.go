package currency

import (
	"context"
	"net/http"
	"strconv"
	"log"

	"github.com/q4Zar/go-rest-api/dto"
	"github.com/q4Zar/go-rest-api/service"
	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/auth"
	"goyave.dev/goyave/v5/util/typeutil"
)

type Service interface {
	Create(ctx context.Context, createDTO *dto.CreateCurrency) error
	Update(ctx context.Context, id uint, updateDTO *dto.UpdateCurrency) error
	Delete(ctx context.Context, id uint) error
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
	
	authRouter := subrouter.Group().SetMeta(auth.MetaAuth, true)
	authRouter.Post("/", ctrl.Create).ValidateBody(ctrl.CreateRequest)
	authRouter.Patch("/{currencyID:[0-9]+}", ctrl.Update).ValidateBody(ctrl.UpdateRequest)
	authRouter.Delete("/{currencyID:[0-9]+}", ctrl.Delete)
}


func (ctrl *Controller) Create(response *goyave.Response, request *goyave.Request) {
	createDTO := typeutil.MustConvert[*dto.CreateCurrency](request.Data)

	err := ctrl.CurrencyService.Create(request.Context(), createDTO)
	if err != nil {
		response.Error(err)
		return
	}
	response.Status(http.StatusCreated)
}

func (ctrl *Controller) Update(response *goyave.Response, request *goyave.Request) {
	id, err := strconv.ParseUint(request.RouteParams["currencyID"], 10, 64)
	log.Println(id)
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