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

// func (r *Asset) Index(ctx context.Context, request *filter.Request) (*database.Paginator[*model.Asset], error) {
// 	settings := &filter.Settings[*model.Asset]{
// 		DefaultSort: []*filter.Sort{
// 			{Field: "created_at", Order: filter.SortDescending},
// 		},
// 		// FieldsSearch: []uint{"userID"},
// 		Blacklist: filter.Blacklist{
// 			FieldsBlacklist: []string{"deleted_at"},
// 			// Relations: map[string]*filter.Blacklist{
// 			// "Author": {IsFinal: true},
// 			// },
// 		},
// 	}
// 	paginator, err := settings.Scope(session.DB(ctx, r.DB), request, &[]*model.Asset{})
// 	return paginator, errors.New(err)
// }

// First returns the user identified by the given ID, or `nil`
func (r *Asset) First(ctx context.Context, id uint) (*model.Asset, error) {
	var asset *model.Asset
	db := session.DB(ctx, r.DB).Where("id", id).First(&asset)
	return asset, errors.New(db.Error)
}

func (r *Asset) GetByID(ctx context.Context, id uint) (*model.Asset, error) {
	var asset *model.Asset
	db := session.DB(ctx, r.DB).Where("id", id).First(&asset)
	return asset, errors.New(db.Error)
}

func (r *Asset) Paginate(ctx context.Context, request *filter.Request) (*database.Paginator[*model.Asset], error) {
	settings := &filter.Settings[*model.Asset]{
		DisableFields: false, // Prevent usage of "fields"
		DisableFilter: false, // Prevent usage of "filter"
		DisableSort:   false, // Prevent usage of "sort"
		DisableJoin:   false, // Prevent usage of "join"

		// If not nil and not empty, and if the request is not providing any
		// sort, the request will be sorted according to the `*Sort` defined in this slice.
		// If `DisableSort` is enabled, this has no effect.
		DefaultSort: []*filter.Sort{{Field: "create_at", Order: filter.SortDescending}},

		// If true, the sort will wrap the value in `LOWER()` if it's a string, resulting in `ORDER BY LOWER(column)`.
		// CaseInsensitiveSort: true,

		// FieldsSearch:   []string{"a", "b"},      // Optional, the fields used for the search feature
		// SearchOperator: filter.Operators["$eq"], // Optional, operator used for the search feature, defaults to "$cont"

		Blacklist: filter.Blacklist{
			// Prevent selecting, sorting and filtering on these fields
			FieldsBlacklist: []string{"deleted_at"},

			// Prevent joining these relations
			// RelationsBlacklist: []string{"Relation"},

			// Relations: map[string]*filter.Blacklist{
			// 	// Blacklist settings to apply to this relation
			// 	"Relation": &filter.Blacklist{
			// 		FieldsBlacklist:    []string{"c", "d"},
			// 		RelationsBlacklist: []string{"Parent"},
			// 		Relations:          map[string]*filter.Blacklist{ /*...*/ },
			// 		IsFinal:            true, // Prevent joining any child relation if true
			// 	},
			// },
		},
	}
	assets := []*model.Asset{}
	// paginator, err := filter.Scope(session.DB(ctx, r.DB), request, &assets)
	paginator, err := settings.Scope(session.DB(ctx, r.DB), request, &assets)
	return paginator, errors.New(err)
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
