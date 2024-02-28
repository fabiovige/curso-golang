package secao1fundamentos

import (
	"fmt"
	"reflect"
)

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

	// operadores

	num1 := 10
	num2 := 20
	soma := num1 + num2
	fmt.Println(soma)

	subtracao := num1 - num2
	fmt.Println(subtracao)

	multiplicacao := num1 * num2
	fmt.Println(multiplicacao)

	divisao := num1 / num2
	fmt.Println(divisao)

	// saber o tipo da variavel com reflect
	fmt.Println(reflect.TypeOf(num1))

	// concatenando strings
	texto1 := "texto1"
	texto2 := "texto2"
	result := texto1 + texto2
	fmt.Println(result)

}
