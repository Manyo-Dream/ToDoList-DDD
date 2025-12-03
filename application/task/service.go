package task

import (
	"context"
	"sync"

	"github.com/manyodream/todolist-ddd/domain/task/entity"
	"github.com/manyodream/todolist-ddd/interfaces/types"
)

type Service interface {
	TaskCreate(ctx context.Context, task *types.TaskCreateReq) (*entity.Task, error)
	TaskList(ctx context.Context, task *types.TaskList) (any, error)
	TaskDetail(ctx context.Context, task *types.TaskDetail) (*entity.Task, error)
	TaskUpdate(ctx context.Context, task *types.TaskUpdate) error
	TaskSearch(ctx context.Context, task *types.TaskSearch) (any, error)
	TaskDetele(ctx context.Context, task *types.TaskDetele) error
}

type ServiceImpl struct {
}

var (
	ServiceImplIns  *ServiceImpl
	ServiceImplOnce sync.Once
)

func NewServiceImpl() *ServiceImpl {
	ServiceImplOnce.Do(func() {
		ServiceImplIns = &ServiceImpl{}
	})
	return ServiceImplIns
}

func (s *ServiceImpl) TaskCreate(ctx context.Context, task *types.TaskCreateReq) (*entity.Task, error) {

}

func (s *ServiceImpl) TaskList(ctx context.Context, task *types.TaskList) (any, error) {

}

func (s *ServiceImpl) TaskDetail(ctx context.Context, task *types.TaskDetail) (*entity.Task, error) {

}

func (s *ServiceImpl) TaskUpdate(ctx context.Context, task *types.TaskUpdate) error {

}

func (s *ServiceImpl) TaskSearch(ctx context.Context, task *types.TaskSearch) (any, error) {

}

func (s *ServiceImpl) TaskDetele(ctx context.Context, task *types.TaskDetele) error {

}
