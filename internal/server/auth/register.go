package auth

import (
	database "TesisMVC/database/sqlc"
	"TesisMVC/internal/server/common_data"
	"TesisMVC/pkg/util"
	"context"
	"database/sql"
	"net/http"
	"strconv"
)

func (s *Server) Register(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
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
	data["prov"] = AllProvMun

	if r.Method == "POST" {
		if RegisterUser(r, w, s.GetStore(), s) {
			return
		}
		w.WriteHeader(401)
		data["status"] = 401
		tpl.ExecuteTemplate(w, "register.html", data)
		return
	}
	tpl.ExecuteTemplate(w, "register.html", data)
}

func RegisterUser(r *http.Request, w http.ResponseWriter, store database.Store, s *Server) bool {
	var SU bool
	username := r.FormValue("nombre")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirm := r.FormValue("confirm-password")
	superuser := r.FormValue("superuser")
	if superuser == "no" {
		SU = false
	} else {
		SU = true
	}
	municipio := r.FormValue("municipio")
	munId, err := strconv.Atoi(municipio)
	if err != nil {
		return false
	}
	if password != confirm {
		return false
	}
	pass, err := util.HashPassword(password)
	if err != nil {
		return false
	}

	param := database.InsertUserParams{
		Username:    username,
		Email:       email,
		Password:    pass,
		SuperUser:   sql.NullBool{Bool: SU, Valid: true},
		IDMunicipio: int64(munId),
	}
	_, err = store.InsertUser(context.Background(), param)
	if err != nil {
		return false
	}
	s.GetCookie().SetCookie(username, w)
	return true

}
