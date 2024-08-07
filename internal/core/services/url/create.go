package url

import "github.com/Llambi/cortito/internal/core/domain"

func (s Service) Create(url domain.CortitoUrl) (domain.CortitoUrl, error) {
	url.Shorten()
	return s.Repo.Insert(url)
}
