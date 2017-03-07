package service

type WriteService interface {
	Write() error
}

type WriteServiceImpl struct {
}

type writeRequest struct {
}

type writeResponse struct {
}

func (ws WriteServiceImpl) Write() error {
	return nil
}
