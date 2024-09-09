package model_app

import uuid "github.com/jackc/pgx-gofrs-uuid"

type App struct {
	Id          string
	Name        string
	Env         string
	Description string
}

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
	Envs []*AppEnvModel
}

type CreateApp struct {
	Name        string
	Description string
}
