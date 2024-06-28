package diccionario

type Diccionario[K comparable, V any] interface {
	Guardar(clave K, dato V)
	Pertenece(clave K) bool
	Obtener(clave K) V
	Borrar(clave K) V
	Cantidad() int
	ImprimirInorder()
	Iterar(func(clave K, dato V) bool)
	Iterador() IterDiccionario[K, V]
}

type IterDiccionario[K comparable, V any] interface {
	// 	//indica si todavia quedan elementos por iterar
	HaySiguiente() bool
	// 	// devuelve la clave y el valor del elemento actual
	VerActual() (K, V)
	// 	//avanza el iterador para que apunte al siguiente elemento
	Siguiente()
}
