package main

import (
	"fmt"

	 "./prefixo"
	 "./operadora"
)
// NomeDousuario  é um nome do usuario
var NomeDoUsuario = "Bob"

func main() {
	fmt.Printf("Nome do usuario: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da capitao: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome da operadora: %s\r\n", operadora.NomeOperadora)
	fmt.Printf("Valor teste: %s\r\n", prefixo.TesteComPrefixo)
}
