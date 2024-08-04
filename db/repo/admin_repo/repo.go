package admin_repo

import (
	"context"
	"log"

	"github.com/MXslade/log_service_go/db"
	uuid "github.com/jackc/pgx-gofrs-uuid"
	"github.com/jackc/pgx/v5"
)

type AdminModel struct {
	ID       uuid.UUID
	Username string
	Password string
}

type AdminSafeModel struct {
	ID       uuid.UUID
	Username string
}

type CreateAdmin struct {
	Username string
	Password string
}

type AdminRepo interface {
	GetAll(ctx context.Context) ([]*AdminSafeModel, error)
	GetById(ctx context.Context, id uuid.UUID) error
	Create(ctx context.Context, data CreateAdmin) (*AdminModel, error)
}

type adminRepo struct {
}

func New() *adminRepo {
	return &adminRepo{}
}

func (a *adminRepo) GetAll(ctx context.Context) ([]*AdminSafeModel, error) {
	conn, err := db.AcquireConnection(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, _ := conn.Query(ctx, "SELECT id, username FROM admins")
	admins, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (*AdminSafeModel, error) {
		var admin AdminSafeModel
		err := row.Scan(&admin.ID, &admin.Username)
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

func (a *adminRepo) Create(ctx context.Context, data CreateAdmin) (*AdminModel, error) {
	conn, err := db.AcquireConnection(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	return nil, nil
}
