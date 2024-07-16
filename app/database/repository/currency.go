package repository

import (
	"context"

	"github.com/q4Zar/go-rest-api/database/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"goyave.dev/goyave/v5/util/errors"
	"goyave.dev/goyave/v5/util/session"
)

type Currency struct {
	DB *gorm.DB
}

func NewCurrency(db *gorm.DB) *Currency {
	return &Currency{
		DB: db,
	}
}


func (r *Currency) GetByID(ctx context.Context, id uint) (*model.Currency, error) {
	var currency *model.Currency
	db := session.DB(ctx, r.DB).Where("id", id).First(&currency)
	return currency, errors.New(db.Error)
}

func (r *Currency) GetByName(ctx context.Context, name string) (*model.Currency, error) {
	var currency *model.Currency
	db := session.DB(ctx, r.DB).Where("name", name).First(&currency)
	return currency, errors.New(db.Error)
}

func (r *Currency) Create(ctx context.Context, Currency *model.Currency) (*model.Currency, error) {
	db := session.DB(ctx, r.DB).Omit(clause.Associations).Create(&Currency)
	return Currency, errors.New(db.Error)
}

func (r *Currency) Update(ctx context.Context, Currency *model.Currency) (*model.Currency, error) {
	db := session.DB(ctx, r.DB).Omit(clause.Associations).Save(&Currency)
	return Currency, errors.New(db.Error)
}

func (r *Currency) Delete(ctx context.Context, id uint) error {
	db := session.DB(ctx, r.DB).Delete(&model.Currency{ID: id})
	if db.RowsAffected == 0 {
		return errors.New(gorm.ErrRecordNotFound)
	}
	return errors.New(db.Error)
}
