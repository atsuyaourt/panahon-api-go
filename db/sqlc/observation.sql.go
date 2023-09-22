// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: observation.sql

package db

import (
	"context"

	"github.com/emiliogozo/panahon-api-go/util"
	"github.com/jackc/pgx/v5/pgtype"
)

const countObservations = `-- name: CountObservations :one
SELECT count(*) FROM observations_observation
WHERE station_id = ANY($1::bigint[])
  AND (CASE WHEN $2::bool THEN timestamp >= $3 ELSE TRUE END)
  AND (CASE WHEN $4::bool THEN timestamp <= $5 ELSE TRUE END)
`

type CountObservationsParams struct {
	StationIds  []int64            `json:"station_ids"`
	IsStartDate bool               `json:"is_start_date"`
	StartDate   pgtype.Timestamptz `json:"start_date"`
	IsEndDate   bool               `json:"is_end_date"`
	EndDate     pgtype.Timestamptz `json:"end_date"`
}

func (q *Queries) CountObservations(ctx context.Context, arg CountObservationsParams) (int64, error) {
	row := q.db.QueryRow(ctx, countObservations,
		arg.StationIds,
		arg.IsStartDate,
		arg.StartDate,
		arg.IsEndDate,
		arg.EndDate,
	)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countStationObservations = `-- name: CountStationObservations :one
SELECT count(*) FROM observations_observation
WHERE station_id = $1
  AND (CASE WHEN $2::bool THEN timestamp >= $3 ELSE TRUE END)
  AND (CASE WHEN $4::bool THEN timestamp <= $5 ELSE TRUE END)
`

type CountStationObservationsParams struct {
	StationID   int64              `json:"station_id"`
	IsStartDate bool               `json:"is_start_date"`
	StartDate   pgtype.Timestamptz `json:"start_date"`
	IsEndDate   bool               `json:"is_end_date"`
	EndDate     pgtype.Timestamptz `json:"end_date"`
}

func (q *Queries) CountStationObservations(ctx context.Context, arg CountStationObservationsParams) (int64, error) {
	row := q.db.QueryRow(ctx, countStationObservations,
		arg.StationID,
		arg.IsStartDate,
		arg.StartDate,
		arg.IsEndDate,
		arg.EndDate,
	)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createStationObservation = `-- name: CreateStationObservation :one
INSERT INTO observations_observation (
  pres,
  rr,
  rh,
  temp,
  td,
  wdir,
  wspd,
  wspdx,
  srad,
  mslp,
  hi,
  wchill,
  timestamp,
  qc_level,
  station_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
) RETURNING id, pres, rr, rh, temp, td, wdir, wspd, wspdx, srad, mslp, hi, station_id, timestamp, wchill, qc_level, created_at, updated_at
`

type CreateStationObservationParams struct {
	Pres      util.NullFloat4    `json:"pres"`
	Rr        util.NullFloat4    `json:"rr"`
	Rh        util.NullFloat4    `json:"rh"`
	Temp      util.NullFloat4    `json:"temp"`
	Td        util.NullFloat4    `json:"td"`
	Wdir      util.NullFloat4    `json:"wdir"`
	Wspd      util.NullFloat4    `json:"wspd"`
	Wspdx     util.NullFloat4    `json:"wspdx"`
	Srad      util.NullFloat4    `json:"srad"`
	Mslp      util.NullFloat4    `json:"mslp"`
	Hi        util.NullFloat4    `json:"hi"`
	Wchill    util.NullFloat4    `json:"wchill"`
	Timestamp pgtype.Timestamptz `json:"timestamp"`
	QcLevel   int32              `json:"qc_level"`
	StationID int64              `json:"station_id"`
}

func (q *Queries) CreateStationObservation(ctx context.Context, arg CreateStationObservationParams) (ObservationsObservation, error) {
	row := q.db.QueryRow(ctx, createStationObservation,
		arg.Pres,
		arg.Rr,
		arg.Rh,
		arg.Temp,
		arg.Td,
		arg.Wdir,
		arg.Wspd,
		arg.Wspdx,
		arg.Srad,
		arg.Mslp,
		arg.Hi,
		arg.Wchill,
		arg.Timestamp,
		arg.QcLevel,
		arg.StationID,
	)
	var i ObservationsObservation
	err := row.Scan(
		&i.ID,
		&i.Pres,
		&i.Rr,
		&i.Rh,
		&i.Temp,
		&i.Td,
		&i.Wdir,
		&i.Wspd,
		&i.Wspdx,
		&i.Srad,
		&i.Mslp,
		&i.Hi,
		&i.StationID,
		&i.Timestamp,
		&i.Wchill,
		&i.QcLevel,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteStationObservation = `-- name: DeleteStationObservation :exec
DELETE FROM observations_observation WHERE station_id = $1 AND id = $2
`

type DeleteStationObservationParams struct {
	StationID int64 `json:"station_id"`
	ID        int64 `json:"id"`
}

func (q *Queries) DeleteStationObservation(ctx context.Context, arg DeleteStationObservationParams) error {
	_, err := q.db.Exec(ctx, deleteStationObservation, arg.StationID, arg.ID)
	return err
}

const getStationObservation = `-- name: GetStationObservation :one
SELECT id, pres, rr, rh, temp, td, wdir, wspd, wspdx, srad, mslp, hi, station_id, timestamp, wchill, qc_level, created_at, updated_at FROM observations_observation
WHERE station_id = $1 AND id = $2 LIMIT 1
`

type GetStationObservationParams struct {
	StationID int64 `json:"station_id"`
	ID        int64 `json:"id"`
}

func (q *Queries) GetStationObservation(ctx context.Context, arg GetStationObservationParams) (ObservationsObservation, error) {
	row := q.db.QueryRow(ctx, getStationObservation, arg.StationID, arg.ID)
	var i ObservationsObservation
	err := row.Scan(
		&i.ID,
		&i.Pres,
		&i.Rr,
		&i.Rh,
		&i.Temp,
		&i.Td,
		&i.Wdir,
		&i.Wspd,
		&i.Wspdx,
		&i.Srad,
		&i.Mslp,
		&i.Hi,
		&i.StationID,
		&i.Timestamp,
		&i.Wchill,
		&i.QcLevel,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listObservations = `-- name: ListObservations :many
SELECT id, pres, rr, rh, temp, td, wdir, wspd, wspdx, srad, mslp, hi, station_id, timestamp, wchill, qc_level, created_at, updated_at FROM observations_observation
WHERE station_id = ANY($1::bigint[])
  AND (CASE WHEN $2::bool THEN timestamp >= $3 ELSE TRUE END)
  AND (CASE WHEN $4::bool THEN timestamp <= $5 ELSE TRUE END)
ORDER BY timestamp DESC
LIMIT $7
OFFSET $6
`

type ListObservationsParams struct {
	StationIds  []int64            `json:"station_ids"`
	IsStartDate bool               `json:"is_start_date"`
	StartDate   pgtype.Timestamptz `json:"start_date"`
	IsEndDate   bool               `json:"is_end_date"`
	EndDate     pgtype.Timestamptz `json:"end_date"`
	Offset      int32              `json:"offset"`
	Limit       util.NullInt4      `json:"limit"`
}

func (q *Queries) ListObservations(ctx context.Context, arg ListObservationsParams) ([]ObservationsObservation, error) {
	rows, err := q.db.Query(ctx, listObservations,
		arg.StationIds,
		arg.IsStartDate,
		arg.StartDate,
		arg.IsEndDate,
		arg.EndDate,
		arg.Offset,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ObservationsObservation{}
	for rows.Next() {
		var i ObservationsObservation
		if err := rows.Scan(
			&i.ID,
			&i.Pres,
			&i.Rr,
			&i.Rh,
			&i.Temp,
			&i.Td,
			&i.Wdir,
			&i.Wspd,
			&i.Wspdx,
			&i.Srad,
			&i.Mslp,
			&i.Hi,
			&i.StationID,
			&i.Timestamp,
			&i.Wchill,
			&i.QcLevel,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listStationObservations = `-- name: ListStationObservations :many
SELECT id, pres, rr, rh, temp, td, wdir, wspd, wspdx, srad, mslp, hi, station_id, timestamp, wchill, qc_level, created_at, updated_at FROM observations_observation
WHERE station_id = $1
  AND (CASE WHEN $2::bool THEN timestamp >= $3 ELSE TRUE END)
  AND (CASE WHEN $4::bool THEN timestamp <= $5 ELSE TRUE END)
ORDER BY timestamp DESC
LIMIT $7
OFFSET $6
`

type ListStationObservationsParams struct {
	StationID   int64              `json:"station_id"`
	IsStartDate bool               `json:"is_start_date"`
	StartDate   pgtype.Timestamptz `json:"start_date"`
	IsEndDate   bool               `json:"is_end_date"`
	EndDate     pgtype.Timestamptz `json:"end_date"`
	Offset      int32              `json:"offset"`
	Limit       util.NullInt4      `json:"limit"`
}

func (q *Queries) ListStationObservations(ctx context.Context, arg ListStationObservationsParams) ([]ObservationsObservation, error) {
	rows, err := q.db.Query(ctx, listStationObservations,
		arg.StationID,
		arg.IsStartDate,
		arg.StartDate,
		arg.IsEndDate,
		arg.EndDate,
		arg.Offset,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ObservationsObservation{}
	for rows.Next() {
		var i ObservationsObservation
		if err := rows.Scan(
			&i.ID,
			&i.Pres,
			&i.Rr,
			&i.Rh,
			&i.Temp,
			&i.Td,
			&i.Wdir,
			&i.Wspd,
			&i.Wspdx,
			&i.Srad,
			&i.Mslp,
			&i.Hi,
			&i.StationID,
			&i.Timestamp,
			&i.Wchill,
			&i.QcLevel,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateStationObservation = `-- name: UpdateStationObservation :one
UPDATE observations_observation
SET
  pres = COALESCE($1, pres),
  rr = COALESCE($2, rr),
  rh = COALESCE($3, rh),
  temp = COALESCE($4, temp),
  td = COALESCE($5, td),
  wdir = COALESCE($6, wdir),
  wspd = COALESCE($7, wspd),
  wspdx = COALESCE($8, wspdx),
  srad = COALESCE($9, srad),
  mslp = COALESCE($10, mslp),
  hi = COALESCE($11, hi),
  wchill = COALESCE($12, wchill),
  timestamp = COALESCE($13, timestamp),
  qc_level = COALESCE($14, qc_level),
  updated_at = now()
WHERE station_id = $15 AND id = $16
RETURNING id, pres, rr, rh, temp, td, wdir, wspd, wspdx, srad, mslp, hi, station_id, timestamp, wchill, qc_level, created_at, updated_at
`

type UpdateStationObservationParams struct {
	Pres      util.NullFloat4    `json:"pres"`
	Rr        util.NullFloat4    `json:"rr"`
	Rh        util.NullFloat4    `json:"rh"`
	Temp      util.NullFloat4    `json:"temp"`
	Td        util.NullFloat4    `json:"td"`
	Wdir      util.NullFloat4    `json:"wdir"`
	Wspd      util.NullFloat4    `json:"wspd"`
	Wspdx     util.NullFloat4    `json:"wspdx"`
	Srad      util.NullFloat4    `json:"srad"`
	Mslp      util.NullFloat4    `json:"mslp"`
	Hi        util.NullFloat4    `json:"hi"`
	Wchill    util.NullFloat4    `json:"wchill"`
	Timestamp pgtype.Timestamptz `json:"timestamp"`
	QcLevel   util.NullInt4      `json:"qc_level"`
	StationID int64              `json:"station_id"`
	ID        int64              `json:"id"`
}

func (q *Queries) UpdateStationObservation(ctx context.Context, arg UpdateStationObservationParams) (ObservationsObservation, error) {
	row := q.db.QueryRow(ctx, updateStationObservation,
		arg.Pres,
		arg.Rr,
		arg.Rh,
		arg.Temp,
		arg.Td,
		arg.Wdir,
		arg.Wspd,
		arg.Wspdx,
		arg.Srad,
		arg.Mslp,
		arg.Hi,
		arg.Wchill,
		arg.Timestamp,
		arg.QcLevel,
		arg.StationID,
		arg.ID,
	)
	var i ObservationsObservation
	err := row.Scan(
		&i.ID,
		&i.Pres,
		&i.Rr,
		&i.Rh,
		&i.Temp,
		&i.Td,
		&i.Wdir,
		&i.Wspd,
		&i.Wspdx,
		&i.Srad,
		&i.Mslp,
		&i.Hi,
		&i.StationID,
		&i.Timestamp,
		&i.Wchill,
		&i.QcLevel,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
