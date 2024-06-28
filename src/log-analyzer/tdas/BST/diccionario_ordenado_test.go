package diccionario_test

import (
	"github.com/stretchr/testify/require"
	TDADiccionario "tdas/BST"
	"testing"
)

func compararCadenas(cadena1, cadena2 string) int {
	if cadena1 > cadena2 {
		return 1
	} else if cadena1 < cadena2 {
		return -1
	} else {
		return 0
	}
}

func TestDiccionarioVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDADiccionario.CrearABB[string, string](compararCadenas)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.Panics(t, func() { dic.Obtener("A") })
	require.Panics(t, func() { dic.Borrar("A") })
}

func TestDiccionarioGuardar(t *testing.T) {
	t.Log("Compruebe que primitiva guardar funciona")
	dic := TDADiccionario.CrearABB[string, string](compararCadenas)
	require.EqualValues(t, 0, dic.Cantidad())
	dic.Guardar("lucas", "lucas")
	require.EqualValues(t, 1, dic.Cantidad())
	dic.Guardar("delfina", "delfina")
	require.EqualValues(t, 2, dic.Cantidad())

}

func TestDiccionarioReemplazar(t *testing.T) {
	t.Log("Compruebe que primitiva guardar funciona correctamente cuando se quiere guardar con la misma clave")
	dic := TDADiccionario.CrearABB[string, string](compararCadenas)
	require.EqualValues(t, 0, dic.Cantidad())
	dic.Guardar("lucas", "lucas")
	require.EqualValues(t, 1, dic.Cantidad())
	dic.Guardar("delfina", "delfina")
	require.EqualValues(t, 2, dic.Cantidad())
	dic.Guardar("lucas", "delfina")
	require.EqualValues(t, 2, dic.Cantidad())
	dic.Guardar("martin", "martin")
	require.EqualValues(t, 3, dic.Cantidad())
}

func TestDiccVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDADiccionario.CrearABB[string, string](compararCadenas)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
}

func TestDiccGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Perro"
	clave2 := "Gato"
	clave3 := "Vaca"
	valor1 := "guau"
	valor2 := "miau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := TDADiccionario.CrearABB[string, string](compararCadenas)
	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[1], valores[1])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestBorrado(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Perro"
	clave2 := "Gato"
	clave3 := "Vaca"
	valor1 := "guau"
	valor2 := "miau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](compararCadenas)

	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])
	dic.ImprimirInorder()
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[2]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[2]))

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[0]) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[0]) })

	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[1]) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[1]) })
}

func TestAbbIteradorInternoClaves(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	abb := TDADiccionario.CrearABB[string, *int](cmpString)
	abb.Guardar(claves[0], nil)
	abb.Guardar(claves[1], nil)
	abb.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad

	abb.Iterar(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
	require.NotEqualValues(t, -1, buscarAbb(cs[0], claves))
	require.NotEqualValues(t, -1, buscarAbb(cs[1], claves))
	require.NotEqualValues(t, -1, buscarAbb(cs[2], claves))
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])
}

func buscarAbb(clave string, claves []string) int {
	for i := 0; i < len(claves); i++ {
		if claves[i] == clave {
			return i
		}
	}
	return -1
}

func TestAbbIteradorInternoValores(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	abb := TDADiccionario.CrearABB[string, int](cmpString)
	abb.Guardar(clave1, 6)
	abb.Guardar(clave2, 2)
	abb.Guardar(clave3, 3)
	abb.Guardar(clave4, 4)
	abb.Guardar(clave5, 5)

	factorial := 1
	ptrFactorial := &factorial
	abb.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestAbbIteradorInternoValoresConBorrados(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno, sin recorrer datos borrados")
	clave0 := "Elefante"
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	abb := TDADiccionario.CrearABB[string, int](cmpString)
	abb.Guardar(clave0, 7)
	abb.Guardar(clave1, 6)
	abb.Guardar(clave2, 2)
	abb.Guardar(clave3, 3)
	abb.Guardar(clave4, 4)
	abb.Guardar(clave5, 5)

	abb.Borrar(clave0)

	factorial := 1
	ptrFactorial := &factorial
	abb.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func cmpString(s1, s2 string) int {
	if s1 == s2 {
		return 0
	} else if s1 < s2 {
		return -1
	}
	return 1
}

func compararNumeros(i1, i2 int) int {
	if i1 == i2 {
		return 0
	} else if i1 < i2 {
		return -1
	}
	return 1
}

func TestIteradorRangoConHasta(t *testing.T) {
	t.Log("Prueba el iterador externo solo con el limite HASTA definido")
	dic := TDADiccionario.CrearABB[int, int](compararNumeros)
	dic.Guardar(7, 7)
	dic.Guardar(3, 3)
	dic.Guardar(8, 8)
	dic.Guardar(2, 2)
	dic.Guardar(5, 5)
	hasta := 7
	iter := dic.IteradorRango(nil, &hasta)

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.EqualValues(t, 2, primero)

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	segundo, _ := iter.VerActual()
	require.EqualValues(t, 3, segundo)

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.EqualValues(t, 5, tercero)

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	cuarto, _ := iter.VerActual()
	require.EqualValues(t, 7, cuarto)

	iter.Siguiente()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })

}

func TestIteradorRangoConDesde(t *testing.T) {
	t.Log("Prueba el iterador externo solo con el limite DESDE definido")
	dic := TDADiccionario.CrearABB[int, int](compararNumeros)
	dic.Guardar(7, 7)
	dic.Guardar(3, 3)
	dic.Guardar(8, 8)
	dic.Guardar(2, 2)
	dic.Guardar(5, 5)
	desde := 3
	iter := dic.IteradorRango(&desde, nil)

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.EqualValues(t, 3, primero)

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	segundo, _ := iter.VerActual()
	require.EqualValues(t, 5, segundo)

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.EqualValues(t, 7, tercero)

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	cuarto, _ := iter.VerActual()
	require.EqualValues(t, 8, cuarto)

	iter.Siguiente()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })

}

func TestIteradorInternoRangos(t *testing.T) {
	t.Log("Prueba el iterador interno con un rango definido")
	dic := TDADiccionario.CrearABB[int, int](compararNumeros)
	dic.Guardar(7, 7)
	dic.Guardar(3, 3)
	dic.Guardar(8, 8)
	dic.Guardar(2, 2)
	dic.Guardar(5, 5)
	desde := 3
	hasta := 7
	factorial := 1
	ptrFactorial := &factorial
	dic.IterarRango(&desde, &hasta, func(_ int, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 105, factorial)
}

func TestIteradorInternoRangosSinHasta(t *testing.T) {
	t.Log("Prueba el iterador interno con un rango sin desde (desde = nil)")
	dic := TDADiccionario.CrearABB[int, int](compararNumeros)
	dic.Guardar(7, 7)
	dic.Guardar(3, 3)
	dic.Guardar(8, 8)
	dic.Guardar(2, 2)
	dic.Guardar(5, 5)
	desde := 5
	factorial := 1
	ptrFactorial := &factorial
	dic.IterarRango(&desde, nil, func(_ int, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 280, factorial)
}
