package model

import uuid "github.com/jackc/pgx-gofrs-uuid"

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
