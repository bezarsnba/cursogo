package main

import "fmt"

var (
	//Endereco é um valor importante
	// Telefone é uma informacao para ligar :D
	Endereco   string
	Telefone   string
	quantidade int
	comprou    bool
	valor      float64
	palavras   rune
)

func main() {
	teste := "Valor de teste"
	fmt.Printf("endereco: %s\r\n", Endereco)
	fmt.Printf("qtd: %d\r\n", quantidade)
	fmt.Printf("comprou: %v\r\n", comprou)
	fmt.Printf("palavras: %v\r\n", palavras)
	fmt.Printf("O valor de teste é: %v\r\n", teste)
}
