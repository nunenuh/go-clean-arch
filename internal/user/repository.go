package user

import (
	"context"
	"database/sql"
)

const (
	QUERY_GET_USERS   = "SELECT * FROM users"
	QUERY_GET_USER    = "SELECT * FROM users WHERE id = ?"
	QUERY_CREATE_USER = "INSERT INTO users (name, username, password, email, phone, created, modified, deleted) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	QUERY_UPDATE_USER = "UPDATE users SET name = ?, username = ?, password = ?, email = ?, phone = ?, modified = ? WHERE id = ?"
	QUERY_DELETE_USER = "DELETE FROM users WHERE id = ?"
)

type mariadbRepository struct {
	mariadb *sql.DB
}

func NewUserRepository(mariadbConnection *sql.DB) UserRepository {
	return &mariadbRepository{mariadb: mariadbConnection}
}

func (r *mariadbRepository) GetUsers(ctx context.Context) (*[]User, error) {
	var users []User

	res, err := r.mariadb.QueryContext(ctx, QUERY_GET_USERS)
	if err != nil {
		return nil, err
	}

	defer res.Close()

	for res.Next() {
		user := &User{}
		err := res.Scan(&user.ID, &user.Name, &user.Username, &user.Password, &user.Email, &user.Phone, &user.Created, &user.Modified, &user.Deleted)
		if err != nil && err == sql.ErrNoRows {
			return nil, nil
		}
		if err != nil {
			return nil, err
		}

		users = append(users, *user)
	}

	return &users, nil
}

func (r *mariadbRepository) GetUser(ctx context.Context, userID int) (*User, error) {
	user := &User{}

	stmt, err := r.mariadb.PrepareContext(ctx, QUERY_GET_USER)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, userID).Scan(&user.ID, &user.Name, &user.Username, &user.Password, &user.Email, &user.Phone, &user.Created, &user.Modified, &user.Deleted)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *mariadbRepository) CreateUser(ctx context.Context, user *User) error {
	stmt, err := r.mariadb.PrepareContext(ctx, QUERY_CREATE_USER)
	if err != nil {
		return err
	}
	defer stmt.Close()

	//Insert one User
	_, err = stmt.ExecContext(ctx, user.Name, user.Username, user.Password, user.Email, user.Phone, user.Created, user.Modified, user.Deleted)
	if err != nil {
		return err
	}

	return nil
}

func (r *mariadbRepository) UpdateUser(ctx context.Context, userID int, user *User) error {
	stmt, err := r.mariadb.PrepareContext(ctx, QUERY_UPDATE_USER)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, user.Name, user.Username, user.Password, user.Email, user.Phone, user.Modified, userID)
	if err != nil {
		return err
	}

	return nil
}

func (r *mariadbRepository) DeleteUser(ctx context.Context, userID int) error {
	stmt, err := r.mariadb.PrepareContext(ctx, QUERY_DELETE_USER)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}
