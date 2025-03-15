package main

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/iawaknahc/gogenwrapper/pkg/gogenwrapper"
)

type deadliner interface {
	SetReadDeadline(deadline time.Time) error
	SetWriteDeadline(deadline time.Time) error
}

func main() {
	packageName := "nethttp"

	g := &gogenwrapper.T{
		PackageName:            packageName,
		StructNameWrapper:      "ResponseWriterWrapper",
		StructNameInterceptor:  "ResponseWriterInterceptor",
		InterfaceNameUnwrap:    "ResponseWriterUnwrapper",
		FunctionNameWrap:       "WrapResponseWriter",
		FunctionNameUnwrap:     "UnwrapResponseWriter",
		FunctionTypeNamePrefix: "",
		BaseInterface:          new(http.ResponseWriter),
		AdditionalInterfaces: []any{
			new(http.Flusher),
			new(http.CloseNotifier),
			new(http.Hijacker),
			new(io.ReaderFrom),
			new(http.Pusher),
			new(deadliner),
			new(interface {
				EnableFullDuplex() error
			}),
		},
	}

	mainFileContents, err := g.GenerateGoSourceFile()
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("response_writer.go", []byte(mainFileContents), 0o600)
	if err != nil {
		panic(err)
	}

	testFileContents, err := g.GenerateGoTestFile()
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("response_writer_test.go", []byte(testFileContents), 0o600)
	if err != nil {
		panic(err)
	}
}
