package ports

import "github.com/Llambi/cortito/internal/core/domain"

type UrlService interface {
	Create(url domain.CreateCortitoUrlRequest) (domain.CortitoUrlResponse, error)

	GetByKey(key string) (domain.CortitoUrlResponse, error)
}

type UrlRepository interface {
	Insert(url domain.CortitoUrl) (domain.CortitoUrl, error)

	GetByKey(key string) (domain.CortitoUrl, error)
}
