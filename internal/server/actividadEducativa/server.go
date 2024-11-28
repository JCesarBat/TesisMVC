package actividadEducativa

import (
	database "TesisMVC/database/sqlc"
	"TesisMVC/internal/server/common_data"
	"context"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	common_data.GinServer
	Data map[string]any
}

func (s *Server) ResetData() {
	s.Data = make(map[string]any)

}

func (s *Server) CrearActividadEducativa(w http.ResponseWriter, r *http.Request) {
	var grado database.UltimoGrado
	ultimoGrado := r.FormValue("grados")
	especialidad := r.FormValue("especialidad")
	centro := r.FormValue("centro")
	fecha := r.FormValue("fecha")
	enseñansa := r.FormValue("enseñansa")
	layout := time.DateOnly
	date, err := time.Parse(layout, fecha)
	if err != nil {
		http.Error(w, "Error al convertir la fecha", 500)
		return
	}
	param := database.CreateEstudiosActualesParams{
		TipoEnseñansa:         enseñansa,
		Centro:                centro,
		EspecialidadGradoOAño: especialidad,
		AñoDelDato:            date,
		FechaDeGraduacion:     date,
	}
	estudiosActuales, err := s.GetStore().CreateEstudiosActuales(context.Background(), param)
	if err != nil {
		http.Error(w, "Error al crear los estudios actuales", 500)
		return
	}
	id, err := strconv.Atoi(r.URL.Query()["id"][0])
	if err != nil {
		http.Error(w, "Error al convertir el id", 500)
		return
	}
	err = grado.Scan(ultimoGrado)
	if err != nil {
		http.Error(w, "Error al obtener el ultimo grado", 500)
		return
	}

	arg := database.CreateActividadEducativaParams{
		IDAsociado:          int64(id),
		IDEstudiosActuales:  estudiosActuales.ID,
		UltimoGradoAprobado: grado,
	}
	_, err = s.GetStore().CreateActividadEducativa(context.Background(), arg)
	if err != nil {
		http.Error(w, "Error al crear la actividad educativa", 500)
		return
	}
	http.Redirect(w, r, "/listar", http.StatusFound)
}
