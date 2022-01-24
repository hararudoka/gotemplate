package handler

type Handler struct {
	HelloHandler HelloHandler
}

// New creates new handler
func New() (*Handler, error) {
	return &Handler{
		//...
	}, nil
}