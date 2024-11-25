package auth

import (
	"TesisMVC/internal/server/common_data"
	"html/template"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("internal/template/login.html", "internal/template/register.html"))

type Server struct {
	common_data.GinServer
}

func (s *Server) CloseSession(w http.ResponseWriter, r *http.Request) {
	s.GetCookie().DeleteCookie(w)
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}
