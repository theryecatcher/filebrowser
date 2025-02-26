package auth

import (
	"net/http"

	"github.com/gtsteffaniak/filebrowser/backend/files"
	"github.com/gtsteffaniak/filebrowser/backend/users"
)

// MethodNoAuth is used to identify no auth.
const MethodNoAuth = "noauth"

// NoAuth is no auth implementation of auther.
type NoAuth struct{}

// Auth uses authenticates user 1.
func (a NoAuth) Auth(r *http.Request, usr *users.Storage) (*users.User, error) {
	return usr.Get(files.RootPaths["default"], uint(1))
}

// LoginPage tells that no auth doesn't require a login page.
func (a NoAuth) LoginPage() bool {
	return false
}
