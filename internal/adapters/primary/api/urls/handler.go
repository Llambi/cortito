package url

import "github.com/Llambi/cortito/internal/ports"

type Handler struct {
	UrlService ports.UrlService

	UserService ports.UserService
}
