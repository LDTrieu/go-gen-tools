package main

import (
	//nolint
	"go-gen-tools/cmd/protoc-gen-custom/internal"

	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	opt := protogen.Options{}
	internal.Run(opt)
}
