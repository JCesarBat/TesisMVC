package asociado

import (
	database "TesisMVC/database/sqlc"
	"context"
	"net/http"
)

func (s *Server) CrearAsociado(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

	}
	tpl.ExecuteTemplate(w, "crearAsociado.html", s.Data)
}

func CreateAsociado(r *http.Request, store database.Store) bool {
	nombre := r.FormValue("nombre")
	param := database.InsertAsoiciadoParams{
		Nombre: nombre,
	}
	_, err := store.InsertAsoiciado(context.Background(), param)
	if err != nil {
		return false
	}
	return true
}
