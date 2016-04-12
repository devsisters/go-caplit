help:
	@echo "This Example offers 2 different examples with 2 capnp package. One is a simple capnp struct, Dog. And the other is 'Cat' package, which is including Dog package."
	@echo "You can learn how to generate and use 'ReadCapLit' function with 2 examples."
	@echo "If you want to know the exact command for running examples, please open 'Makefile'."
	@echo ""
	@echo " make singlepackage - it runs simple example with single capnp package 'Dog'."
	@echo " make multipackage - it deals with 'Cat' package"

singlepackage:clean
	capnp compile -ogo --src-prefix=import capnp/dog/*.capnp
	cd singlepackage && go generate singlepackage.go && go run singlepackage.go

multipackage:clean
	capnp compile -ogo --src-prefix=import capnp/dog/*.capnp capnp/cat/*.capnp
	cd multipackage && go generate multipackage.go && go run multipackage.go

clean:
	rm -f capnp/dog/*.go
	rm -f capnp/cat/*.go


