package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

// os elementos do struct possuem valores iniciais chamados de zero value

func main() { // Função Principal/Inicial
	contaDoJoão := ContaCorrente{titular: "João", numeroAgencia: 589, numeroConta: 12345, saldo: 125.5} // atribuição curta (n precisa informar que e var) o codigo sem atribuição curta ficaria = var contaDoJoao ContaCorrente = ContaCorrente{}
	fmt.Println(contaDoJoão)
	// fmt é um pacote(?) que imprime no console
	contaDoVinicius := ContaCorrente{"Vinícius", 589, 12345, 125.5} // atribuição curta (n precisa informar que e var)
	fmt.Println(contaDoVinicius)                                    // fmt é um pacote(?) que imprime no console

}
