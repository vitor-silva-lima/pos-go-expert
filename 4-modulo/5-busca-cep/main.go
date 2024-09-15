package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCepOutput struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		resp, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			fmt.Fprint(os.Stderr, "cep not found")
			continue
		}
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}
		var viaCepOutput ViaCepOutput
		err = json.Unmarshal(data, &viaCepOutput)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}
		fmt.Println(viaCepOutput)
		file, err := os.Create(fmt.Sprintf("./%s.json", cep))
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}
		defer file.Close()
		err = json.NewEncoder(file).Encode(viaCepOutput)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}
	}
}
