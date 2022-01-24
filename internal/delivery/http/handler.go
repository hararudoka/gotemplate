package handler

type Handler struct {
	// here some useful variables - *sqlx.DB or some Configs
}

// New creates new handler
func New() (*Handler, error) {
	return &Handler{
		//...
	}, nil
}