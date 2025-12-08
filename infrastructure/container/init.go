package container

import (
	"github.com/manyodream/todolist-ddd/application/task"
	"github.com/manyodream/todolist-ddd/application/user"
	taskSrv "github.com/manyodream/todolist-ddd/domain/task/service"
	"github.com/manyodream/todolist-ddd/domain/user/service"
	"github.com/manyodream/todolist-ddd/infrastructure/auth"
	"github.com/manyodream/todolist-ddd/infrastructure/encrypt"
	"github.com/manyodream/todolist-ddd/infrastructure/persistence"
	"github.com/manyodream/todolist-ddd/infrastructure/persistence/dbs"
)

func LoadingDomain() {
	repos := persistence.NewRepositories(dbs.DB)
	pwdEncryptService := encrypt.NewPwdEncryptService()

	jwtToken := auth.NewJWTTokenService()
	userDomain := service.NewUserDomainImpl(repos.User, pwdEncryptService)
	user.NewServiceImpl(userDomain, jwtToken)

	taskDomain := taskSrv.NewTaskDomainImpl(repos.Task)
	task.NewServiceImpl(taskDomain)
}
