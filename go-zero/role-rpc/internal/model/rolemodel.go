package model

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type Role struct {
	Id          int64     `db:"id"`
	Name        string    `db:"name"`
	Code        string    `db:"code"`
	Description string    `db:"description"`
	Status      int64     `db:"status"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type UserRole struct {
	Id        int64     `db:"id"`
	UserId    int64     `db:"user_id"`
	RoleId    int64     `db:"role_id"`
	CreatedAt time.Time `db:"created_at"`
}

type RoleModel interface {
	Insert(ctx context.Context, data *Role) (sql.Result, error)
	FindOne(ctx context.Context, id int64) (*Role, error)
	FindOneByCode(ctx context.Context, code string) (*Role, error)
	Update(ctx context.Context, data *Role) error
	Delete(ctx context.Context, id int64) error
	FindList(ctx context.Context, page, pageSize int64) ([]*Role, error)
	Count(ctx context.Context) (int64, error)
	AssignRoles(ctx context.Context, userId int64, roleIds []int64) error
	GetUserRoles(ctx context.Context, userId int64) ([]*Role, error)
}

type defaultRoleModel struct {
	conn      sqlx.SqlConn
	table     string
	userRoles string
}

func NewRoleModel(conn sqlx.SqlConn) RoleModel {
	return &defaultRoleModel{
		conn:      conn,
		table:     "roles",
		userRoles: "user_roles",
	}
}

func (m *defaultRoleModel) Insert(ctx context.Context, data *Role) (sql.Result, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, code, description, status) VALUES (?, ?, ?, ?)", m.table)
	return m.conn.ExecCtx(ctx, query, data.Name, data.Code, data.Description, data.Status)
}

func (m *defaultRoleModel) FindOne(ctx context.Context, id int64) (*Role, error) {
	query := fmt.Sprintf("SELECT id, name, code, description, status, created_at, updated_at FROM %s WHERE id = ? LIMIT 1", m.table)
	var role Role
	err := m.conn.QueryRowCtx(ctx, &role, query, id)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (m *defaultRoleModel) FindOneByCode(ctx context.Context, code string) (*Role, error) {
	query := fmt.Sprintf("SELECT id, name, code, description, status, created_at, updated_at FROM %s WHERE code = ? LIMIT 1", m.table)
	var role Role
	err := m.conn.QueryRowCtx(ctx, &role, query, code)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (m *defaultRoleModel) Update(ctx context.Context, data *Role) error {
	query := fmt.Sprintf("UPDATE %s SET name = ?, description = ?, status = ? WHERE id = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.Description, data.Status, data.Id)
	return err
}

func (m *defaultRoleModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultRoleModel) FindList(ctx context.Context, page, pageSize int64) ([]*Role, error) {
	offset := (page - 1) * pageSize
	query := fmt.Sprintf("SELECT id, name, code, description, status, created_at, updated_at FROM %s ORDER BY id LIMIT ? OFFSET ?", m.table)
	var roles []*Role
	err := m.conn.QueryRowsCtx(ctx, &roles, query, pageSize, offset)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (m *defaultRoleModel) Count(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", m.table)
	var count int64
	err := m.conn.QueryRowCtx(ctx, &count, query)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (m *defaultRoleModel) AssignRoles(ctx context.Context, userId int64, roleIds []int64) error {
	// 先删除用户的所有角色
	deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE user_id = ?", m.userRoles)
	_, err := m.conn.ExecCtx(ctx, deleteQuery, userId)
	if err != nil {
		return err
	}

	// 插入新的角色关联
	for _, roleId := range roleIds {
		insertQuery := fmt.Sprintf("INSERT INTO %s (user_id, role_id) VALUES (?, ?)", m.userRoles)
		_, err = m.conn.ExecCtx(ctx, insertQuery, userId, roleId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *defaultRoleModel) GetUserRoles(ctx context.Context, userId int64) ([]*Role, error) {
	query := fmt.Sprintf(`
		SELECT r.id, r.name, r.code, r.description, r.status, r.created_at, r.updated_at
		FROM %s r
		INNER JOIN %s ur ON r.id = ur.role_id
		WHERE ur.user_id = ?
	`, m.table, m.userRoles)
	var roles []*Role
	err := m.conn.QueryRowsCtx(ctx, &roles, query, userId)
	if err != nil {
		return nil, err
	}
	return roles, nil
}
