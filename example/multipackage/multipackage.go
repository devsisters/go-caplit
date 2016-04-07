package main
//go:generate go run gen/generate.go
import (
	"fmt"
	"github.com/devsisters/go-caplit/example/capnp/cat"
	"github.com/glycerine/go-capnproto"
	"strings"
)

func main() {
	fmt.Println("-- Making Dummy Cat Struct")
	buf := capn.NewBuffer(nil)
	myCat := cat.NewRootCat(buf)

	catCaplit := `(id=1,name="Navi",partner=(id=2,name="Lolo",age=16,toys=[(id=1,price=42,toyType=doll),]))`
	fmt.Println("-- ReadCapLit with data: ")
	fmt.Println(catCaplit)
	myCat.ReadCapLit(strings.NewReader(catCaplit))
	fmt.Printf("#%v cat '%v' has a partner dog '%v'\n", myCat.Id(), myCat.Name(), myCat.Partner().Name())
}
