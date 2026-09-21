package user

import "context"

type Repository interface {
	GetByID(ctx context.Context, id string) (User, error)
	Create(ctx context.Context, u User) (User, error)
	Update(ctx context.Context, u User) (User, error)
	Delete(ctx context.Context, id string) error
}
