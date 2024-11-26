package asociado

import "net/http"

func (s *Server) Listar(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "listarAsociado.html", s.Data)
}
