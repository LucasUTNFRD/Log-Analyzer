package cola

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	cola := new(colaEnlazada[T])
	cola.primero = nil
	cola.ultimo = nil
	return cola
}

func (cola *colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil && cola.ultimo == nil
}

func (cola *colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("la cola esta vacia")
	}
	return cola.primero.dato
}

func (cola *colaEnlazada[T]) Encolar(elem T) {
	nodo := crearNodo(elem)
	if cola.EstaVacia() {
		cola.primero = nodo
		cola.ultimo = nodo
	} else {
		cola.ultimo.prox = nodo
		cola.ultimo = nodo
	}
}

func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("la cola esta vacia")
	}
	dato := cola.primero.dato
	cola.primero = cola.primero.prox
	if cola.primero == nil {
		cola.ultimo = nil
	}
	return dato

}

func crearNodo[T any](dato T) *nodoCola[T] {
	nodo := new(nodoCola[T])
	nodo.prox = nil
	nodo.dato = dato
	return nodo
}
