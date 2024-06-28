package pila_test

import (
	"github.com/stretchr/testify/require"
	TDAPila "tdas/pila"
	"testing"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	// mas pruebas para este caso...
}

func TestApilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	// Apilar algunos elementos y luego verificar si la pila no está vacía
	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(3)
	require.False(t, pila.EstaVacia())

	// Verificar que el tope de la pila sea 3
	require.Equal(t, 3, pila.VerTope())

	// Desapilar elementos y verificar que el tope cambia correctamente
	require.Equal(t, 3, pila.Desapilar())
	require.Equal(t, 2, pila.VerTope())
	require.Equal(t, 2, pila.Desapilar())
	require.Equal(t, 1, pila.VerTope())
	require.Equal(t, 1, pila.Desapilar())
	require.True(t, pila.EstaVacia())
}
func TestDesapilar(t *testing.T) {
	// Crear una nueva pila dinámica de enteros
	pila := TDAPila.CrearPilaDinamica[int]()

	// Apilar algunos elementos
	pila.Apilar(10)
	pila.Apilar(20)
	pila.Apilar(30)

	// Verificar que la pila no está vacía
	require.False(t, pila.EstaVacia())

	// Verificar el tope de la pila
	require.Equal(t, 30, pila.VerTope())

	// Desapilar un elemento y verificar
	require.Equal(t, 30, pila.Desapilar())
	require.Equal(t, 20, pila.VerTope())

	// Desapilar otro elemento y verificar
	require.Equal(t, 20, pila.Desapilar())
	require.Equal(t, 10, pila.VerTope())

	// Desapilar el último elemento y verificar que la pila esté vacía
	require.Equal(t, 10, pila.Desapilar())
	require.True(t, pila.EstaVacia())

	// Intentar desapilar un elemento de una pila vacía debería causar un pánico
	require.Panics(t, func() { pila.Desapilar() })
}
