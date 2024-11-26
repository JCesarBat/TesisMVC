package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

func init() {

	rand.Seed(time.Now().UnixNano())

}
func generateRandomAsociado() InsertAsoiciadoParams {
	nombres := []string{"Juan", "María", "Pedro", "Ana", "Luis", "Carmen"}
	apellidos := []string{"García", "Rodríguez", "Martínez", "López", "González", "Pérez"}

	return InsertAsoiciadoParams{

		Nombre: nombres[rand.Intn(len(nombres))],

		Apellido1:           apellidos[rand.Intn(len(apellidos))],
		Apellido2:           apellidos[rand.Intn(len(apellidos))],
		Activo:              rand.Intn(2) == 1,
		Carnet:              "01241321531",
		Sexo:                rand.Intn(2) == 1,
		NumeroT:             sql.NullString{String: "dqw34135246", Valid: rand.Intn(2) == 1},
		NumeroPerteneciente: sql.NullString{String: fmt.Sprintf("NP%d", rand.Intn(99999999-10000000+1)+10000000), Valid: rand.Intn(2) == 1},
		Direccion:           fmt.Sprintf("Calle %d, Ciudad %d", rand.Intn(100), rand.Intn(100)),
		IDMunicipio:         rand.Int63n(140),
	}
}

func TestCreateAsociado(t *testing.T) {
	param := generateRandomAsociado()
	for i := 0; i < 10; i++ {
		_, err := testStore.InsertAsoiciado(context.Background(), param)
		require.NoError(t, err)
	}

}
