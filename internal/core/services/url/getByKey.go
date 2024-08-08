package url

import (
	"github.com/Llambi/cortito/internal/core/domain"
)

func (s Service) GetByKey(key string) (domain.CortitoUrl, error) {
	return s.Repo.GetByKey(key)
}
