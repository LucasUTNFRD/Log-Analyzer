package diccionario

import (
	TDAPila "TP2-Analisis-Logs/src/log-analyzer/tdas/pila"
	"fmt"
)

// Definición del nodo del ABB
type nodoABB[K comparable, V any] struct {
	izq   *nodoABB[K, V]
	der   *nodoABB[K, V]
	clave K
	dato  V
}

// Crear un nuevo nodo del ABB
func crearNodoABB[K comparable, V any](clave K, dato V) *nodoABB[K, V] {
	return &nodoABB[K, V]{clave: clave, dato: dato}
}

// Definición de la función de comparación
type funcCmp[K comparable] func(K, K) int

// Definición del ABB
type abb[K comparable, V any] struct {
	raiz     *nodoABB[K, V]
	cantidad int
	cmp      funcCmp[K]
}

// Crear un nuevo ABB
func CrearABB[K comparable, V any](funcionCmp func(K, K) int) *abb[K, V] {
	return &abb[K, V]{cmp: funcionCmp}
}

// funcion auxiliar para buscarClave
func (arbol *abb[K, V]) buscarClave(clave K, nodo *nodoABB[K, V], padre *nodoABB[K, V]) (*nodoABB[K, V], *nodoABB[K, V]) {
	if nodo == nil {
		return nodo, padre
	}
	if arbol.cmp(nodo.clave, clave) < 0 {
		return arbol.buscarClave(clave, nodo.der, nodo)
	} else if arbol.cmp(nodo.clave, clave) > 0 {
		return arbol.buscarClave(clave, nodo.izq, nodo)
	} else {
		return nodo, padre
	}
}

// Guardar un dato en el ABB
func (arbol *abb[K, V]) Guardar(clave K, dato V) {
	nuevo := crearNodoABB(clave, dato)
	if arbol.raiz == nil {
		arbol.raiz = nuevo
		arbol.cantidad++
		return
	}
	nodo, padre := arbol.buscarClave(clave, arbol.raiz, nil)
	if nodo == nil {
		if arbol.cmp(padre.clave, clave) < 0 {
			padre.der = nuevo
		} else {
			padre.izq = nuevo
		}
		arbol.cantidad++
	} else {
		nodo.dato = dato
	}
}

func (arbol *abb[K, V]) Obtener(clave K) V {
	nodo, _ := arbol.buscarClave(clave, arbol.raiz, nil)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	} else {
		return nodo.dato
	}
}

func (arbol *abb[K, V]) Pertenece(clave K) bool {
	nodo, _ := arbol.buscarClave(clave, arbol.raiz, nil)
	return nodo != nil
}

func (arbol *abb[K, V]) inorder(nodo *nodoABB[K, V]) {
	if nodo == nil {
		return
	}
	arbol.inorder(nodo.izq)
	fmt.Printf("Dato: %v,Clave: %v \n", nodo.clave, nodo.dato)
	arbol.inorder(nodo.der)

}

func (nodo *nodoABB[K, V]) cantHijos() int {
	if nodo.izq == nil && nodo.der == nil {
		return 0
	}
	if (nodo.izq == nil && nodo.der != nil) || (nodo.izq != nil && nodo.der == nil) {
		return 1
	}
	return 2
}

func (arbol *abb[K, V]) Borrar(clave K) V {
	nodo, padre := arbol.buscarClave(clave, arbol.raiz, nil)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	cant_hijos := nodo.cantHijos()
	borrado := nodo.dato

	if cant_hijos == 0 {
		if padre == nil {
			arbol.raiz = nil
		} else if arbol.cmp(nodo.clave, padre.clave) < 0 {
			padre.izq = nil
		} else {
			padre.der = nil
		}
	} else if cant_hijos == 1 {
		if nodo.izq != nil {
			if padre == nil {
				arbol.raiz = nodo.izq
			} else if arbol.cmp(nodo.clave, padre.clave) < 0 {
				padre.izq = nodo.izq
			} else {
				padre.der = nodo.izq
			}
		} else {
			if padre == nil {
				arbol.raiz = nodo.der
			} else if arbol.cmp(nodo.clave, padre.clave) < 0 {
				padre.izq = nodo.der
			} else {
				padre.der = nodo.der
			}
		}
	} else {
		reemplazante := nodo.izq.sucesor()
		claveRemplazante := reemplazante.clave
		valorReemplazante := arbol.Borrar(reemplazante.clave)
		arbol.cantidad++ //este llamado de Borrar reemplazante nos restó 1 en cantidad
		nodo.clave, nodo.dato = claveRemplazante, valorReemplazante
	}
	arbol.cantidad--
	return borrado
}

func (arbol *abb[K, V]) ImprimirInorder() {
	arbol.inorder(arbol.raiz)
}

// funcion auxiliar para encontrar con que nodo debo remplazar en el borrado
func (nodo *nodoABB[K, V]) sucesor() *nodoABB[K, V] {
	if nodo.der == nil {
		return nodo
	}
	return nodo.der.sucesor()
}

// Obtener la cantidad de nodos en el ABB
func (arbol *abb[K, V]) Cantidad() int {
	return arbol.cantidad
}

// definicion del iterador del ABB
type iterABB[K comparable, V any] struct {
	pila         *TDAPila.Pila[*nodoABB[K, V]]
	actual       *nodoABB[K, V]
	abb          *abb[K, V]
	desde, hasta *K
	cmp          funcCmp[K]
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return abb.IteradorRango(nil, nil)
}

func (abb *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	iterarwrapper(abb.raiz, visitar)
}

func iterarwrapper[K comparable, V any](nodo *nodoABB[K, V], visitar func(clave K, dato V) bool) bool {
	if nodo == nil {
		return true
	}
	if !iterarwrapper(nodo.izq, visitar) || !visitar(nodo.clave, nodo.dato) || !iterarwrapper(nodo.der, visitar) {
		return false
	}
	return true
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	iterarRangoWrapper(abb.raiz, desde, hasta, visitar, abb.cmp)
}

func iterarRangoWrapper[K comparable, V any](nodo *nodoABB[K, V], desde, hasta *K, visitar func(clave K, dato V) bool, cmp func(K, K) int) bool {
	if nodo == nil {
		return true
	}
	siguiente := true
	if desde == nil || cmp(*desde, nodo.clave) < 0 {
		siguiente = iterarRangoWrapper(nodo.izq, desde, hasta, visitar, cmp)
	}
	if siguiente && (desde == nil || cmp(*desde, nodo.clave) <= 0) && (hasta == nil || cmp(*hasta, nodo.clave) >= 0) {
		siguiente = visitar(nodo.clave, nodo.dato)
	}
	if !siguiente {
		return false
	}
	if hasta == nil || cmp(*hasta, nodo.clave) > 0 {
		siguiente = iterarRangoWrapper(nodo.der, desde, hasta, visitar, cmp)
	}
	return siguiente
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	pila := TDAPila.CrearPilaDinamica[*nodoABB[K, V]]()
	abb.raiz.rango(&pila, desde, hasta, abb.cmp)
	iterador := &iterABB[K, V]{abb: abb, pila: &pila, desde: desde, hasta: hasta, cmp: abb.cmp}
	return iterador
}

func (nodo *nodoABB[K, V]) rango(pila *TDAPila.Pila[*nodoABB[K, V]], desde, hasta *K, cmp func(K, K) int) {
	if nodo == nil {
		return
	}
	if desde == nil || cmp(*desde, nodo.clave) <= 0 {
		(*pila).Apilar(nodo)
		nodo.izq.rango(pila, desde, hasta, cmp)
	} else if hasta == nil || cmp(*hasta, nodo.clave) > 0 {
		nodo.der.rango(pila, desde, hasta, cmp)
	}
}

// Primitivas del iterador
func (iter *iterABB[K, V]) HaySiguiente() bool {
	return !(*iter.pila).EstaVacia() && (iter.hasta == nil || iter.cmp(*iter.hasta, (*iter.pila).VerTope().clave) >= 0)
}

func (iter *iterABB[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := (*iter.pila).VerTope()
	return nodo.clave, nodo.dato
}

func (iter *iterABB[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := (*iter.pila).Desapilar()
	if nodo.der == nil {
		return
	}
	nodo.der.rango(iter.pila, iter.desde, iter.hasta, iter.cmp)
}
