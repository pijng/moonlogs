package access

import (
	"moonlogs/internal/api/server/session"
	"moonlogs/internal/usecases"
	"net/http"
	"slices"
)

func IsSchemaForbiddenForUser(schemaUseCase *usecases.SchemaUseCase, schemaName string, r *http.Request) bool {
	user := session.GetUserFromContext(r)
	if user == nil {
		return false
	}

	schema, err := schemaUseCase.GetSchemaByName(r.Context(), schemaName)
	if err != nil {
		return true
	}

	if len(user.Tags) > 0 && schema.TagID != 0 && !slices.Contains(user.Tags, schema.TagID) {
		return true
	}

	return false
}
