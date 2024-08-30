package user

import (
	"errors"
	"time"

	"github.com/Llambi/cortito/internal/core/domain"
)

func (s Service) Upsert(CreateCortitoUserRequest domain.CreateCortitoUserRequest) (domain.CortitoUserResponse, error) {
	cortitoUser := CreateCortitoUserRequest.Into()

	if foundCortitoUser, err := s.Repo.GetByKey(cortitoUser.Ip); err == nil {
		// Quota is over
		if foundCortitoUser.Quota == 0 {
			// And RefreshAt is over too
			if foundCortitoUser.RefreshAt.Before(time.Now()) {
				// Refresh all
				foundCortitoUser.Quota = 30
				foundCortitoUser.RefreshAt = time.Now().Add(30 * time.Minute)
			} else { // RefreshAt is not over
				// The user must wait
				return domain.CortitoUserResponse{}, errors.New("the quota has been exhausted")
			}
		} else { // Quota is not over
			foundCortitoUser.Quota -= 1
		}
		cortitoUser, err = s.Repo.Update(foundCortitoUser)
		if err != nil {
			return domain.CortitoUserResponse{}, errors.New("can not update the user")
		}
	} else {
		cortitoUser.Quota = 30
		cortitoUser.RefreshAt = time.Now().Add(30 * time.Minute)
		cortitoUser, err = s.Repo.Insert(cortitoUser)
		if err != nil {
			return domain.CortitoUserResponse{}, errors.New("can not create the user")
		}
	}
	return cortitoUser.Into(), nil
}
