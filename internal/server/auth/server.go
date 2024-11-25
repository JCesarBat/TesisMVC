package auth

import (
	"TesisMVC/internal/server/common_data"
	"html/template"
)

var tpl = template.Must(template.ParseFiles("internal/template/login.html", "internal/template/register.html"))

type Server struct {
	common_data.GinServer
}
