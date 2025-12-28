package model

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type User struct {
	Id        int64     `db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	Email     string    `db:"email"`
	Phone     string    `db:"phone"`
	Avatar    string    `db:"avatar"`
	Status    int64     `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type UserModel interface {
	Insert(ctx context.Context, data *User) (sql.Result, error)
	FindOne(ctx context.Context, id int64) (*User, error)
	FindOneByUsername(ctx context.Context, username string) (*User, error)
	Update(ctx context.Context, data *User) error
	Delete(ctx context.Context, id int64) error
	FindList(ctx context.Context, page, pageSize int64) ([]*User, error)
	Count(ctx context.Context) (int64, error)
}

type defaultUserModel struct {
	conn  sqlx.SqlConn
	table string
}

func NewUserModel(conn sqlx.SqlConn) UserModel {
	return &defaultUserModel{
		conn:  conn,
		table: "users",
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, data *User) (sql.Result, error) {
	query := fmt.Sprintf("INSERT INTO %s (username, password, email, phone, avatar, status) VALUES (?, ?, ?, ?, ?, ?)", m.table)
	return m.conn.ExecCtx(ctx, query, data.Username, data.Password, data.Email, data.Phone, data.Avatar, data.Status)
}

func (m *defaultUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	query := fmt.Sprintf("SELECT id, username, password, email, phone, avatar, status, created_at, updated_at FROM %s WHERE id = ? LIMIT 1", m.table)
	var user User
	err := m.conn.QueryRowCtx(ctx, &user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *defaultUserModel) FindOneByUsername(ctx context.Context, username string) (*User, error) {
	query := fmt.Sprintf("SELECT id, username, password, email, phone, avatar, status, created_at, updated_at FROM %s WHERE username = ? LIMIT 1", m.table)
	var user User
	err := m.conn.QueryRowCtx(ctx, &user, query, username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *defaultUserModel) Update(ctx context.Context, data *User) error {
	query := fmt.Sprintf("UPDATE %s SET email = ?, phone = ?, avatar = ? WHERE id = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, data.Email, data.Phone, data.Avatar, data.Id)
	return err
}

func (m *defaultUserModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUserModel) FindList(ctx context.Context, page, pageSize int64) ([]*User, error) {
	offset := (page - 1) * pageSize
	query := fmt.Sprintf("SELECT id, username, password, email, phone, avatar, status, created_at, updated_at FROM %s ORDER BY id DESC LIMIT ? OFFSET ?", m.table)
	var users []*User
	err := m.conn.QueryRowsCtx(ctx, &users, query, pageSize, offset)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (m *defaultUserModel) Count(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", m.table)
	var count int64
	err := m.conn.QueryRowCtx(ctx, &count, query)
	if err != nil {
		return 0, err
	}
	return count, nil
}
