package models

import (
	"slices"
	"time"

	db "github.com/emiliogozo/panahon-api-go/internal/db/sqlc"
	"github.com/emiliogozo/panahon-api-go/internal/util"
	"github.com/jackc/pgx/v5/pgtype"
)

type BaseStationObs struct {
	Pres      *float32  `json:"pres,omitempty" fake:"{float32range:990,1100}"`
	Rr        *float32  `json:"rr,omitempty"`
	Rh        *float32  `json:"rh,omitempty"`
	Temp      *float32  `json:"temp,omitempty" fake:"{float32range:25,35}"`
	Td        *float32  `json:"td,omitempty"`
	Wdir      *float32  `json:"wdir,omitempty"`
	Wspd      *float32  `json:"wspd,omitempty"`
	Wspdx     *float32  `json:"wspdx,omitempty"`
	Srad      *float32  `json:"srad,omitempty"`
	Mslp      *float32  `json:"mslp,omitempty"`
	Hi        *float32  `json:"hi,omitempty"`
	Wchill    *float32  `json:"wchill,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

type StationObservation struct {
	ID        int64 `json:"id" fake:"{number:1,1000}"`
	StationID int64 `json:"station_id" fake:"{number:1,250}"`
	QcLevel   int32 `json:"qc_level"`
	BaseStationObs
} //@name StationObservation

func shouldInclude(varNames []string, varName string) bool {
	return varNames == nil || slices.Contains(varNames, varName)
}

func setFloat32(dest **float32, src pgtype.Float4, varNames []string, varName string) {
	if src.Valid && shouldInclude(varNames, varName) {
		*dest = &src.Float32
	}
}

// NewStationObservation creates new StationObservation from db.ObservationsObservation
func NewStationObservation(obs db.ObservationsObservation, varNames []string) StationObservation {
	res := StationObservation{
		ID:        obs.ID,
		StationID: obs.StationID,
		QcLevel:   obs.QcLevel,
	}

	setFloat32(&res.Pres, obs.Pres, varNames, "pres")
	setFloat32(&res.Rr, obs.Rr, varNames, "rain")
	setFloat32(&res.Rh, obs.Rh, varNames, "rh")
	setFloat32(&res.Temp, obs.Temp, varNames, "temp")
	setFloat32(&res.Td, obs.Td, varNames, "td")
	setFloat32(&res.Wdir, obs.Wdir, varNames, "wind")
	setFloat32(&res.Wspd, obs.Wspd, varNames, "wind")
	setFloat32(&res.Wspdx, obs.Wspdx, varNames, "windx")
	setFloat32(&res.Srad, obs.Srad, varNames, "srad")
	setFloat32(&res.Mslp, obs.Mslp, varNames, "pres")
	setFloat32(&res.Hi, obs.Hi, varNames, "hi")
	setFloat32(&res.Wchill, obs.Wchill, varNames, "wchill")

	if obs.Timestamp.Valid {
		res.Timestamp = obs.Timestamp.Time
	}

	return res
}

type CreateStationObsReq struct {
	StationID int64 `json:"station_id"`
	QcLevel   int32 `json:"qc_level"`
	BaseStationObs
} //@name CreateStationObservationReq

func (r CreateStationObsReq) Transform() db.CreateStationObservationParams {
	return transformStationObs(
		r.BaseStationObs,
		db.CreateStationObservationParams{
			StationID: r.StationID,
			QcLevel:   r.QcLevel,
		})
}

type UpdateStationObsReq struct {
	ID        int64  `json:"id"`
	StationID int64  `json:"station_id"`
	QcLevel   *int32 `json:"qc_level"`
	BaseStationObs
} //@name UpdateStationObservationParams

func (r UpdateStationObsReq) Transform() db.UpdateStationObservationParams {
	return transformStationObs(
		r.BaseStationObs,
		db.UpdateStationObservationParams{
			ID:        r.ID,
			StationID: r.StationID,
			QcLevel:   util.ToInt4(r.QcLevel),
		})
}

type StationObsParams interface {
	db.CreateStationObservationParams | db.UpdateStationObservationParams
}

func transformStationObs[T StationObsParams](req BaseStationObs, extraParams T) T {
	arg := db.CreateStationObservationParams{
		Pres:   util.ToFloat4(req.Pres),
		Rr:     util.ToFloat4(req.Rr),
		Rh:     util.ToFloat4(req.Rh),
		Temp:   util.ToFloat4(req.Temp),
		Td:     util.ToFloat4(req.Td),
		Wdir:   util.ToFloat4(req.Wdir),
		Wspd:   util.ToFloat4(req.Wspd),
		Wspdx:  util.ToFloat4(req.Wspdx),
		Srad:   util.ToFloat4(req.Srad),
		Mslp:   util.ToFloat4(req.Mslp),
		Hi:     util.ToFloat4(req.Hi),
		Wchill: util.ToFloat4(req.Wchill),
		Timestamp: pgtype.Timestamptz{
			Time:  req.Timestamp,
			Valid: !req.Timestamp.IsZero(),
		},
	}

	switch v := any(extraParams).(type) {
	case db.CreateStationObservationParams:
		arg.StationID = v.StationID
		arg.QcLevel = v.QcLevel
		return any(arg).(T)
	case db.UpdateStationObservationParams:
		return any(db.UpdateStationObservationParams{
			ID:        v.ID,
			StationID: v.StationID,
			Pres:      arg.Pres,
			Rr:        arg.Rr,
			Rh:        arg.Rh,
			Temp:      arg.Temp,
			Td:        arg.Td,
			Wdir:      arg.Wdir,
			Wspd:      arg.Wspd,
			Wspdx:     arg.Wspdx,
			Srad:      arg.Srad,
			Mslp:      arg.Mslp,
			Hi:        arg.Hi,
			Wchill:    arg.Wchill,
			Timestamp: arg.Timestamp,
			QcLevel:   v.QcLevel,
		}).(T)
	default:
		panic("Unsupported type")
	}
}
