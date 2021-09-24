package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	rt "gosoft.co.th/mygo-lab/router"
)

func main() {
	log.Println("API Port 8080 Start!!")

	godotenv.Load()
	log.Fatal(http.ListenAndServe(":8080", rt.HandlerRouter()))
}
