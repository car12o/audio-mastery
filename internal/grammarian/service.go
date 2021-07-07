package grammarian

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) Process(record []byte) ([]byte, error) {
	return nil, nil
}
