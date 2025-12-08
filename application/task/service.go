package task

import (
	"context"
	"errors"
	"sync"

	ctl "github.com/manyodream/todolist-ddd/infrastructure/context"

	"github.com/manyodream/todolist-ddd/domain/task/entity"
	"github.com/manyodream/todolist-ddd/domain/task/service"
	"github.com/manyodream/todolist-ddd/interfaces/types"
)

type Service interface {
	TaskCreate(ctx context.Context, taskReq *types.TaskCreateReq) (*entity.Task, error)
	TaskList(ctx context.Context, taskReq *types.TaskListReq) (any, error)
	TaskDetail(ctx context.Context, taskReq *types.TaskDetailReq) (*entity.Task, error)
	TaskUpdate(ctx context.Context, taskReq *types.TaskUpdateReq) error
	TaskSearch(ctx context.Context, taskReq *types.TaskSearchReq) (any, error)
	TaskDetele(ctx context.Context, taskReq *types.TaskDeteleReq) error
}

type ServiceImpl struct {
	td service.TaskDomain
}

var (
	ServiceImplIns  *ServiceImpl
	ServiceImplOnce sync.Once
)

func NewServiceImpl(td service.TaskDomain) *ServiceImpl {
	ServiceImplOnce.Do(func() {
		ServiceImplIns = &ServiceImpl{
			td: td,
		}
	})
	return ServiceImplIns
}

func (s *ServiceImpl) TaskCreate(ctx context.Context, taskReq *types.TaskCreateReq) (*entity.Task, error) {
	userInfo, err := ctl.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	task, err := entity.NewTask(userInfo.Id, userInfo.Name, taskReq.Title, taskReq.Content)
	if err != nil {
		return nil, err
	}
	return s.td.CreateTask(ctx, task)
}

func (s *ServiceImpl) TaskList(ctx context.Context, taskReq *types.TaskListReq) (any, error) {
	userInfo, err := ctl.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	list, count, err := s.td.ListTaskByUid(ctx, userInfo.Id, taskReq.Pagination)
	if err != nil {
		return nil, err
	}
	return ListResponse(list, count), nil
}

func (s *ServiceImpl) TaskDetail(ctx context.Context, taskReq *types.TaskDetailReq) (*entity.Task, error) {
	userInfo, err := ctl.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	task, err := s.td.FindTaskByTid(ctx, taskReq.ID, userInfo.Id)
	if err != nil {
		return nil, err
	}
	if !task.IsExist() {
		return nil, errors.New("task not exist")
	}
	if !task.Belong2User(userInfo.Id) {
		return nil, errors.New("task not belong to user")
	}
	return task, s.td.UpdateTask(ctx, task)
}

func (s *ServiceImpl) TaskUpdate(ctx context.Context, taskReq *types.TaskUpdateReq) error {
	userInfo, err := ctl.GetUserInfo(ctx)
	if err != nil {
		return err
	}
	task := UpdateRep2TaskEntity(taskReq.ID, userInfo.Id, userInfo.Name, taskReq)
	return s.td.UpdateTask(ctx, task)
}

func (s *ServiceImpl) TaskSearch(ctx context.Context, taskReq *types.TaskSearchReq) (any, error) {
	userInfo, err := ctl.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}
	tasks, count, err := s.td.SearchTask(ctx, userInfo.Id, taskReq.Info, taskReq.Pagination)
	if err != nil {
		return nil, err
	}
	return ListResponse(tasks, count), nil
}

func (s *ServiceImpl) TaskDetele(ctx context.Context, taskReq *types.TaskDeteleReq) error {
	userInfo, err := ctl.GetUserInfo(ctx)
	if err != nil {
		return err
	}
	return s.td.DeleteTask(ctx, userInfo.Id, taskReq.ID)
}
