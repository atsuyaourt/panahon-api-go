// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	BatchCreateUserRoles(ctx context.Context, arg []BatchCreateUserRolesParams) *BatchCreateUserRolesBatchResults
	BatchDeleteUserRoles(ctx context.Context, arg []BatchDeleteUserRolesParams) *BatchDeleteUserRolesBatchResults
	CountLufftStationMsg(ctx context.Context, stationID int64) (int64, error)
	CountObservations(ctx context.Context, arg CountObservationsParams) (int64, error)
	CountRoles(ctx context.Context) (int64, error)
	CountStationObservations(ctx context.Context, arg CountStationObservationsParams) (int64, error)
	CountStations(ctx context.Context) (int64, error)
	CountStationsWithinBBox(ctx context.Context, arg CountStationsWithinBBoxParams) (int64, error)
	CountStationsWithinRadius(ctx context.Context, arg CountStationsWithinRadiusParams) (int64, error)
	CountUsers(ctx context.Context) (int64, error)
	CreateCurrentObservation(ctx context.Context, arg CreateCurrentObservationParams) (ObservationsCurrent, error)
	CreateGLabsLoad(ctx context.Context, arg CreateGLabsLoadParams) (GlabsLoad, error)
	CreateRole(ctx context.Context, arg CreateRoleParams) (Role, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateSimAccessToken(ctx context.Context, arg CreateSimAccessTokenParams) (SimAccessToken, error)
	CreateSimCard(ctx context.Context, arg CreateSimCardParams) (SimCard, error)
	CreateStation(ctx context.Context, arg CreateStationParams) (ObservationsStation, error)
	CreateStationHealth(ctx context.Context, arg CreateStationHealthParams) (ObservationsStationhealth, error)
	CreateStationObservation(ctx context.Context, arg CreateStationObservationParams) (ObservationsObservation, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteRole(ctx context.Context, id int64) error
	DeleteSimAccessToken(ctx context.Context, accessToken string) error
	DeleteStation(ctx context.Context, id int64) error
	DeleteStationHealth(ctx context.Context, arg DeleteStationHealthParams) error
	DeleteStationObservation(ctx context.Context, arg DeleteStationObservationParams) error
	DeleteUser(ctx context.Context, id int64) error
	GetLatestStationObservation(ctx context.Context, id int64) (GetLatestStationObservationRow, error)
	GetNearestLatestStationObservation(ctx context.Context, arg GetNearestLatestStationObservationParams) (GetNearestLatestStationObservationRow, error)
	GetRole(ctx context.Context, id int64) (Role, error)
	GetRoleByName(ctx context.Context, name string) (Role, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetSimAccessToken(ctx context.Context, accessToken string) (SimAccessToken, error)
	GetSimCard(ctx context.Context, mobileNumber string) (SimCard, error)
	GetStation(ctx context.Context, id int64) (ObservationsStation, error)
	GetStationByMobileNumber(ctx context.Context, mobileNumber pgtype.Text) (ObservationsStation, error)
	GetStationHealth(ctx context.Context, arg GetStationHealthParams) (ObservationsStationhealth, error)
	GetStationObservation(ctx context.Context, arg GetStationObservationParams) (ObservationsObservation, error)
	GetUser(ctx context.Context, id int64) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	InsertCurrentObservations(ctx context.Context) ([]ObservationsCurrent, error)
	ListLatestObservations(ctx context.Context) ([]ListLatestObservationsRow, error)
	ListLufftStationMsg(ctx context.Context, arg ListLufftStationMsgParams) ([]ListLufftStationMsgRow, error)
	ListObservations(ctx context.Context, arg ListObservationsParams) ([]ObservationsObservation, error)
	ListRoles(ctx context.Context, arg ListRolesParams) ([]Role, error)
	ListStationHealths(ctx context.Context, arg ListStationHealthsParams) ([]ObservationsStationhealth, error)
	ListStationObservations(ctx context.Context, arg ListStationObservationsParams) ([]ObservationsObservation, error)
	ListStations(ctx context.Context, arg ListStationsParams) ([]ObservationsStation, error)
	ListStationsWithinBBox(ctx context.Context, arg ListStationsWithinBBoxParams) ([]ObservationsStation, error)
	ListStationsWithinRadius(ctx context.Context, arg ListStationsWithinRadiusParams) ([]ObservationsStation, error)
	ListUserRoles(ctx context.Context, userID int64) ([]string, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	UpdateRole(ctx context.Context, arg UpdateRoleParams) (Role, error)
	UpdateStation(ctx context.Context, arg UpdateStationParams) (ObservationsStation, error)
	UpdateStationHealth(ctx context.Context, arg UpdateStationHealthParams) (ObservationsStationhealth, error)
	UpdateStationObservation(ctx context.Context, arg UpdateStationObservationParams) (ObservationsObservation, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
