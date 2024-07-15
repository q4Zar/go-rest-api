package user

import (
	"context"
	"net/http"

	"github.com/q4Zar/go-rest-api/dto"
	"github.com/q4Zar/go-rest-api/service"
	"gorm.io/gorm"
	"goyave.dev/goyave/v5"
	"goyave.dev/goyave/v5/auth"
	"goyave.dev/goyave/v5/util/typeutil"
)

type Service interface {
	UniqueScope() func(db *gorm.DB, val any) *gorm.DB
	GetByID(ctx context.Context, id uint) (*dto.InternalUser, error)
	Register(ctx context.Context, registerDTO *dto.RegisterUser) error
	Update(ctx context.Context, id uint, updateDTO *dto.UpdateUser) error
}


type Controller struct {
	goyave.Component
	UserService    Service
}

func NewController() *Controller {
	return &Controller{}
}

func (ctrl *Controller) Init(server *goyave.Server) {
	ctrl.Component.Init(server)
	ctrl.UserService = server.Service(service.User).(Service)
}

func (ctrl *Controller) RegisterRoutes(router *goyave.Router) {
	subrouter := router.Subrouter("/users")
	subrouter.Post("/", ctrl.Register).ValidateBody(ctrl.RegisterRequest)

	authRouter := subrouter.Group().SetMeta(auth.MetaAuth, true)
	authRouter.Get("/profile", ctrl.ShowProfile)
	authRouter.Patch("/", ctrl.Update).ValidateBody(ctrl.UpdateRequest)
}

func (ctrl *Controller) ShowProfile(response *goyave.Response, request *goyave.Request) {
	userDTO := typeutil.MustConvert[*dto.User](request.User)
	response.JSON(http.StatusOK, userDTO)
}

func (ctrl *Controller) Register(response *goyave.Response, request *goyave.Request) {
	registerDTO := typeutil.MustConvert[*dto.RegisterUser](request.Data)

	err := ctrl.UserService.Register(request.Context(), registerDTO)
	if err != nil {
		response.Error(err)
		return
	}
	response.Status(http.StatusCreated)
}

func (ctrl *Controller) Update(response *goyave.Response, request *goyave.Request) {
	updateDTO := typeutil.MustConvert[*dto.UpdateUser](request.Data)
	id := request.User.(*dto.InternalUser).ID

	err := ctrl.UserService.Update(request.Context(), id, updateDTO)
	if err != nil {
		response.Error(err)
		return
	}
}