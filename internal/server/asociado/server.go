package asociado

import (
	"TesisMVC/internal/server/common_data"
	"html/template"
)

var tpl = template.Must(template.ParseFiles("internal/template/asociado/listarAsociado.html"))

type Server struct {
	common_data.GinServer
	Data map[string]any
}

func (s *Server) ResetData() {
	s.Data = make(map[string]any)

}
