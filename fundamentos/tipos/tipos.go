package main

import (
	"fmt"
	"reflect"
	"math"
)
func main(){
	//números inteiros
	fmt.Println(1, 2, 2000)
	fmt.Println("Liberal inteiro é", reflect.TypeOf(32000))

	//sem final (só positivos)... uint8 uint16 uint32 uint64
	
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo do int é", reflect.TypeOf(i1))

	//rune representa um mapeamento da tabela unicode (int32)
	var i2 rune = 'a'
	fmt.Println("O rune é", reflect.TypeOf(i2))

	//números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo liberal de 49.99 é", reflect.TypeOf(49.99))

	//boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	//string
	s1 := "Olá meu nome é Leo"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	s2 := `Olá
	meu 
	nome
	é 
	Pedro`

	fmt.Println("O tamanho da string é", len(s2))
}