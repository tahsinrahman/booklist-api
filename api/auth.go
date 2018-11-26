package api

import (
	"errors"
	"net/http"
)

// User struct defines a user
type User struct {
	username string
	password string
}

var users []User

func init() {
	admin := User{
		username: "admin",
		password: "admin",
	}
	user := User{
		username: "user",
		password: "user",
	}

	users = append(users, admin, user)
}

// CheckAuth takes http.Request as parameter and checks requests's authorization
// header. For invalid username/password, it returns error
func CheckAuth(r *http.Request) error {
	username, password, ok := r.BasicAuth()
	if !ok {
		return errors.New("unauthorized")
	}

	for _, user := range users {
		if user.username == username && user.password == password {
			return nil
		} else if user.username == username && user.password != password {
			return errors.New("invalid password")
		}
	}
	return errors.New("user not found")
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := CheckAuth(r); err != nil {
			writeError(w, http.StatusUnauthorized, err.Error())
			return
		}
		next.ServeHTTP(w, r)
	})
}
