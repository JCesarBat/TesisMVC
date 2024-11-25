package util

import "net/http"

type Cookie struct {
}

func (c Cookie) SetCookie(value string, w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:     "miCookie",
		Value:    value,
		Path:     "/",
		MaxAge:   3600, // La cookie expirará en una hora
		HttpOnly: true, // Solo accesible a través de HTTP(S), no JavaScript
	}
	http.SetCookie(w, &cookie)

}

func (c Cookie) GetCookie(r *http.Request) string {
	cookie, err := r.Cookie("miCookie")
	if err != nil || cookie.Value == "" {
		return ""
	}
	return cookie.Value
}

func (c Cookie) DeleteCookie(w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:     "miCookie",
		Value:    "",
		Path:     "/",
		MaxAge:   -1, // La cookie se eliminará al expirar
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}
