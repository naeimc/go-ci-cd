package main

import (
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/sha3"
)

func main() {
	fmt.Println(base64.StdEncoding.EncodeToString(sha3.NewShake256().Sum([]byte{})))
}
