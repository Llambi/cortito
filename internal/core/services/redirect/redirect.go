package redirect

import "github.com/Llambi/cortito/internal/core/domain"

func (s Service) Redirect(key string) (domain.CortitoUrlResponse, error) {
	cortitoUrl, err := s.Repo.GetByKey(key)
	return cortitoUrl.IntoCortitoUrlResponse(), err
}
