package a

import "github.com/golang/protobuf/ptypes" // want `use the the packages from "google.golang.org/protobuf/types/known" instead of the deprecated package "github.com/golang/protobuf/ptypes" \(see https://github.com/golang/protobuf/releases#v1.4-well-known-types\)`

func IllegalImports() {
	_ = ptypes.TimestampNow()
}
