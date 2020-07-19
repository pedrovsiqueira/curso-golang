package main

import "fmt"

func main(){

	//Print serve para escrever no console na mesma linha
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	//Println serve para escrever no console em uma nova linha
	fmt.Println(" Nova")
	fmt.Println("linha")

	
	x := 3.141516
	
	/*
		fmt.Println("O valor de x é " + x)
		a linha acima não funciona pq o go não permite concatenar string com float
	*/
	//transformar em string basta utilizar Sprint.
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)

	//imprimindo valores com string sem precisar alterar a sua tipagem
	fmt.Println("O valor de x é", x)
	
	//%f significa que vai ser impresso uma variável do tipo float
	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "Pedro"

	/*
		%d = inteiro
		%f = float
		%t = boolean
		%s = string
		%v = todos os tipos
	*/

	fmt.Printf("\n%d %f %t %s", a, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}