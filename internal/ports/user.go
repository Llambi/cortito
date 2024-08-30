package ports

import "github.com/Llambi/cortito/internal/core/domain"

type UserService interface {
	Create(user domain.CreateCortitoUserRequest) (domain.CortitoUserResponse, error)

	Upsert(user domain.CreateCortitoUserRequest) (domain.CortitoUserResponse, error)

	GetByKey(key string) (domain.CortitoUserResponse, error)
}

type UserRepository interface {
	Insert(user domain.CortitoUser) (domain.CortitoUser, error)

	Update(user domain.CortitoUser) (domain.CortitoUser, error)

	GetByKey(key string) (domain.CortitoUser, error)
}
