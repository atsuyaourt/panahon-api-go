// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: glabs.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createGLabsLoad = `-- name: CreateGLabsLoad :one
INSERT INTO glabs_load (
  promo,
  transaction_id,
  status,
  mobile_number
) VALUES (
  $1, $2, $3, $4
) RETURNING id, status, promo, transaction_id, mobile_number, created_at, updated_at
`

type CreateGLabsLoadParams struct {
	Promo         pgtype.Text `json:"promo"`
	TransactionID pgtype.Int4 `json:"transaction_id"`
	Status        pgtype.Text `json:"status"`
	MobileNumber  string      `json:"mobile_number"`
}

func (q *Queries) CreateGLabsLoad(ctx context.Context, arg CreateGLabsLoadParams) (GlabsLoad, error) {
	row := q.db.QueryRow(ctx, createGLabsLoad,
		arg.Promo,
		arg.TransactionID,
		arg.Status,
		arg.MobileNumber,
	)
	var i GlabsLoad
	err := row.Scan(
		&i.ID,
		&i.Status,
		&i.Promo,
		&i.TransactionID,
		&i.MobileNumber,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
