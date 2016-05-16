package main

//go:generate go run gen/generate.go
import (
	"fmt"
	"github.com/devsisters/go-caplit/example/capnp/dog"
	"github.com/glycerine/go-capnproto"
	"strings"
)

func main() {
	fmt.Println("-- Making Dummy Dog Struct")
	buf := capn.NewBuffer(nil)
	myDog := dog.NewRootDog(buf)

	dogCaplit := `(id=2,name="Lolo",age=16,toys=[(id=1,price=42,toyType=doll),])`
	fmt.Println("-- ReadCapLit with data: ")
	fmt.Println(dogCaplit)
	err := myDog.ReadCapLit(strings.NewReader(dogCaplit))
	if err != nil {
		panic(err)
	}
	fmt.Printf("#%v dog '%v' is %v years old\n", myDog.Id(), myDog.Name(), myDog.Age())
	fmt.Printf("He has %v toy. It's price is %v\n", myDog.Toys().Len(), myDog.Toys().At(0).Price())
}
