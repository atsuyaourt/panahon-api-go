// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"context"

	"github.com/emiliogozo/panahon-api-go/util"
	"github.com/google/uuid"
)

type Querier interface {
	BatchCreateUserRoles(ctx context.Context, arg []BatchCreateUserRolesParams) *BatchCreateUserRolesBatchResults
	BatchDeleteUserRoles(ctx context.Context, arg []BatchDeleteUserRolesParams) *BatchDeleteUserRolesBatchResults
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
	DeleteStation(ctx context.Context, id int64) error
	DeleteStationHealth(ctx context.Context, arg DeleteStationHealthParams) error
	DeleteStationObservation(ctx context.Context, arg DeleteStationObservationParams) error
	DeleteUser(ctx context.Context, id int64) error
	GetRole(ctx context.Context, id int64) (Role, error)
	GetRoleByName(ctx context.Context, name string) (Role, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetSimAccessToken(ctx context.Context, accessToken string) (SimAccessToken, error)
	GetSimCard(ctx context.Context, mobileNumber string) (SimCard, error)
	GetStation(ctx context.Context, id int64) (ObservationsStation, error)
	GetStationByMobileNumber(ctx context.Context, mobileNumber util.NullString) (ObservationsStation, error)
	GetStationHealth(ctx context.Context, arg GetStationHealthParams) (ObservationsStationhealth, error)
	GetStationObservation(ctx context.Context, arg GetStationObservationParams) (ObservationsObservation, error)
	GetUser(ctx context.Context, id int64) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	ListLufftStationMsg(ctx context.Context, arg ListLufftStationMsgParams) ([]ListLufftStationMsgRow, error)
	ListRoles(ctx context.Context, arg ListRolesParams) ([]Role, error)
	ListStationHealths(ctx context.Context, arg ListStationHealthsParams) ([]ObservationsStationhealth, error)
	ListStationObservations(ctx context.Context, arg ListStationObservationsParams) ([]ObservationsObservation, error)
	ListStations(ctx context.Context, arg ListStationsParams) ([]ObservationsStation, error)
	ListUserRoles(ctx context.Context, userID int64) ([]string, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	UpdateRole(ctx context.Context, arg UpdateRoleParams) (Role, error)
	UpdateStation(ctx context.Context, arg UpdateStationParams) (ObservationsStation, error)
	UpdateStationHealth(ctx context.Context, arg UpdateStationHealthParams) (ObservationsStationhealth, error)
	UpdateStationObservation(ctx context.Context, arg UpdateStationObservationParams) (ObservationsObservation, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
