package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/horacerta", horaCerta)
	log.Println("Executando...")
	log.Println(http.ListenAndServe(":3000", nil))
}

func horaCerta(w http.ResponseWriter, r *http.Request) {

	s := time.Now().Format("02/01/2006 03:04:05") //formatação de data e hora em Go é mapeada ex: 02 igual a dia

	fmt.Fprintf(w, "<h1>Hora certa: %s</h1>", s)

}
