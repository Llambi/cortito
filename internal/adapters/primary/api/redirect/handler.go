package redirect

import "github.com/Llambi/cortito/internal/ports"

type Handler struct {
	RedirectService ports.RedirectService
}
