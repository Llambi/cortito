package ports

type StatusService interface {
	Get() (string, error)
}
