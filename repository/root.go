package repository

import (
	"sync"
)

//network, repository 의 가교

var (
	repositoryInit     sync.Once
	repositoryInstance *Repository
)

type Repository struct {
	User *UserRepository
}

func NewRepository() *Repository {
	repositoryInit.Do(func() {
		repositoryInstance = &Repository{
			User: NewUserRepository(),
		}
	})

	return repositoryInstance
}
