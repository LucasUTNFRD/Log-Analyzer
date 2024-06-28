package diccionario

type Diccionario[K comparable, V any] interface {
	Guardar(clave K, dato V)
	Pertenece(clave K) bool
	Obtener(clave K) V
	Borrar(clave K) V
	Cantidad() int
	Iterar(func(clave K, dato V) bool)
	Iterador() IterDiccionario[K, V]
	//agregada por un ejercicio ce la guia
	//Claves() TDALista.Lista[K]
}

type IterDiccionario[K comparable, V any] interface {
	//indica si todavia quedan elementos por iterar
	HaySiguiente() bool
	// devuelve la clave y el valor del elemento actual
	VerActual() (K, V)
	//avanza el iterador para que apunte al siguiente elemento
	Siguiente()
}
