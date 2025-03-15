package gogenwrapper

import (
	"database/sql/driver"
	"encoding"
	"io"
	"reflect"
	"testing"
)

func TestGetAllPackagePaths(t *testing.T) {
	test := func(pointerToType any, expected []string) {
		type_pointerToType := reflect.TypeOf(pointerToType)
		if type_pointerToType.Kind() != reflect.Pointer {
			t.Errorf("%#v (%T) is not a pointer to an interface", pointerToType, pointerToType)
			return
		}
		type_type := type_pointerToType.Elem()

		actual := getAllPackagePaths(type_type)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("unexpected: %v != %v", actual, expected)
		}
	}

	// builtin type
	test(new(int), nil)
	test(new(bool), nil)
	test(new(string), nil)

	// named type
	test(new(io.Closer), []string{"io"})

	// array
	test(new([1]int), nil)
	test(new([1]io.Closer), []string{"io"})

	// slice
	test(new([]int), nil)
	test(new([]io.Closer), []string{"io"})

	// chan
	test(new(chan int), nil)
	test(new(chan<- int), nil)
	test(new(<-chan int), nil)
	test(new(chan io.Closer), []string{"io"})
	test(new(chan<- io.Closer), []string{"io"})
	test(new(<-chan io.Closer), []string{"io"})

	// pointer
	test(new(*int), nil)
	test(new(*io.Closer), []string{"io"})

	// map
	test(new(map[int]int), nil)
	test(new(map[reflect.Type]io.Closer), []string{"io", "reflect"})

	// func
	test(new(func()), nil)
	test(new(func(int)), nil)
	test(new(func(int) int), nil)
	test(new(func(int) (int, error)), nil)
	test(new(func(io.Reader, reflect.Type) (int, error)), []string{"io", "reflect"})

	// struct
	test(new(struct{}), nil)
	test(new(struct {
		Int    int
		Closer io.Closer
	}), []string{"io"})

	// interface
	test(new(interface {
	}), nil)
	test(new(interface {
		A()
	}), nil)
	test(new(interface {
		A(int)
	}), nil)
	test(new(interface {
		A(io.Closer)
	}), []string{"io"})
	test(new(interface {
		A(io.Closer) reflect.Type
	}), []string{"io", "reflect"})
	test(new(interface {
		A(io.Closer) reflect.Type
		B(encoding.BinaryMarshaler) encoding.TextMarshaler
	}), []string{"encoding", "io", "reflect"})
}

func TestFixSameNameInterface(t *testing.T) {
	typ := reflect.TypeOf(new(driver.ColumnConverter)).Elem()
	methods, ok := fixSameNameInterface(typ)
	if !ok {
		t.Errorf("unexpected")
	}
	if len(methods) != 1 {
		t.Errorf("unexpected: %v", len(methods))
	}
}

func TestFixPrivateInterface(t *testing.T) {
	type privateInterface interface {
		A()
	}

	typ := reflect.TypeOf(new(privateInterface)).Elem()
	methods, ok := fixPrivateInterface(typ)
	if !ok {
		t.Errorf("unexpected")
	}
	if len(methods) != 1 {
		t.Errorf("unexpected: %v", len(methods))
	}
}

func TestPackagesImports(t *testing.T) {
	packages := newPackages()

	if packages.PackageName("io") != "io" {
		t.Errorf("unexpected")
	}
	if packages.PackageName("io") != "io" {
		t.Errorf("unexpected")
	}
	if packages.PackageName("net/http") != "http" {
		t.Errorf("unexpected")
	}
	if packages.PackageName("net/http") != "http" {
		t.Errorf("unexpected")
	}
	if packages.PackageName("github.com/owner/repo/pkg/http") != "http1" {
		t.Errorf("unexpected")
	}
	if packages.PackageName("github.com/owner/repo/pkg/http") != "http1" {
		t.Errorf("unexpected")
	}

	actual := packages.Imports()
	expected := []import_{
		{
			PackageName: "http1",
			PackagePath: "github.com/owner/repo/pkg/http",
		},
		{
			PackageName: "io",
			PackagePath: "io",
		},
		{
			PackageName: "http",
			PackagePath: "net/http",
		},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("unexpected: %v != %v", actual, expected)
	}
}

func TestPackagesTypeName(t *testing.T) {
	packages := newPackages()

	test := func(pointerToType any, expected string) {
		typ := reflect.TypeOf(pointerToType).Elem()
		actual := packages.TypeName(typ)
		if actual != expected {
			t.Errorf("unexpected: %v != %v", actual, expected)
		}
	}

	// builtin
	test(new(int), "int")

	// error
	test(new(error), "error")

	// named type.
	test(new(io.Closer), "io.Closer")

	// array
	test(new([1]int), "[1]int")

	// chan
	test(new(<-chan int), "<-chan int")
	test(new(chan<- int), "chan<- int")
	test(new(chan int), "chan int")

	// func
	test(new(func()), "func()")
	test(new(func(int)), "func(int)")
	test(new(func(int) error), "func(int) error")
	test(new(func(...int) error), "func(...int) error")
	test(new(func(int, ...string) error), "func(int, ...string) error")
	test(new(func(int, ...string) (int, error)), "func(int, ...string) (int, error)")

	// interface
	test(new(interface{}), "any")
	test(new(interface{ A() }), "interface{ A() }")
	test(new(interface{ A(int) }), "interface{ A(int) }")
	test(new(interface{ A(int) error }), "interface{ A(int) error }")
	test(new(interface {
		A(int) error
		B(string) error
	}), "interface{ A(int) error; B(string) error }")
	test(new(interface {
		A(func() int) error
	}), "interface{ A(func() int) error }")
	test(new(interface {
		A(func() func()) error
	}), "interface{ A(func() func()) error }")
	test(new(interface {
		A(func() func() int) error
	}), "interface{ A(func() func() int) error }")

	// map
	test(new(map[string]interface{}), "map[string]any")
	test(new(map[string]struct{}), "map[string]struct{}")

	// pointer
	test(new(*int), "*int")

	// slice
	test(new([]int), "[]int")

	// struct
	test(new(struct{}), "struct{}")
	test(new(struct{ A int }), "struct{ A int }")
	test(new(struct {
		A int
		B string
	}), "struct{ A int; B string }")
}

func TestInterfaceMethod(t *testing.T) {
	type iface1 interface {
		A() int
		B() string
	}
	type iface2 interface {
		A() int
		B() bool
	}

	x, err := inspectInterface(new(iface1))
	if err != nil {
		t.Fatalf("expected iface1 to be an interface: %v", err)
	}

	y, err := inspectInterface(new(iface2))
	if err != nil {
		t.Fatalf("expected iface1 to be an interface: %v", err)
	}

	b := newInterfaceMethod(x.Method(0)) == newInterfaceMethod(y.Method(0))
	if !b {
		t.Errorf("unexpected: %v != %v", x.Method(0).Name, y.Method(0).Name)
	}

	b = newInterfaceMethod(x.Method(1)) == newInterfaceMethod(y.Method(1))
	if b {
		t.Errorf("unexpected: %v = %v", x.Method(1).Name, y.Method(1).Name)
	}
}

func TestInspectInterface(t *testing.T) {
	expected := `"" (string) is not a pointer to an interface`
	_, err := inspectInterface("")
	if err == nil {
		t.Errorf("expected InspectInterface to return error")
	} else {
		actual := err.Error()
		if actual != expected {
			t.Errorf("unexpected: %#v != %#v", actual, expected)
		}
	}
}
