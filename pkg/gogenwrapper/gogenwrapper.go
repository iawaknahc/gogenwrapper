package gogenwrapper

import (
	"bytes"
	"cmp"
	"database/sql/driver"
	"fmt"
	"go/format"
	"net/http"
	"reflect"
	"slices"
	"sort"
	"strings"
	"unicode"
)

var errorType = reflect.TypeOf(new(error)).Elem()

func getAllPackagePaths(typ reflect.Type) []string {
	seen := make(map[string]struct{})
	var out []string

	var recur func(reflect.Type)
	recur = func(typ reflect.Type) {
		pkgPath := typ.PkgPath()
		name := typ.Name()

		switch {
		case pkgPath != "" && name != "":
			// e.g. io.Closer
			_, ok := seen[pkgPath]
			if !ok {
				seen[pkgPath] = struct{}{}
				out = append(out, pkgPath)
			}
		default:
			switch typ.Kind() {
			case reflect.Array:
				recur(typ.Elem())
			case reflect.Chan:
				recur(typ.Elem())
			case reflect.Func:
				for i := 0; i < typ.NumIn(); i += 1 {
					recur(typ.In(i))
				}
				for i := 0; i < typ.NumOut(); i += 1 {
					recur(typ.Out(i))
				}
			case reflect.Interface:
				numMethod := typ.NumMethod()
				for i := 0; i < numMethod; i += 1 {
					recur(typ.Method(i).Type)
				}
			case reflect.Map:
				recur(typ.Key())
				recur(typ.Elem())
			case reflect.Pointer:
				recur(typ.Elem())
			case reflect.Slice:
				recur(typ.Elem())
			case reflect.Struct:
				for i := 0; i < typ.NumField(); i += 1 {
					recur(typ.Field(i).Type)
				}
			}
		}
	}

	recur(typ)

	sort.Strings(out)
	return out
}

// fixOverlappingInterface handles cases like database/sql/driver.Rows and friends,
// where the optional interfaces embed the main interface, thus the methods are overlapping with each other.
// Overlapping methods causes the Go compiler to think there is an ambiguity when
// we embed these interfaces in a struct.
func fixOverlappingInterface(allInterfaces []reflect.Type, iface reflect.Type) ([]interfaceMethod, bool) {
	// true means keep
	interfaceMethods := make(map[interfaceMethod]bool)
	for i := 0; i < iface.NumMethod(); i += 1 {
		method := iface.Method(i)
		interfaceMethod := newInterfaceMethod(method)
		interfaceMethods[interfaceMethod] = true
	}

	// Empty interface cannot be disjointed.
	if len(interfaceMethods) == 0 {
		return nil, false
	}

	for _, anotherIface := range allInterfaces {
		// Exclude self.
		if iface != anotherIface {
			for i := 0; i < anotherIface.NumMethod(); i += 1 {
				method := anotherIface.Method(i)
				interfaceMethod := newInterfaceMethod(method)

				_, ok := interfaceMethods[interfaceMethod]
				if ok {
					interfaceMethods[interfaceMethod] = false
				}
			}
		}
	}

	var disjointedMethods []interfaceMethod
	disjointed := false
	for interfaceMethod, keep := range interfaceMethods {
		if !keep {
			disjointed = true
		} else {
			disjointedMethods = append(disjointedMethods, interfaceMethod)
		}
	}
	slices.SortStableFunc(disjointedMethods, func(a interfaceMethod, b interfaceMethod) int {
		return cmp.Compare(a.Name, b.Name)
	})

	if disjointed {
		return disjointedMethods, true
	}
	return nil, false
}

// fixSameNameInterface handles cases like database/sql/driver.ColumnConverter, where
// The name of the interface collides with the method name.
// This name collision causes the Go compiler to think a struct does not implement
// the interface because the field name has precedence over the method.
func fixSameNameInterface(iface reflect.Type) ([]interfaceMethod, bool) {
	hasCollision := false
	interfaceName := iface.Name()

	var interfaceMethods []interfaceMethod
	for i := 0; i < iface.NumMethod(); i += 1 {
		method := iface.Method(i)
		interfaceMethod := newInterfaceMethod(method)
		interfaceMethods = append(interfaceMethods, interfaceMethod)
		if interfaceMethod.Name == interfaceName {
			hasCollision = true
		}
	}
	slices.SortStableFunc(interfaceMethods, func(a interfaceMethod, b interfaceMethod) int {
		return cmp.Compare(a.Name, b.Name)
	})

	if hasCollision {
		return interfaceMethods, true
	}

	return nil, false
}

func nameIsPrivate(name string) bool {
	if name == "" {
		return false
	}
	isPrivate := false
	for idx, rune_ := range name {
		if idx == 0 {
			isPrivate = unicode.IsLower(rune_)
			break
		}
	}
	return isPrivate
}

// fixPrivateInterface handles cases where the user pass in a private interface.
func fixPrivateInterface(iface reflect.Type) ([]interfaceMethod, bool) {
	name := iface.Name()
	isPrivate := nameIsPrivate(name)
	if !isPrivate {
		return nil, false
	}

	var interfaceMethods []interfaceMethod
	for i := 0; i < iface.NumMethod(); i += 1 {
		method := iface.Method(i)
		interfaceMethod := newInterfaceMethod(method)
		interfaceMethods = append(interfaceMethods, interfaceMethod)
	}
	slices.SortStableFunc(interfaceMethods, func(a interfaceMethod, b interfaceMethod) int {
		return cmp.Compare(a.Name, b.Name)
	})

	return interfaceMethods, true
}

// fixAnonymousInterface handles cases where the interface is anonymous.
// It is a syntax error to embed an anonymous interface is a struct.
//
//	struct {
//		interface { A () }
//	}
//
// is a syntax error.
func fixAnonymousInterface(iface reflect.Type) ([]interfaceMethod, bool) {
	if iface.Name() != "" {
		return nil, false
	}

	var interfaceMethods []interfaceMethod
	for i := 0; i < iface.NumMethod(); i += 1 {
		method := iface.Method(i)
		interfaceMethod := newInterfaceMethod(method)
		interfaceMethods = append(interfaceMethods, interfaceMethod)
	}
	slices.SortStableFunc(interfaceMethods, func(a interfaceMethod, b interfaceMethod) int {
		return cmp.Compare(a.Name, b.Name)
	})

	return interfaceMethods, true
}

type import_ struct {
	PackageName string
	PackagePath string
}

type packages_ struct {
	packageNameToPackagePath map[string]string
	packagePathToPackageName map[string]string
}

func newPackages() *packages_ {
	return &packages_{
		packageNameToPackagePath: make(map[string]string),
		packagePathToPackageName: make(map[string]string),
	}
}

func (p *packages_) PackageName(pkgPath string) string {
	if pkgName, ok := p.packagePathToPackageName[pkgPath]; ok {
		return pkgName
	}

	parts := strings.Split(pkgPath, "/")
	lastPart := parts[len(parts)-1]

	suffix := 0
	pkgName := lastPart
	for {
		_, ok := p.packageNameToPackagePath[pkgName]
		if !ok {
			p.packagePathToPackageName[pkgPath] = pkgName
			p.packageNameToPackagePath[pkgName] = pkgPath
			return pkgName
		}

		suffix += 1
		pkgName = fmt.Sprintf("%v%v", lastPart, suffix)
	}
}

func (p *packages_) Imports() []import_ {
	var imports []import_
	for pkgPath, pkgName := range p.packagePathToPackageName {
		imports = append(imports, import_{
			PackageName: pkgName,
			PackagePath: pkgPath,
		})
	}
	slices.SortStableFunc(imports, func(a import_, b import_) int {
		return cmp.Compare(a.PackagePath, b.PackagePath)
	})
	return imports
}

func (p *packages_) typeNameFunc(typ reflect.Type, funcName string) string {
	var buf strings.Builder
	isVariadic := typ.IsVariadic()
	numIn := typ.NumIn()
	numOut := typ.NumOut()
	_, _ = buf.WriteString(fmt.Sprintf("%v(", funcName))
	for i := 0; i < numIn; i += 1 {
		if i != 0 {
			_, _ = buf.WriteString(", ")
		}
		if isVariadic && i == numIn-1 {
			_, _ = buf.WriteString(fmt.Sprintf("...%v", p.TypeName(typ.In(i).Elem())))
		} else {
			_, _ = buf.WriteString(fmt.Sprintf("%v", p.TypeName(typ.In(i))))
		}
	}
	_, _ = buf.WriteString(")")
	switch numOut {
	case 0:
		break
	case 1:
		_, _ = buf.WriteString(" ")
		buf.WriteString(fmt.Sprintf("%v", p.TypeName(typ.Out(0))))
	default:
		_, _ = buf.WriteString(" (")
		for i := 0; i < numOut; i += 1 {
			if i != 0 {
				_, _ = buf.WriteString(", ")
			}
			_, _ = buf.WriteString(fmt.Sprintf("%v", p.TypeName(typ.Out(i))))
		}
		_, _ = buf.WriteString(")")
	}
	return buf.String()
}

func (p *packages_) typeNameMethods(methods []interfaceMethod) string {
	// Return any instead of interface{} so that it can be embedded in a struct without syntax error.
	if len(methods) == 0 {
		return "any"
	}

	var buf strings.Builder
	_, _ = buf.WriteString("interface{ ")

	for i, method := range methods {
		if i != 0 {
			_, _ = buf.WriteString("; ")
		}
		_, _ = buf.WriteString(fmt.Sprintf("%v", p.typeNameFunc(method.Type, method.Name)))
	}
	_, _ = buf.WriteString(" }")
	return buf.String()
}

func (p *packages_) TypeName(typ reflect.Type) string {
	pkgPath := typ.PkgPath()
	name := typ.Name()
	switch {
	case typ == errorType:
		// Special case: error
		return "error"
	case pkgPath != "" && !nameIsPrivate(name):
		pkgName := p.PackageName(pkgPath)
		return fmt.Sprintf("%v.%v", pkgName, name)
	default:
		switch typ.Kind() {
		case reflect.Array:
			return fmt.Sprintf("[%v]%v", typ.Len(), p.TypeName(typ.Elem()))
		case reflect.Chan:
			switch typ.ChanDir() {
			case reflect.RecvDir:
				return fmt.Sprintf("<-chan %v", p.TypeName(typ.Elem()))
			case reflect.SendDir:
				return fmt.Sprintf("chan<- %v", p.TypeName(typ.Elem()))
			case reflect.BothDir:
				return fmt.Sprintf("chan %v", p.TypeName(typ.Elem()))
			default:
				panic(fmt.Errorf("unknown chan direction: %v", typ))
			}
		case reflect.Func:
			return p.typeNameFunc(typ, "func")
		case reflect.Interface:
			var methods []interfaceMethod
			for i := 0; i < typ.NumMethod(); i += 1 {
				method := typ.Method(i)
				methods = append(methods, newInterfaceMethod(method))
			}
			return p.typeNameMethods(methods)
		case reflect.Map:
			return fmt.Sprintf("map[%v]%v", p.TypeName(typ.Key()), p.TypeName(typ.Elem()))
		case reflect.Pointer:
			return fmt.Sprintf("*%v", p.TypeName(typ.Elem()))
		case reflect.Slice:
			return fmt.Sprintf("[]%v", p.TypeName(typ.Elem()))
		case reflect.Struct:
			var buf strings.Builder
			if typ.NumField() != 0 {
				_, _ = buf.WriteString("struct{ ")
			} else {
				_, _ = buf.WriteString("struct{")
			}
			for i := 0; i < typ.NumField(); i += 1 {
				field := typ.Field(i)
				if i != 0 {
					_, _ = buf.WriteString("; ")
				}
				_, _ = buf.WriteString(fmt.Sprintf("%v %v", field.Name, p.TypeName(field.Type)))
			}
			if typ.NumField() != 0 {
				_, _ = buf.WriteString(" }")
			} else {
				_, _ = buf.WriteString("}")
			}
			return buf.String()
		default:
			return fmt.Sprintf("%v", typ)
		}
	}
}

// interfaceMethod is comparable with ==.
type interfaceMethod struct {
	// Name is the name of the interface method.
	Name string
	// Type is the type of the interface method.
	Type reflect.Type
}

func newInterfaceMethod(method reflect.Method) interfaceMethod {
	return interfaceMethod{
		Name: method.Name,
		Type: method.Type,
	}
}

func (m interfaceMethod) IsIncompatible(that interfaceMethod) bool {
	return m.Name == that.Name && m.Type != that.Type
}

// inspectInterface takes a pointer to an interface, and return an reflect.Type for the interface.
// It is expected you use the builtin new function to create a pointer to an interface.
// For example, inspectInterface(new(io.Closer)).
func inspectInterface(pointerToInterface any) (reflect.Type, error) {
	type_pointerToInterface := reflect.TypeOf(pointerToInterface)
	if type_pointerToInterface.Kind() != reflect.Pointer {
		return nil, fmt.Errorf("%#v (%T) is not a pointer to an interface", pointerToInterface, pointerToInterface)
	}

	type_interface := type_pointerToInterface.Elem()
	if type_interface.Kind() != reflect.Interface {
		return nil, fmt.Errorf("%#v (%T) is not a pointer to an interface", pointerToInterface, pointerToInterface)
	}

	return type_interface, nil
}

type wrapperGeneratorMethod struct {
	Name      string
	Type      reflect.Type
	Interface reflect.Type
	In        []reflect.Type
	Out       []reflect.Type
}

var _ http.ResponseWriter = nil
var _ driver.Conn = nil

// T is a wrapper generator for complex interfaces likes [http.ResponseWriter] and [driver.Conn].
type T struct {
	// PackageName defines the package name of the generated Go source file.
	// It must be specified.
	PackageName string

	// StructNameWrapper defines the name of the wrapper struct.
	// For example, if you are wrapping [http.ResponseWriter], then ResponseWriterWrapper is a common choice.
	// It must be specified.
	StructNameWrapper string

	// StructNameInterceptor defines the name of the interceptor struct.
	// For example, if you are wrapping [http.ResponseWriter], then ResponseWriterInterceptor is a common choice.
	// It must be specified.
	StructNameInterceptor string

	// FunctionNameWrap defines the name of the wrap function.
	// It is the main entrypoint to wrap the interface.
	// For example, if you are wrapping [http.ResponseWriter], then WrapResponseWriter is a common choice.
	// It must be specified.
	FunctionNameWrap string

	// FunctionNameUnwrap defines the name of the unwrap function.
	// For example, if you are wrapping [http.ResponseWriter], then UnwrapResponseWriter is a common choice.
	// It must be specified.
	FunctionNameUnwrap string

	// InterfaceNameUnwrap defines the name of the unwrap interface.
	// For example, if you are wrapping [http.ResponseWriter], then ResponseWriterUnwrapper is a common choice.
	// It must be specified.
	InterfaceNameUnwrap string

	// FunctionTypeNamePrefix defines the prefix of the function type for each interface method.
	// For example, if you are wrapping [http.ResponseWriter], then ResponseWriter_ is a common choice.
	// It can be an empty string.
	FunctionTypeNamePrefix string

	// BaseInterface is the base interface.
	// It must be a pointer to the interface.
	// For example, if you are wrapping [http.ResponseWriter], pass in new(http.ResponseWriter).
	BaseInterface any

	// AdditionalInterfaces are the additional interfaces that the wrapped interface may implement.
	// The value must be a pointer to an interface.
	// For example, if you are wrapping [http.ResponseWriter], you need to pass in all possible interfaces
	// that can be implemented by [http.ResponseWriter], like [http.Fluster].
	// It can be unspecified. In this case, the wrapper only implements the base interface.
	AdditionalInterfaces []any
}

func (g *T) validate() (err error) {
	if g.PackageName == "" {
		err = fmt.Errorf("PackageName must be specified")
		return
	}
	if g.StructNameWrapper == "" {
		err = fmt.Errorf("StructNameWrapper must be specified")
		return
	}
	if g.StructNameInterceptor == "" {
		err = fmt.Errorf("StructNameInterceptor must be specified")
		return
	}
	if g.FunctionNameWrap == "" {
		err = fmt.Errorf("FunctionNameWrap must be specified")
		return
	}
	if g.FunctionNameUnwrap == "" {
		err = fmt.Errorf("FunctionNameUnwrap must be specified")
		return
	}
	if g.InterfaceNameUnwrap == "" {
		err = fmt.Errorf("InterfaceNameUnwrap must be specified")
		return
	}

	return
}

func (g *T) prepareInterfaces() (baseInterface reflect.Type, additionalInterfaces []reflect.Type, err error) {
	baseInterface, err = inspectInterface(g.BaseInterface)
	if err != nil {
		return
	}

	for i, pointerToInterface := range g.AdditionalInterfaces {
		var iface reflect.Type
		iface, err = inspectInterface(pointerToInterface)
		if err != nil {
			err = fmt.Errorf("AdditionalInterfaces[%v]: %w", i, err)
			return
		}
		additionalInterfaces = append(additionalInterfaces, iface)
	}

	return
}

func (g *T) prepareMethods(baseInterface reflect.Type, additionalInterfaces []reflect.Type) (methods []wrapperGeneratorMethod, err error) {
	addMethod := func(iface reflect.Type, m interfaceMethod) error {
		skip := false
		for _, existing := range methods {
			that := interfaceMethod{
				Name: existing.Name,
				Type: existing.Type,
			}
			if m.IsIncompatible(that) {
				return fmt.Errorf("%v (%v) is incompatible with %v (%v)", m.Name, m.Type, that.Name, that.Type)
			}
			if m == that {
				skip = true
			}
		}
		if !skip {
			var in []reflect.Type
			var out []reflect.Type
			for i := 0; i < m.Type.NumIn(); i += 1 {
				in = append(in, m.Type.In(i))
			}
			for i := 0; i < m.Type.NumOut(); i += 1 {
				out = append(out, m.Type.Out(i))
			}
			methods = append(methods, wrapperGeneratorMethod{
				Name:      m.Name,
				Type:      m.Type,
				Interface: iface,
				In:        in,
				Out:       out,
			})
		}
		return nil
	}

	for i := 0; i < baseInterface.NumMethod(); i += 1 {
		m := newInterfaceMethod(baseInterface.Method(i))
		err = addMethod(baseInterface, m)
		if err != nil {
			return
		}
	}
	for _, iface := range additionalInterfaces {
		for i := 0; i < iface.NumMethod(); i += 1 {
			m := newInterfaceMethod(iface.Method(i))
			err = addMethod(iface, m)
			if err != nil {
				return
			}
		}
	}

	return
}

func (g *T) fixInterface(allInterfaces []reflect.Type, iface reflect.Type) ([]interfaceMethod, bool) {
	// The order here is important.
	// 1. No matter whether the interface is anonymous / private / name collision with its method, if
	//    it overlaps with other interfaces, it has to be disjointed.
	// 2. Otherwise, the interface is disjointed already.
	// 3. Anonymous interface (without name).
	// 4. Private interface (by definition, it has a name).
	// 5. Name collision interface.
	if interfaceMethods, ok := fixOverlappingInterface(allInterfaces, iface); ok {
		return interfaceMethods, ok
	} else if interfaceMethods, ok := fixAnonymousInterface(iface); ok {
		return interfaceMethods, ok
	} else if interfaceMethods, ok := fixPrivateInterface(iface); ok {
		return interfaceMethods, ok
	} else if interfaceMethods, ok := fixSameNameInterface(iface); ok {
		return interfaceMethods, ok
	}
	return nil, false
}

// GenerateGoTestFile generates the Go test file for the wrapper.
// You are supposed to write it to a file.
func (g *T) GenerateGoTestFile() (code string, err error) {
	err = g.validate()
	if err != nil {
		return
	}

	baseInterface, additionalInterfaces, err := g.prepareInterfaces()
	if err != nil {
		return
	}

	packages := newPackages()
	// import "testing"
	packages.PackageName("testing")

	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	var buf bytes.Buffer
	printf := func(format string, values ...any) {
		_, err := fmt.Fprintf(&buf, format, values...)
		if err != nil {
			panic(err)
		}
	}

	printf("func Test%v(t *testing.T) {\n", g.FunctionNameWrap)

	numBit := len(additionalInterfaces)
	var numCombinations uint64 = uint64(1) << numBit
	for i := uint64(0); i < numCombinations; i += 1 {
		printf("	{\n")

		var allInterfaces = []reflect.Type{baseInterface}
		for bit := 0; bit < numBit; bit += 1 {
			val := uint64(1) << bit
			b := (i & val) == 0
			if !b {
				allInterfaces = append(allInterfaces, additionalInterfaces[bit])
			}
		}

		interfaceMethodsByBit := make([][]interfaceMethod, numBit)
		for bit := 0; bit < numBit; bit += 1 {
			val := uint64(1) << bit
			b := (i & val) == 0
			if !b {
				original := additionalInterfaces[bit]
				interfaceMethods, ok := g.fixInterface(allInterfaces, original)
				if ok {
					interfaceMethodsByBit[bit] = interfaceMethods
				}
			}
		}

		printf("		t.Log(\"combination %v/%v: %v", i+1, numCombinations, packages.TypeName(baseInterface))
		for bit := 0; bit < numBit; bit += 1 {
			val := uint64(1) << bit
			b := (i & val) == 0
			if !b {
				printf(" %v", packages.TypeName(additionalInterfaces[bit]))
			}
		}
		printf("\")\n")

		for bit, interfaceMethods := range interfaceMethodsByBit {
			if interfaceMethods != nil {
				printf("		type combination_%v_%v %v\n", i, bit, packages.typeNameMethods(interfaceMethods))
			}
		}

		printf("		wrapped := struct {\n")
		printf("			%v\n", packages.TypeName(baseInterface))
		for bit := 0; bit < numBit; bit += 1 {
			val := uint64(1) << bit
			b := (i & val) == 0
			if !b {
				original := additionalInterfaces[bit]
				interfaceMethods := interfaceMethodsByBit[bit]
				if interfaceMethods == nil {
					printf("			%v\n", packages.TypeName(original))
				} else {
					printf("			combination_%v_%v\n", i, bit)
				}
			}
		}
		printf("		}{}\n")
		printf("		w := %v(wrapped, %v{})\n", g.FunctionNameWrap, g.StructNameInterceptor)
		printf("\n")

		for bit := 0; bit < numBit; bit += 1 {
			val := uint64(1) << bit
			b := (i & val) == 0
			printf("		if _, ok := w.(%v); ok != %v {\n", packages.TypeName(additionalInterfaces[bit]), !b)
			printf("			t.Errorf(\"combination %v/%v: unexpected interface %v\")\n", i+1, numCombinations, packages.TypeName(additionalInterfaces[bit]))
			printf("		}\n")
		}
		printf("\n")

		printf("		if w, ok := w.(%v); ok {\n", g.InterfaceNameUnwrap)
		printf("			if w.%v() != wrapped {\n", g.FunctionNameUnwrap)
		printf("				t.Errorf(\"combination %v/%v: %v() failed\")\n", i+1, numCombinations, g.FunctionNameUnwrap)
		printf("			}\n")
		printf("		} else {\n")
		printf("			t.Errorf(\"combination %v/%v: %v interface not implemented\")\n", i+1, numCombinations, g.InterfaceNameUnwrap)
		printf("		}\n")

		printf("	}\n")
	}
	printf("}\n")

	// Write the header at last to ensure we have captured all imports.
	writtenSoFar := buf.Bytes()
	buf = bytes.Buffer{}
	printf("package %v\n", g.PackageName)
	printf("\n")

	printf("import (\n")
	for _, i := range packages.Imports() {
		printf("	%v %q\n", i.PackageName, i.PackagePath)
	}
	printf(")\n")
	printf("\n")
	_, _ = buf.Write(writtenSoFar)

	codeBytes, err := format.Source(buf.Bytes())
	if err != nil {
		return
	}
	code = string(codeBytes)
	return
}

// GenerateGoSourceFile generates the Go source file for the wrapper.
// You are supposed to write it to a file.
func (g *T) GenerateGoSourceFile() (code string, err error) {
	err = g.validate()
	if err != nil {
		return
	}

	baseInterface, additionalInterfaces, err := g.prepareInterfaces()
	if err != nil {
		return
	}

	allMethods, err := g.prepareMethods(baseInterface, additionalInterfaces)
	if err != nil {
		return
	}

	packages := newPackages()

	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	var buf bytes.Buffer
	printf := func(format string, values ...any) {
		_, err := fmt.Fprintf(&buf, format, values...)
		if err != nil {
			panic(err)
		}
	}

	printf("type %v interface {\n", g.InterfaceNameUnwrap)
	printf("	%v() %v", g.FunctionNameUnwrap, packages.TypeName(baseInterface))
	printf("}\n")
	printf("\n")

	printf("func %v(w %v) %v {\n", g.FunctionNameUnwrap, packages.TypeName(baseInterface), packages.TypeName(baseInterface))
	printf("	if ww, ok := w.(%v); ok {\n", g.InterfaceNameUnwrap)
	printf("		return %v(ww.%v())\n", g.FunctionNameUnwrap, g.FunctionNameUnwrap)
	printf("	} else {\n")
	printf("		return w\n")
	printf("	}\n")
	printf("}\n")
	printf("\n")

	for _, m := range allMethods {
		printf("type %v%v %v\n", g.FunctionTypeNamePrefix, m.Name, packages.TypeName(m.Type))
	}
	printf("\n")

	printf("type %v struct {\n", g.StructNameInterceptor)
	for _, m := range allMethods {
		printf("	%v func(%v%v) %v%v\n", m.Name, g.FunctionTypeNamePrefix, m.Name, g.FunctionTypeNamePrefix, m.Name)
	}
	printf("}\n")
	printf("\n")

	printf("type %v struct {\n", g.StructNameWrapper)
	printf("	wrapped	%v\n", packages.TypeName(baseInterface))
	printf("	interceptor	%v\n", g.StructNameInterceptor)
	printf("}\n")
	printf("\n")

	printf("func (w *%v) %v() %v {\n", g.StructNameWrapper, g.FunctionNameUnwrap, packages.TypeName(baseInterface))
	printf("	return w.wrapped\n")
	printf("}\n")
	printf("\n")

	for mIdx, m := range allMethods {
		if mIdx != 0 {
			printf("\n")
		}

		printf("func (w *%v) %v(", g.StructNameWrapper, m.Name)
		isVariadic := m.Type.IsVariadic()
		for idx, arg := range m.In {
			if idx != 0 {
				printf(", ")
			}
			if idx == len(m.In)-1 && isVariadic {
				printf("a%v ...%v", idx, packages.TypeName(arg.Elem()))
			} else {
				printf("a%v %v", idx, packages.TypeName(arg))
			}
		}
		printf(")")

		switch len(m.Out) {
		case 0:
			break
		case 1:
			printf(" %v", packages.TypeName(m.Out[0]))
		default:
			printf(" (")
			for idx, ret := range m.Out {
				if idx != 0 {
					printf(", ")
				}
				printf("%v", packages.TypeName(ret))
			}
			printf(")")
		}
		printf(" {\n")
		printf("	f := w.wrapped.(%v).%v\n", packages.TypeName(m.Interface), m.Name)
		printf("	if w.interceptor.%v != nil {\n", m.Name)
		printf("		f = w.interceptor.%v(f)\n", m.Name)
		printf("	}\n")
		switch len(m.Out) {
		case 0:
			printf("	f(")
		default:
			printf("	return f(")
		}
		for idx := range m.In {
			if idx != 0 {
				printf(", ")
			}
			printf("a%v", idx)
		}
		printf(")\n")

		printf("}\n")
	}
	printf("\n")

	numBit := len(additionalInterfaces)
	var numCombinations uint64 = uint64(1) << numBit
	printf("func %v(wrapped %v, interceptor %v) %v {\n",
		g.FunctionNameWrap,
		packages.TypeName(baseInterface),
		g.StructNameInterceptor,
		packages.TypeName(baseInterface),
	)
	printf("	w := &%v{wrapped: wrapped, interceptor: interceptor}\n", g.StructNameWrapper)
	for idx, iface := range additionalInterfaces {
		printf("	_, ok%v := wrapped.(%v)\n", idx, packages.TypeName(iface))
	}
	if numCombinations <= 1 {
		printf("	return struct {\n")
		printf("		%v\n", g.InterfaceNameUnwrap)
		printf("		%v\n", packages.TypeName(baseInterface))
		printf("	}{w, w}\n")
	} else {
		printf("	switch {\n")
		for i := uint64(0); i < numCombinations; i += 1 {
			printf("	// combination %v/%v\n", i+1, numCombinations)

			allInterfaces := []reflect.Type{baseInterface}
			for bit := 0; bit < numBit; bit += 1 {
				val := uint64(1) << bit
				b := (i & val) == 0
				if !b {
					allInterfaces = append(allInterfaces, additionalInterfaces[bit])
				}
			}

			interfaceMethodsByBit := make([][]interfaceMethod, numBit)
			for bit := 0; bit < numBit; bit += 1 {
				val := uint64(1) << bit
				b := (i & val) == 0
				if !b {
					original := additionalInterfaces[bit]
					interfaceMethods, ok := g.fixInterface(allInterfaces, original)
					if ok {
						interfaceMethodsByBit[bit] = interfaceMethods
					}
				}
			}

			printf("	case ")
			for bit := 0; bit < numBit; bit += 1 {
				if bit != 0 {
					printf(" && ")
				}
				val := uint64(1) << bit
				b := (i & val) == 0
				if !b {
					printf("ok%v", bit)
				} else {
					printf("!ok%v", bit)
				}
			}
			printf(":\n")

			for bit, interfaceMethods := range interfaceMethodsByBit {
				if interfaceMethods != nil {
					printf("		type combination_%v_%v %v\n", i, bit, packages.typeNameMethods(interfaceMethods))
				}
			}

			printf("		return struct {\n")
			printf("			%v\n", g.InterfaceNameUnwrap)
			printf("			%v\n", packages.TypeName(baseInterface))
			for bit := 0; bit < numBit; bit += 1 {
				val := uint64(1) << bit
				b := (i & val) == 0
				if !b {
					original := additionalInterfaces[bit]
					interfaceMethods := interfaceMethodsByBit[bit]
					if interfaceMethods == nil {
						printf("			%v\n", packages.TypeName(original))
					} else {
						printf("			combination_%v_%v\n", i, bit)
					}
				}
			}
			printf("		}{w, w")
			for bit := 0; bit < numBit; bit += 1 {
				val := uint64(1) << bit
				b := (i & val) == 0
				if !b {
					printf(", w")
				}
			}
			printf("}\n")
		}
		printf("	}\n")
		printf("	panic(\"unreachable\")\n")
	}
	printf("}\n")

	// Write the header at last to ensure we have captured all imports.
	writtenSoFar := buf.Bytes()
	buf = bytes.Buffer{}
	printf("package %v\n", g.PackageName)
	printf("\n")

	printf("import (\n")
	for _, i := range packages.Imports() {
		printf("	%v %q\n", i.PackageName, i.PackagePath)
	}
	printf(")\n")
	printf("\n")
	_, _ = buf.Write(writtenSoFar)

	codeBytes, err := format.Source(buf.Bytes())
	if err != nil {
		return
	}
	code = string(codeBytes)
	return
}
