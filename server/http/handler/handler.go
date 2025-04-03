package handler

type Service interface {
	Draw(list []string) (string, error)
}

type Handler struct {
	service Service
}

func New(service Service) (*Handler, error) {
	return &Handler{
		service: service,
	}, nil
}
