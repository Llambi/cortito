package user

import "github.com/Llambi/cortito/internal/core/domain"

func (s Service) Create(createCortitoUserRequest domain.CreateCortitoUserRequest) (domain.CortitoUserResponse, error) {
	cortitoUser := createCortitoUserRequest.Into()

	if cortitoUser, err := s.Repo.Insert(cortitoUser); err != nil {
		return domain.CortitoUserResponse{}, err
	} else {
		return cortitoUser.Into(), nil
	}
}
