package admin_repo

import (
	"context"
	"log"

	"github.com/MXslade/log_service_go/db"
	uuid "github.com/jackc/pgx-gofrs-uuid"
	"github.com/jackc/pgx/v5"
)

type AdminModel struct {
	ID   uuid.UUID
	Name string
}

type CreateAdmin struct {
	Name     string
	Password string
}

type AdminRepo interface {
	GetAll(ctx context.Context) ([]*AdminModel, error)
	GetById(ctx context.Context, id uuid.UUID) error
	Create(ctx context.Context, data CreateAdmin) error
}

type adminRepo struct {
}

func New() *adminRepo {
	return &adminRepo{}
}

func (a *adminRepo) GetAll(ctx context.Context) ([]*AdminModel, error) {
	conn, err := db.AcquireConnection(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, _ := conn.Query(ctx, "SELECT id, name FROM admins")
	admins, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (*AdminModel, error) {
		var admin AdminModel
		err := row.Scan(&admin.ID, &admin.Name)
		return &admin, err
	})
	if err != nil {
		log.Printf("err: %v", err)
		return nil, err
	}

	return admins, nil
}

func (a *adminRepo) GetById(ctx context.Context, id uuid.UUID) error {
	return nil
}

func (a *adminRepo) Create(ctx context.Context, data CreateAdmin) error {
	return nil
}
