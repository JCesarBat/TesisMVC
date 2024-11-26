package asociado

import (
	database "TesisMVC/database/sqlc"
	"context"
	"net/http"
)

type CulturalContent struct {
	Actividad     database.ActividadCultural
	Participacion []database.ParticipacionC
}

type DeportivaContent struct {
	Actividad     database.ActividadDeportiva
	Participacion []database.ParticipacionD
}

type EducativaContent struct {
	Actividad database.ActividadEducativa
	Estudios  database.EstudiosActuale
}
type AsociadosContent struct {
	database.Asociado
	database.DatosSociale
	Cultural  CulturalContent
	Deportiva DeportivaContent
	Educativa EducativaContent
}

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
