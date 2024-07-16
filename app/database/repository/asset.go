package repository

import (
	"context"

	"github.com/q4Zar/go-rest-api/database/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"goyave.dev/filter"
	"goyave.dev/goyave/v5/database"
	"goyave.dev/goyave/v5/util/errors"
	"goyave.dev/goyave/v5/util/session"
)

type Asset struct {
	DB *gorm.DB
}

func NewAsset(db *gorm.DB) *Asset {
	return &Asset{
		DB: db,
	}
}

func (r *Asset) Index(ctx context.Context, request *filter.Request) (*database.Paginator[*model.Asset], error) {
	settings := &filter.Settings[*model.Asset]{
		DefaultSort: []*filter.Sort{
			{Field: "created_at", Order: filter.SortDescending},
		},
		// FieldsSearch: []uint{"userID"},
		Blacklist: filter.Blacklist{
			FieldsBlacklist: []string{"deleted_at", "created_at", "updated_at"},
			// Relations: map[string]*filter.Blacklist{
			// 	"currencyID": {IsFinal: true},
			// },
		},
	}
	paginator, err := settings.Scope(session.DB(ctx, r.DB), request, &[]*model.Asset{})
	return paginator, errors.New(err)
}

func (r *Asset) GetByID(ctx context.Context, id uint) (*model.Asset, error) {
	var asset *model.Asset
	db := session.DB(ctx, r.DB).Where("id", id).First(&asset)
	return asset, errors.New(db.Error)
}

// func (r *Asset) GetBySlug(ctx context.Context, slug string) (*model.Asset, error) {
// 	var asset *model.Asset
// 	db := session.DB(ctx, r.DB).Where("slug", slug).First(&asset)
// 	return asset, errors.New(db.Error)
// }

func (r *Asset) Create(ctx context.Context, asset *model.Asset) (*model.Asset, error) {
	db := session.DB(ctx, r.DB).Omit(clause.Associations).Create(&asset)
	return asset, errors.New(db.Error)
}

func (r *Asset) Update(ctx context.Context, asset *model.Asset) (*model.Asset, error) {
	db := session.DB(ctx, r.DB).Omit(clause.Associations).Save(&asset)
	return asset, errors.New(db.Error)
}

func (r *Asset) Delete(ctx context.Context, id uint) error {
	db := session.DB(ctx, r.DB).Delete(&model.Asset{ID: id})
	if db.RowsAffected == 0 {
		return errors.New(gorm.ErrRecordNotFound)
	}
	return errors.New(db.Error)
}

func (r *User) Paginate(ctx context.Context, request *filter.Request) (*database.Paginator[*model.Asset], error) {
	assets := []*model.Asset{}
	paginator, err := filter.Scope(session.DB(ctx, r.DB), request, &assets)
	return paginator, errors.New(err)
}
	