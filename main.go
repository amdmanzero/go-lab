package main

import (
	"log"
	"net/http"

	rt "gosoft.co.th/mygo-lab/router"
)

func main() {
	log.Println("API Port 8080 Start!!")
	log.Fatal(http.ListenAndServe(":8080", rt.HandlerRouter()))
}
