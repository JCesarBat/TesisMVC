package auth

import (
	database "TesisMVC/database/sqlc"
	"TesisMVC/pkg/util"
	"context"
	"database/sql"
	"net/http"
	"strconv"
)

type ProvinciaComplete struct {
	Provincia  database.Provincium
	Municipios []database.Municipio
}

func (s *Server) Register(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	AllProvMun := []ProvinciaComplete{}
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
		AllProvMun = append(AllProvMun, ProvinciaComplete{
			Provincia:  p,
			Municipios: mun,
		})
	}
	data["prov"] = AllProvMun

	if r.Method == "POST" {
		if RegisterUser(r, s.GetStore()) {
			http.Redirect(w, r, "/home", http.StatusFound)
			return
		}
		http.Error(w, "Invalid user information", http.StatusBadRequest)
	}
	tpl.ExecuteTemplate(w, "register.html", data)
}

func RegisterUser(r *http.Request, store database.Store) bool {
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
	return true
}
