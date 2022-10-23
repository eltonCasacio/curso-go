package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {

	http.HandleFunc("/", BuscaCEPHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao startar servidor...")
	}
}

func BuscaCEPHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cepParams := r.URL.Query().Get("cep")
	if cepParams == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := BuscaCEP(cepParams)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func BuscaCEP(cep string) (ViaCEP, error) {
	v := ViaCEP{}
	req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return v, err
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return v, err
	}

	var data ViaCEP
	err = json.Unmarshal(body, &data)
	if err != nil {
		return v, err
	}

	return data, nil
}
