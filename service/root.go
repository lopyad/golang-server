package service

import (
	"golang-server/repository"
	"sync"
)

//network, repository 의 가교

var (
	serviceInit     sync.Once
	serviceInstance *Service
)

type Service struct {
	repository *repository.Repository

	User *User
}

func NewService(repo *repository.Repository) *Service {
	serviceInit.Do(func() {
		serviceInstance = &Service{
			repository: repo,
		}

		serviceInstance.User = newUserService(repo.User)
	})

	return serviceInstance
}
