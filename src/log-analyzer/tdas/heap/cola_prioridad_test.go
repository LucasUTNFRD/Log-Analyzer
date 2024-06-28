package cola_prioridad_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"strings"
	TDAHeap "tdas/heap"
	"testing"
)

func cmpInt(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

func TestCrearHeap(t *testing.T) {
	t.Log("Test que crear heap funciona correctamente")
	h := TDAHeap.CrearHeap[int](cmpInt)
	require.True(t, h.EstaVacia())
	require.Panics(t, func() { h.Desencolar() })
	require.Panics(t, func() { h.VerMax() })
	h.Encolar(5)
	require.EqualValues(t, 5, h.VerMax())
}

func TestCrearHeapArr(t *testing.T) {
	t.Log("Test que crear heap arr funciona correctamente")
	arr := []int{1, 3, 2, 5, 4}
	// 5 ,4 , 2 , 3 ,1 -> heapified
	h := TDAHeap.CrearHeapArr(arr, cmpInt)
	require.False(t, h.EstaVacia())
	require.EqualValues(t, 5, h.VerMax())
}

func TestHeap_Desencolar(t *testing.T) {
	t.Log("Test de metodo desencolar")
	arr := []int{1, 3, 2, 5, 4}
	// 5 ,4 , 2 , 3 ,1 -> heapified
	h := TDAHeap.CrearHeapArr(arr, cmpInt)
	require.False(t, h.EstaVacia())
	require.EqualValues(t, 5, h.VerMax())
	require.EqualValues(t, 5, h.Desencolar())
	require.EqualValues(t, 4, h.Cantidad())
	require.EqualValues(t, 4, h.Desencolar())
	require.EqualValues(t, 3, h.Cantidad())
}

func TestHeap_Encolar(t *testing.T) {
	t.Log("Test de metodo encolar")
	arr := []int{1, 3, 2, 5, 4}
	// 5 ,4 , 2 , 3 ,1 -> heapified
	h := TDAHeap.CrearHeapArr(arr, cmpInt)
	require.False(t, h.EstaVacia())

	elementosParaEncolar := []int{30, 10, 0, 20, 42, 100}
	for i := range elementosParaEncolar {
		h.Encolar(elementosParaEncolar[i])
	}

	require.EqualValues(t, (len(elementosParaEncolar) + len(arr)), h.Cantidad())
	require.EqualValues(t, 100, h.VerMax())
}

func TestHeapDadoSliceDeStrings(t *testing.T) {
	t.Log("test construir heap dado un slice de strings funciona correctamente")
	arr := []string{"hola", "mundo", "cruel", "soy", "una", "cola", "con", "prioridad"}
	h := TDAHeap.CrearHeapArr[string](arr, strings.Compare)
	require.False(t, h.EstaVacia())
	require.EqualValues(t, len(arr), h.Cantidad())
	require.EqualValues(t, "una", h.VerMax())
	require.EqualValues(t, "una", h.Desencolar())
	require.EqualValues(t, "soy", h.VerMax())
}

func TestHeapCreadoConStrings(t *testing.T) {
	t.Log("test de construir heap con elementos del tipo string")
	//arr := []string{"hola", "mundo", "cruel", "soy", "una", "cola", "con", "prioridad"}
	h := TDAHeap.CrearHeap[string](strings.Compare)
	require.True(t, h.EstaVacia())
	require.Panics(t, func() { h.Desencolar() })
	require.Panics(t, func() { h.VerMax() })
	h.Encolar("hola")
	require.EqualValues(t, "hola", h.VerMax())
}

func TestVolumen(t *testing.T) {
	t.Log("Realiza pruebas de volumen para un heap. Encola y desencola una cantidad de elementos grande")
	heap := TDAHeap.CrearHeap(cmpInt)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())

	for i := 0; i < 10000; i++ {
		require.EqualValues(t, i, heap.Cantidad())
		heap.Encolar(i)
		require.False(t, heap.EstaVacia())
		require.EqualValues(t, i, heap.VerMax())
		require.EqualValues(t, i+1, heap.Cantidad())
	}

	for i := 10000; i > 0; i-- {
		require.EqualValues(t, i, heap.Cantidad())
		require.EqualValues(t, i-1, heap.VerMax())
		require.False(t, heap.EstaVacia())
		heap.Desencolar()
	}

	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())
}
func TestHeapSort(t *testing.T) {
	t.Log("Comprueba que el HeapSort ordene un arreglo correctamente")
	arreglo := []int{100, -10, 7, 3, 90, 1}
	esperado := []int{-10, 1, 3, 7, 90, 100}
	TDAHeap.HeapSort(arreglo, cmpInt)
	fmt.Println(arreglo)
	require.EqualValues(t, esperado, arreglo)
}
