package auth

import (
	"TesisMVC/internal/server/common_data"
	"html/template"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("internal/template/auth/login.html", "internal/template/auth/register.html", "internal/template/auth/home.html"))

type Server struct {
	common_data.GinServer
	Data map[string]any
}

func (s *Server) ResetData() {
	s.Data = make(map[string]any)

}

func (s *Server) CloseSession(w http.ResponseWriter, r *http.Request) {
	s.GetCookie().DeleteCookie(w)
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}
