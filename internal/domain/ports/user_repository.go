package ports

import "github.com/bperezgo/go-template/internal/domain/entities"

type WhereFindUserRepoCriteria struct {
	Email string
}

type FindUserRepoCriteria struct {
	Where WhereFindUserRepoCriteria
}

type UserRepository interface {
	Find(criteria FindUserRepoCriteria) (*entities.User, error)
}
