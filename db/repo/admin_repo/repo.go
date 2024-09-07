package admin_repo

import (
	"context"
	"log"

	"github.com/MXslade/log_service_go/db"
	"github.com/MXslade/log_service_go/model"
	"github.com/MXslade/log_service_go/service/auth_service"
	uuid "github.com/jackc/pgx-gofrs-uuid"
	"github.com/jackc/pgx/v5"
)

type AdminRepo struct {
	authService auth_service.AuthService
}

func New() (*AdminRepo, error) {
	authService, err := auth_service.New()
	if err != nil {
		return nil, err
	}
	return &AdminRepo{authService: authService}, nil
}

func (a *AdminRepo) GetAll(ctx context.Context) ([]*model.AdminSafeModel, error) {
	conn, err := db.AcquireConnection(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, _ := conn.Query(ctx, "SELECT id, username FROM admins")
	admins, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (*model.AdminSafeModel, error) {
		var admin model.AdminSafeModel
		err := row.Scan(&admin.ID, &admin.Username)
		return &admin, err
	})
	if err != nil {
		return nil, err
	}

	return admins, nil
}

func (a *AdminRepo) GetByID(ctx context.Context, id uuid.UUID) error {
	return nil
}

func (a *AdminRepo) GetByUsername(ctx context.Context, username string) (*model.AdminModel, error) {
	conn, err := db.AcquireConnection(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	var admin model.AdminModel
	err = conn.QueryRow(
		ctx,
		"SELECT id, username, password FROM admins WHERE username=$1",
		username,
	).Scan(&admin.ID, &admin.Username, &admin.Password)
	if err != nil {
		return nil, err
	}

	return &admin, nil
}

func (a *AdminRepo) Create(ctx context.Context, data model.CreateAdmin) (*model.AdminSafeModel, error) {
	conn, err := db.AcquireConnection(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	hashedPassword := a.authService.HashPassword(data.Password)

	var admin model.AdminSafeModel
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

func (a *AdminRepo) Delete(ctx context.Context, id uuid.UUID) error {
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
