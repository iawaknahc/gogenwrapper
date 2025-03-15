package main

import (
	"io"
	"os"

	"github.com/iawaknahc/gogenwrapper/pkg/gogenwrapper"
)

func main() {
	packageName := "stdlibio"

	g := &gogenwrapper.T{
		PackageName:            packageName,
		StructNameWrapper:      "Wrapper",
		StructNameInterceptor:  "Interceptor",
		InterfaceNameUnwrap:    "Unwrapper",
		FunctionNameWrap:       "Wrap",
		FunctionNameUnwrap:     "Unwrap",
		FunctionTypeNamePrefix: "",
		BaseInterface:          new(interface{}),
		AdditionalInterfaces: []any{
			// It is observed that when the number of AdditionalInterfaces > 9,
			// the compilation consumes a lot of memory and runs very slow.
			new(io.Closer),
			new(io.Reader),
			new(io.ReaderAt),
			new(io.ReaderFrom),
			new(io.Seeker),

			new(io.Writer),
			new(io.WriterAt),
			new(io.WriterTo),

			// new(io.ByteReader),
			// new(interface {
			// 	UnreadByte() error
			// }),
			// new(io.ByteWriter),
			// new(io.RuneReader),
			// new(interface {
			// 	UnreadRune() error
			// }),
			// new(io.StringWriter),
		},
	}

	mainFileContents, err := g.GenerateGoSourceFile()
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("io.go", []byte(mainFileContents), 0o600)
	if err != nil {
		panic(err)
	}

	testFileContents, err := g.GenerateGoTestFile()
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("io_test.go", []byte(testFileContents), 0o600)
	if err != nil {
		panic(err)
	}
}
