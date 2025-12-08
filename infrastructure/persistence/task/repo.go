package task

import (
	"context"

	"github.com/manyodream/todolist-ddd/domain/infrastructure/persistence/user"
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

func Paginate(p types.Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if p.Page <= 0 {
			p.Page = 1
		}
		switch {
		case p.PageSize > 100:
			p.PageSize = 100
		case p.PageSize <= 0:
			p.PageSize = 10
		}
		offset := (p.Page - 1) * p.PageSize
		return db.Offset(offset).Limit(p.PageSize)
	}
}

func (r RepositoryImpl) CreateTask(ctx context.Context, task *entity.Task) (*entity.Task, error) {
	taskPO := Entity2PO(task)
	db := r.db.WithContext(ctx)
	err := db.Model(&Task{}).Create(&taskPO).Error
	if err != nil {
		return nil, err
	}
	var u *user.User
	err = db.Model(&user.User{}).Where("id = ?", task.UID).Find(&u).Error
	if err != nil {
		return nil, err
	}
	return PO2Entity(taskPO, u), nil
}

func (r RepositoryImpl) DeteleTask(ctx context.Context, uid, tid uint) error {
	err := r.db.WithContext(ctx).Model(&Task{}).
		Where("id = ? AND uid = ?", tid, uid).
		Delete(&Task{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r RepositoryImpl) FindTaskByTid(ctx context.Context, tid, uid uint) (*entity.Task, error) {
	task := &entity.Task{}
	err := r.db.WithContext(ctx).Model(&Task{}).
		Joins("AS task LEFT JOIN user AS u ON task.uid = u.id").
		Where("task.id = ? AND u.id = ?", tid, uid).
		Select("u.id AS uid,u.user_name AS user_name, task.*").
		Find(&task).Error
	return task, err
}

func (r RepositoryImpl) ListTaskByUid(ctx context.Context, uid uint, p types.Pagination) ([]*entity.Task, int64, error) {
	var tasks []*entity.Task
	var count int64
	err := r.db.WithContext(ctx).Model(&Task{}).
		Joins("AS task LEFT JOIN user AS u ON task.uid = u.id").
		Where("u.id = ?", uid).Count(&count).
		Scopes(Paginate(p)).
		Select("u.id AS uid,u.user_name AS user_name, task.*").
		Find(&tasks).Error
	if err != nil {
		return nil, 0, err
	}
	return tasks, count, err
}

func (r RepositoryImpl) SearchTaskByTid(ctx context.Context, uid uint, keyword string, p types.Pagination) ([]*entity.Task, int64, error) {
	var tasks []*entity.Task
	var count int64
	err := r.db.WithContext(ctx).Model(&Task{}).
		Where("u.id = ?", uid).
		Where("title LIKE ? OR content LIKE ?", "%"+keyword+"%", "%"+keyword+"%").
		Count(&count).
		Scopes(Paginate(p)).
		Select("id, uid, title, content, status, start_time, end_time").
		Find(&tasks).Error
	if err != nil {
		return nil, 0, err
	}
	return tasks, count, err
}

func (r RepositoryImpl) UpdateTask(ctx context.Context, task *entity.Task) error {
	update := make(map[string]any)
	if task.Content != "" {
		update["content"] = task.Content
	}
	if task.Status != 0 {
		update["status"] = task.Status
	}
	if task.Title != "" {
		update["title"] = task.Title
	}
	return r.db.WithContext(ctx).Model(&Task{}).
		Where("id = ? AND uid = ?", task.ID, task.UID).Updates(update).Error
}
