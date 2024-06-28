package pila

const capacidadInicial = 10 // Tamaño inicial de la pila

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = nil
	pila.cantidad = 0
	return pila
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	return p.datos[p.cantidad-1]
}

func (p *pilaDinamica[T]) Apilar(valor T) {
	if p.datos == nil {
		p.datos = make([]T, capacidadInicial)
	}
	if p.cantidad >= len(p.datos) {
		// Resize the slice
		newDatos := make([]T, p.cantidad*2) // Double the capacity
		copy(newDatos, p.datos)
		p.datos = newDatos
	}
	p.datos[p.cantidad] = valor
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	valor := p.datos[p.cantidad-1]
	p.cantidad--
	if len(p.datos)/2 >= p.cantidad && len(p.datos) > capacidadInicial {
		// Si la cantidad de elementos es la mitad de la capacidad actual o menos, y
		// la capacidad actual es mayor que la capacidad inicial, se reduce a la mitad
		nuevaCapacidad := len(p.datos) / 2
		nuevosDatos := make([]T, len(p.datos), nuevaCapacidad)
		copy(nuevosDatos, p.datos)
		p.datos = nuevosDatos
	}
	return valor
}
