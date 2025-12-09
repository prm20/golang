package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/avukadin/goapi/api"
	"github.com/avukadin/goapi/internal/tools"
	log "github.con/sirupsen/logrus"
)

var UnAuthorizedError = errors.new("Invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}
		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
		}
		var loginDetails *tools.loginDetails
		loginDetails = (*database).GetUserLoginDetails(username)
		if (loginDetails == nil || (token != (*loginDetails).AuthToken)) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}
		next.ServeHTTP(w, r)
	})
}