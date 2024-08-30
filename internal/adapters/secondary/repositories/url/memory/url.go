package url

import (
	"errors"

	"github.com/Llambi/cortito/internal/core/domain"
)

func (repo MemoryRepository) Insert(url domain.CortitoUrl) (domain.CortitoUrl, error) {
	repo.Store.Set(url.ShortUrl, url.Url)
	return url, nil
}

func (repo MemoryRepository) GetByKey(key string) (domain.CortitoUrl, error) {
	if url, exist := repo.Store.Get(key); exist {
		return domain.CortitoUrl{Url: url, ShortUrl: key}, nil
	}
	return domain.CortitoUrl{}, errors.New("url key not exists") //TODO: Add Error
}
