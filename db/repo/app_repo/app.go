package app_repo

import (
	"context"

	"github.com/MXslade/log_service_go/db"
	uuid "github.com/jackc/pgx-gofrs-uuid"
	"github.com/jackc/pgx/v5"
)

type AppModel struct {
	ID          uuid.UUID
	Name        string
	Description string
}

type AppEnvModel struct {
	ID    uuid.UUID
	Name  string
	AppID uuid.UUID
}

type AppWithEnvs struct {
	AppModel
	envs []*AppEnvModel
}

type CreateApp struct {
	Name        string
	Description string
}

type AppRepo interface {
	GetAll(ctx context.Context) ([]*AppModel, error)
	GetAllWithEnvs(ctx context.Context) ([]*AppWithEnvs, error)
	Create(ctx context.Context, data CreateApp) (*AppModel, error)
	GetByID(ctx context.Context, id uuid.UUID) (*AppModel, error)
	GetByIDWithEnvs(ctx context.Context, id uuid.UUID) (*AppWithEnvs, error)
}

type appRepo struct {
}

func New() AppRepo {
	return &appRepo{}
}

func (a *appRepo) GetAll(ctx context.Context) ([]*AppModel, error) {
	conn, err := db.AcquireConnection(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, _ := conn.Query(ctx, "SELECT id, name, description FROM apps")
	apps, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (*AppModel, error) {
		var app AppModel
		err := row.Scan(&app.ID, &app.Name, &app.Description)
		return &app, err
	})
	if err != nil {
		return nil, err
	}

	return apps, nil
}

func (a *appRepo) GetAllWithEnvs(ctx context.Context) ([]*AppWithEnvs, error) {
	conn, err := db.AcquireConnection(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, _ := conn.Query(
		ctx,
		`SELECT apps.id, apps.name, apps.description, app_envs.id, app_envs.name, app_envs.app_id
        FROM apps
        LEFT JOIN app_envs ON apps.id = app_envs.app_id
        ORDER BY apps.id
    `)
    
    pgx.ForEachRow(rows)

	return nil, nil
}

func (a *appRepo) Create(ctx context.Context, data CreateApp) (*AppModel, error) {
	return nil, nil
}

func (a *appRepo) GetByID(ctx context.Context, id uuid.UUID) (*AppModel, error) {
	return nil, nil
}

func (a *appRepo) GetByIDWithEnvs(ctx context.Context, id uuid.UUID) (*AppWithEnvs, error) {
	return nil, nil
}
