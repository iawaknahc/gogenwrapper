package nethttp

import (
	bufio "bufio"
	io "io"
	net "net"
	http "net/http"
	time "time"
)

type ResponseWriterUnwrapper interface {
	UnwrapResponseWriter() http.ResponseWriter
}

func UnwrapResponseWriter(w http.ResponseWriter) http.ResponseWriter {
	if ww, ok := w.(ResponseWriterUnwrapper); ok {
		return UnwrapResponseWriter(ww.UnwrapResponseWriter())
	} else {
		return w
	}
}

type Header func() http.Header
type Write func([]uint8) (int, error)
type WriteHeader func(int)
type Flush func()
type CloseNotify func() <-chan bool
type Hijack func() (net.Conn, *bufio.ReadWriter, error)
type ReadFrom func(io.Reader) (int64, error)
type Push func(string, *http.PushOptions) error
type SetReadDeadline func(time.Time) error
type SetWriteDeadline func(time.Time) error
type EnableFullDuplex func() error

type ResponseWriterInterceptor struct {
	Header           func(Header) Header
	Write            func(Write) Write
	WriteHeader      func(WriteHeader) WriteHeader
	Flush            func(Flush) Flush
	CloseNotify      func(CloseNotify) CloseNotify
	Hijack           func(Hijack) Hijack
	ReadFrom         func(ReadFrom) ReadFrom
	Push             func(Push) Push
	SetReadDeadline  func(SetReadDeadline) SetReadDeadline
	SetWriteDeadline func(SetWriteDeadline) SetWriteDeadline
	EnableFullDuplex func(EnableFullDuplex) EnableFullDuplex
}

type ResponseWriterWrapper struct {
	wrapped     http.ResponseWriter
	interceptor ResponseWriterInterceptor
}

func (w *ResponseWriterWrapper) UnwrapResponseWriter() http.ResponseWriter {
	return w.wrapped
}

func (w *ResponseWriterWrapper) Header() http.Header {
	f := w.wrapped.(http.ResponseWriter).Header
	if w.interceptor.Header != nil {
		f = w.interceptor.Header(f)
	}
	return f()
}

func (w *ResponseWriterWrapper) Write(a0 []uint8) (int, error) {
	f := w.wrapped.(http.ResponseWriter).Write
	if w.interceptor.Write != nil {
		f = w.interceptor.Write(f)
	}
	return f(a0)
}

func (w *ResponseWriterWrapper) WriteHeader(a0 int) {
	f := w.wrapped.(http.ResponseWriter).WriteHeader
	if w.interceptor.WriteHeader != nil {
		f = w.interceptor.WriteHeader(f)
	}
	f(a0)
}

func (w *ResponseWriterWrapper) Flush() {
	f := w.wrapped.(http.Flusher).Flush
	if w.interceptor.Flush != nil {
		f = w.interceptor.Flush(f)
	}
	f()
}

func (w *ResponseWriterWrapper) CloseNotify() <-chan bool {
	f := w.wrapped.(http.CloseNotifier).CloseNotify
	if w.interceptor.CloseNotify != nil {
		f = w.interceptor.CloseNotify(f)
	}
	return f()
}

func (w *ResponseWriterWrapper) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	f := w.wrapped.(http.Hijacker).Hijack
	if w.interceptor.Hijack != nil {
		f = w.interceptor.Hijack(f)
	}
	return f()
}

func (w *ResponseWriterWrapper) ReadFrom(a0 io.Reader) (int64, error) {
	f := w.wrapped.(io.ReaderFrom).ReadFrom
	if w.interceptor.ReadFrom != nil {
		f = w.interceptor.ReadFrom(f)
	}
	return f(a0)
}

func (w *ResponseWriterWrapper) Push(a0 string, a1 *http.PushOptions) error {
	f := w.wrapped.(http.Pusher).Push
	if w.interceptor.Push != nil {
		f = w.interceptor.Push(f)
	}
	return f(a0, a1)
}

func (w *ResponseWriterWrapper) SetReadDeadline(a0 time.Time) error {
	f := w.wrapped.(interface {
		SetReadDeadline(time.Time) error
		SetWriteDeadline(time.Time) error
	}).SetReadDeadline
	if w.interceptor.SetReadDeadline != nil {
		f = w.interceptor.SetReadDeadline(f)
	}
	return f(a0)
}

func (w *ResponseWriterWrapper) SetWriteDeadline(a0 time.Time) error {
	f := w.wrapped.(interface {
		SetReadDeadline(time.Time) error
		SetWriteDeadline(time.Time) error
	}).SetWriteDeadline
	if w.interceptor.SetWriteDeadline != nil {
		f = w.interceptor.SetWriteDeadline(f)
	}
	return f(a0)
}

func (w *ResponseWriterWrapper) EnableFullDuplex() error {
	f := w.wrapped.(interface{ EnableFullDuplex() error }).EnableFullDuplex
	if w.interceptor.EnableFullDuplex != nil {
		f = w.interceptor.EnableFullDuplex(f)
	}
	return f()
}

func WrapResponseWriter(wrapped http.ResponseWriter, interceptor ResponseWriterInterceptor) http.ResponseWriter {
	w := &ResponseWriterWrapper{wrapped: wrapped, interceptor: interceptor}
	_, ok0 := wrapped.(http.Flusher)
	_, ok1 := wrapped.(http.CloseNotifier)
	_, ok2 := wrapped.(http.Hijacker)
	_, ok3 := wrapped.(io.ReaderFrom)
	_, ok4 := wrapped.(http.Pusher)
	_, ok5 := wrapped.(interface {
		SetReadDeadline(time.Time) error
		SetWriteDeadline(time.Time) error
	})
	_, ok6 := wrapped.(interface{ EnableFullDuplex() error })
	switch {
	// combination 1/128
	case !ok0 && !ok1 && !ok2 && !ok3 && !ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
		}{w, w}
	// combination 2/128
	case ok0 && !ok1 && !ok2 && !ok3 && !ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
		}{w, w, w}
	// combination 3/128
	case !ok0 && ok1 && !ok2 && !ok3 && !ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
		}{w, w, w}
	// combination 4/128
	case ok0 && ok1 && !ok2 && !ok3 && !ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
		}{w, w, w, w}
	// combination 5/128
	case !ok0 && !ok1 && ok2 && !ok3 && !ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Hijacker
		}{w, w, w}
	// combination 6/128
	case ok0 && !ok1 && ok2 && !ok3 && !ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
		}{w, w, w, w}
	// combination 7/128
	case !ok0 && ok1 && ok2 && !ok3 && !ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
		}{w, w, w, w}
	// combination 8/128
	case ok0 && ok1 && ok2 && !ok3 && !ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
		}{w, w, w, w, w}
	// combination 9/128
	case !ok0 && !ok1 && !ok2 && ok3 && !ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			io.ReaderFrom
		}{w, w, w}
	// combination 10/128
	case ok0 && !ok1 && !ok2 && ok3 && !ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
		}{w, w, w, w}
	// combination 11/128
	case !ok0 && ok1 && !ok2 && ok3 && !ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
		}{w, w, w, w}
	// combination 12/128
	case ok0 && ok1 && !ok2 && ok3 && !ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
		}{w, w, w, w, w}
	// combination 13/128
	case !ok0 && !ok1 && ok2 && ok3 && !ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
		}{w, w, w, w}
	// combination 14/128
	case ok0 && !ok1 && ok2 && ok3 && !ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
		}{w, w, w, w, w}
	// combination 15/128
	case !ok0 && ok1 && ok2 && ok3 && !ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
		}{w, w, w, w, w}
	// combination 16/128
	case ok0 && ok1 && ok2 && ok3 && !ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
		}{w, w, w, w, w, w}
	// combination 17/128
	case !ok0 && !ok1 && !ok2 && !ok3 && ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Pusher
		}{w, w, w}
	// combination 18/128
	case ok0 && !ok1 && !ok2 && !ok3 && ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.Pusher
		}{w, w, w, w}
	// combination 19/128
	case !ok0 && ok1 && !ok2 && !ok3 && ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Pusher
		}{w, w, w, w}
	// combination 20/128
	case ok0 && ok1 && !ok2 && !ok3 && ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Pusher
		}{w, w, w, w, w}
	// combination 21/128
	case !ok0 && !ok1 && ok2 && !ok3 && ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Hijacker
			http.Pusher
		}{w, w, w, w}
	// combination 22/128
	case ok0 && !ok1 && ok2 && !ok3 && ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			http.Pusher
		}{w, w, w, w, w}
	// combination 23/128
	case !ok0 && ok1 && ok2 && !ok3 && ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			http.Pusher
		}{w, w, w, w, w}
	// combination 24/128
	case ok0 && ok1 && ok2 && !ok3 && ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			http.Pusher
		}{w, w, w, w, w, w}
	// combination 25/128
	case !ok0 && !ok1 && !ok2 && ok3 && ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			io.ReaderFrom
			http.Pusher
		}{w, w, w, w}
	// combination 26/128
	case ok0 && !ok1 && !ok2 && ok3 && ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
			http.Pusher
		}{w, w, w, w, w}
	// combination 27/128
	case !ok0 && ok1 && !ok2 && ok3 && ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
			http.Pusher
		}{w, w, w, w, w}
	// combination 28/128
	case ok0 && ok1 && !ok2 && ok3 && ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
			http.Pusher
		}{w, w, w, w, w, w}
	// combination 29/128
	case !ok0 && !ok1 && ok2 && ok3 && ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
			http.Pusher
		}{w, w, w, w, w}
	// combination 30/128
	case ok0 && !ok1 && ok2 && ok3 && ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
			http.Pusher
		}{w, w, w, w, w, w}
	// combination 31/128
	case !ok0 && ok1 && ok2 && ok3 && ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			http.Pusher
		}{w, w, w, w, w, w}
	// combination 32/128
	case ok0 && ok1 && ok2 && ok3 && ok4 && !ok5 && !ok6:
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			http.Pusher
		}{w, w, w, w, w, w, w}
	// combination 33/128
	case !ok0 && !ok1 && !ok2 && !ok3 && !ok4 && ok5 && !ok6:
		type combination_32_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			combination_32_5
		}{w, w, w}
	// combination 34/128
	case ok0 && !ok1 && !ok2 && !ok3 && !ok4 && ok5 && !ok6:
		type combination_33_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			combination_33_5
		}{w, w, w, w}
	// combination 35/128
	case !ok0 && ok1 && !ok2 && !ok3 && !ok4 && ok5 && !ok6:
		type combination_34_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
			combination_34_5
		}{w, w, w, w}
	// combination 36/128
	case ok0 && ok1 && !ok2 && !ok3 && !ok4 && ok5 && !ok6:
		type combination_35_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			combination_35_5
		}{w, w, w, w, w}
	// combination 37/128
	case !ok0 && !ok1 && ok2 && !ok3 && !ok4 && ok5 && !ok6:
		type combination_36_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Hijacker
			combination_36_5
		}{w, w, w, w}
	// combination 38/128
	case ok0 && !ok1 && ok2 && !ok3 && !ok4 && ok5 && !ok6:
		type combination_37_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			combination_37_5
		}{w, w, w, w, w}
	// combination 39/128
	case !ok0 && ok1 && ok2 && !ok3 && !ok4 && ok5 && !ok6:
		type combination_38_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			combination_38_5
		}{w, w, w, w, w}
	// combination 40/128
	case ok0 && ok1 && ok2 && !ok3 && !ok4 && ok5 && !ok6:
		type combination_39_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			combination_39_5
		}{w, w, w, w, w, w}
	// combination 41/128
	case !ok0 && !ok1 && !ok2 && ok3 && !ok4 && ok5 && !ok6:
		type combination_40_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			io.ReaderFrom
			combination_40_5
		}{w, w, w, w}
	// combination 42/128
	case ok0 && !ok1 && !ok2 && ok3 && !ok4 && ok5 && !ok6:
		type combination_41_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
			combination_41_5
		}{w, w, w, w, w}
	// combination 43/128
	case !ok0 && ok1 && !ok2 && ok3 && !ok4 && ok5 && !ok6:
		type combination_42_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
			combination_42_5
		}{w, w, w, w, w}
	// combination 44/128
	case ok0 && ok1 && !ok2 && ok3 && !ok4 && ok5 && !ok6:
		type combination_43_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
			combination_43_5
		}{w, w, w, w, w, w}
	// combination 45/128
	case !ok0 && !ok1 && ok2 && ok3 && !ok4 && ok5 && !ok6:
		type combination_44_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
			combination_44_5
		}{w, w, w, w, w}
	// combination 46/128
	case ok0 && !ok1 && ok2 && ok3 && !ok4 && ok5 && !ok6:
		type combination_45_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
			combination_45_5
		}{w, w, w, w, w, w}
	// combination 47/128
	case !ok0 && ok1 && ok2 && ok3 && !ok4 && ok5 && !ok6:
		type combination_46_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			combination_46_5
		}{w, w, w, w, w, w}
	// combination 48/128
	case ok0 && ok1 && ok2 && ok3 && !ok4 && ok5 && !ok6:
		type combination_47_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			combination_47_5
		}{w, w, w, w, w, w, w}
	// combination 49/128
	case !ok0 && !ok1 && !ok2 && !ok3 && ok4 && ok5 && !ok6:
		type combination_48_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Pusher
			combination_48_5
		}{w, w, w, w}
	// combination 50/128
	case ok0 && !ok1 && !ok2 && !ok3 && ok4 && ok5 && !ok6:
		type combination_49_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.Pusher
			combination_49_5
		}{w, w, w, w, w}
	// combination 51/128
	case !ok0 && ok1 && !ok2 && !ok3 && ok4 && ok5 && !ok6:
		type combination_50_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Pusher
			combination_50_5
		}{w, w, w, w, w}
	// combination 52/128
	case ok0 && ok1 && !ok2 && !ok3 && ok4 && ok5 && !ok6:
		type combination_51_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Pusher
			combination_51_5
		}{w, w, w, w, w, w}
	// combination 53/128
	case !ok0 && !ok1 && ok2 && !ok3 && ok4 && ok5 && !ok6:
		type combination_52_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Hijacker
			http.Pusher
			combination_52_5
		}{w, w, w, w, w}
	// combination 54/128
	case ok0 && !ok1 && ok2 && !ok3 && ok4 && ok5 && !ok6:
		type combination_53_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			http.Pusher
			combination_53_5
		}{w, w, w, w, w, w}
	// combination 55/128
	case !ok0 && ok1 && ok2 && !ok3 && ok4 && ok5 && !ok6:
		type combination_54_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			http.Pusher
			combination_54_5
		}{w, w, w, w, w, w}
	// combination 56/128
	case ok0 && ok1 && ok2 && !ok3 && ok4 && ok5 && !ok6:
		type combination_55_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			http.Pusher
			combination_55_5
		}{w, w, w, w, w, w, w}
	// combination 57/128
	case !ok0 && !ok1 && !ok2 && ok3 && ok4 && ok5 && !ok6:
		type combination_56_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			io.ReaderFrom
			http.Pusher
			combination_56_5
		}{w, w, w, w, w}
	// combination 58/128
	case ok0 && !ok1 && !ok2 && ok3 && ok4 && ok5 && !ok6:
		type combination_57_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
			http.Pusher
			combination_57_5
		}{w, w, w, w, w, w}
	// combination 59/128
	case !ok0 && ok1 && !ok2 && ok3 && ok4 && ok5 && !ok6:
		type combination_58_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
			http.Pusher
			combination_58_5
		}{w, w, w, w, w, w}
	// combination 60/128
	case ok0 && ok1 && !ok2 && ok3 && ok4 && ok5 && !ok6:
		type combination_59_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
			http.Pusher
			combination_59_5
		}{w, w, w, w, w, w, w}
	// combination 61/128
	case !ok0 && !ok1 && ok2 && ok3 && ok4 && ok5 && !ok6:
		type combination_60_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
			http.Pusher
			combination_60_5
		}{w, w, w, w, w, w}
	// combination 62/128
	case ok0 && !ok1 && ok2 && ok3 && ok4 && ok5 && !ok6:
		type combination_61_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
			http.Pusher
			combination_61_5
		}{w, w, w, w, w, w, w}
	// combination 63/128
	case !ok0 && ok1 && ok2 && ok3 && ok4 && ok5 && !ok6:
		type combination_62_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			http.Pusher
			combination_62_5
		}{w, w, w, w, w, w, w}
	// combination 64/128
	case ok0 && ok1 && ok2 && ok3 && ok4 && ok5 && !ok6:
		type combination_63_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			http.Pusher
			combination_63_5
		}{w, w, w, w, w, w, w, w}
	// combination 65/128
	case !ok0 && !ok1 && !ok2 && !ok3 && !ok4 && !ok5 && ok6:
		type combination_64_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			combination_64_6
		}{w, w, w}
	// combination 66/128
	case ok0 && !ok1 && !ok2 && !ok3 && !ok4 && !ok5 && ok6:
		type combination_65_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			combination_65_6
		}{w, w, w, w}
	// combination 67/128
	case !ok0 && ok1 && !ok2 && !ok3 && !ok4 && !ok5 && ok6:
		type combination_66_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
			combination_66_6
		}{w, w, w, w}
	// combination 68/128
	case ok0 && ok1 && !ok2 && !ok3 && !ok4 && !ok5 && ok6:
		type combination_67_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			combination_67_6
		}{w, w, w, w, w}
	// combination 69/128
	case !ok0 && !ok1 && ok2 && !ok3 && !ok4 && !ok5 && ok6:
		type combination_68_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Hijacker
			combination_68_6
		}{w, w, w, w}
	// combination 70/128
	case ok0 && !ok1 && ok2 && !ok3 && !ok4 && !ok5 && ok6:
		type combination_69_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			combination_69_6
		}{w, w, w, w, w}
	// combination 71/128
	case !ok0 && ok1 && ok2 && !ok3 && !ok4 && !ok5 && ok6:
		type combination_70_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			combination_70_6
		}{w, w, w, w, w}
	// combination 72/128
	case ok0 && ok1 && ok2 && !ok3 && !ok4 && !ok5 && ok6:
		type combination_71_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			combination_71_6
		}{w, w, w, w, w, w}
	// combination 73/128
	case !ok0 && !ok1 && !ok2 && ok3 && !ok4 && !ok5 && ok6:
		type combination_72_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			io.ReaderFrom
			combination_72_6
		}{w, w, w, w}
	// combination 74/128
	case ok0 && !ok1 && !ok2 && ok3 && !ok4 && !ok5 && ok6:
		type combination_73_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
			combination_73_6
		}{w, w, w, w, w}
	// combination 75/128
	case !ok0 && ok1 && !ok2 && ok3 && !ok4 && !ok5 && ok6:
		type combination_74_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
			combination_74_6
		}{w, w, w, w, w}
	// combination 76/128
	case ok0 && ok1 && !ok2 && ok3 && !ok4 && !ok5 && ok6:
		type combination_75_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
			combination_75_6
		}{w, w, w, w, w, w}
	// combination 77/128
	case !ok0 && !ok1 && ok2 && ok3 && !ok4 && !ok5 && ok6:
		type combination_76_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
			combination_76_6
		}{w, w, w, w, w}
	// combination 78/128
	case ok0 && !ok1 && ok2 && ok3 && !ok4 && !ok5 && ok6:
		type combination_77_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
			combination_77_6
		}{w, w, w, w, w, w}
	// combination 79/128
	case !ok0 && ok1 && ok2 && ok3 && !ok4 && !ok5 && ok6:
		type combination_78_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			combination_78_6
		}{w, w, w, w, w, w}
	// combination 80/128
	case ok0 && ok1 && ok2 && ok3 && !ok4 && !ok5 && ok6:
		type combination_79_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			combination_79_6
		}{w, w, w, w, w, w, w}
	// combination 81/128
	case !ok0 && !ok1 && !ok2 && !ok3 && ok4 && !ok5 && ok6:
		type combination_80_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Pusher
			combination_80_6
		}{w, w, w, w}
	// combination 82/128
	case ok0 && !ok1 && !ok2 && !ok3 && ok4 && !ok5 && ok6:
		type combination_81_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.Pusher
			combination_81_6
		}{w, w, w, w, w}
	// combination 83/128
	case !ok0 && ok1 && !ok2 && !ok3 && ok4 && !ok5 && ok6:
		type combination_82_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Pusher
			combination_82_6
		}{w, w, w, w, w}
	// combination 84/128
	case ok0 && ok1 && !ok2 && !ok3 && ok4 && !ok5 && ok6:
		type combination_83_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Pusher
			combination_83_6
		}{w, w, w, w, w, w}
	// combination 85/128
	case !ok0 && !ok1 && ok2 && !ok3 && ok4 && !ok5 && ok6:
		type combination_84_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Hijacker
			http.Pusher
			combination_84_6
		}{w, w, w, w, w}
	// combination 86/128
	case ok0 && !ok1 && ok2 && !ok3 && ok4 && !ok5 && ok6:
		type combination_85_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			http.Pusher
			combination_85_6
		}{w, w, w, w, w, w}
	// combination 87/128
	case !ok0 && ok1 && ok2 && !ok3 && ok4 && !ok5 && ok6:
		type combination_86_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			http.Pusher
			combination_86_6
		}{w, w, w, w, w, w}
	// combination 88/128
	case ok0 && ok1 && ok2 && !ok3 && ok4 && !ok5 && ok6:
		type combination_87_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			http.Pusher
			combination_87_6
		}{w, w, w, w, w, w, w}
	// combination 89/128
	case !ok0 && !ok1 && !ok2 && ok3 && ok4 && !ok5 && ok6:
		type combination_88_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			io.ReaderFrom
			http.Pusher
			combination_88_6
		}{w, w, w, w, w}
	// combination 90/128
	case ok0 && !ok1 && !ok2 && ok3 && ok4 && !ok5 && ok6:
		type combination_89_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
			http.Pusher
			combination_89_6
		}{w, w, w, w, w, w}
	// combination 91/128
	case !ok0 && ok1 && !ok2 && ok3 && ok4 && !ok5 && ok6:
		type combination_90_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
			http.Pusher
			combination_90_6
		}{w, w, w, w, w, w}
	// combination 92/128
	case ok0 && ok1 && !ok2 && ok3 && ok4 && !ok5 && ok6:
		type combination_91_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
			http.Pusher
			combination_91_6
		}{w, w, w, w, w, w, w}
	// combination 93/128
	case !ok0 && !ok1 && ok2 && ok3 && ok4 && !ok5 && ok6:
		type combination_92_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
			http.Pusher
			combination_92_6
		}{w, w, w, w, w, w}
	// combination 94/128
	case ok0 && !ok1 && ok2 && ok3 && ok4 && !ok5 && ok6:
		type combination_93_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
			http.Pusher
			combination_93_6
		}{w, w, w, w, w, w, w}
	// combination 95/128
	case !ok0 && ok1 && ok2 && ok3 && ok4 && !ok5 && ok6:
		type combination_94_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			http.Pusher
			combination_94_6
		}{w, w, w, w, w, w, w}
	// combination 96/128
	case ok0 && ok1 && ok2 && ok3 && ok4 && !ok5 && ok6:
		type combination_95_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			http.Pusher
			combination_95_6
		}{w, w, w, w, w, w, w, w}
	// combination 97/128
	case !ok0 && !ok1 && !ok2 && !ok3 && !ok4 && ok5 && ok6:
		type combination_96_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_96_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			combination_96_5
			combination_96_6
		}{w, w, w, w}
	// combination 98/128
	case ok0 && !ok1 && !ok2 && !ok3 && !ok4 && ok5 && ok6:
		type combination_97_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_97_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			combination_97_5
			combination_97_6
		}{w, w, w, w, w}
	// combination 99/128
	case !ok0 && ok1 && !ok2 && !ok3 && !ok4 && ok5 && ok6:
		type combination_98_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_98_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
			combination_98_5
			combination_98_6
		}{w, w, w, w, w}
	// combination 100/128
	case ok0 && ok1 && !ok2 && !ok3 && !ok4 && ok5 && ok6:
		type combination_99_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_99_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			combination_99_5
			combination_99_6
		}{w, w, w, w, w, w}
	// combination 101/128
	case !ok0 && !ok1 && ok2 && !ok3 && !ok4 && ok5 && ok6:
		type combination_100_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_100_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Hijacker
			combination_100_5
			combination_100_6
		}{w, w, w, w, w}
	// combination 102/128
	case ok0 && !ok1 && ok2 && !ok3 && !ok4 && ok5 && ok6:
		type combination_101_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_101_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			combination_101_5
			combination_101_6
		}{w, w, w, w, w, w}
	// combination 103/128
	case !ok0 && ok1 && ok2 && !ok3 && !ok4 && ok5 && ok6:
		type combination_102_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_102_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			combination_102_5
			combination_102_6
		}{w, w, w, w, w, w}
	// combination 104/128
	case ok0 && ok1 && ok2 && !ok3 && !ok4 && ok5 && ok6:
		type combination_103_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_103_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			combination_103_5
			combination_103_6
		}{w, w, w, w, w, w, w}
	// combination 105/128
	case !ok0 && !ok1 && !ok2 && ok3 && !ok4 && ok5 && ok6:
		type combination_104_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_104_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			io.ReaderFrom
			combination_104_5
			combination_104_6
		}{w, w, w, w, w}
	// combination 106/128
	case ok0 && !ok1 && !ok2 && ok3 && !ok4 && ok5 && ok6:
		type combination_105_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_105_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
			combination_105_5
			combination_105_6
		}{w, w, w, w, w, w}
	// combination 107/128
	case !ok0 && ok1 && !ok2 && ok3 && !ok4 && ok5 && ok6:
		type combination_106_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_106_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
			combination_106_5
			combination_106_6
		}{w, w, w, w, w, w}
	// combination 108/128
	case ok0 && ok1 && !ok2 && ok3 && !ok4 && ok5 && ok6:
		type combination_107_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_107_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
			combination_107_5
			combination_107_6
		}{w, w, w, w, w, w, w}
	// combination 109/128
	case !ok0 && !ok1 && ok2 && ok3 && !ok4 && ok5 && ok6:
		type combination_108_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_108_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
			combination_108_5
			combination_108_6
		}{w, w, w, w, w, w}
	// combination 110/128
	case ok0 && !ok1 && ok2 && ok3 && !ok4 && ok5 && ok6:
		type combination_109_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_109_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
			combination_109_5
			combination_109_6
		}{w, w, w, w, w, w, w}
	// combination 111/128
	case !ok0 && ok1 && ok2 && ok3 && !ok4 && ok5 && ok6:
		type combination_110_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_110_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			combination_110_5
			combination_110_6
		}{w, w, w, w, w, w, w}
	// combination 112/128
	case ok0 && ok1 && ok2 && ok3 && !ok4 && ok5 && ok6:
		type combination_111_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_111_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			combination_111_5
			combination_111_6
		}{w, w, w, w, w, w, w, w}
	// combination 113/128
	case !ok0 && !ok1 && !ok2 && !ok3 && ok4 && ok5 && ok6:
		type combination_112_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_112_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Pusher
			combination_112_5
			combination_112_6
		}{w, w, w, w, w}
	// combination 114/128
	case ok0 && !ok1 && !ok2 && !ok3 && ok4 && ok5 && ok6:
		type combination_113_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_113_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.Pusher
			combination_113_5
			combination_113_6
		}{w, w, w, w, w, w}
	// combination 115/128
	case !ok0 && ok1 && !ok2 && !ok3 && ok4 && ok5 && ok6:
		type combination_114_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_114_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Pusher
			combination_114_5
			combination_114_6
		}{w, w, w, w, w, w}
	// combination 116/128
	case ok0 && ok1 && !ok2 && !ok3 && ok4 && ok5 && ok6:
		type combination_115_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_115_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Pusher
			combination_115_5
			combination_115_6
		}{w, w, w, w, w, w, w}
	// combination 117/128
	case !ok0 && !ok1 && ok2 && !ok3 && ok4 && ok5 && ok6:
		type combination_116_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_116_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Hijacker
			http.Pusher
			combination_116_5
			combination_116_6
		}{w, w, w, w, w, w}
	// combination 118/128
	case ok0 && !ok1 && ok2 && !ok3 && ok4 && ok5 && ok6:
		type combination_117_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_117_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			http.Pusher
			combination_117_5
			combination_117_6
		}{w, w, w, w, w, w, w}
	// combination 119/128
	case !ok0 && ok1 && ok2 && !ok3 && ok4 && ok5 && ok6:
		type combination_118_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_118_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			http.Pusher
			combination_118_5
			combination_118_6
		}{w, w, w, w, w, w, w}
	// combination 120/128
	case ok0 && ok1 && ok2 && !ok3 && ok4 && ok5 && ok6:
		type combination_119_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_119_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			http.Pusher
			combination_119_5
			combination_119_6
		}{w, w, w, w, w, w, w, w}
	// combination 121/128
	case !ok0 && !ok1 && !ok2 && ok3 && ok4 && ok5 && ok6:
		type combination_120_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_120_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			io.ReaderFrom
			http.Pusher
			combination_120_5
			combination_120_6
		}{w, w, w, w, w, w}
	// combination 122/128
	case ok0 && !ok1 && !ok2 && ok3 && ok4 && ok5 && ok6:
		type combination_121_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_121_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
			http.Pusher
			combination_121_5
			combination_121_6
		}{w, w, w, w, w, w, w}
	// combination 123/128
	case !ok0 && ok1 && !ok2 && ok3 && ok4 && ok5 && ok6:
		type combination_122_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_122_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
			http.Pusher
			combination_122_5
			combination_122_6
		}{w, w, w, w, w, w, w}
	// combination 124/128
	case ok0 && ok1 && !ok2 && ok3 && ok4 && ok5 && ok6:
		type combination_123_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_123_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
			http.Pusher
			combination_123_5
			combination_123_6
		}{w, w, w, w, w, w, w, w}
	// combination 125/128
	case !ok0 && !ok1 && ok2 && ok3 && ok4 && ok5 && ok6:
		type combination_124_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_124_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
			http.Pusher
			combination_124_5
			combination_124_6
		}{w, w, w, w, w, w, w}
	// combination 126/128
	case ok0 && !ok1 && ok2 && ok3 && ok4 && ok5 && ok6:
		type combination_125_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_125_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
			http.Pusher
			combination_125_5
			combination_125_6
		}{w, w, w, w, w, w, w, w}
	// combination 127/128
	case !ok0 && ok1 && ok2 && ok3 && ok4 && ok5 && ok6:
		type combination_126_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_126_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			http.Pusher
			combination_126_5
			combination_126_6
		}{w, w, w, w, w, w, w, w}
	// combination 128/128
	case ok0 && ok1 && ok2 && ok3 && ok4 && ok5 && ok6:
		type combination_127_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_127_6 interface{ EnableFullDuplex() error }
		return struct {
			ResponseWriterUnwrapper
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			http.Pusher
			combination_127_5
			combination_127_6
		}{w, w, w, w, w, w, w, w, w}
	}
	panic("unreachable")
}
