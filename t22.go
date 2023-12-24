package main

import (
	"fmt"
	"math/big"
)

func mul(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func div(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}

func add(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func sub(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func main() {
	a := big.NewInt(1 << 40)
	b := big.NewInt(1 << 30)

	fmt.Println(mul(a, b))
	fmt.Println(div(a, b))
	fmt.Println(add(a, b))
	fmt.Println(sub(a, b))
}
