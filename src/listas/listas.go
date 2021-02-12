package listas

import "fmt"

type Msg struct {
	Fecha string
	Texto string
}

type Mensaje struct {
	Origen  string
	Destino string
	Msgs    []Msg
}

type Nodo struct {
	Info Mensaje

	Next *Nodo
	Prev *Nodo
}

type Lista struct {
	first, last *Nodo
	size        int
}

func NuevaLista() *Lista {
	return &Lista{nil, nil, 0}
}

func (this *Lista) Insertar(nuevo *Nodo) {
	if this.first == nil {
		this.first = nuevo
		this.last = nuevo
	} else {
		this.last.Next = nuevo
		nuevo.Prev = this.last
		this.last = nuevo
	}
}

func (this *Lista) ToString() string {
	var texto string
	aux := this.first
	for aux != nil {
		texto += aux.ToString()
		aux = aux.Next
	}
	return texto
}

func (this *Nodo) ToString() string {
	var texto string
	texto = "Origen: " + this.Info.Origen + "\n" +
		"Destino: " + this.Info.Destino + "\n" + "Mensajes:\n"
	for i := 0; i < len(this.Info.Msgs); i++ {
		texto += "[" + this.Info.Msgs[i].Fecha + "] " + this.Info.Msgs[i].Texto + "\n\n"
	}
	return texto
}

func (this *Lista) Print() {
	fmt.Println("---LISTA DE MENSAJES---\n")
	this.ToString()
}
