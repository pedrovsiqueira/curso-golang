package main

import (
	"fmt"
	"math"
)
/*
Definir um tipo de variável: const ou var.
Nome da variável
Tipo da variável (não é obrigatório se eu especificar o valor da variável)
e o que ela recebe
*/

func main(){
	const PI float64 = 3.1415
	var raio = 3.2 //tipo float64 inferido pelo compilador.

	//forma reduzida de criar uma var

	area := PI * math.Pow(raio, 2)
	fmt.Println("A área da circunferência é", area)

	//declarando variáveis em blocos

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	//declarando variaveis com tipagens seguidas
	var e, f bool = true, false

	//reduzida
	g, h, i := 2, false, "epa!"

	fmt.Println(a, b, c, d, e, f)
	fmt.Println(g, h, i)
}