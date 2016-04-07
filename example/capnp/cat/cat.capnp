@0xd04259d6323c70ed;

using Go = import "../import/go.capnp";
$Go.package("cat");
$Go.import("github.com/devsisters/go-caplit/example/capnp/cat");

using Dog = import "../dog/dog.capnp".Dog;

struct Cat {
  id      @0 : Int32;
  name    @1 : Text;
  partner @2 : Dog;
}