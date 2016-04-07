singlepackage:clean
	capnp compile -ogo --src-prefix=import capnp/dog/*.capnp
	cd singlepackage && go generate singlepackage.go && go run singlepackage.go

multipackage:clean
	capnp compile -ogo --src-prefix=import capnp/dog/*.capnp capnp/cat/*.capnp
	cd multipackage && go generate multipackage.go && go run multipackage.go

clean:
	rm -f capnp/dog/*.go
	rm -f capnp/cat/*.go


