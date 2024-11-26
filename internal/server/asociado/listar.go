package asociado

import (
	database "TesisMVC/database/sqlc"
	"context"
	"net/http"
)

func (s *Server) Listar(w http.ResponseWriter, r *http.Request) {
	arg := database.ListAsociadoParams{
		Limit:  10,
		Offset: (1 - 1) * 10,
	}
	asociados, err := s.GetStore().ListAsociado(context.Background(), arg)
	if err != nil {
		http.Error(w, "error al leer los asociados", 500)
		return
	}
	s.Data["asociados"] = asociados

	tpl.ExecuteTemplate(w, "listarAsociado.html", s.Data)
}
