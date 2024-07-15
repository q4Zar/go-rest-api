package repository

import (
	"context"
	"log"

	"github.com/q4Zar/go-rest-api/database/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"goyave.dev/filter"
	"goyave.dev/goyave/v5/database"
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

func (r *Currency) Index(ctx context.Context, request *filter.Request) (*database.Paginator[*model.Currency], error) {
	settings := &filter.Settings[*model.Currency]{
		DefaultSort: []*filter.Sort{
			{Field: "created_at", Order: filter.SortDescending},
		},
		FieldsSearch: []string{"name"},
		Blacklist: filter.Blacklist{
			FieldsBlacklist: []string{"deleted_at"},
			Relations: map[string]*filter.Blacklist{
				"Owner": {IsFinal: true},
			},
		},
	}
	paginator, err := settings.Scope(session.DB(ctx, r.DB), request, &[]*model.Currency{})
	return paginator, errors.New(err)
}

func (r *Currency) GetByID(ctx context.Context, id uint) (*model.Currency, error) {
	var Currency *model.Currency
	db := session.DB(ctx, r.DB).Where("id", id).First(&Currency)
	log.Println(Currency)
	return Currency, errors.New(db.Error)
}

func (r *Currency) GetBySlug(ctx context.Context, slug string) (*model.Currency, error) {
	var Currency *model.Currency
	db := session.DB(ctx, r.DB).Where("slug", slug).First(&Currency)
	return Currency, errors.New(db.Error)
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

func (r *Currency) IsOwner(ctx context.Context, resourceID, ownerID uint) (bool, error) {
	var one int64
	db := session.DB(ctx, r.DB).
		Table(model.Currency{}.TableName()).
		Select("1").
		Where("id", resourceID).
		Where("owner_id", ownerID).
		Where("deleted_at IS NULL").
		Find(&one)
	return one == 1, errors.New(db.Error)
}