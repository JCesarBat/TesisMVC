package auth

import (
	database "TesisMVC/database/sqlc"
	"TesisMVC/pkg/util"
	"context"
	"net/http"
)

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	var data map[string]int
	data = make(map[string]int)
	data["status"] = 0
	if r.Method == http.MethodPost {
		// Process the form data
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Validate the credentials
		if validCredentials(username, password, s.GetStore()) {
			// Redirect to the home page
			s.GetCookie().SetCookie(username, w) // Set the cookie with the username for future authentication
			http.Redirect(w, r, "/home/", http.StatusFound)
			return
		} else {
			// Display an error message

			w.WriteHeader(401)
			data["status"] = 401
			tpl.ExecuteTemplate(w, "login.html", data)
			return
		}
	}
	tpl.ExecuteTemplate(w, "login.html", data)
}

func validCredentials(username, password string, store database.Store) bool {
	user, err := store.GetUserByUsername(context.Background(), username)
	if err != nil {
		return false
	}
	err = util.CheckPassword(user.Password, password)
	if err != nil {
		return false
	}
	return true
}
