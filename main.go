package main

import (
	"log"
	"net/http"
	"fmt"
)

func main() {
	fmt.Print("Serveur lancé")

	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Fatal(http.ListenAndServe(":3000", nil))
}
