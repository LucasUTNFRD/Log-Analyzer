package diccionario_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	TDADiccionario "tdas/hash"
	"testing"
)

var TAMS_VOLUMEN = []int{12500, 25000, 50000, 100000, 200000, 400000}

func TestDiccionarioVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDADiccionario.CrearHash[string, string]()
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.Panics(t, func() { dic.Obtener("A") })
	require.Panics(t, func() { dic.Borrar("A") })
}

func TestDiccionarioAgregar(t *testing.T) {
	t.Log("Comprueba que se agrega elemento correctamente al diccionario")
	dic := TDADiccionario.CrearHash[string, uint]()
	require.EqualValues(t, 0, dic.Cantidad())
	dic.Guardar("lucas", 18)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("lucas"))
	require.EqualValues(t, 18, dic.Obtener("lucas"))
}

func TestDiccionarioBorrar(t *testing.T) {
	t.Log("Comprueba que se borra correctamente un elemento en el diccionario")
	dic := TDADiccionario.CrearHash[string, uint]()
	require.EqualValues(t, 0, dic.Cantidad())
	dic.Guardar("lucas", 18)
	require.EqualValues(t, 1, dic.Cantidad())
	fmt.Println(dic.Cantidad())
	elementoBorrado := dic.Borrar("lucas")
	require.EqualValues(t, 18, elementoBorrado)
	require.False(t, dic.Pertenece("lucas"))
	require.Panics(t, func() { dic.Obtener("lucas") })
	fmt.Println(dic.Cantidad())

}
func TestDiccionarioClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un Hash vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")
	dic := TDADiccionario.CrearHash[string, string]()
	require.False(t, dic.Pertenece(""))
	require.PanicsWithValue(t, "La clave no se encuentra en el diccionario", func() { dic.Obtener("") })
	require.PanicsWithValue(t, "La clave no se encuentra en el diccionario", func() { dic.Borrar("") })

	dicNum := TDADiccionario.CrearHash[int, string]()
	require.False(t, dicNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no se encuentra en el diccionario", func() { dicNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no se encuentra en el diccionario", func() { dicNum.Borrar(0) })
}

func TestDiccionarioGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "lucas"
	clave2 := "trini"
	clave3 := "marcelo"
	valor1 := "hijo"
	valor2 := "hija"
	valor3 := "padre"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := TDADiccionario.CrearHash[string, string]()
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

func TestReemplazoDato(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	dic := TDADiccionario.CrearHash[string, string]()
	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "miau", dic.Obtener(clave))
	require.EqualValues(t, "guau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}

func TestDiccionarioDeStruct(t *testing.T) {
	t.Log("valida que el diccionario funcione con structs")
	type structTest struct {
		a string
		b string
	}
	dic := TDADiccionario.CrearHash[structTest, string]()
	test1 := structTest{a: "hola", b: "mundo"}
	test2 := structTest{a: "hola", b: "lucas"}
	test3 := structTest{a: "hola", b: "soy un test"}

	dic.Guardar(test1, "test 1 ")
	dic.Guardar(test2, "test 2 ")
	dic.Guardar(test3, "test 3 ")

	require.True(t, dic.Pertenece(test1))
	require.True(t, dic.Pertenece(test2))
	require.True(t, dic.Pertenece(test3))
}

func TestClaveVacia(t *testing.T) {
	t.Log("Pruebo que el diccionario puede tomar como clave una clave vacia")
	dic := TDADiccionario.CrearHash[string, string]()
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
}

func TestValorNil(t *testing.T) {
	t.Log("Pruebo que  se puede asociar como dato un valor nil")
	dic := TDADiccionario.CrearHash[string, *string]()
	clave := "test"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, (*string)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*string)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestValorSlice(t *testing.T) {
	t.Log("Pruebo que se puede asociar como valor un slice")
	dic := TDADiccionario.CrearHash[string, []string]()
	testString := []string{"hola", "mundo", "cruel", "vacio"}
	dic.Guardar("hola mundo", testString)
	require.True(t, dic.Pertenece("hola mundo"))
	require.EqualValues(t, testString, dic.Obtener("hola mundo"))
}
func TestCadenaLargaParticular(t *testing.T) {
	t.Log("Se han visto casos problematicos al utilizar la funcion de hashing de K&R, por lo que " +
		"se agrega una prueba con dicha funcion de hashing y una cadena muy larga")
	// El caracter '~' es el de mayor valor en ASCII (126).
	claves := make([]string, 10)
	cadena := "%d~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~" +
		"~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"
	dic := TDADiccionario.CrearHash[string, string]()
	valores := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	for i := 0; i < 10; i++ {
		claves[i] = fmt.Sprintf(cadena, i)
		dic.Guardar(claves[i], valores[i])
	}
	require.EqualValues(t, 10, dic.Cantidad())

	ok := true
	for i := 0; i < 10 && ok; i++ {
		ok = dic.Obtener(claves[i]) == valores[i]
	}

	require.True(t, ok, "Obtener clave larga funciona")
}

func TestIteradorConDiccVacio(t *testing.T) {
	t.Log("Si creo un iterador y el hash esta vacio, el iterador va a estar al final")
	dic := TDADiccionario.CrearHash[string, int]()
	require.EqualValues(t, 0, dic.Cantidad())
	iterDic := dic.Iterador()
	require.False(t, iterDic.HaySiguiente())
	require.Panics(t, func() { iterDic.Siguiente() })
	require.Panics(t, func() { iterDic.VerActual() })

}
func ejecutarPruebaVolumen(b *testing.B, n int) {
	dic := TDADiccionario.CrearHash[string, int]()

	claves := make([]string, n)
	valores := make([]int, n)

	/* Inserta 'n' parejas en el hash */
	for i := 0; i < n; i++ {
		valores[i] = i
		claves[i] = fmt.Sprintf("%08d", i)
		dic.Guardar(claves[i], valores[i])
	}

	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que devuelva los valores correctos */
	ok := true
	for i := 0; i < n; i++ {
		ok = dic.Pertenece(claves[i])
		if !ok {
			break
		}
		ok = dic.Obtener(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que borre y devuelva los valores correctos */
	for i := 0; i < n; i++ {
		ok = dic.Borrar(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, dic.Cantidad())
}

func BenchmarkDiccionario(b *testing.B) {
	b.Log("Prueba de stress del Diccionario. Prueba guardando distinta cantidad de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad " +
		"sea la adecuada. Luego validamos que podemos obtener y ver si pertenece cada una de las claves geeneradas, " +
		"y que luego podemos borrar sin problemas")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumen(b, n)
			}
		})
	}
}
