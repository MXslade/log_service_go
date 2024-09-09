package app_repo

import (
	"context"
	"fmt"

	"github.com/MXslade/log_service_go/db"
	model_app "github.com/MXslade/log_service_go/model/app_model"
	uuid "github.com/jackc/pgx-gofrs-uuid"
	"github.com/jackc/pgx/v5"
)

type AppRepo struct {
}

func New() *AppRepo {
	return &AppRepo{}
}

func (a *AppRepo) GetAll(ctx context.Context) ([]*model_app.AppModel, error) {
	conn, err := db.AcquireConnection(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, _ := conn.Query(ctx, "SELECT id, name, description FROM apps")
	apps, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (*model_app.AppModel, error) {
		var app model_app.AppModel
		err := row.Scan(&app.ID, &app.Name, &app.Description)
		return &app, err
	})
	if err != nil {
		return nil, err
	}

	return apps, nil
}

func (a *AppRepo) GetAllWithEnvs(ctx context.Context) ([]*model_app.AppWithEnvs, error) {
	conn, err := db.AcquireConnection(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, _ := conn.Query(
		ctx,
		`SELECT apps.id, apps.name, apps.description, app_envs.id, app_envs.name
        FROM apps
        LEFT JOIN app_envs ON apps.id = app_envs.app_id
        ORDER BY apps.id
    `)

	result := make([]*model_app.AppWithEnvs, 0)
	scans := make([]any, 5)
	_, err = pgx.ForEachRow(rows, scans, func() error {
		appID, ok := scans[0].(uuid.UUID)
		if !ok {
			return fmt.Errorf("cannot convert: %v to appID uuid", scans[0])
		}
		appName, ok := scans[1].(string)
		if !ok {
			return fmt.Errorf("cannot convert: %v to appName string", scans[1])
		}
		appDescription, ok := scans[2].(string)
		if !ok {
			return fmt.Errorf("cannot convert: %v to appDescription string", scans[2])
		}
		envID, ok := scans[3].(uuid.UUID)
		if !ok {
			return fmt.Errorf("cannot convert: %v to envID uuid", scans[3])
		}
		envName, ok := scans[4].(string)
		if !ok {
			return fmt.Errorf("cannot convert: %v to envName string", scans[4])
		}
		if len(result) > 0 && appID == result[len(result)-1].ID {
			result[len(result)-1].Envs = append(result[len(result)-1].Envs, &model_app.AppEnvModel{ID: envID, Name: envName, AppID: appID})
		} else {
			result = append(result, &model_app.AppWithEnvs{
				AppModel: model_app.AppModel{
					ID:          appID,
					Name:        appName,
					Description: appDescription,
				},
				Envs: []*model_app.AppEnvModel{
					{
						ID:    envID,
						Name:  envName,
						AppID: appID,
					},
				},
			})
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *AppRepo) Create(ctx context.Context, data model_app.CreateApp) (*model_app.AppModel, error) {
	return nil, nil
}

func (a *AppRepo) GetByID(ctx context.Context, id uuid.UUID) (*model_app.AppModel, error) {
	return nil, nil
}

func (a *AppRepo) GetByIDWithEnvs(ctx context.Context, id uuid.UUID) (*model_app.AppWithEnvs, error) {
	return nil, nil
}
