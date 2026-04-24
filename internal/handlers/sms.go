package handlers

import (
	"errors"
	"fmt"
	"net/http"

	db "github.com/emiliogozo/panahon-api-go/internal/db/sqlc"
	"github.com/emiliogozo/panahon-api-go/internal/sensor"
	"github.com/emiliogozo/panahon-api-go/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type smsStoreLufftReq struct {
	Number string `json:"number" binding:"required"`
	Msg    string `json:"msg" binding:"required"`
} //@name LufftSMSParams

// SMSStoreLufft
//
//	@Summary	Store Lufft observation and health
//	@Tags		sms
//	@Accept		json
//	@Produce	json
//	@Param		req	body		smsStoreLufftReq	true	"SMS parameters"
//	@Success	200	{object}	observationRes
//	@Router		/sms [post]
func (h *DefaultHandler) SMSStoreLufft(ctx *gin.Context) {
	var req smsStoreLufftReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.logger.Error().Err(err).
			Msg("[SMSStoreLufft] Bad request")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	lufft, err := sensor.NewLufftFromString(req.Msg)
	if err != nil {
		h.logger.Error().Err(err).
			Str("sender", req.Number).
			Str("msg", req.Msg).
			Msg("[SMSStoreLufft] Invalid string")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	mobileNumber, ok := util.ParseMobileNumber(req.Number)
	if !ok {
		err := fmt.Errorf("invalid mobile number: %s", req.Number)
		h.logger.Error().Err(err).
			Str("sender", req.Number).
			Str("msg", req.Msg).
			Msg("[SMSStoreLufft] Invalid mobile number")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	station, err := h.store.GetStationByMobileNumber(ctx, pgtype.Text{
		String: mobileNumber,
		Valid:  true,
	})
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			h.logger.Error().Err(err).
				Str("sender", req.Number).
				Str("msg", req.Msg).
				Msg("[SMSStoreLufft] No station found")
			ctx.JSON(http.StatusNotFound, errorResponse(errors.New("station not found")))
			return
		}
		h.logger.Error().Err(err).
			Str("sender", req.Number).
			Str("msg", req.Msg).
			Msg("[SMSStoreLufft] An error occured")
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	obsArg := db.CreateStationObservationParams{
		StationID: station.ID,
		Pres:      util.ToFloat4(lufft.Obs.Pres),
		Rr:        util.ToFloat4(lufft.Obs.Rr),
		Rh:        util.ToFloat4(lufft.Obs.Rh),
		Temp:      util.ToFloat4(lufft.Obs.Temp),
		Td:        util.ToFloat4(lufft.Obs.Td),
		Wdir:      util.ToFloat4(lufft.Obs.Wdir),
		Wspd:      util.ToFloat4(lufft.Obs.Wspd),
		Wspdx:     util.ToFloat4(lufft.Obs.Wspdx),
		Srad:      util.ToFloat4(lufft.Obs.Srad),
		Mslp:      util.ToFloat4(lufft.Obs.Mslp),
		Hi:        util.ToFloat4(lufft.Obs.Hi),
		Wchill:    util.ToFloat4(lufft.Obs.Wchill),
		Timestamp: pgtype.Timestamptz{
			Time:  lufft.Obs.Timestamp,
			Valid: true,
		},
	}

	obs, err := h.store.CreateStationObservation(ctx, obsArg)
	if err != nil {
		h.logger.Error().Err(err).
			Str("sender", req.Number).
			Str("msg", req.Msg).
			Msg("[SMSStoreLufft] Cannot store station observation")
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	healthArg := db.CreateStationHealthParams{
		StationID:         station.ID,
		Vb1:               util.ToFloat4(lufft.Health.Vb1),
		Vb2:               util.ToFloat4(lufft.Health.Vb2),
		Curr:              util.ToFloat4(lufft.Health.Curr),
		Bp1:               util.ToFloat4(lufft.Health.Bp1),
		Bp2:               util.ToFloat4(lufft.Health.Bp2),
		Cm:                util.ToPgText(lufft.Health.Cm),
		Ss:                util.ToInt4(lufft.Health.Ss),
		TempArq:           util.ToFloat4(lufft.Health.TempArq),
		RhArq:             util.ToFloat4(lufft.Health.RhArq),
		Fpm:               util.ToPgText(lufft.Health.Fpm),
		MinutesDifference: util.ToInt4(&lufft.Health.MinutesDifference),
		DataCount:         util.ToInt4(&lufft.Health.DataCount),
		DataStatus:        util.ToPgText(lufft.Health.DataStatus),
		Timestamp: pgtype.Timestamptz{
			Time:  lufft.Health.Timestamp,
			Valid: true,
		},
		Message:  util.ToPgText(lufft.Health.Message),
		ErrorMsg: util.ToPgText(lufft.Health.ErrorMsg),
	}

	health, err := h.store.CreateStationHealth(ctx, healthArg)
	if err != nil {
		h.logger.Error().Err(err).
			Str("sender", req.Number).
			Str("msg", req.Msg).
			Msg("[SMSStoreLufft] Cannot store station status")
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	res := newObservationResponse(station, obs, health)

	h.logger.Debug().
		Str("sender", req.Number).
		Str("msg", req.Msg).
		Msg("[SMSStoreLufft] Data saved successfully")
	ctx.JSON(http.StatusCreated, res)
}
