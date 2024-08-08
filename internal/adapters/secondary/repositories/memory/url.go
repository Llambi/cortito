package url

import "github.com/Llambi/cortito/internal/core/domain"

func (repo Repository) Insert(url domain.CortitoUrl) (domain.CortitoUrl, error) {
	repo.Store[url.ShortUrl] = url.Url
	return url, nil
}

func (repo Repository) GetByKey(key string) (domain.CortitoUrl, error) {
	if url, exist := repo.Store[key]; exist {
		return domain.CortitoUrl{Url: url, ShortUrl: key}, nil
	}
	return domain.CortitoUrl{}, nil //TODO: Add Error
}
