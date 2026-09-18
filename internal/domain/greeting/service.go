package greeting

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Get() string {
	return "from Service"
}
