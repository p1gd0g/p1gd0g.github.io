package main

import (
	"crypto/dsa"
	"crypto/rand"
	"crypto/sha1"
	"fmt"
)

func main() {
	var sk dsa.PrivateKey
	dsa.GenerateParameters(&sk.Parameters, rand.Reader, dsa.L1024N160)
	dsa.GenerateKey(&sk, rand.Reader)

	fmt.Println(sk.X)
	fmt.Println(sk.PublicKey.Y)

	h := sha1.New().Sum([]byte("Hello world!"))

	r, s, _ := dsa.Sign(rand.Reader, &sk, h)

	fmt.Println(r, s)

	fmt.Println(dsa.Verify(&sk.PublicKey, h, r, s))
}
