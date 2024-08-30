package user

import "github.com/scalalang2/golang-fifo/sieve"

type Repository struct {
	Store *sieve.Sieve[string, string]
}
