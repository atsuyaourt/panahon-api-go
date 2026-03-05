package models

import (
	"encoding/json"
	"time"

	db "github.com/emiliogozo/panahon-api-go/internal/db/sqlc"
)

const (
	APITokenAuthKey       = "api_token_auth"
	APIKeyQueryParam      = "api_key"
	DefaultMaxDaysHistory = 365
)

// APITokenPermissions defines access restrictions for an API token
type APITokenPermissions struct {
	StationIDs     []int64  `json:"station_ids"`      // null means all stations
	Variables      []string `json:"variables"`        // null means all variables
	MaxDaysHistory int      `json:"max_days_history"` // 0 means unlimited
}

// APIToken represents an API token for programmatic access
type APIToken struct {
	ID          int64               `json:"id"`
	UserID      int64               `json:"user_id"`
	Name        string              `json:"name"`
	TokenPrefix string              `json:"token_prefix"`
	Permissions APITokenPermissions `json:"permissions"`
	LastUsedAt  *time.Time          `json:"last_used_at,omitempty"`
	CreatedAt   time.Time           `json:"created_at"`
}

// APITokenWithSecret includes the full token (only returned on creation)
type APITokenWithSecret struct {
	APIToken
	Token string `json:"token"`
}

// CreateAPITokenReq is the request body for creating an API token
type CreateAPITokenReq struct {
	Name        string              `json:"name" binding:"required"`
	ExpiresAt   *time.Time          `json:"expires_at,omitempty"`
	Permissions APITokenPermissions `json:"permissions"`
}

// NewAPIToken creates model from db type
func NewAPIToken(dbToken db.APIToken) APIToken {
	var permissions APITokenPermissions
	if dbToken.Permissions != nil {
		if err := json.Unmarshal(dbToken.Permissions, &permissions); err != nil {
			permissions = APITokenPermissions{
				MaxDaysHistory: 365,
			}
		}
	} else {
		permissions = APITokenPermissions{
			MaxDaysHistory: 365,
		}
	}

	var lastUsedAt *time.Time
	if dbToken.LastUsedAt.Valid {
		lastUsedAt = &dbToken.LastUsedAt.Time
	}

	return APIToken{
		ID:          dbToken.ID,
		UserID:      dbToken.UserID,
		Name:        dbToken.Name,
		TokenPrefix: dbToken.TokenPrefix,
		Permissions: permissions,
		LastUsedAt:  lastUsedAt,
		CreatedAt:   dbToken.CreatedAt.Time,
	}
}
