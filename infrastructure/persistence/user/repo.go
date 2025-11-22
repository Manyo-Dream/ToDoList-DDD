package user

import (
	"context"

	"github.com/manyodream/todolist-ddd/domain/user/entity"
	"github.com/manyodream/todolist-ddd/domain/user/repository"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repository.User {
	return &RepositoryImpl{db: db}
}

func (r *RepositoryImpl) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	userPO := Entity2PO(user)
	err := r.db.WithContext(ctx).Model(&User{}).Create(userPO).Error
	if err != nil {
		return nil, err
	}
	return PO2Entity(userPO), nil
}

func (r *RepositoryImpl) GetUserByName(ctx context.Context, name string) (*entity.User, error) {
	var userPO User
	err := r.db.WithContext(ctx).Model(&User{}).
		Where("user_name = ?", name).Find(&userPO).Error
	if err != nil {
		return nil, err
	}
	return PO2Entity(&userPO), nil
}

func (r *RepositoryImpl) GetUserByID(ctx context.Context, id int64) (*entity.User, error) {
	var userPO User
	err := r.db.WithContext(ctx).Model(&User{}).
		Where("id = ?", id).Find(&userPO).Error
	if err != nil {
		return nil, err
	}
	return PO2Entity(&userPO), nil
}
