package asociado

import (
	database "TesisMVC/database/sqlc"
	"TesisMVC/internal/server/common_data"
	"context"
	"database/sql"
	"net/http"
	"strconv"
)

func (s *Server) CrearAsociado(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if CreateAsociado(r, s.GetStore()) {
			http.Redirect(w, r, "/listar", http.StatusSeeOther)
			return
		}
		http.Error(w, "POST", http.StatusInternalServerError)
		return
	}
	AllProvMun := []common_data.ProvinciaComplete{}
	prov, err := s.GetStore().GetAllProv(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, p := range prov {
		mun, err := s.GetStore().GetAllMunicipio(context.Background(), p.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		AllProvMun = append(AllProvMun, common_data.ProvinciaComplete{
			Provincia:  p,
			Municipios: mun,
		})
	}
	s.Data["prov"] = AllProvMun

	tpl.ExecuteTemplate(w, "crearAsociado.html", s.Data)
}

func CreateAsociado(r *http.Request, store database.Store) bool {
	nombre := r.FormValue("nombre")
	apellido1 := r.FormValue("apellido")
	apellido2 := r.FormValue("segundoApellido")
	activo := r.FormValue("activo")
	carnet := r.FormValue("carnet")
	sexo := r.FormValue("sexo")
	telefono := r.FormValue("telefono")
	otroTelefono := r.FormValue("otroTelefono")
	direccion := r.FormValue("direccion")
	municipio := r.FormValue("municipio")
	IDmun, err := strconv.Atoi(municipio)
	param := database.InsertAsoiciadoParams{
		Nombre:              nombre,
		Apellido1:           apellido1,
		Apellido2:           apellido2,
		Activo:              activo == "on",
		Carnet:              carnet,
		Sexo:                sexo == "on",
		NumeroT:             sql.NullString{String: telefono, Valid: true},
		NumeroPerteneciente: sql.NullString{String: otroTelefono, Valid: true},
		Direccion:           direccion,
		IDMunicipio:         int64(IDmun),
	}
	_, err = store.InsertAsoiciado(context.Background(), param)
	if err != nil {
		return false
	}
	return true
}
