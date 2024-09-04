package user

import (
	"github.com/Llambi/cortito/internal/ports"
	"github.com/Llambi/cortito/platform/config"
)

type Service struct {
	Repo   ports.UserRepository
	Config *config.Config
}
