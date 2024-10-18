package repository

import "golang-server/types"
import "golang-server/types/errors"

type UserRepository struct {
	userMap []*types.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		userMap: []*types.User{},
	}
}

func (u *UserRepository) Create(newUser *types.User) error {
	u.userMap = append(u.userMap, newUser)
	return nil
}

func (u *UserRepository) Update(name string, newAge int64) error {
	var isExisted bool = false
	for _, user := range u.userMap {
		if name == user.Name {
			user.Age = newAge
			isExisted = true
			continue
		}
	}

	if !isExisted {
		return errors.Errorf(errors.NotFoundUser, nil)
	} else {
		return nil
	}
}

func (u *UserRepository) Delete(userName string) error {
	var isExisted bool = false
	for index, user := range u.userMap {
		if userName == user.Name {
			u.userMap = append(u.userMap[:index], u.userMap[index+1:]...)
			isExisted = true
			continue
		}
	}

	if !isExisted {
		return errors.Errorf(errors.NotFoundUser, nil)
	} else {
		return nil
	}

}

func (u *UserRepository) Get() []*types.User {
	return u.userMap
}
