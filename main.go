package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

type nodo struct {
	nombre, apellido, apodo, favoritos string
	Siguiente, Anterior                *nodo
}

type lista struct {
	cabeza *nodo
	cola   *nodo
}

func (this *lista) Insertar(nuevo *nodo) {
	if this.cabeza == nil {
		this.cabeza = nuevo
		this.cola = nuevo
	} else {
		this.cola.Siguiente = nuevo
		nuevo.Anterior = this.cola
		this.cola = nuevo
	}
}

func main() {
	li := lista{nil, nil}
	a := nodo{"Marvin", "Martinez", "Marvin25ronal", "Jugar apex", nil, nil}
	b := nodo{"Yaiza", "Pineda", "Bambi", "Patinar", nil, nil}
	c := nodo{"Jonathan", "Lopez", "Pancho", "Comer", nil, nil}
	d := nodo{"usuario1", "bla", "bla", "Jugar apex", nil, nil}
	e := nodo{"usuario2", "bla", "bla", "Jugar apex", nil, nil}
	f := nodo{"usuario3", "sale edd", "vamos con todo", "100 en la fase 1", nil, nil}
	li.Insertar(&a)
	li.Insertar(&b)
	li.Insertar(&c)
	li.Insertar(&d)
	li.Insertar(&e)
	li.Insertar(&f)
	li.graficar()
}

func (this *lista) graficar() {
	var codigo, conexiones string
	var contador int = 0
	codigo = "digraph G{\n"
	if this.cabeza != nil {
		aux := this.cabeza
		for aux != nil {
			codigo += "nodo" + strconv.Itoa(contador) + "[shape=Mrecord, color=" +
				"\"#00bf0d\" label=\"{{" + aux.nombre + " " + aux.apellido + "|" +
				aux.apodo + "}|" + aux.favoritos + "}\" color=\"#80036b\"" + "];\n"
			if contador > 0 {
				conexiones += "nodo" + strconv.Itoa(contador-1) + "->nodo" +
					strconv.Itoa(contador) + "[arrowhead=rvee, color=yellow]\n"
				conexiones += "nodo" + strconv.Itoa(contador) + "->nodo" +
					strconv.Itoa(contador-1) + "[arrowhead=rvee, color=orange]\n"
			}
			contador++
			aux = aux.Siguiente
		}
		codigo += "head [shape = oval label = \"Cabeza\" color  = \"#000dff\"]\n" +
			"cola [shape = oval, label = \"Cola\" color  = \"#000dff\"]\n"
		conexiones += "head->nodo0 [arrowhead=vee, color=red]\n" +
			"nodo" + strconv.Itoa(contador-1) + "-> cola [dir=back arrowtail=vee color=red]\n}"
		data := []byte(codigo + conexiones)
		_ = ioutil.WriteFile("Grafica.dot", data, 0644)
		path, _ := exec.LookPath("dot")
		cmd, _ := exec.Command(path, "-Tpdf", "Grafica.dot").Output()
		_ = ioutil.WriteFile("Grafica.pdf", cmd, os.FileMode(0777))
	} else {
		fmt.Println("No se pudo graficar, la lista está vacía")
	}

}
