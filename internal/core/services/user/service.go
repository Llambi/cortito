package user

import "github.com/Llambi/cortito/internal/ports"

type Service struct {
	Repo ports.UserRepository
}
