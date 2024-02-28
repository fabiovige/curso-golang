package secao1fundamentos

import "fmt"

func main() {
	// variaveis e atribuições de tipos
	var name string
	var age int
	var mensagem string
	mensagem = "Funcionalidades"
	fmt.Println(name, age, mensagem)

	nova_mensagem := "Novamensagem"
	nova_idade := 10
	preco := 20.5
	fmt.Println(nova_mensagem, nova_idade, preco)

	// tipo de variavel
	teste := "teste"
	teste = "teste 2"
	fmt.Println(teste)

	var texto string
	var numero int
	var metro float32
	var genero bool

	fmt.Println(texto)
	fmt.Println(numero)
	fmt.Println(metro)
	fmt.Println(genero)

	var joao string
	joao = "João"
}
