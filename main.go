package main

import (
	"fmt"

	"github.com/alura/banco/clientes"
	"github.com/alura/banco/contas"
)

func main() {
	clienteBruno := clientes.Titular{"Bruno", "123.123.123-12", "Desenvolvedor"}
	contaDoBruno := contas.ContaCorrente{clienteBruno, 123, 123456, 100}


	println(contaDoBruno.Saldo)
	fmt.Println(contaDoBruno.Sacar(50)) 

	println(contaDoBruno.ObterSaldo())
}