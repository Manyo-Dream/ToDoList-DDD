package container

import (
	"github.com/manyodream/todolist-ddd/application/user"
	"github.com/manyodream/todolist-ddd/domain/user/service"
	"github.com/manyodream/todolist-ddd/infrastructure/persistence"
	"github.com/manyodream/todolist-ddd/infrastructure/persistence/dbs"
)

func LoadingDomain() {
	repos := persistence.NewRepositories(dbs.DB)
	userDomain := service.NewUserDomainImpl(repos.User, nil)
	user.NewServiceImpl(userDomain)
}
