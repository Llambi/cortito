package ports

import "github.com/Llambi/cortito/internal/core/domain"

type RedirectService interface {
	Redirect(key string) (domain.CortitoUrlResponse, error)
}
