// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: shopkeeper.sql

package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const createNewShopKeeperUser = `-- name: CreateNewShopKeeperUser :one
WITH user_sub AS ( INSERT INTO common_user (full_name, cpf, email) values ($1, $2, $3) RETURNING id, full_name, cpf, email, created_at, deleted_at),
     shopkeeper_sub
         AS ( INSERT INTO shoopkeepers (common_user_id, cnpj) values ((SELECT id FROM user_sub), $4) RETURNING common_user_id, cnpj)
SELECT user_sub.id, user_sub.full_name, user_sub.cpf, user_sub.email, user_sub.created_at, user_sub.deleted_at, shopkeeper_sub.cnpj
from user_sub,
     shopkeeper_sub
`

type CreateNewShopKeeperUserParams struct {
	FullName string `json:"full_name"`
	Cpf      string `json:"cpf"`
	Email    string `json:"email"`
	Cnpj     string `json:"cnpj"`
}

type CreateNewShopKeeperUserRow struct {
	ID        int64              `json:"id"`
	FullName  string             `json:"full_name"`
	Cpf       string             `json:"cpf"`
	Email     string             `json:"email"`
	CreatedAt time.Time          `json:"created_at"`
	DeletedAt pgtype.Timestamptz `json:"deleted_at"`
	Cnpj      string             `json:"cnpj"`
}

func (q *Queries) CreateNewShopKeeperUser(ctx context.Context, arg CreateNewShopKeeperUserParams) (CreateNewShopKeeperUserRow, error) {
	row := q.db.QueryRow(ctx, createNewShopKeeperUser,
		arg.FullName,
		arg.Cpf,
		arg.Email,
		arg.Cnpj,
	)
	var i CreateNewShopKeeperUserRow
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Cpf,
		&i.Email,
		&i.CreatedAt,
		&i.DeletedAt,
		&i.Cnpj,
	)
	return i, err
}

const getShopKeeperUser = `-- name: GetShopKeeperUser :one
SELECT cu.id, cu.full_name, cu.cpf, cu.email, cu.created_at, cu.deleted_at, su.cnpj
FROM shoopkeepers su
         JOIN public.common_user cu on cu.id = su.common_user_id
WHERE common_user_id = $1
LIMIT 1
`

type GetShopKeeperUserRow struct {
	ID        int64              `json:"id"`
	FullName  string             `json:"full_name"`
	Cpf       string             `json:"cpf"`
	Email     string             `json:"email"`
	CreatedAt time.Time          `json:"created_at"`
	DeletedAt pgtype.Timestamptz `json:"deleted_at"`
	Cnpj      string             `json:"cnpj"`
}

func (q *Queries) GetShopKeeperUser(ctx context.Context, commonUserID int64) (GetShopKeeperUserRow, error) {
	row := q.db.QueryRow(ctx, getShopKeeperUser, commonUserID)
	var i GetShopKeeperUserRow
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Cpf,
		&i.Email,
		&i.CreatedAt,
		&i.DeletedAt,
		&i.Cnpj,
	)
	return i, err
}
