// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"github.com/emiliogozo/panahon-api-go/util"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type GlabsLoad struct {
	ID            int64              `json:"id"`
	Status        util.NullString    `json:"status"`
	Promo         util.NullString    `json:"promo"`
	TransactionID util.NullInt4      `json:"transaction_id"`
	MobileNumber  string             `json:"mobile_number"`
	CreatedAt     pgtype.Timestamptz `json:"created_at"`
	UpdatedAt     pgtype.Timestamptz `json:"updated_at"`
}

type ObservationsMoObservation struct {
	ID        int64              `json:"id"`
	Pres      util.NullFloat4    `json:"pres"`
	Rr        util.NullFloat4    `json:"rr"`
	Rh        util.NullFloat4    `json:"rh"`
	Temp      util.NullFloat4    `json:"temp"`
	Td        util.NullFloat4    `json:"td"`
	Wdir      util.NullFloat4    `json:"wdir"`
	Wspd      util.NullFloat4    `json:"wspd"`
	Wspdx     util.NullFloat4    `json:"wspdx"`
	Srad      util.NullFloat4    `json:"srad"`
	Hi        util.NullFloat4    `json:"hi"`
	StationID int64              `json:"station_id"`
	Timestamp pgtype.Timestamptz `json:"timestamp"`
	Wchill    util.NullFloat4    `json:"wchill"`
	Rain      util.NullFloat4    `json:"rain"`
	Tx        util.NullFloat4    `json:"tx"`
	Tn        util.NullFloat4    `json:"tn"`
	Wrun      util.NullFloat4    `json:"wrun"`
	Thwi      util.NullFloat4    `json:"thwi"`
	Thswi     util.NullFloat4    `json:"thswi"`
	Senergy   util.NullFloat4    `json:"senergy"`
	Sradx     util.NullFloat4    `json:"sradx"`
	Uvi       util.NullFloat4    `json:"uvi"`
	Uvdose    util.NullFloat4    `json:"uvdose"`
	Uvx       util.NullFloat4    `json:"uvx"`
	Hdd       util.NullFloat4    `json:"hdd"`
	Cdd       util.NullFloat4    `json:"cdd"`
	Et        util.NullFloat4    `json:"et"`
	QcLevel   int32              `json:"qc_level"`
	Wdirx     util.NullFloat4    `json:"wdirx"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

type ObservationsObservation struct {
	ID        int64              `json:"id"`
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
	StationID int64              `json:"station_id"`
	Timestamp pgtype.Timestamptz `json:"timestamp"`
	Wchill    util.NullFloat4    `json:"wchill"`
	QcLevel   int32              `json:"qc_level"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

type ObservationsStation struct {
	ID            int64              `json:"id"`
	Name          string             `json:"name"`
	Lat           util.NullFloat4    `json:"lat"`
	Lon           util.NullFloat4    `json:"lon"`
	Elevation     util.NullFloat4    `json:"elevation"`
	DateInstalled pgtype.Date        `json:"date_installed"`
	MoStationID   util.NullString    `json:"mo_station_id"`
	SmsSystemType util.NullString    `json:"sms_system_type"`
	MobileNumber  util.NullString    `json:"mobile_number"`
	StationType   util.NullString    `json:"station_type"`
	StationType2  util.NullString    `json:"station_type2"`
	StationUrl    util.NullString    `json:"station_url"`
	Status        util.NullString    `json:"status"`
	LoggerVersion util.NullString    `json:"logger_version"`
	PriorityLevel util.NullString    `json:"priority_level"`
	ProviderID    util.NullString    `json:"provider_id"`
	Province      util.NullString    `json:"province"`
	Region        util.NullString    `json:"region"`
	Address       util.NullString    `json:"address"`
	CreatedAt     pgtype.Timestamptz `json:"created_at"`
	UpdatedAt     pgtype.Timestamptz `json:"updated_at"`
	DeletedAt     pgtype.Timestamptz `json:"deleted_at"`
}

type ObservationsStationhealth struct {
	ID                int64              `json:"id"`
	Vb1               util.NullFloat4    `json:"vb1"`
	Vb2               util.NullFloat4    `json:"vb2"`
	Curr              util.NullFloat4    `json:"curr"`
	Bp1               util.NullFloat4    `json:"bp1"`
	Bp2               util.NullFloat4    `json:"bp2"`
	Cm                util.NullString    `json:"cm"`
	Ss                util.NullInt4      `json:"ss"`
	TempArq           util.NullFloat4    `json:"temp_arq"`
	RhArq             util.NullFloat4    `json:"rh_arq"`
	Fpm               util.NullString    `json:"fpm"`
	ErrorMsg          util.NullString    `json:"error_msg"`
	Message           util.NullString    `json:"message"`
	DataCount         util.NullInt4      `json:"data_count"`
	DataStatus        util.NullString    `json:"data_status"`
	Timestamp         pgtype.Timestamptz `json:"timestamp"`
	StationID         int64              `json:"station_id"`
	MinutesDifference util.NullInt4      `json:"minutes_difference"`
	CreatedAt         pgtype.Timestamptz `json:"created_at"`
	UpdatedAt         pgtype.Timestamptz `json:"updated_at"`
}

type Role struct {
	ID          int64              `json:"id"`
	Name        string             `json:"name"`
	Description util.NullString    `json:"description"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
}

type RoleUser struct {
	RoleID    int64              `json:"role_id"`
	UserID    int64              `json:"user_id"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

type Session struct {
	ID           uuid.UUID          `json:"id"`
	UserID       int64              `json:"user_id"`
	RefreshToken string             `json:"refresh_token"`
	UserAgent    string             `json:"user_agent"`
	ClientIp     string             `json:"client_ip"`
	IsBlocked    bool               `json:"is_blocked"`
	ExpiresAt    pgtype.Timestamptz `json:"expires_at"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
}

type SimAccessToken struct {
	AccessToken  string             `json:"access_token"`
	Type         string             `json:"type"`
	MobileNumber string             `json:"mobile_number"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
	UpdatedAt    pgtype.Timestamptz `json:"updated_at"`
}

type SimCard struct {
	MobileNumber string             `json:"mobile_number"`
	Type         util.NullString    `json:"type"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
	UpdatedAt    pgtype.Timestamptz `json:"updated_at"`
}

type User struct {
	ID                int64              `json:"id"`
	Username          string             `json:"username"`
	FullName          string             `json:"full_name"`
	Email             string             `json:"email"`
	Password          string             `json:"password"`
	EmailVerifiedAt   pgtype.Timestamptz `json:"email_verified_at"`
	PasswordChangedAt pgtype.Timestamptz `json:"password_changed_at"`
	CreatedAt         pgtype.Timestamptz `json:"created_at"`
	UpdatedAt         pgtype.Timestamptz `json:"updated_at"`
}
