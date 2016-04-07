package main

import (
	"fmt"
	"github.com/devsisters/go-caplit"
)

func main() {
	fmt.Println("Generate ReadCaplit function...")
	caplit.GenCapnpReadCapLit("../capnp/dog", "../capnp/dog/generated_readcaplit.go", "dog", caplit.GetEnumList("../capnp/dog"))
}
