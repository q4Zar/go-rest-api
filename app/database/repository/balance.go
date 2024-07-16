package repository

import (
	"context"

	"github.com/q4Zar/go-rest-api/database/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	// "goyave.dev/filter"
	// "goyave.dev/goyave/v5/database"
	"goyave.dev/goyave/v5/util/errors"
	"goyave.dev/goyave/v5/util/session"
)

type Balance struct {
	DB *gorm.DB
}

func NewBalance(db *gorm.DB) *Balance {
	return &Balance{
		DB: db,
	}
}

// func (r *Balance) Index(ctx context.Context, request *filter.Request) (*database.Paginator[*model.Balance], error) {
// 	settings := &filter.Settings[*model.Balance]{
// 		DefaultSort: []*filter.Sort{
// 			{Field: "created_at", Order: filter.SortDescending},
// 		},
// 		FieldsSearch: []string{"title"},
// 		Blacklist: filter.Blacklist{
// 			FieldsBlacklist: []string{"deleted_at"},
// 			Relations: map[string]*filter.Blacklist{
// 				"Author": {IsFinal: true},
// 			},
// 		},
// 	}
// 	paginator, err := settings.Scope(session.DB(ctx, r.DB), request, &[]*model.Balance{})
// 	return paginator, errors.New(err)
// }

// func (r *Balance) GetByID(ctx context.Context, id uint) (*model.Balance, error) {
// 	var article *model.Balance
// 	db := session.DB(ctx, r.DB).Where("id", id).First(&article)
// 	return article, errors.New(db.Error)
// }

// func (r *Balance) GetBySlug(ctx context.Context, slug string) (*model.Balance, error) {
// 	var article *model.Balance
// 	db := session.DB(ctx, r.DB).Where("slug", slug).First(&article)
// 	return article, errors.New(db.Error)
// }

func (r *Balance) Create(ctx context.Context, article *model.Balance) (*model.Balance, error) {
	db := session.DB(ctx, r.DB).Omit(clause.Associations).Create(&article)
	return article, errors.New(db.Error)
}

func (r *Balance) Update(ctx context.Context, article *model.Balance) (*model.Balance, error) {
	db := session.DB(ctx, r.DB).Omit(clause.Associations).Save(&article)
	return article, errors.New(db.Error)
}

func (r *Balance) Delete(ctx context.Context, id uint) error {
	db := session.DB(ctx, r.DB).Delete(&model.Balance{ID: id})
	if db.RowsAffected == 0 {
		return errors.New(gorm.ErrRecordNotFound)
	}
	return errors.New(db.Error)
}