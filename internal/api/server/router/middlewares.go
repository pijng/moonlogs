package router

import (
	"context"
	"moonlogs/internal/api/server/response"
	"moonlogs/internal/api/server/session"
	"moonlogs/internal/entities"
	"moonlogs/internal/shared"
	"moonlogs/internal/usecases"
	"net/http"
	"slices"
)

type Middlewares struct {
	userUseCase     *usecases.UserUseCase
	apiTokenUseCase *usecases.ApiTokenUseCase
}

func InitMiddlewares(uc *usecases.UseCases) *Middlewares {
	return &Middlewares{
		userUseCase:     uc.UserUseCase,
		apiTokenUseCase: uc.ApiTokenUseCase,
	}
}

func (mw *Middlewares) SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionCookie, err := session.GetSessionStore().Get(r, session.NAME)
		if err != nil {
			response.Return(w, false, http.StatusInternalServerError, err, nil, response.Meta{})
			return
		}

		if sessionCookie != nil {
			sessionToken, tokenOk := sessionCookie.Values["token"].(string)
			sessionUserID, userIDOk := sessionCookie.Values["userID"].(int)

			if tokenOk && sessionToken != "" && userIDOk && sessionUserID != 0 {
				ctx := context.WithValue(r.Context(), session.KEY, sessionCookie)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
		}

		bearerToken := shared.ExtractBearerToken(r)

		ok, err := mw.apiTokenUseCase.IsTokenValid(r.Context(), bearerToken)
		if err != nil {
			response.Return(w, false, http.StatusInternalServerError, err, nil, response.Meta{})
			return
		}

		if ok {
			next.ServeHTTP(w, r)
			return
		}

		token, ok := sessionCookie.Values["token"].(string)
		if bearerToken == "" && ok {
			bearerToken = token
		}

		user, _ := mw.userUseCase.GetUserByToken(r.Context(), bearerToken)

		if !ok || user == nil || bool(user.IsRevoked) {
			response.Return(w, false, http.StatusUnauthorized, nil, nil, response.Meta{})
			return
		}

		if token == "" {
			token = user.Token
			sessionCookie.Values["token"] = token
			sessionCookie.Values["userID"] = user.ID
			sessionCookie.Values["role"] = string(user.Role)
		}

		ctx := context.WithValue(r.Context(), session.KEY, sessionCookie)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (mw *Middlewares) RoleMiddleware(next http.HandlerFunc, requiredRoles ...entities.Role) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionCookie, err := session.GetSessionStore().Get(r, session.NAME)
		if err != nil {
			response.Return(w, false, http.StatusInternalServerError, err, nil, response.Meta{})
			return
		}

		if sessionCookie != nil {
			role, ok := sessionCookie.Values["role"].(string)

			if ok && slices.Contains(requiredRoles, entities.Role(role)) {
				ctx := context.WithValue(r.Context(), session.KEY, sessionCookie)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
		}

		bearerToken := shared.ExtractBearerToken(r)

		validAPIToken, err := mw.apiTokenUseCase.IsTokenValid(r.Context(), bearerToken)
		if err != nil {
			response.Return(w, false, http.StatusInternalServerError, err, nil, response.Meta{})
			return
		}

		// allow access by API token for certain actions only
		if validAPIToken {
			if !slices.Contains(requiredRoles, entities.TokenRole) {
				response.Return(w, false, http.StatusForbidden, nil, nil, response.Meta{})
				return
			}

			next.ServeHTTP(w, r)
			return
		}

		user, _ := mw.userUseCase.GetUserByToken(r.Context(), bearerToken)

		if !slices.Contains(requiredRoles, user.Role) {
			response.Return(w, false, http.StatusForbidden, nil, nil, response.Meta{})
			return
		}

		next(w, r)
	}
}
