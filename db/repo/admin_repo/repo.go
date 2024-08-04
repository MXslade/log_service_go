package admin_repo

import (
	"context"
	"log"

	"github.com/MXslade/log_service_go/db"
	"github.com/MXslade/log_service_go/service/auth_service"
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
	Create(ctx context.Context, data CreateAdmin) (*AdminSafeModel, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type adminRepo struct {
	authService auth_service.AuthService
}

func New() (*adminRepo, error) {
	authService, err := auth_service.New()
	if err != nil {
		return nil, err
	}
	return &adminRepo{authService: authService}, nil
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
		log.Printf("err: %v\n", err)
		return nil, err
	}

	return admins, nil
}

func (a *adminRepo) GetById(ctx context.Context, id uuid.UUID) error {
	return nil
}

func (a *adminRepo) Create(ctx context.Context, data CreateAdmin) (*AdminSafeModel, error) {
	conn, err := db.AcquireConnection(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	hashedPassword := a.authService.HashPassword(data.Password)

	var admin AdminSafeModel
	err = conn.QueryRow(
		ctx,
		"INSERT INTO admins (username, password) VALUES($1, $2) RETURNING id, username;",
		data.Username, hashedPassword,
	).Scan(&admin.ID, &admin.Username)

	if err != nil {
		log.Printf("Admin Repo Create Error: %v\n", err)
		return nil, err
	}

	return &admin, nil
}

func (a *adminRepo) Delete(ctx context.Context, id uuid.UUID) error {
	conn, err := db.AcquireConnection(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, "DELETE FROM admins WHERE id=$1", id)

	if err != nil {
		return err
	}

	return nil
}
