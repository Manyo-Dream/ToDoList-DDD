package service

import (
	"context"

	"github.com/manyodream/todolist-ddd/domain/task/entity"
	"github.com/manyodream/todolist-ddd/domain/task/repository"
	"github.com/manyodream/todolist-ddd/interfaces/types"
)

type TaskDomain interface {
	CreateTask(ctx context.Context, task *entity.Task) (*entity.Task, error)
	FindTaskByTid(ctx context.Context, tid, uid uint) (*entity.Task, error)
	ListTaskByUid(ctx context.Context, uid uint, p types.Pagination) ([]*entity.Task, int64, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	SearchTask(ctx context.Context, uid uint, keyword string, p types.Pagination) ([]*entity.Task, int64, error)
	DeleteTask(ctx context.Context, uid, tid uint) error
}

type TaskDomainImpl struct {
	repo repository.Task
}

func NewTaskDomainImpl(repo repository.Task) TaskDomain {
	return &TaskDomainImpl{repo: repo}
}

// CreateTask implements TaskDomain.
func (t *TaskDomainImpl) CreateTask(ctx context.Context, task *entity.Task) (*entity.Task, error) {
	return t.repo.CreateTask(ctx, task)
}

// DeleteTask implements TaskDomain.
func (t *TaskDomainImpl) DeleteTask(ctx context.Context, uid uint, tid uint) error {
	return t.repo.DeteleTask(ctx, uid, tid)
}

// FindTaskByTid implements TaskDomain.
func (t *TaskDomainImpl) FindTaskByTid(ctx context.Context, tid uint, uid uint) (*entity.Task, error) {
	return t.repo.FindTaskByTid(ctx, tid, uid)
}

// ListTaskByUid implements TaskDomain.
func (t *TaskDomainImpl) ListTaskByUid(ctx context.Context, uid uint, p types.Pagination) ([]*entity.Task, int64, error) {
	return t.repo.ListTaskByUid(ctx, uid, p)
}

// SearchTask implements TaskDomain.
func (t *TaskDomainImpl) SearchTask(ctx context.Context, uid uint, keyword string, p types.Pagination) ([]*entity.Task, int64, error) {
	return t.repo.SearchTaskByTid(ctx, uid, keyword, p)
}

// UpdateTask implements TaskDomain.
func (t *TaskDomainImpl) UpdateTask(ctx context.Context, task *entity.Task) error {
	return t.repo.UpdateTask(ctx, task)
}
