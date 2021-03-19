package controllers

import (
	"CNPJ/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetCNPJ(cnpj string) *models.CNPJ {

	resp, err := http.Get("https://www.receitaws.com.br/v1/cnpj/" + cnpj)
	if err != nil {
		log.Fatalf("Erro: %s", err)
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Erro: %s", err)
		return nil
	}

	defer resp.Body.Close()
	fmt.Printf("%s\n", body)

	var cnpjInfo models.CNPJ
	if err := json.Unmarshal([]byte(body), &cnpjInfo); err != nil {
		log.Fatalf("Erro: %s", err)
		return nil
	}

	return &cnpjInfo

}
