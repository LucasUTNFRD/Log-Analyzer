package cola_prioridad

type Heap[T any] struct {
	datos    []T
	cantidad int
	cmp      func(T, T) int
}

//La función de comparación, recibe dos claves y devuelve:
//
//Un entero menor que 0 si la primera clave es menor que la segunda.
//Un entero mayor que 0 si la primera clave es mayor que la segunda.
//0 si ambas claves son iguales

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &Heap[T]{datos: []T{}, cmp: funcion_cmp}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	h := &Heap[T]{datos: arreglo, cantidad: len(arreglo), cmp: funcion_cmp}
	h.heapify() // permite darle propiedad de heap en tiempo lineal
	return h
}

func padre(i int) int {
	return i / 2
}
func derecha(i int) int {
	return 2*i + 1
}
func izquierda(i int) int {
	return 2 * i
}

func (h *Heap[T]) swap(i, j int) {
	h.datos[i], h.datos[j] = h.datos[j], h.datos[i]
}

func (h *Heap[T]) downHeap(i int) {
	hijoIzq, hijoDer := izquierda(i), derecha(i)
	maximo := i
	if hijoIzq < h.cantidad && h.cmp(h.datos[hijoIzq], h.datos[maximo]) > 0 {
		maximo = hijoIzq
	}
	if hijoDer < h.cantidad && h.cmp(h.datos[hijoDer], h.datos[maximo]) > 0 {
		maximo = hijoDer
	}
	if maximo != i {
		h.swap(i, maximo)
		h.downHeap(maximo)
	}
}

func (h *Heap[T]) upHeap(index int) {
	parentIndex := padre(index)
	for index > 0 && h.cmp(h.datos[index], h.datos[parentIndex]) > 0 {
		h.swap(index, parentIndex)
		index = parentIndex
		parentIndex = padre(index)
	}
}

func (h *Heap[T]) heapify() {
	for i := (h.cantidad / 2) - 1; i >= 0; i-- {
		h.downHeap(i)
	}
}

// Inteface methods
func (h *Heap[T]) EstaVacia() bool {
	return h.cantidad == 0
}

func (h *Heap[T]) Encolar(elemento T) {
	h.datos = append(h.datos, elemento)
	h.cantidad++
	h.upHeap(h.cantidad - 1)
}

func (h *Heap[T]) VerMax() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	return h.datos[0]
}

func (h *Heap[T]) Desencolar() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	maximo := h.datos[0]
	h.cantidad--
	h.datos[0] = h.datos[h.cantidad]
	h.datos = h.datos[:h.cantidad]
	h.downHeap(0)
	return maximo
}

func (h *Heap[T]) Cantidad() int {
	return h.cantidad
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	h := &Heap[T]{datos: elementos, cantidad: len(elementos), cmp: funcion_cmp}
	h.heapify() // Convert the array into a max-heap

	for i := h.cantidad - 1; i > 0; i-- {
		h.swap(0, i)  // Move current max to the end
		h.cantidad--  // Reduce the heap size
		h.downHeap(0) // Restore the heap property
	}
	h.cantidad = len(elementos) // Restore the original heap size
}
