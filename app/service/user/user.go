package user

import (
	"context"
	"fmt"

	"github.com/q4Zar/go-rest-api/database/model"
	"github.com/q4Zar/go-rest-api/dto"
	"github.com/q4Zar/go-rest-api/service"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"goyave.dev/goyave/v5/slog"
	"goyave.dev/goyave/v5/util/errors"
	"goyave.dev/goyave/v5/util/session"
	"goyave.dev/goyave/v5/util/typeutil"
)

type Repository interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
	GetByID(ctx context.Context, id uint) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	UniqueScope() func(db *gorm.DB, val any) *gorm.DB
}

type StorageService interface {
	Delete(string) error
}

type Service struct {
	Session        session.Session
	Repository     Repository
	Logger         *slog.Logger
}

func NewService(session session.Session, logger *slog.Logger, repository Repository, ) *Service {
	return &Service{
		Session:        session,
		Logger:         logger,
		Repository:     repository,
	}
}

func (s *Service) UniqueScope() func(db *gorm.DB, val any) *gorm.DB {
	return s.Repository.UniqueScope()
}

func (s *Service) GetByID(ctx context.Context, id uint) (*dto.InternalUser, error) {
	user, err := s.Repository.GetByID(ctx, id)
	if err != nil {
		return nil, errors.New(err)
	}
	return typeutil.MustConvert[*dto.InternalUser](user), nil
}

func (s *Service) FindByUsername(ctx context.Context, username any) (*dto.InternalUser, error) {
	user, err := s.Repository.GetByUsername(ctx, fmt.Sprintf("%v", username))
	if err != nil {
		return nil, errors.New(err)
	}
	return typeutil.MustConvert[*dto.InternalUser](user), nil
}

func (s *Service) Register(ctx context.Context, registerDTO *dto.RegisterUser) error {
	err := s.Session.Transaction(ctx, func(ctx context.Context) error {
		user := typeutil.Copy(&model.User{}, registerDTO)

		b, err := bcrypt.GenerateFromPassword([]byte(registerDTO.Password), bcrypt.DefaultCost)
		if err != nil {
			return errors.New(err)
		}
		user.Password = string(b)

		_, err = s.Repository.Create(ctx, user)
		// deleted stuff here
		return errors.New(err)
	})
	return errors.New(err)
}

func (s *Service) Update(ctx context.Context, id uint, updateDTO *dto.UpdateUser) error {
	err := s.Session.Transaction(ctx, func(ctx context.Context) error {
		var err error
		user, err := s.Repository.GetByID(ctx, id)
		if err != nil {
			return errors.New(err)
		}

		user = typeutil.Copy(user, updateDTO)

		_, err = s.Repository.Update(ctx, user)
		return errors.New(err)
	})

	return errors.New(err)
}

func (s *Service) Name() string {
	return service.User
}