@0x854c63acad565e2d;

using Go = import "../import/go.capnp";
$Go.package("dog");
$Go.import("github.com/devsisters/go-caplit/example/capnp/dog");

struct DogToy {
  id      @0 :Int32;
  price   @1 :Int64;
  toyType @2 :ToyType;
}

struct Dog {
  id     @0 :Int32;
  name   @1 :Text;
  age    @2 :Int8;
  toys   @3 :List(DogToy);
}

enum ToyType {
  none @0;
  doll @1;
  asdf @2;
}