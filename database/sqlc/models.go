// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type UltimoGrado string

const (
	UltimoGradoPrimaria         UltimoGrado = "Primaria"
	UltimoGradoSecundaria       UltimoGrado = "Secundaria"
	UltimoGradoPreUniversitario UltimoGrado = "Pre Universitario"
	UltimoGradoUniversitario    UltimoGrado = "Universitario"
)

func (e *UltimoGrado) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UltimoGrado(s)
	case string:
		*e = UltimoGrado(s)
	default:
		return fmt.Errorf("unsupported scan type for UltimoGrado: %T", src)
	}
	return nil
}

type NullUltimoGrado struct {
	UltimoGrado UltimoGrado `json:"ultimo_grado"`
	Valid       bool        `json:"valid"` // Valid is true if UltimoGrado is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUltimoGrado) Scan(value interface{}) error {
	if value == nil {
		ns.UltimoGrado, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UltimoGrado.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUltimoGrado) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UltimoGrado), nil
}

type ActividadCultural struct {
	ID           int64    `json:"id"`
	IDAsociado   int64    `json:"id_asociado"`
	Especialidad []string `json:"especialidad"`
}

type ActividadDeportiva struct {
	ID                int64    `json:"id"`
	IDAsociado        int64    `json:"id_asociado"`
	AficcionOPractica []string `json:"aficcion_o_practica"`
}

type ActividadEducativa struct {
	ID                  int64       `json:"id"`
	IDAsociado          int64       `json:"id_asociado"`
	IDEstudiosActuales  int64       `json:"id_estudios_actuales"`
	UltimoGradoAprobado UltimoGrado `json:"ultimo_grado_aprobado"`
}

type Asociado struct {
	ID                  int64          `json:"id"`
	Nombre              string         `json:"nombre"`
	Apellido1           string         `json:"apellido1"`
	Apellido2           string         `json:"apellido2"`
	Activo              bool           `json:"activo"`
	Carnet              string         `json:"carnet"`
	Sexo                bool           `json:"sexo"`
	NumeroT             sql.NullString `json:"numeroT"`
	NumeroPerteneciente sql.NullString `json:"numeroPerteneciente"`
	Direccion           string         `json:"direccion"`
	IDMunicipio         int64          `json:"id_municipio"`
}

type DatosSociale struct {
	ID                        int64          `json:"id"`
	IDAsociado                int64          `json:"id_asociado"`
	Ocupacion                 sql.NullString `json:"ocupacion"`
	EstadoCivil               sql.NullString `json:"estado_civil"`
	IntegracionRevolucionaria sql.NullString `json:"integracion_revolucionaria"`
}

type EstudiosActuale struct {
	ID                    int64     `json:"id"`
	TipoEnseñansa         string    `json:"tipo_enseñansa"`
	Centro                string    `json:"centro"`
	EspecialidadGradoOAño string    `json:"especialidad_grado_o_año"`
	AñoDelDato            time.Time `json:"año_del_dato"`
	FechaDeGraduacion     time.Time `json:"fecha_de_graduacion"`
}

type File struct {
	ID       int64  `json:"ID"`
	ProvID   int64  `json:"prov_id"`
	FileName string `json:"file_name"`
	FileType string `json:"file_type"`
	FilePath string `json:"file_path"`
}

type Municipio struct {
	ID          int64  `json:"id"`
	IDProvincia int64  `json:"id_provincia"`
	Name        string `json:"name"`
}

type ParticipacionC struct {
	ID                  int64         `json:"id"`
	IDActividadCultural sql.NullInt64 `json:"id_actividad_cultural"`
	Especialidad        string        `json:"especialidad"`
	Fecha               time.Time     `json:"fecha"`
	LugarAlcanzado      sql.NullInt32 `json:"lugar_alcanzado"`
	DondeSeDesarrollo   string        `json:"donde_se_desarrollo"`
}

type ParticipacionD struct {
	ID                   int64         `json:"id"`
	IDActividadDeportiva sql.NullInt64 `json:"id_actividad_deportiva"`
	Deporte              string        `json:"deporte"`
	Fecha                time.Time     `json:"fecha"`
	LugarAlcanzado       sql.NullInt32 `json:"lugar_alcanzado"`
	DondeSeDesarrollo    string        `json:"donde_se_desarrollo"`
}

type Provincium struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Rol struct {
	ID  int64  `json:"id"`
	Rol string `json:"rol"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	UserID       int64     `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type User struct {
	ID          int64        `json:"id"`
	IDMunicipio int64        `json:"id_municipio"`
	Username    string       `json:"username"`
	Password    string       `json:"password"`
	Email       string       `json:"email"`
	SuperUser   sql.NullBool `json:"super_user"`
	CreatedAt   time.Time    `json:"created_at"`
}

type UserRole struct {
	ID     int64 `json:"id"`
	UserID int64 `json:"user_id"`
	RolID  int64 `json:"rol_id"`
}
