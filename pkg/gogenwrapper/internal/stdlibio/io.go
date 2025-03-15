package stdlibio

import (
	io "io"
)

type Unwrapper interface {
	Unwrap() any
}

func Unwrap(w any) any {
	if ww, ok := w.(Unwrapper); ok {
		return Unwrap(ww.Unwrap())
	} else {
		return w
	}
}

type Close func() error
type Read func([]uint8) (int, error)
type ReadAt func([]uint8, int64) (int, error)
type ReadFrom func(io.Reader) (int64, error)
type Seek func(int64, int) (int64, error)
type Write func([]uint8) (int, error)
type WriteAt func([]uint8, int64) (int, error)
type WriteTo func(io.Writer) (int64, error)

type Interceptor struct {
	Close    func(Close) Close
	Read     func(Read) Read
	ReadAt   func(ReadAt) ReadAt
	ReadFrom func(ReadFrom) ReadFrom
	Seek     func(Seek) Seek
	Write    func(Write) Write
	WriteAt  func(WriteAt) WriteAt
	WriteTo  func(WriteTo) WriteTo
}

type Wrapper struct {
	wrapped     any
	interceptor Interceptor
}

func (w *Wrapper) Unwrap() any {
	return w.wrapped
}

func (w *Wrapper) Close() error {
	f := w.wrapped.(io.Closer).Close
	if w.interceptor.Close != nil {
		f = w.interceptor.Close(f)
	}
	return f()
}

func (w *Wrapper) Read(a0 []uint8) (int, error) {
	f := w.wrapped.(io.Reader).Read
	if w.interceptor.Read != nil {
		f = w.interceptor.Read(f)
	}
	return f(a0)
}

func (w *Wrapper) ReadAt(a0 []uint8, a1 int64) (int, error) {
	f := w.wrapped.(io.ReaderAt).ReadAt
	if w.interceptor.ReadAt != nil {
		f = w.interceptor.ReadAt(f)
	}
	return f(a0, a1)
}

func (w *Wrapper) ReadFrom(a0 io.Reader) (int64, error) {
	f := w.wrapped.(io.ReaderFrom).ReadFrom
	if w.interceptor.ReadFrom != nil {
		f = w.interceptor.ReadFrom(f)
	}
	return f(a0)
}

func (w *Wrapper) Seek(a0 int64, a1 int) (int64, error) {
	f := w.wrapped.(io.Seeker).Seek
	if w.interceptor.Seek != nil {
		f = w.interceptor.Seek(f)
	}
	return f(a0, a1)
}

func (w *Wrapper) Write(a0 []uint8) (int, error) {
	f := w.wrapped.(io.Writer).Write
	if w.interceptor.Write != nil {
		f = w.interceptor.Write(f)
	}
	return f(a0)
}

func (w *Wrapper) WriteAt(a0 []uint8, a1 int64) (int, error) {
	f := w.wrapped.(io.WriterAt).WriteAt
	if w.interceptor.WriteAt != nil {
		f = w.interceptor.WriteAt(f)
	}
	return f(a0, a1)
}

func (w *Wrapper) WriteTo(a0 io.Writer) (int64, error) {
	f := w.wrapped.(io.WriterTo).WriteTo
	if w.interceptor.WriteTo != nil {
		f = w.interceptor.WriteTo(f)
	}
	return f(a0)
}

func Wrap(wrapped any, interceptor Interceptor) any {
	w := &Wrapper{wrapped: wrapped, interceptor: interceptor}
	_, ok0 := wrapped.(io.Closer)
	_, ok1 := wrapped.(io.Reader)
	_, ok2 := wrapped.(io.ReaderAt)
	_, ok3 := wrapped.(io.ReaderFrom)
	_, ok4 := wrapped.(io.Seeker)
	_, ok5 := wrapped.(io.Writer)
	_, ok6 := wrapped.(io.WriterAt)
	_, ok7 := wrapped.(io.WriterTo)
	switch {
	// combination 1/256
	case !ok0 && !ok1 && !ok2 && !ok3 && !ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
		}{w, w}
	// combination 2/256
	case ok0 && !ok1 && !ok2 && !ok3 && !ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
		}{w, w, w}
	// combination 3/256
	case !ok0 && ok1 && !ok2 && !ok3 && !ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
		}{w, w, w}
	// combination 4/256
	case ok0 && ok1 && !ok2 && !ok3 && !ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
		}{w, w, w, w}
	// combination 5/256
	case !ok0 && !ok1 && ok2 && !ok3 && !ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
		}{w, w, w}
	// combination 6/256
	case ok0 && !ok1 && ok2 && !ok3 && !ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
		}{w, w, w, w}
	// combination 7/256
	case !ok0 && ok1 && ok2 && !ok3 && !ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
		}{w, w, w, w}
	// combination 8/256
	case ok0 && ok1 && ok2 && !ok3 && !ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
		}{w, w, w, w, w}
	// combination 9/256
	case !ok0 && !ok1 && !ok2 && ok3 && !ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.ReaderFrom
		}{w, w, w}
	// combination 10/256
	case ok0 && !ok1 && !ok2 && ok3 && !ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderFrom
		}{w, w, w, w}
	// combination 11/256
	case !ok0 && ok1 && !ok2 && ok3 && !ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderFrom
		}{w, w, w, w}
	// combination 12/256
	case ok0 && ok1 && !ok2 && ok3 && !ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderFrom
		}{w, w, w, w, w}
	// combination 13/256
	case !ok0 && !ok1 && ok2 && ok3 && !ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
			io.ReaderFrom
		}{w, w, w, w}
	// combination 14/256
	case ok0 && !ok1 && ok2 && ok3 && !ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
		}{w, w, w, w, w}
	// combination 15/256
	case !ok0 && ok1 && ok2 && ok3 && !ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
		}{w, w, w, w, w}
	// combination 16/256
	case ok0 && ok1 && ok2 && ok3 && !ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
		}{w, w, w, w, w, w}
	// combination 17/256
	case !ok0 && !ok1 && !ok2 && !ok3 && ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Seeker
		}{w, w, w}
	// combination 18/256
	case ok0 && !ok1 && !ok2 && !ok3 && ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Seeker
		}{w, w, w, w}
	// combination 19/256
	case !ok0 && ok1 && !ok2 && !ok3 && ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.Seeker
		}{w, w, w, w}
	// combination 20/256
	case ok0 && ok1 && !ok2 && !ok3 && ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.Seeker
		}{w, w, w, w, w}
	// combination 21/256
	case !ok0 && !ok1 && ok2 && !ok3 && ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
			io.Seeker
		}{w, w, w, w}
	// combination 22/256
	case ok0 && !ok1 && ok2 && !ok3 && ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
			io.Seeker
		}{w, w, w, w, w}
	// combination 23/256
	case !ok0 && ok1 && ok2 && !ok3 && ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
			io.Seeker
		}{w, w, w, w, w}
	// combination 24/256
	case ok0 && ok1 && ok2 && !ok3 && ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.Seeker
		}{w, w, w, w, w, w}
	// combination 25/256
	case !ok0 && !ok1 && !ok2 && ok3 && ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.ReaderFrom
			io.Seeker
		}{w, w, w, w}
	// combination 26/256
	case ok0 && !ok1 && !ok2 && ok3 && ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderFrom
			io.Seeker
		}{w, w, w, w, w}
	// combination 27/256
	case !ok0 && ok1 && !ok2 && ok3 && ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderFrom
			io.Seeker
		}{w, w, w, w, w}
	// combination 28/256
	case ok0 && ok1 && !ok2 && ok3 && ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderFrom
			io.Seeker
		}{w, w, w, w, w, w}
	// combination 29/256
	case !ok0 && !ok1 && ok2 && ok3 && ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
		}{w, w, w, w, w}
	// combination 30/256
	case ok0 && !ok1 && ok2 && ok3 && ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
		}{w, w, w, w, w, w}
	// combination 31/256
	case !ok0 && ok1 && ok2 && ok3 && ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
		}{w, w, w, w, w, w}
	// combination 32/256
	case ok0 && ok1 && ok2 && ok3 && ok4 && !ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
		}{w, w, w, w, w, w, w}
	// combination 33/256
	case !ok0 && !ok1 && !ok2 && !ok3 && !ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Writer
		}{w, w, w}
	// combination 34/256
	case ok0 && !ok1 && !ok2 && !ok3 && !ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Writer
		}{w, w, w, w}
	// combination 35/256
	case !ok0 && ok1 && !ok2 && !ok3 && !ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.Writer
		}{w, w, w, w}
	// combination 36/256
	case ok0 && ok1 && !ok2 && !ok3 && !ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.Writer
		}{w, w, w, w, w}
	// combination 37/256
	case !ok0 && !ok1 && ok2 && !ok3 && !ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
			io.Writer
		}{w, w, w, w}
	// combination 38/256
	case ok0 && !ok1 && ok2 && !ok3 && !ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
			io.Writer
		}{w, w, w, w, w}
	// combination 39/256
	case !ok0 && ok1 && ok2 && !ok3 && !ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
			io.Writer
		}{w, w, w, w, w}
	// combination 40/256
	case ok0 && ok1 && ok2 && !ok3 && !ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.Writer
		}{w, w, w, w, w, w}
	// combination 41/256
	case !ok0 && !ok1 && !ok2 && ok3 && !ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.ReaderFrom
			io.Writer
		}{w, w, w, w}
	// combination 42/256
	case ok0 && !ok1 && !ok2 && ok3 && !ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderFrom
			io.Writer
		}{w, w, w, w, w}
	// combination 43/256
	case !ok0 && ok1 && !ok2 && ok3 && !ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderFrom
			io.Writer
		}{w, w, w, w, w}
	// combination 44/256
	case ok0 && ok1 && !ok2 && ok3 && !ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderFrom
			io.Writer
		}{w, w, w, w, w, w}
	// combination 45/256
	case !ok0 && !ok1 && ok2 && ok3 && !ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
			io.ReaderFrom
			io.Writer
		}{w, w, w, w, w}
	// combination 46/256
	case ok0 && !ok1 && ok2 && ok3 && !ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
			io.Writer
		}{w, w, w, w, w, w}
	// combination 47/256
	case !ok0 && ok1 && ok2 && ok3 && !ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Writer
		}{w, w, w, w, w, w}
	// combination 48/256
	case ok0 && ok1 && ok2 && ok3 && !ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Writer
		}{w, w, w, w, w, w, w}
	// combination 49/256
	case !ok0 && !ok1 && !ok2 && !ok3 && ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Seeker
			io.Writer
		}{w, w, w, w}
	// combination 50/256
	case ok0 && !ok1 && !ok2 && !ok3 && ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Seeker
			io.Writer
		}{w, w, w, w, w}
	// combination 51/256
	case !ok0 && ok1 && !ok2 && !ok3 && ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.Seeker
			io.Writer
		}{w, w, w, w, w}
	// combination 52/256
	case ok0 && ok1 && !ok2 && !ok3 && ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.Seeker
			io.Writer
		}{w, w, w, w, w, w}
	// combination 53/256
	case !ok0 && !ok1 && ok2 && !ok3 && ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
			io.Seeker
			io.Writer
		}{w, w, w, w, w}
	// combination 54/256
	case ok0 && !ok1 && ok2 && !ok3 && ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
			io.Seeker
			io.Writer
		}{w, w, w, w, w, w}
	// combination 55/256
	case !ok0 && ok1 && ok2 && !ok3 && ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
			io.Seeker
			io.Writer
		}{w, w, w, w, w, w}
	// combination 56/256
	case ok0 && ok1 && ok2 && !ok3 && ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.Seeker
			io.Writer
		}{w, w, w, w, w, w, w}
	// combination 57/256
	case !ok0 && !ok1 && !ok2 && ok3 && ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.ReaderFrom
			io.Seeker
			io.Writer
		}{w, w, w, w, w}
	// combination 58/256
	case ok0 && !ok1 && !ok2 && ok3 && ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderFrom
			io.Seeker
			io.Writer
		}{w, w, w, w, w, w}
	// combination 59/256
	case !ok0 && ok1 && !ok2 && ok3 && ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderFrom
			io.Seeker
			io.Writer
		}{w, w, w, w, w, w}
	// combination 60/256
	case ok0 && ok1 && !ok2 && ok3 && ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderFrom
			io.Seeker
			io.Writer
		}{w, w, w, w, w, w, w}
	// combination 61/256
	case !ok0 && !ok1 && ok2 && ok3 && ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
		}{w, w, w, w, w, w}
	// combination 62/256
	case ok0 && !ok1 && ok2 && ok3 && ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
		}{w, w, w, w, w, w, w}
	// combination 63/256
	case !ok0 && ok1 && ok2 && ok3 && ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
		}{w, w, w, w, w, w, w}
	// combination 64/256
	case ok0 && ok1 && ok2 && ok3 && ok4 && ok5 && !ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
		}{w, w, w, w, w, w, w, w}
	// combination 65/256
	case !ok0 && !ok1 && !ok2 && !ok3 && !ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.WriterAt
		}{w, w, w}
	// combination 66/256
	case ok0 && !ok1 && !ok2 && !ok3 && !ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.WriterAt
		}{w, w, w, w}
	// combination 67/256
	case !ok0 && ok1 && !ok2 && !ok3 && !ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.WriterAt
		}{w, w, w, w}
	// combination 68/256
	case ok0 && ok1 && !ok2 && !ok3 && !ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.WriterAt
		}{w, w, w, w, w}
	// combination 69/256
	case !ok0 && !ok1 && ok2 && !ok3 && !ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
			io.WriterAt
		}{w, w, w, w}
	// combination 70/256
	case ok0 && !ok1 && ok2 && !ok3 && !ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
			io.WriterAt
		}{w, w, w, w, w}
	// combination 71/256
	case !ok0 && ok1 && ok2 && !ok3 && !ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
			io.WriterAt
		}{w, w, w, w, w}
	// combination 72/256
	case ok0 && ok1 && ok2 && !ok3 && !ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.WriterAt
		}{w, w, w, w, w, w}
	// combination 73/256
	case !ok0 && !ok1 && !ok2 && ok3 && !ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.ReaderFrom
			io.WriterAt
		}{w, w, w, w}
	// combination 74/256
	case ok0 && !ok1 && !ok2 && ok3 && !ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderFrom
			io.WriterAt
		}{w, w, w, w, w}
	// combination 75/256
	case !ok0 && ok1 && !ok2 && ok3 && !ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderFrom
			io.WriterAt
		}{w, w, w, w, w}
	// combination 76/256
	case ok0 && ok1 && !ok2 && ok3 && !ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderFrom
			io.WriterAt
		}{w, w, w, w, w, w}
	// combination 77/256
	case !ok0 && !ok1 && ok2 && ok3 && !ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
			io.ReaderFrom
			io.WriterAt
		}{w, w, w, w, w}
	// combination 78/256
	case ok0 && !ok1 && ok2 && ok3 && !ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
			io.WriterAt
		}{w, w, w, w, w, w}
	// combination 79/256
	case !ok0 && ok1 && ok2 && ok3 && !ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.WriterAt
		}{w, w, w, w, w, w}
	// combination 80/256
	case ok0 && ok1 && ok2 && ok3 && !ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.WriterAt
		}{w, w, w, w, w, w, w}
	// combination 81/256
	case !ok0 && !ok1 && !ok2 && !ok3 && ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Seeker
			io.WriterAt
		}{w, w, w, w}
	// combination 82/256
	case ok0 && !ok1 && !ok2 && !ok3 && ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Seeker
			io.WriterAt
		}{w, w, w, w, w}
	// combination 83/256
	case !ok0 && ok1 && !ok2 && !ok3 && ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.Seeker
			io.WriterAt
		}{w, w, w, w, w}
	// combination 84/256
	case ok0 && ok1 && !ok2 && !ok3 && ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.Seeker
			io.WriterAt
		}{w, w, w, w, w, w}
	// combination 85/256
	case !ok0 && !ok1 && ok2 && !ok3 && ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
			io.Seeker
			io.WriterAt
		}{w, w, w, w, w}
	// combination 86/256
	case ok0 && !ok1 && ok2 && !ok3 && ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
			io.Seeker
			io.WriterAt
		}{w, w, w, w, w, w}
	// combination 87/256
	case !ok0 && ok1 && ok2 && !ok3 && ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
			io.Seeker
			io.WriterAt
		}{w, w, w, w, w, w}
	// combination 88/256
	case ok0 && ok1 && ok2 && !ok3 && ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.Seeker
			io.WriterAt
		}{w, w, w, w, w, w, w}
	// combination 89/256
	case !ok0 && !ok1 && !ok2 && ok3 && ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.ReaderFrom
			io.Seeker
			io.WriterAt
		}{w, w, w, w, w}
	// combination 90/256
	case ok0 && !ok1 && !ok2 && ok3 && ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderFrom
			io.Seeker
			io.WriterAt
		}{w, w, w, w, w, w}
	// combination 91/256
	case !ok0 && ok1 && !ok2 && ok3 && ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderFrom
			io.Seeker
			io.WriterAt
		}{w, w, w, w, w, w}
	// combination 92/256
	case ok0 && ok1 && !ok2 && ok3 && ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderFrom
			io.Seeker
			io.WriterAt
		}{w, w, w, w, w, w, w}
	// combination 93/256
	case !ok0 && !ok1 && ok2 && ok3 && ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.WriterAt
		}{w, w, w, w, w, w}
	// combination 94/256
	case ok0 && !ok1 && ok2 && ok3 && ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.WriterAt
		}{w, w, w, w, w, w, w}
	// combination 95/256
	case !ok0 && ok1 && ok2 && ok3 && ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.WriterAt
		}{w, w, w, w, w, w, w}
	// combination 96/256
	case ok0 && ok1 && ok2 && ok3 && ok4 && !ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.WriterAt
		}{w, w, w, w, w, w, w, w}
	// combination 97/256
	case !ok0 && !ok1 && !ok2 && !ok3 && !ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Writer
			io.WriterAt
		}{w, w, w, w}
	// combination 98/256
	case ok0 && !ok1 && !ok2 && !ok3 && !ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Writer
			io.WriterAt
		}{w, w, w, w, w}
	// combination 99/256
	case !ok0 && ok1 && !ok2 && !ok3 && !ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.Writer
			io.WriterAt
		}{w, w, w, w, w}
	// combination 100/256
	case ok0 && ok1 && !ok2 && !ok3 && !ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.Writer
			io.WriterAt
		}{w, w, w, w, w, w}
	// combination 101/256
	case !ok0 && !ok1 && ok2 && !ok3 && !ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
			io.Writer
			io.WriterAt
		}{w, w, w, w, w}
	// combination 102/256
	case ok0 && !ok1 && ok2 && !ok3 && !ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
			io.Writer
			io.WriterAt
		}{w, w, w, w, w, w}
	// combination 103/256
	case !ok0 && ok1 && ok2 && !ok3 && !ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
			io.Writer
			io.WriterAt
		}{w, w, w, w, w, w}
	// combination 104/256
	case ok0 && ok1 && ok2 && !ok3 && !ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.Writer
			io.WriterAt
		}{w, w, w, w, w, w, w}
	// combination 105/256
	case !ok0 && !ok1 && !ok2 && ok3 && !ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.ReaderFrom
			io.Writer
			io.WriterAt
		}{w, w, w, w, w}
	// combination 106/256
	case ok0 && !ok1 && !ok2 && ok3 && !ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderFrom
			io.Writer
			io.WriterAt
		}{w, w, w, w, w, w}
	// combination 107/256
	case !ok0 && ok1 && !ok2 && ok3 && !ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderFrom
			io.Writer
			io.WriterAt
		}{w, w, w, w, w, w}
	// combination 108/256
	case ok0 && ok1 && !ok2 && ok3 && !ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderFrom
			io.Writer
			io.WriterAt
		}{w, w, w, w, w, w, w}
	// combination 109/256
	case !ok0 && !ok1 && ok2 && ok3 && !ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
			io.ReaderFrom
			io.Writer
			io.WriterAt
		}{w, w, w, w, w, w}
	// combination 110/256
	case ok0 && !ok1 && ok2 && ok3 && !ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
			io.Writer
			io.WriterAt
		}{w, w, w, w, w, w, w}
	// combination 111/256
	case !ok0 && ok1 && ok2 && ok3 && !ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Writer
			io.WriterAt
		}{w, w, w, w, w, w, w}
	// combination 112/256
	case ok0 && ok1 && ok2 && ok3 && !ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Writer
			io.WriterAt
		}{w, w, w, w, w, w, w, w}
	// combination 113/256
	case !ok0 && !ok1 && !ok2 && !ok3 && ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Seeker
			io.Writer
			io.WriterAt
		}{w, w, w, w, w}
	// combination 114/256
	case ok0 && !ok1 && !ok2 && !ok3 && ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Seeker
			io.Writer
			io.WriterAt
		}{w, w, w, w, w, w}
	// combination 115/256
	case !ok0 && ok1 && !ok2 && !ok3 && ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.Seeker
			io.Writer
			io.WriterAt
		}{w, w, w, w, w, w}
	// combination 116/256
	case ok0 && ok1 && !ok2 && !ok3 && ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.Seeker
			io.Writer
			io.WriterAt
		}{w, w, w, w, w, w, w}
	// combination 117/256
	case !ok0 && !ok1 && ok2 && !ok3 && ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
			io.Seeker
			io.Writer
			io.WriterAt
		}{w, w, w, w, w, w}
	// combination 118/256
	case ok0 && !ok1 && ok2 && !ok3 && ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
			io.Seeker
			io.Writer
			io.WriterAt
		}{w, w, w, w, w, w, w}
	// combination 119/256
	case !ok0 && ok1 && ok2 && !ok3 && ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
			io.Seeker
			io.Writer
			io.WriterAt
		}{w, w, w, w, w, w, w}
	// combination 120/256
	case ok0 && ok1 && ok2 && !ok3 && ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.Seeker
			io.Writer
			io.WriterAt
		}{w, w, w, w, w, w, w, w}
	// combination 121/256
	case !ok0 && !ok1 && !ok2 && ok3 && ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
		}{w, w, w, w, w, w}
	// combination 122/256
	case ok0 && !ok1 && !ok2 && ok3 && ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
		}{w, w, w, w, w, w, w}
	// combination 123/256
	case !ok0 && ok1 && !ok2 && ok3 && ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
		}{w, w, w, w, w, w, w}
	// combination 124/256
	case ok0 && ok1 && !ok2 && ok3 && ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
		}{w, w, w, w, w, w, w, w}
	// combination 125/256
	case !ok0 && !ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
		}{w, w, w, w, w, w, w}
	// combination 126/256
	case ok0 && !ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
		}{w, w, w, w, w, w, w, w}
	// combination 127/256
	case !ok0 && ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
		}{w, w, w, w, w, w, w, w}
	// combination 128/256
	case ok0 && ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && !ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
		}{w, w, w, w, w, w, w, w, w}
	// combination 129/256
	case !ok0 && !ok1 && !ok2 && !ok3 && !ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.WriterTo
		}{w, w, w}
	// combination 130/256
	case ok0 && !ok1 && !ok2 && !ok3 && !ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.WriterTo
		}{w, w, w, w}
	// combination 131/256
	case !ok0 && ok1 && !ok2 && !ok3 && !ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.WriterTo
		}{w, w, w, w}
	// combination 132/256
	case ok0 && ok1 && !ok2 && !ok3 && !ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.WriterTo
		}{w, w, w, w, w}
	// combination 133/256
	case !ok0 && !ok1 && ok2 && !ok3 && !ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
			io.WriterTo
		}{w, w, w, w}
	// combination 134/256
	case ok0 && !ok1 && ok2 && !ok3 && !ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
			io.WriterTo
		}{w, w, w, w, w}
	// combination 135/256
	case !ok0 && ok1 && ok2 && !ok3 && !ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
			io.WriterTo
		}{w, w, w, w, w}
	// combination 136/256
	case ok0 && ok1 && ok2 && !ok3 && !ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 137/256
	case !ok0 && !ok1 && !ok2 && ok3 && !ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.ReaderFrom
			io.WriterTo
		}{w, w, w, w}
	// combination 138/256
	case ok0 && !ok1 && !ok2 && ok3 && !ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderFrom
			io.WriterTo
		}{w, w, w, w, w}
	// combination 139/256
	case !ok0 && ok1 && !ok2 && ok3 && !ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderFrom
			io.WriterTo
		}{w, w, w, w, w}
	// combination 140/256
	case ok0 && ok1 && !ok2 && ok3 && !ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderFrom
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 141/256
	case !ok0 && !ok1 && ok2 && ok3 && !ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
			io.ReaderFrom
			io.WriterTo
		}{w, w, w, w, w}
	// combination 142/256
	case ok0 && !ok1 && ok2 && ok3 && !ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 143/256
	case !ok0 && ok1 && ok2 && ok3 && !ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 144/256
	case ok0 && ok1 && ok2 && ok3 && !ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 145/256
	case !ok0 && !ok1 && !ok2 && !ok3 && ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Seeker
			io.WriterTo
		}{w, w, w, w}
	// combination 146/256
	case ok0 && !ok1 && !ok2 && !ok3 && ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Seeker
			io.WriterTo
		}{w, w, w, w, w}
	// combination 147/256
	case !ok0 && ok1 && !ok2 && !ok3 && ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.Seeker
			io.WriterTo
		}{w, w, w, w, w}
	// combination 148/256
	case ok0 && ok1 && !ok2 && !ok3 && ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.Seeker
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 149/256
	case !ok0 && !ok1 && ok2 && !ok3 && ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
			io.Seeker
			io.WriterTo
		}{w, w, w, w, w}
	// combination 150/256
	case ok0 && !ok1 && ok2 && !ok3 && ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
			io.Seeker
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 151/256
	case !ok0 && ok1 && ok2 && !ok3 && ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
			io.Seeker
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 152/256
	case ok0 && ok1 && ok2 && !ok3 && ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.Seeker
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 153/256
	case !ok0 && !ok1 && !ok2 && ok3 && ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.ReaderFrom
			io.Seeker
			io.WriterTo
		}{w, w, w, w, w}
	// combination 154/256
	case ok0 && !ok1 && !ok2 && ok3 && ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderFrom
			io.Seeker
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 155/256
	case !ok0 && ok1 && !ok2 && ok3 && ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderFrom
			io.Seeker
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 156/256
	case ok0 && ok1 && !ok2 && ok3 && ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderFrom
			io.Seeker
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 157/256
	case !ok0 && !ok1 && ok2 && ok3 && ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 158/256
	case ok0 && !ok1 && ok2 && ok3 && ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 159/256
	case !ok0 && ok1 && ok2 && ok3 && ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 160/256
	case ok0 && ok1 && ok2 && ok3 && ok4 && !ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.WriterTo
		}{w, w, w, w, w, w, w, w}
	// combination 161/256
	case !ok0 && !ok1 && !ok2 && !ok3 && !ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Writer
			io.WriterTo
		}{w, w, w, w}
	// combination 162/256
	case ok0 && !ok1 && !ok2 && !ok3 && !ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Writer
			io.WriterTo
		}{w, w, w, w, w}
	// combination 163/256
	case !ok0 && ok1 && !ok2 && !ok3 && !ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.Writer
			io.WriterTo
		}{w, w, w, w, w}
	// combination 164/256
	case ok0 && ok1 && !ok2 && !ok3 && !ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.Writer
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 165/256
	case !ok0 && !ok1 && ok2 && !ok3 && !ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
			io.Writer
			io.WriterTo
		}{w, w, w, w, w}
	// combination 166/256
	case ok0 && !ok1 && ok2 && !ok3 && !ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
			io.Writer
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 167/256
	case !ok0 && ok1 && ok2 && !ok3 && !ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
			io.Writer
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 168/256
	case ok0 && ok1 && ok2 && !ok3 && !ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.Writer
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 169/256
	case !ok0 && !ok1 && !ok2 && ok3 && !ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.ReaderFrom
			io.Writer
			io.WriterTo
		}{w, w, w, w, w}
	// combination 170/256
	case ok0 && !ok1 && !ok2 && ok3 && !ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderFrom
			io.Writer
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 171/256
	case !ok0 && ok1 && !ok2 && ok3 && !ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderFrom
			io.Writer
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 172/256
	case ok0 && ok1 && !ok2 && ok3 && !ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderFrom
			io.Writer
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 173/256
	case !ok0 && !ok1 && ok2 && ok3 && !ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
			io.ReaderFrom
			io.Writer
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 174/256
	case ok0 && !ok1 && ok2 && ok3 && !ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
			io.Writer
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 175/256
	case !ok0 && ok1 && ok2 && ok3 && !ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Writer
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 176/256
	case ok0 && ok1 && ok2 && ok3 && !ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Writer
			io.WriterTo
		}{w, w, w, w, w, w, w, w}
	// combination 177/256
	case !ok0 && !ok1 && !ok2 && !ok3 && ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Seeker
			io.Writer
			io.WriterTo
		}{w, w, w, w, w}
	// combination 178/256
	case ok0 && !ok1 && !ok2 && !ok3 && ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Seeker
			io.Writer
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 179/256
	case !ok0 && ok1 && !ok2 && !ok3 && ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.Seeker
			io.Writer
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 180/256
	case ok0 && ok1 && !ok2 && !ok3 && ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.Seeker
			io.Writer
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 181/256
	case !ok0 && !ok1 && ok2 && !ok3 && ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
			io.Seeker
			io.Writer
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 182/256
	case ok0 && !ok1 && ok2 && !ok3 && ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
			io.Seeker
			io.Writer
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 183/256
	case !ok0 && ok1 && ok2 && !ok3 && ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
			io.Seeker
			io.Writer
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 184/256
	case ok0 && ok1 && ok2 && !ok3 && ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.Seeker
			io.Writer
			io.WriterTo
		}{w, w, w, w, w, w, w, w}
	// combination 185/256
	case !ok0 && !ok1 && !ok2 && ok3 && ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 186/256
	case ok0 && !ok1 && !ok2 && ok3 && ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 187/256
	case !ok0 && ok1 && !ok2 && ok3 && ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 188/256
	case ok0 && ok1 && !ok2 && ok3 && ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterTo
		}{w, w, w, w, w, w, w, w}
	// combination 189/256
	case !ok0 && !ok1 && ok2 && ok3 && ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 190/256
	case ok0 && !ok1 && ok2 && ok3 && ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterTo
		}{w, w, w, w, w, w, w, w}
	// combination 191/256
	case !ok0 && ok1 && ok2 && ok3 && ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterTo
		}{w, w, w, w, w, w, w, w}
	// combination 192/256
	case ok0 && ok1 && ok2 && ok3 && ok4 && ok5 && !ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterTo
		}{w, w, w, w, w, w, w, w, w}
	// combination 193/256
	case !ok0 && !ok1 && !ok2 && !ok3 && !ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.WriterAt
			io.WriterTo
		}{w, w, w, w}
	// combination 194/256
	case ok0 && !ok1 && !ok2 && !ok3 && !ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w}
	// combination 195/256
	case !ok0 && ok1 && !ok2 && !ok3 && !ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w}
	// combination 196/256
	case ok0 && ok1 && !ok2 && !ok3 && !ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 197/256
	case !ok0 && !ok1 && ok2 && !ok3 && !ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w}
	// combination 198/256
	case ok0 && !ok1 && ok2 && !ok3 && !ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 199/256
	case !ok0 && ok1 && ok2 && !ok3 && !ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 200/256
	case ok0 && ok1 && ok2 && !ok3 && !ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 201/256
	case !ok0 && !ok1 && !ok2 && ok3 && !ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.ReaderFrom
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w}
	// combination 202/256
	case ok0 && !ok1 && !ok2 && ok3 && !ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderFrom
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 203/256
	case !ok0 && ok1 && !ok2 && ok3 && !ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderFrom
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 204/256
	case ok0 && ok1 && !ok2 && ok3 && !ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderFrom
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 205/256
	case !ok0 && !ok1 && ok2 && ok3 && !ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
			io.ReaderFrom
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 206/256
	case ok0 && !ok1 && ok2 && ok3 && !ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 207/256
	case !ok0 && ok1 && ok2 && ok3 && !ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 208/256
	case ok0 && ok1 && ok2 && ok3 && !ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w, w}
	// combination 209/256
	case !ok0 && !ok1 && !ok2 && !ok3 && ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w}
	// combination 210/256
	case ok0 && !ok1 && !ok2 && !ok3 && ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 211/256
	case !ok0 && ok1 && !ok2 && !ok3 && ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 212/256
	case ok0 && ok1 && !ok2 && !ok3 && ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 213/256
	case !ok0 && !ok1 && ok2 && !ok3 && ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 214/256
	case ok0 && !ok1 && ok2 && !ok3 && ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 215/256
	case !ok0 && ok1 && ok2 && !ok3 && ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 216/256
	case ok0 && ok1 && ok2 && !ok3 && ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w, w}
	// combination 217/256
	case !ok0 && !ok1 && !ok2 && ok3 && ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.ReaderFrom
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 218/256
	case ok0 && !ok1 && !ok2 && ok3 && ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderFrom
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 219/256
	case !ok0 && ok1 && !ok2 && ok3 && ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderFrom
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 220/256
	case ok0 && ok1 && !ok2 && ok3 && ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderFrom
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w, w}
	// combination 221/256
	case !ok0 && !ok1 && ok2 && ok3 && ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 222/256
	case ok0 && !ok1 && ok2 && ok3 && ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w, w}
	// combination 223/256
	case !ok0 && ok1 && ok2 && ok3 && ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w, w}
	// combination 224/256
	case ok0 && ok1 && ok2 && ok3 && ok4 && !ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w, w, w}
	// combination 225/256
	case !ok0 && !ok1 && !ok2 && !ok3 && !ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w}
	// combination 226/256
	case ok0 && !ok1 && !ok2 && !ok3 && !ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 227/256
	case !ok0 && ok1 && !ok2 && !ok3 && !ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 228/256
	case ok0 && ok1 && !ok2 && !ok3 && !ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 229/256
	case !ok0 && !ok1 && ok2 && !ok3 && !ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 230/256
	case ok0 && !ok1 && ok2 && !ok3 && !ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 231/256
	case !ok0 && ok1 && ok2 && !ok3 && !ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 232/256
	case ok0 && ok1 && ok2 && !ok3 && !ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w, w}
	// combination 233/256
	case !ok0 && !ok1 && !ok2 && ok3 && !ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.ReaderFrom
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 234/256
	case ok0 && !ok1 && !ok2 && ok3 && !ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderFrom
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 235/256
	case !ok0 && ok1 && !ok2 && ok3 && !ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderFrom
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 236/256
	case ok0 && ok1 && !ok2 && ok3 && !ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderFrom
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w, w}
	// combination 237/256
	case !ok0 && !ok1 && ok2 && ok3 && !ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
			io.ReaderFrom
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 238/256
	case ok0 && !ok1 && ok2 && ok3 && !ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w, w}
	// combination 239/256
	case !ok0 && ok1 && ok2 && ok3 && !ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w, w}
	// combination 240/256
	case ok0 && ok1 && ok2 && ok3 && !ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w, w, w}
	// combination 241/256
	case !ok0 && !ok1 && !ok2 && !ok3 && ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w}
	// combination 242/256
	case ok0 && !ok1 && !ok2 && !ok3 && ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 243/256
	case !ok0 && ok1 && !ok2 && !ok3 && ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 244/256
	case ok0 && ok1 && !ok2 && !ok3 && ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w, w}
	// combination 245/256
	case !ok0 && !ok1 && ok2 && !ok3 && ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 246/256
	case ok0 && !ok1 && ok2 && !ok3 && ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w, w}
	// combination 247/256
	case !ok0 && ok1 && ok2 && !ok3 && ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w, w}
	// combination 248/256
	case ok0 && ok1 && ok2 && !ok3 && ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w, w, w}
	// combination 249/256
	case !ok0 && !ok1 && !ok2 && ok3 && ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w}
	// combination 250/256
	case ok0 && !ok1 && !ok2 && ok3 && ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w, w}
	// combination 251/256
	case !ok0 && ok1 && !ok2 && ok3 && ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w, w}
	// combination 252/256
	case ok0 && ok1 && !ok2 && ok3 && ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w, w, w}
	// combination 253/256
	case !ok0 && !ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w, w}
	// combination 254/256
	case ok0 && !ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w, w, w}
	// combination 255/256
	case !ok0 && ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w, w, w}
	// combination 256/256
	case ok0 && ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && ok7:
		return struct {
			Unwrapper
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{w, w, w, w, w, w, w, w, w, w}
	}
	panic("unreachable")
}
