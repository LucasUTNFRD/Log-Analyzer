package cola_test

import (
	"github.com/stretchr/testify/require"
	TDACola "tdas/cola"
	"testing"
)

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	// mas pruebas para este caso...
}

func TestColaCargar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	// mas pruebas para este caso...
	for i := range 10 {
		cola.Encolar(i)
	}
	elem := cola.VerPrimero()
	require.True(t, elem == 0)
}

func TestColaDesencolarPanic(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.Panics(t, func() { cola.Desencolar() })
}
