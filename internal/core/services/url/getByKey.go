package url

import (
	"github.com/Llambi/cortito/internal/core/domain"
)

func (s Service) GetByKey(key string) (domain.CortitoUrlResponse, error) {
	cortitoUrl, err := s.Repo.GetByKey(key)
	return cortitoUrl.IntoCortitoUrlResponse(), err
}
