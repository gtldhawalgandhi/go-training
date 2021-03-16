package db

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/georgysavva/scany/pgxscan"
	l "github.com/gtldhawalgandhi/go-training/3.Intermediate/logger"
	"github.com/gtldhawalgandhi/go-training/4.Advanced/util"
)

type UpdateUserRequest struct {
	UserName  string `json:"user_name" binding:"alphanum"`
	Email     string `json:"email" binding:"email"`
	Password  string `json:"password" binding:"min=6"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type CreateUserRequest struct {
	UserName  string    `json:"user_name" binding:"required,alphanum"`
	Email     string    `json:"email" binding:"required,email"`
	Password  string    `json:"password" binding:"required,min=6"`
	FullName  string    `json:"full_name"`
	FirstName string    `json:"first_name" binding:"required"`
	LastName  string    `json:"last_name" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

// UserResponse ...
type UserResponse struct {
	UserName       string    `json:"user_name"`
	Email          string    `json:"email"`
	FullName       string    `json:"full_name"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	HashedPassword string    `json:"-"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	Updated        time.Time `json:"updated"`
}

// GetUserByUserName ...
func (pg *PGStore) GetUserByUserName(ctx context.Context, userName string) (UserResponse, error) {
	var ur UserResponse
	err := pg.db.QueryRow(context.Background(), "select user_name, first_name, last_name, created_at, pass_hash from users where user_name=$1", userName).Scan(&ur.UserName, &ur.FirstName, &ur.LastName, &ur.CreatedAt, &ur.HashedPassword)
	if err != nil {
		return UserResponse{}, err
	}

	return ur, nil
}

// GetUserByEmail ..
func (pg *PGStore) GetUserByEmail(ctx context.Context, email string) (UserResponse, error) {
	var ur UserResponse
	err := pg.db.QueryRow(context.Background(), "select user_name, first_name, last_name, created_at from users where email=$1", email).Scan(&ur.UserName, &ur.FirstName, &ur.LastName, &ur.CreatedAt)
	if err != nil {
		return UserResponse{}, err
	}

	return ur, nil
}

// CreateUser ..
func (pg *PGStore) CreateUser(ctx context.Context, user CreateUserRequest) (ur UserResponse, err error) {
	passHash, err := util.HashPassword(user.Password)

	if err != nil {
		return
	}

	query := `
	insert into users (user_name, first_name, last_name, email, pass_hash, created_at) values 
		($1,$2,$3,$4,$5,$6)
	RETURNING user_name, created_at;
	`

	args := []interface{}{user.UserName, user.FirstName, user.LastName, user.Email, passHash, time.Now()}

	err = pgxscan.Get(ctx, pg.db, &ur, query, args...)
	if err != nil {
		return
	}

	return
}

// UpdateUser ..
func (pg *PGStore) UpdateUser(ctx context.Context, user UpdateUserRequest) (UserResponse, error) {
	var ur UserResponse
	var passHash string
	var err error

	if len(strings.TrimSpace(user.Password)) > 0 {
		passHash, err = util.HashPassword(user.Password)
	}

	if err != nil {
		return UserResponse{}, err
	}

	err = pg.db.QueryRow(context.Background(), `
	update users
		set user_name = $1,
			first_name = $2,
			last_name = $3,
			email = $4,
			created_at = users.created_at,
			pass_hash =  coalesce($5, users.pass_hash),
			updated = now()
		where user_name = $1
	RETURNING user_name, updated;
	`, user.UserName, user.FirstName, user.LastName, user.Email, passHash).Scan(&ur.UserName, &ur.Updated)

	if err != nil {
		return UserResponse{}, err
	}

	return ur, nil
}

// GetUsers ...
func (pg *PGStore) GetUsers(ctx context.Context) ([]UserResponse, error) {
	rows, err := pg.db.Query(context.Background(), "select user_name, first_name, last_name, email, created_at from users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users = make([]UserResponse, 0)
	for rows.Next() {
		var ur UserResponse
		rows.Scan(&ur.UserName, &ur.FirstName, &ur.LastName, &ur.Email, &ur.CreatedAt)
		ur.FullName = ur.FirstName + " " + ur.LastName
		users = append(users, ur)
	}
	l.D(fmt.Sprintf("%+v\n", users))
	return users, nil
}
