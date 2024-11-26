package auth

import "net/http"

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	s.Data["username"] = s.GetCookie().GetCookie(r)
	if s.Data["username"] == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
}
