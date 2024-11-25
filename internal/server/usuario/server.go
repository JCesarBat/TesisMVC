package usuario

import (
	"TesisMVC/internal/server/common_data"
	"html/template"
)

var tpl = template.Must(template.ParseFiles(""))

type Server struct {
	common_data.GinServer
}
