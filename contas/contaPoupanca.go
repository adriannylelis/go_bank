package contas

import "github.com/alura/banco/clientes"

type ContaPopanca struct {
	Titular       clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo         float64
}



func (c *ContaPopanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaPopanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor do deposito menor que zero", c.saldo
	}
}

func (c *ContaPopanca) ObterSaldo() float64 {
	return c.saldo
}