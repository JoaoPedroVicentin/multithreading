package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Address struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	API          string
}

type BrasilCep struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

type ViaCep struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
}

func main() {
	cep := "01153000"

	ch := make(chan Address, 2)

	go buscaBrasilCep(cep, ch)
	go buscaViaCep(cep, ch)

	select {
	case res := <-ch:
		fmt.Printf("Resposta recebida da %s:\n%+v\n", res.API, res)
	case <-time.After(1 * time.Second):
		fmt.Println("Erro: timeout após 1 segundo")
	}
}

func buscaBrasilCep(cep string, ch chan Address) {
	url := "https://brasilapi.com.br/api/cep/v1/" + cep
	res, err := http.Get(url)
	if err != nil {
		log.Println("Erro BrasilAPI:", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Erro ao ler resposta BrasilAPI:", err)
		return
	}

	var data BrasilCep
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println("Erro ao decodificar BrasilAPI:", err)
		return
	}

	ch <- Address{
		Cep:          data.Cep,
		State:        data.State,
		City:         data.City,
		Neighborhood: data.Neighborhood,
		Street:       data.Street,
		API:          "BrasilAPI",
	}
}

func buscaViaCep(cep string, ch chan Address) {
	url := "http://viacep.com.br/ws/" + cep + "/json/"
	res, err := http.Get(url)
	if err != nil {
		log.Println("Erro ViaCEP:", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Erro ao ler resposta ViaCEP:", err)
		return
	}

	var data ViaCep
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println("Erro ao decodificar ViaCEP:", err)
		return
	}

	ch <- Address{
		Cep:          data.Cep,
		State:        data.Uf,
		City:         data.Localidade,
		Neighborhood: data.Bairro,
		Street:       data.Logradouro,
		API:          "ViaCEP",
	}
}
