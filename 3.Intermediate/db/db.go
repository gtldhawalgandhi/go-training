package db

import (
	"context"
)

// DB_DRIVER=postgres
// DB_SOURCE=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable
// SERVER_ADDRESS=0.0.0.0:8080
// TOKEN_SYMMETRIC_KEY=12345678901234567890123456789012
// ACCESS_TOKEN_DURATION=15m

// Store ...
type Store interface {
	UserGetter
	UserUpdater
	UserCreator
}

type UserGetter interface {
	GetUserByEmail(ctx context.Context, email string) (UserResponse, error)
	GetUsers(ctx context.Context) ([]UserResponse, error)
}

type UserCreator interface {
	CreateUser(ctx context.Context, user UserRequest) (UserResponse, error)
}

type UserUpdater interface {
	UpdateUser(ctx context.Context, user UserRequest) (UserResponse, error)
}
