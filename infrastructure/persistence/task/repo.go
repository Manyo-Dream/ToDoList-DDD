package task

import (
	"context"

	"github.com/manyodream/todolist-ddd/domain/task/entity"
	"github.com/manyodream/todolist-ddd/domain/task/repository"
	"github.com/manyodream/todolist-ddd/interfaces/types"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repository.Task {
	return &RepositoryImpl{
		db: db,
	}
}

func (r RepositoryImpl) CreateTask(ctx context.Context, task *entity.Task) (*entity.Task, error) {

}

func (r RepositoryImpl) DeteleTask(ctx context.Context, uid uint, tid uint) error {

}

func (r RepositoryImpl) FindTaskByTid(ctx context.Context, tid uint) (*entity.Task, error) {

}

func (r RepositoryImpl) ListTaskByUid(ctx context.Context, uid uint, p *types.Pagination) ([]*entity.Task, error) {

}

func (r RepositoryImpl) SearchTaskByTid(ctx context.Context, uid uint, keyword string, p *types.Pagination) ([]*entity.Task, error) {

}

func (r RepositoryImpl) UpdateTask(ctx context.Context, task *entity.Task) error {

}
