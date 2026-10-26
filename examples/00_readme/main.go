package main

import (
	"context"
	"log"
	"os"

	nfe "github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfe/nfe_v4_00"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("uso: go run . documento.xml")
	}
	dados, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var documento nfe.NFe
	if err := nfe.Parse(context.Background(), dados, &documento); err != nil {
		log.Fatal(err)
	}
	log.Print("XML validado e documento decodificado")
}
