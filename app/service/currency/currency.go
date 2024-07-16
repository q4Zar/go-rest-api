package currency

import (
	"context"

	"github.com/q4Zar/go-rest-api/database/model"
	"github.com/q4Zar/go-rest-api/dto"
	"github.com/q4Zar/go-rest-api/service"
	"goyave.dev/goyave/v5/util/errors"
	"goyave.dev/goyave/v5/util/session"
	"goyave.dev/goyave/v5/util/typeutil"
)

// func init() {
// }

type Repository interface {
	GetByID(ctx context.Context, id uint) (*model.Currency, error)
	Create(ctx context.Context, currency *model.Currency) (*model.Currency, error)
	Update(ctx context.Context, currency *model.Currency) (*model.Currency, error)
	Delete(ctx context.Context, id uint) error
}

type Service struct {
	Session    session.Session
	Repository Repository
}

func NewService(session session.Session, repository Repository) *Service {
	return &Service{
		Session:    session,
		Repository: repository,
	}
}

func (s *Service) Create(ctx context.Context, createDTO *dto.CreateCurrency) error {
	currency := typeutil.Copy(&model.Currency{}, createDTO)
	_, err := s.Repository.Create(ctx, currency)
	return errors.New(err)
}

func (s *Service) Update(ctx context.Context, id uint, updateDTO *dto.UpdateCurrency) error {
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

func (s *Service) Delete(ctx context.Context, id uint) error {
	return s.Repository.Delete(ctx, id)
}

func (s *Service) Name() string {
	return service.Currency
}