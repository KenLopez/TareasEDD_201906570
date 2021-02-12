package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func inicial(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Working... :D")

}

func agregar(w http.ResponseWriter, r *http.Request) {
	var data Message
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(w, "No puedo, me da amsiedad")
	}
	json.Unmarshal(reqBody, &data)
	json.NewEncoder(w).Encode(data)
}

type Message struct {
	datos []Datos `json:Mensajes`
}

type Datos struct {
	Origen  string     `json:Origen`
	Destino string     `json: Destino`
	Msg     []DatosMsg `json: Msg`
}

type DatosMsg struct {
	Fecha string `json: Fecha`
	Texto string `json: Texto`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", inicial).Methods("GET")
	router.HandleFunc("/agregar", agregar).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))
}
