package sensor

import (
	"time"

	"github.com/emiliogozo/panahon-api-go/internal/util"
	"github.com/jackc/pgx/v5/pgtype"
)

type davisRawCurrentResponseV1 struct {
	Location   string                       `json:"location"`
	Lat        util.JSONFloat4              `json:"latitude"`
	Lon        util.JSONFloat4              `json:"longitude"`
	Time       string                       `json:"observation_time_rfc822"`
	PressureMb util.JSONFloat4              `json:"pressure_mb"`
	Rh         util.JSONFloat4              `json:"relative_humidity"`
	TempC      util.JSONFloat4              `json:"temp_c"`
	TdC        util.JSONFloat4              `json:"dewpoint_c"`
	WindDeg    util.JSONFloat4              `json:"wind_degrees"`
	WindMPH    util.JSONFloat4              `json:"wind_mph"`
	HeatIndexC util.JSONFloat4              `json:"heat_index_c"`
	Obs        davisRawCurrentObservationV1 `json:"davis_current_observation"`
}

type davisRawCurrentObservationV1 struct {
	RRInPerHr       util.JSONFloat4 `json:"rain_rate_in_per_hr"`
	RainDayIn       util.JSONFloat4 `json:"rain_day_in"`
	Srad            util.JSONFloat4 `json:"solar_radiation"`
	UVIndex         util.JSONFloat4 `json:"uv_index"`
	TempDayHighF    util.JSONFloat4 `json:"temp_day_high_f"`
	TempDayLowF     util.JSONFloat4 `json:"temp_day_low_f"`
	WindDayHighMPH  util.JSONFloat4 `json:"wind_day_high_mph"`
	TempDayHighTime string          `json:"temp_day_high_time"`
	TempDayLowTime  string          `json:"temp_day_low_time"`
	WindDayHighTime string          `json:"wind_day_high_time"`
}

func (r davisRawCurrentResponseV1) ToDavisCurrentObservation() *DavisCurrentObservation {
	obs := DavisCurrentObservation{
		Rr:            r.Obs.RRInPerHr.ToFloat4(),
		RainAccum:     r.Obs.RainDayIn.ToFloat4(),
		Temp:          r.TempC.ToFloat4(),
		Rh:            r.Rh.ToFloat4(),
		Wdir:          r.WindDeg.ToFloat4(),
		Wspd:          r.WindMPH.ToFloat4(),
		Srad:          r.Obs.Srad.ToFloat4(),
		Pres:          r.PressureMb.ToFloat4(),
		Tx:            r.Obs.TempDayHighF.ToFloat4(),
		Tn:            r.Obs.TempDayLowF.ToFloat4(),
		Wspdx:         r.Obs.WindDayHighMPH.ToFloat4(),
		Hi:            r.HeatIndexC.ToFloat4(),
		TxTimestamp:   pgtype.Timestamptz{Time: time.Time{}, Valid: true},
		TnTimestamp:   pgtype.Timestamptz{Time: time.Time{}, Valid: true},
		GustTimestamp: pgtype.Timestamptz{Time: time.Time{}, Valid: true},
	}

	if obs.Rr.Valid {
		obs.Rr.Float32 = obs.Rr.Float32 * 25.4
	}
	if obs.RainAccum.Valid {
		obs.RainAccum.Float32 = obs.Rr.Float32 * 25.4
	}
	if obs.Wspd.Valid {
		obs.Wspd.Float32 = obs.Wspd.Float32 * 0.44704
	}
	if obs.Tx.Valid {
		obs.Tx.Float32 = util.FahrenheitToCelsius(obs.Tx.Float32)
		if dt, err := parseTimeStrToDateTime(r.Obs.TempDayHighTime); err == nil {
			obs.TxTimestamp.Time = dt
		}
	}
	if obs.Tn.Valid {
		obs.Tn.Float32 = util.FahrenheitToCelsius(obs.Tn.Float32)
		if dt, err := parseTimeStrToDateTime(r.Obs.TempDayLowTime); err == nil {
			obs.TnTimestamp.Time = dt
		}
	}
	if obs.Wspdx.Valid {
		obs.Wspdx.Float32 = obs.Wspdx.Float32 * 0.44704
		if dt, err := parseTimeStrToDateTime(r.Obs.WindDayHighTime); err == nil {
			obs.GustTimestamp.Time = dt
		}
	}

	layout := "Mon, 2 Jan 2006 15:04:05 -0700"
	if dt, err := time.Parse(layout, r.Time); err == nil {
		obs.Timestamp = pgtype.Timestamptz{Time: dt, Valid: true}
	}

	return &obs
}
