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

type Order struct {
	DB *gorm.DB
}

func NewOrder(db *gorm.DB) *Order {
	return &Order{
		DB: db,
	}
}

func (r *Order) Index(ctx context.Context, request *filter.Request) (*database.Paginator[*model.Order], error) {
	settings := &filter.Settings[*model.Order]{
		DisableFields: false, // Prevent usage of "fields"
		DisableFilter: false, // Prevent usage of "filter"
		DisableSort:   false, // Prevent usage of "sort"
		DisableJoin:   false, // Prevent usage of "join"

		DefaultSort: []*filter.Sort{
			{Field: "created_at", Order: filter.SortDescending},
		},
		// FieldsSearch: []string{"title"},
		Blacklist: filter.Blacklist{
			FieldsBlacklist: []string{"deleted_at"},
			Relations: map[string]*filter.Blacklist{
				"User": {IsFinal: true},
			},
		},
	}
	paginator, err := settings.Scope(session.DB(ctx, r.DB), request, &[]*model.Order{})
	return paginator, errors.New(err)
}

func (r *Order) GetByID(ctx context.Context, id uint) (*model.Order, error) {
	var Order *model.Order
	db := session.DB(ctx, r.DB).Where("id", id).First(&Order)
	return Order, errors.New(db.Error)
}

func (r *Order) Create(ctx context.Context, Order *model.Order) (*model.Order, error) {
	db := session.DB(ctx, r.DB).Omit(clause.Associations).Create(&Order)
	return Order, errors.New(db.Error)
}

func (r *Order) Update(ctx context.Context, order *model.Order) (*model.Order, error) {
	db := session.DB(ctx, r.DB).Omit(clause.Associations).Save(&order)
	return order, errors.New(db.Error)
}

func (r *Order) Delete(ctx context.Context, id uint) error {
	db := session.DB(ctx, r.DB).Delete(&model.Order{ID: id})
	if db.RowsAffected == 0 {
		return errors.New(gorm.ErrRecordNotFound)
	}
	return errors.New(db.Error)
}

func (r *Order) IsOwner(ctx context.Context, resourceID, ownerID uint) (bool, error) {
	var one int64
	db := session.DB(ctx, r.DB).
		Table(model.Order{}.TableName()).
		Select("1").
		Where("id", resourceID).
		Where("author_id", ownerID).
		Where("deleted_at IS NULL").
		Find(&one)
	return one == 1, errors.New(db.Error)
}
