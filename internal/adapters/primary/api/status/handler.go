package status

import "github.com/Llambi/cortito/internal/ports"

type Handler struct {
	StatusService ports.StatusService
}