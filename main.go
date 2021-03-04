package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Arbol struct {
	raiz *Nodo
}

type Nodo struct {
	dato     int
	der, izq *Nodo
}

func NewArbol() *Arbol {
	return &Arbol{raiz: nil}
}

func NewNodo(dato int) *Nodo {
	return &Nodo{dato: dato, der: nil, izq: nil}
}

func (arbol *Arbol) Insertar(dato int) {
	arbol.raiz = Insertar(dato, arbol.raiz)
}

func Insertar(dato int, nodo *Nodo) *Nodo {
	if nodo == nil {
		return NewNodo(dato)
	}
	if dato < nodo.dato {
		nodo.izq = Insertar(dato, nodo.izq)
	} else {
		nodo.der = Insertar(dato, nodo.der)
	}
	return nodo
}

func PreOrden(raiz *Nodo) string {
	var cadena strings.Builder
	if raiz == nil {
		return ""
	}
	cadena.WriteString(strconv.Itoa(raiz.dato))
	cadena.WriteString(", ")
	cadena.WriteString(PreOrden(raiz.izq))
	cadena.WriteString(PreOrden(raiz.der))
	return cadena.String()
}

func InOrden(raiz *Nodo) string {
	var cadena strings.Builder
	if raiz == nil {
		return ""
	}
	cadena.WriteString(InOrden(raiz.izq))
	cadena.WriteString(strconv.Itoa(raiz.dato))
	cadena.WriteString(", ")
	cadena.WriteString(InOrden(raiz.der))
	return cadena.String()
}

func PostOrden(raiz *Nodo) string {
	var cadena strings.Builder
	if raiz == nil {
		return ""
	}
	cadena.WriteString(PostOrden(raiz.izq))
	cadena.WriteString(PostOrden(raiz.der))
	cadena.WriteString(strconv.Itoa(raiz.dato))
	cadena.WriteString(", ")
	return cadena.String()
}

func (arbol *Arbol) Recorrer() string {
	var cadena strings.Builder
	cadena.WriteString("***RECORRIDO PREORDEN***\n")
	cadena.WriteString(PreOrden(arbol.raiz))
	cadena.WriteString("\n\n")
	cadena.WriteString("***RECORRIDO INORDEN***\n")
	cadena.WriteString(InOrden(arbol.raiz))
	cadena.WriteString("\n\n")
	cadena.WriteString("***RECORRIDO POSTORDEN***\n")
	cadena.WriteString(PostOrden(arbol.raiz))
	cadena.WriteString("\n\n")
	return cadena.String()
}

func main() {
	arbol := NewArbol()
	arbol.Insertar(5)
	arbol.Insertar(8)
	arbol.Insertar(2)
	arbol.Insertar(1)
	arbol.Insertar(4)
	arbol.Insertar(3)
	arbol.Insertar(9)
	arbol.Insertar(10)
	arbol.Insertar(6)
	arbol.Insertar(7)
	fmt.Println(arbol.Recorrer())

}
