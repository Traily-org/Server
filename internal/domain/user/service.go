package user

import "context"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Get(ctx context.Context, id string) (User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) Create(ctx context.Context, u User) (User, error) {
	return s.repo.Create(ctx, u)
}

func (s *Service) Update(ctx context.Context, u User) (User, error) {
	return s.repo.Update(ctx, u)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
