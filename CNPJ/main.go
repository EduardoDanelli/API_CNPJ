package main

import (
	"CNPJ/busca"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/cnpj", busca.CNPJ)

	if err := r.Run(":9000"); err != nil {
		log.Fatal(err.Error())
	}
}
