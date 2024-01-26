package access

import (
	"moonlogs/internal/api/server/session"
	"moonlogs/internal/config"
	"moonlogs/internal/storage"
	"moonlogs/internal/usecases"
	"net/http"
	"slices"
)

func IsSchemaForbiddenForUser(schemaName string, r *http.Request) bool {
	schemaStorage := storage.NewSchemaStorage(r.Context(), config.Get().DBAdapter)
	schema, err := usecases.NewSchemaUseCase(schemaStorage).GetSchemaByName(schemaName)
	if err != nil || schema.ID == 0 {
		return true
	}

	user := session.GetUserFromContext(r)
	if len(user.Tags) > 0 && schema.TagID != 0 && !slices.Contains(user.Tags, schema.TagID) {
		return true
	}

	return false
}
