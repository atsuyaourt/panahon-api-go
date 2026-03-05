package handlers

import (
	"encoding/json"

	db "github.com/emiliogozo/panahon-api-go/internal/db/sqlc"
	"github.com/emiliogozo/panahon-api-go/internal/models"
	"github.com/gin-gonic/gin"
)

func getAccessibleVariables(ctx *gin.Context) []string {
	key, exists := ctx.Get(models.APITokenAuthKey)
	if !exists {
		return nil
	}

	dbToken, ok := key.(*db.APIToken)
	if !ok || dbToken == nil {
		return nil
	}

	var perm models.APITokenPermissions
	err := json.Unmarshal(dbToken.Permissions, &perm)
	if err != nil {
		return nil
	}

	if perm.Variables == nil {
		return nil
	}

	return perm.Variables
}
