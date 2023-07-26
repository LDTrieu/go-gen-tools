package repositories

import (
	"context"
	"go-gen-tools/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
}
