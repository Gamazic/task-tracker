package microframework

import (
	"encoding/base64"
	"errors"
	"net/http"
	"strings"
)

const AuthHeaderKey = "Authorization"

func BasicAuthentication(next http.Handler, protectedPathPrefix string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, protectedPathPrefix) {
			next.ServeHTTP(w, r)
			return
		}
		authHeader := r.Header.Get(AuthHeaderKey)
		if authHeader == "" {
			w.Header().Set("WWW-Authenticate", "basic")
			NewResponseBuilder(w).
				BuildStatus(http.StatusUnauthorized).
				BuildBodyPlainMsg("Authorization required").
				Send()
			return
		}
		authProvide := strings.HasPrefix(authHeader, "Basic ")
		if !authProvide {
			SendValidationError(w, errors.New("only basic auth is supported"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

type Credentials struct {
	Username string
	Password string
}

func GetCredentials(r *http.Request) (Credentials, error) {
	authHeader := r.Header.Get(AuthHeaderKey)
	if authHeader == "" {
		return Credentials{}, errors.New("no auth header")
	}
	authSchema := strings.Split(authHeader, " ")
	if len(authSchema) != 2 {
		return Credentials{}, errors.New("bad auth header value")
	}
	encCreds, err := base64.StdEncoding.DecodeString(authSchema[1])
	if err != nil {
		return Credentials{}, errors.New("bad base64 encoding in auth header")
	}
	creds := strings.Split(string(encCreds), ":")
	if len(creds) != 2 {
		return Credentials{}, errors.New("auth header does not follow user:password scheme")
	}
	return Credentials{
		Username: creds[0],
		Password: creds[1],
	}, nil
}
