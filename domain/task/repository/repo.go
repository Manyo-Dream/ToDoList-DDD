package repository

import (
	"context"

	"github.com/manyodream/todolist-ddd/domain/task/entity"
	"github.com/manyodream/todolist-ddd/interfaces/types"
)

type Task interface {
	TaskBase
	TaskQuery
}

type TaskBase interface {
	CreateTask(ctx context.Context, task *entity.Task) (*entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	ListTaskByUid(ctx context.Context, uid uint, p *types.Pagination) ([]*entity.Task, error)
	FindTaskByTid(ctx context.Context, tid uint) (*entity.Task, error)
	SearchTaskByTid(ctx context.Context, uid uint, keyword string, p *types.Pagination) ([]*entity.Task, error)
	DeteleTask(ctx context.Context, uid, tid uint) error
}

type TaskQuery interface {
}
