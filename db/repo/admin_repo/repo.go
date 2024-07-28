package admin_repo

import (
	"context"

	uuid "github.com/jackc/pgx-gofrs-uuid"
)

type CreateAdmin struct {
	Username string
	Password string
}

type AdminRepo interface {
	GetAll(ctx context.Context)
	GetById(ctx context.Context, id uuid.UUID)
	Create(ctx context.Context, data CreateAdmin)
}

type adminRepo struct {
}

func New() *adminRepo {
	return &adminRepo{}
}

func (a *adminRepo) GetAll(ctx context.Context) {
}

func (a *adminRepo) GetById(ctx context.Context, id uuid.UUID) {
}

func (a *adminRepo) Create(ctx context.Context, data CreateAdmin) {
}
