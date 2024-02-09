package basic

import (
	"log"
	"my-protobuf/protogen/basic"
)

// new instance of protobuf message
func BasicHello() {
	h := basic.Hello{
		Name: "Clark Kent",
	}

	log.Println(&h)
}
