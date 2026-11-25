package store

import (
	"context"
	"database/sql"
	"time"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{db: db}
}

type User struct {
	ID           int64
	Email        string
	Username     string
	DOB          time.Time
	PasswordHash string
	CreatedAt    time.Time
}

func (u *UserStore) CreateUser(ctx context.Context, user User) (User, error) {
	row := u.db.QueryRowContext(ctx, `Insert into users(email,username,dob,password_hash)
	values($1,$2,$3,$4) returning id,created_at`, user.Email, user.Username, user.DOB, user.PasswordHash)
	err := row.Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (u *UserStore) GetUser(ctx context.Context, id int64) (User, error) {
	row := u.db.QueryRowContext(ctx, "Select * from Users where id = $1", id)
	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.DOB, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

// func(u *UserStore)GetUserByEmail(ctx context.Context,email string)(User,error){

// }
