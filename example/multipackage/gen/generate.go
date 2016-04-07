package main

import (
	"fmt"
	"github.com/devsisters/go-caplit"
)

func main() {
	fmt.Println("-- Generate ReadCaplit function...")
	dogEnumList := caplit.GetEnumList("../capnp/dog")
	caplit.GenCapnpReadCapLit("../capnp/dog", "../capnp/dog/generated_readcaplit.go", "dog", dogEnumList, []string{})
	allEnumList := append(caplit.GetEnumList("../capnp/cat"), dogEnumList...)
	caplit.GenCapnpReadCapLit("../capnp/cat", "../capnp/cat/generated_readcaplit.go", "cat", allEnumList, []string{"github.com/devsisters/go-caplit/example/capnp/dog"})
}
