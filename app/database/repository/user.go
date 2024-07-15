package repository

import (
	"context"

	"github.com/q4Zar/go-rest-api/database/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"goyave.dev/goyave/v5/util/errors"
	"goyave.dev/goyave/v5/util/session"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{
		DB: db,
	}
}

func (r *User) GetByID(ctx context.Context, id uint) (*model.User, error) {
	var user *model.User
	db := session.DB(ctx, r.DB).Where("id", id).First(&user)
	return user, errors.New(db.Error)
}

func (r *User) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	var user *model.User
	db := session.DB(ctx, r.DB).Where("username", username).First(&user)
	return user, errors.New(db.Error)
}

func (r *User) Create(ctx context.Context, user *model.User) (*model.User, error) {
	db := session.DB(ctx, r.DB).Omit(clause.Associations).Create(&user)
	return user, errors.New(db.Error)
}

func (r *User) Update(ctx context.Context, user *model.User) (*model.User, error) {
	db := session.DB(ctx, r.DB).Omit(clause.Associations).Save(&user)
	return user, errors.New(db.Error)
}

func (r *User) UniqueScope() func(db *gorm.DB, val any) *gorm.DB {
	return func(db *gorm.DB, val any) *gorm.DB {
		return db.Table(model.User{}.TableName()).Where("username", val)
	}
}