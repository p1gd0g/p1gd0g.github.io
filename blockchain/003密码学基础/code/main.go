package main

import (
	"crypto/dsa"
	"crypto/rand"
	"fmt"
)

func main() {
	var sk dsa.PrivateKey
	dsa.GenerateParameters(&sk.Parameters, rand.Reader, dsa.L1024N160)
	dsa.GenerateKey(&sk, rand.Reader)

	fmt.Println(sk.X)
	fmt.Println(sk.PublicKey.Y)

	r, s, _ := dsa.Sign(rand.Reader, &sk, []byte("p1gd0g"))

	fmt.Println(r, s)

	fmt.Println(dsa.Verify(&sk.PublicKey, []byte("p1gd0g"), r, s))
}
