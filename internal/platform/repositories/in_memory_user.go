package repositories

import (
	"github.com/bperezgo/go-template/internal/domain/entities"
	"github.com/bperezgo/go-template/internal/domain/ports"
)

type InMemoryUserRepository struct {
	users []entities.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: []entities.User{
			// Change for encrypted password
			entities.NewUser("1", "example1@gmail.com", "Password123"),
			entities.NewUser("2", "example2@gmail.com", "Password123456"),
		},
	}
}

func (r *InMemoryUserRepository) Find(criteria ports.FindUserRepoCriteria) (*entities.User, error) {
	for _, u := range r.users {
		userDto := u.ToDto()
		if userDto.Email == criteria.Where.Email {
			return &u, nil
		}
	}

	return nil, nil
}
