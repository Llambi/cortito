package user

import "github.com/Llambi/cortito/internal/core/domain"

func (s Service) GetByKey(key string) (domain.CortitoUserResponse, error) {
	cortitoUser, err := s.Repo.GetByKey(key)
	return cortitoUser.Into(), err
}
