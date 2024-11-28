package asociado

import (
	database "TesisMVC/database/sqlc"
	"context"
	"fmt"
	"net/http"
	"strconv"
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
	Actividad  database.ActividadEducativa
	Estudios   database.EstudiosActuale
	FormatDate string
	Editar     bool
}
type AsociadosContent struct {
	database.Asociado
	database.DatosSociale
	Cultural  CulturalContent
	Deportiva DeportivaContent
	Educativa EducativaContent
}

func (s *Server) Listar(w http.ResponseWriter, r *http.Request) {
	var resp []AsociadosContent
	paginacion, err := PaginacionNumber(s.GetStore())
	if err != nil {
		http.Error(w, "error al leer la paginacion", 500)
		return
	}
	s.Data["paginacion"] = paginacion
	Offset := r.URL.Query()["offset"]
	arg := database.ListAsociadoParams{}
	if Offset != nil {
		offset, err := strconv.Atoi(Offset[0])
		if err != nil {
			http.Error(w, "error al obtener el offset", 500)
			return
		}
		arg = database.ListAsociadoParams{
			Limit:  10,
			Offset: (int32(offset) - 1) * 10,
		}
	} else {
		arg = database.ListAsociadoParams{
			Limit:  10,
			Offset: (0) * 10,
		}
	}

	asociados, err := s.GetStore().ListAsociado(context.Background(), arg)
	if err != nil {
		http.Error(w, "error al leer los asociados", 500)
		return
	}
	if err != nil {
		http.Error(w, "error al leer los estudios actuales", 500)
		return
	}

	for _, v := range asociados {
		datosSociales, _ := s.GetStore().GetDatosSociales(context.Background(), v.ID)
		cultural, _ := s.GetStore().GetActividadCultural(context.Background(), v.ID)
		actividadDeportiva, _ := s.GetStore().GetActividadDeportiva(context.Background(), v.ID)
		actividadEducativa, _ := s.GetStore().GetActividadEducativa(context.Background(), v.ID)
		estudiosActuales, err := s.GetStore().GetEstudiosActuales(context.Background(), actividadEducativa.IDEstudiosActuales)
		permitir := true
		if err != nil {
			permitir = false
		}
		year, month, day := estudiosActuales.FechaDeGraduacion.Date()
		formattedDate := fmt.Sprintf("%04d-%02d-%02d", year, month, day)
		resp = append(resp, AsociadosContent{
			Asociado:     v,
			DatosSociale: datosSociales,
			Cultural: CulturalContent{
				Actividad:     cultural,
				Participacion: []database.ParticipacionC{},
			},
			Deportiva: DeportivaContent{
				Actividad:     actividadDeportiva,
				Participacion: []database.ParticipacionD{},
			},
			Educativa: EducativaContent{
				Actividad:  actividadEducativa,
				Estudios:   estudiosActuales,
				FormatDate: formattedDate,
				Editar:     permitir,
			},
		})
	}
	s.Data["resp"] = resp

	tpl.ExecuteTemplate(w, "listarAsociado.html", s.Data)
}

func PaginacionNumber(store database.Store) ([]int, error) {
	asociados, err := store.ListarAsociadoAll(context.Background())
	if err != nil {
		return nil, err
	}
	return incrementalSum(len(asociados)), nil
}

func incrementalSum(value int) []int {
	bucle := true
	num := 10
	for bucle {

		if float64(value)/float64(num) > 1 {
			num += 10
		} else {
			bucle = false
		}
	}
	retornar := []int{}
	for i := 0; i < num/10; i++ {
		retornar = append(retornar, i+1)
	}
	return retornar
}
