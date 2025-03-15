package nethttp

import (
	io "io"
	http "net/http"
	testing "testing"
	time "time"
)

func TestWrapResponseWriter(t *testing.T) {
	{
		t.Log("combination 1/128: http.ResponseWriter")
		wrapped := struct {
			http.ResponseWriter
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 1/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 1/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 1/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 1/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 1/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 1/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 1/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 1/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 1/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 2/128: http.ResponseWriter http.Flusher")
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 2/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 2/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 2/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 2/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 2/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 2/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 2/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 2/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 2/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 3/128: http.ResponseWriter http.CloseNotifier")
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 3/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 3/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 3/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 3/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 3/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 3/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 3/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 3/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 3/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 4/128: http.ResponseWriter http.Flusher http.CloseNotifier")
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 4/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 4/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 4/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 4/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 4/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 4/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 4/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 4/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 4/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 5/128: http.ResponseWriter http.Hijacker")
		wrapped := struct {
			http.ResponseWriter
			http.Hijacker
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 5/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 5/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 5/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 5/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 5/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 5/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 5/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 5/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 5/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 6/128: http.ResponseWriter http.Flusher http.Hijacker")
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.Hijacker
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 6/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 6/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 6/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 6/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 6/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 6/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 6/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 6/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 6/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 7/128: http.ResponseWriter http.CloseNotifier http.Hijacker")
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 7/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 7/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 7/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 7/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 7/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 7/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 7/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 7/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 7/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 8/128: http.ResponseWriter http.Flusher http.CloseNotifier http.Hijacker")
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 8/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 8/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 8/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 8/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 8/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 8/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 8/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 8/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 8/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 9/128: http.ResponseWriter io.ReaderFrom")
		wrapped := struct {
			http.ResponseWriter
			io.ReaderFrom
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 9/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 9/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 9/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 9/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 9/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 9/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 9/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 9/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 9/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 10/128: http.ResponseWriter http.Flusher io.ReaderFrom")
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 10/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 10/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 10/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 10/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 10/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 10/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 10/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 10/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 10/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 11/128: http.ResponseWriter http.CloseNotifier io.ReaderFrom")
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 11/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 11/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 11/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 11/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 11/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 11/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 11/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 11/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 11/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 12/128: http.ResponseWriter http.Flusher http.CloseNotifier io.ReaderFrom")
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 12/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 12/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 12/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 12/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 12/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 12/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 12/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 12/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 12/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 13/128: http.ResponseWriter http.Hijacker io.ReaderFrom")
		wrapped := struct {
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 13/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 13/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 13/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 13/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 13/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 13/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 13/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 13/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 13/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 14/128: http.ResponseWriter http.Flusher http.Hijacker io.ReaderFrom")
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 14/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 14/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 14/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 14/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 14/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 14/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 14/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 14/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 14/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 15/128: http.ResponseWriter http.CloseNotifier http.Hijacker io.ReaderFrom")
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 15/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 15/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 15/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 15/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 15/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 15/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 15/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 15/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 15/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 16/128: http.ResponseWriter http.Flusher http.CloseNotifier http.Hijacker io.ReaderFrom")
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 16/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 16/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 16/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 16/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 16/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 16/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 16/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 16/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 16/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 17/128: http.ResponseWriter http.Pusher")
		wrapped := struct {
			http.ResponseWriter
			http.Pusher
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 17/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 17/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 17/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 17/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 17/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 17/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 17/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 17/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 17/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 18/128: http.ResponseWriter http.Flusher http.Pusher")
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.Pusher
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 18/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 18/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 18/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 18/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 18/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 18/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 18/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 18/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 18/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 19/128: http.ResponseWriter http.CloseNotifier http.Pusher")
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Pusher
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 19/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 19/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 19/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 19/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 19/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 19/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 19/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 19/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 19/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 20/128: http.ResponseWriter http.Flusher http.CloseNotifier http.Pusher")
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Pusher
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 20/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 20/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 20/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 20/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 20/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 20/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 20/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 20/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 20/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 21/128: http.ResponseWriter http.Hijacker http.Pusher")
		wrapped := struct {
			http.ResponseWriter
			http.Hijacker
			http.Pusher
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 21/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 21/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 21/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 21/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 21/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 21/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 21/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 21/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 21/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 22/128: http.ResponseWriter http.Flusher http.Hijacker http.Pusher")
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			http.Pusher
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 22/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 22/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 22/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 22/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 22/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 22/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 22/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 22/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 22/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 23/128: http.ResponseWriter http.CloseNotifier http.Hijacker http.Pusher")
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			http.Pusher
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 23/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 23/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 23/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 23/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 23/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 23/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 23/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 23/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 23/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 24/128: http.ResponseWriter http.Flusher http.CloseNotifier http.Hijacker http.Pusher")
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			http.Pusher
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 24/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 24/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 24/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 24/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 24/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 24/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 24/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 24/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 24/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 25/128: http.ResponseWriter io.ReaderFrom http.Pusher")
		wrapped := struct {
			http.ResponseWriter
			io.ReaderFrom
			http.Pusher
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 25/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 25/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 25/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 25/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 25/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 25/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 25/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 25/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 25/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 26/128: http.ResponseWriter http.Flusher io.ReaderFrom http.Pusher")
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
			http.Pusher
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 26/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 26/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 26/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 26/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 26/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 26/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 26/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 26/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 26/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 27/128: http.ResponseWriter http.CloseNotifier io.ReaderFrom http.Pusher")
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
			http.Pusher
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 27/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 27/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 27/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 27/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 27/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 27/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 27/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 27/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 27/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 28/128: http.ResponseWriter http.Flusher http.CloseNotifier io.ReaderFrom http.Pusher")
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
			http.Pusher
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 28/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 28/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 28/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 28/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 28/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 28/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 28/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 28/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 28/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 29/128: http.ResponseWriter http.Hijacker io.ReaderFrom http.Pusher")
		wrapped := struct {
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
			http.Pusher
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 29/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 29/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 29/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 29/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 29/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 29/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 29/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 29/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 29/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 30/128: http.ResponseWriter http.Flusher http.Hijacker io.ReaderFrom http.Pusher")
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
			http.Pusher
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 30/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 30/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 30/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 30/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 30/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 30/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 30/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 30/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 30/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 31/128: http.ResponseWriter http.CloseNotifier http.Hijacker io.ReaderFrom http.Pusher")
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			http.Pusher
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 31/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 31/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 31/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 31/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 31/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 31/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 31/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 31/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 31/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 32/128: http.ResponseWriter http.Flusher http.CloseNotifier http.Hijacker io.ReaderFrom http.Pusher")
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			http.Pusher
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 32/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 32/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 32/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 32/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 32/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 32/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 32/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 32/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 32/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 33/128: http.ResponseWriter interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_32_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			combination_32_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 33/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 33/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 33/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 33/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 33/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 33/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 33/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 33/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 33/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 34/128: http.ResponseWriter http.Flusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_33_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			combination_33_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 34/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 34/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 34/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 34/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 34/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 34/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 34/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 34/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 34/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 35/128: http.ResponseWriter http.CloseNotifier interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_34_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
			combination_34_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 35/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 35/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 35/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 35/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 35/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 35/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 35/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 35/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 35/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 36/128: http.ResponseWriter http.Flusher http.CloseNotifier interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_35_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			combination_35_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 36/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 36/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 36/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 36/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 36/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 36/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 36/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 36/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 36/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 37/128: http.ResponseWriter http.Hijacker interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_36_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			http.Hijacker
			combination_36_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 37/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 37/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 37/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 37/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 37/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 37/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 37/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 37/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 37/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 38/128: http.ResponseWriter http.Flusher http.Hijacker interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_37_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			combination_37_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 38/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 38/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 38/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 38/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 38/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 38/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 38/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 38/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 38/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 39/128: http.ResponseWriter http.CloseNotifier http.Hijacker interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_38_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			combination_38_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 39/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 39/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 39/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 39/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 39/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 39/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 39/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 39/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 39/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 40/128: http.ResponseWriter http.Flusher http.CloseNotifier http.Hijacker interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_39_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			combination_39_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 40/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 40/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 40/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 40/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 40/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 40/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 40/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 40/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 40/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 41/128: http.ResponseWriter io.ReaderFrom interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_40_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			io.ReaderFrom
			combination_40_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 41/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 41/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 41/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 41/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 41/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 41/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 41/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 41/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 41/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 42/128: http.ResponseWriter http.Flusher io.ReaderFrom interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_41_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
			combination_41_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 42/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 42/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 42/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 42/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 42/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 42/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 42/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 42/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 42/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 43/128: http.ResponseWriter http.CloseNotifier io.ReaderFrom interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_42_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
			combination_42_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 43/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 43/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 43/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 43/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 43/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 43/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 43/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 43/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 43/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 44/128: http.ResponseWriter http.Flusher http.CloseNotifier io.ReaderFrom interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_43_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
			combination_43_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 44/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 44/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 44/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 44/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 44/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 44/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 44/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 44/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 44/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 45/128: http.ResponseWriter http.Hijacker io.ReaderFrom interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_44_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
			combination_44_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 45/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 45/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 45/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 45/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 45/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 45/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 45/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 45/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 45/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 46/128: http.ResponseWriter http.Flusher http.Hijacker io.ReaderFrom interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_45_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
			combination_45_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 46/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 46/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 46/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 46/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 46/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 46/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 46/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 46/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 46/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 47/128: http.ResponseWriter http.CloseNotifier http.Hijacker io.ReaderFrom interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_46_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			combination_46_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 47/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 47/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 47/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 47/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 47/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 47/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 47/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 47/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 47/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 48/128: http.ResponseWriter http.Flusher http.CloseNotifier http.Hijacker io.ReaderFrom interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_47_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			combination_47_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 48/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 48/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 48/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 48/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 48/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 48/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 48/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 48/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 48/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 49/128: http.ResponseWriter http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_48_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			http.Pusher
			combination_48_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 49/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 49/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 49/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 49/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 49/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 49/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 49/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 49/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 49/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 50/128: http.ResponseWriter http.Flusher http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_49_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.Pusher
			combination_49_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 50/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 50/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 50/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 50/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 50/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 50/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 50/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 50/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 50/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 51/128: http.ResponseWriter http.CloseNotifier http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_50_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Pusher
			combination_50_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 51/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 51/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 51/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 51/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 51/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 51/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 51/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 51/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 51/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 52/128: http.ResponseWriter http.Flusher http.CloseNotifier http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_51_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Pusher
			combination_51_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 52/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 52/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 52/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 52/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 52/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 52/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 52/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 52/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 52/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 53/128: http.ResponseWriter http.Hijacker http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_52_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			http.Hijacker
			http.Pusher
			combination_52_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 53/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 53/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 53/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 53/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 53/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 53/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 53/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 53/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 53/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 54/128: http.ResponseWriter http.Flusher http.Hijacker http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_53_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			http.Pusher
			combination_53_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 54/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 54/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 54/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 54/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 54/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 54/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 54/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 54/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 54/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 55/128: http.ResponseWriter http.CloseNotifier http.Hijacker http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_54_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			http.Pusher
			combination_54_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 55/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 55/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 55/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 55/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 55/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 55/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 55/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 55/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 55/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 56/128: http.ResponseWriter http.Flusher http.CloseNotifier http.Hijacker http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_55_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			http.Pusher
			combination_55_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 56/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 56/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 56/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 56/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 56/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 56/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 56/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 56/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 56/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 57/128: http.ResponseWriter io.ReaderFrom http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_56_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			io.ReaderFrom
			http.Pusher
			combination_56_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 57/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 57/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 57/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 57/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 57/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 57/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 57/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 57/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 57/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 58/128: http.ResponseWriter http.Flusher io.ReaderFrom http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_57_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
			http.Pusher
			combination_57_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 58/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 58/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 58/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 58/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 58/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 58/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 58/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 58/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 58/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 59/128: http.ResponseWriter http.CloseNotifier io.ReaderFrom http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_58_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
			http.Pusher
			combination_58_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 59/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 59/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 59/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 59/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 59/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 59/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 59/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 59/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 59/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 60/128: http.ResponseWriter http.Flusher http.CloseNotifier io.ReaderFrom http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_59_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
			http.Pusher
			combination_59_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 60/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 60/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 60/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 60/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 60/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 60/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 60/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 60/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 60/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 61/128: http.ResponseWriter http.Hijacker io.ReaderFrom http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_60_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
			http.Pusher
			combination_60_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 61/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 61/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 61/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 61/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 61/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 61/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 61/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 61/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 61/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 62/128: http.ResponseWriter http.Flusher http.Hijacker io.ReaderFrom http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_61_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
			http.Pusher
			combination_61_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 62/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 62/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 62/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 62/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 62/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 62/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 62/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 62/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 62/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 63/128: http.ResponseWriter http.CloseNotifier http.Hijacker io.ReaderFrom http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_62_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			http.Pusher
			combination_62_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 63/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 63/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 63/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 63/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 63/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 63/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 63/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 63/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 63/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 64/128: http.ResponseWriter http.Flusher http.CloseNotifier http.Hijacker io.ReaderFrom http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		type combination_63_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			http.Pusher
			combination_63_5
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 64/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 64/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 64/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 64/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 64/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 64/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != false {
			t.Errorf("combination 64/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 64/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 64/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 65/128: http.ResponseWriter interface{ EnableFullDuplex() error }")
		type combination_64_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			combination_64_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 65/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 65/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 65/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 65/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 65/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 65/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 65/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 65/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 65/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 66/128: http.ResponseWriter http.Flusher interface{ EnableFullDuplex() error }")
		type combination_65_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			combination_65_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 66/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 66/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 66/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 66/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 66/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 66/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 66/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 66/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 66/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 67/128: http.ResponseWriter http.CloseNotifier interface{ EnableFullDuplex() error }")
		type combination_66_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
			combination_66_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 67/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 67/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 67/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 67/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 67/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 67/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 67/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 67/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 67/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 68/128: http.ResponseWriter http.Flusher http.CloseNotifier interface{ EnableFullDuplex() error }")
		type combination_67_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			combination_67_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 68/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 68/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 68/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 68/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 68/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 68/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 68/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 68/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 68/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 69/128: http.ResponseWriter http.Hijacker interface{ EnableFullDuplex() error }")
		type combination_68_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Hijacker
			combination_68_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 69/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 69/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 69/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 69/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 69/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 69/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 69/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 69/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 69/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 70/128: http.ResponseWriter http.Flusher http.Hijacker interface{ EnableFullDuplex() error }")
		type combination_69_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			combination_69_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 70/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 70/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 70/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 70/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 70/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 70/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 70/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 70/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 70/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 71/128: http.ResponseWriter http.CloseNotifier http.Hijacker interface{ EnableFullDuplex() error }")
		type combination_70_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			combination_70_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 71/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 71/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 71/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 71/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 71/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 71/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 71/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 71/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 71/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 72/128: http.ResponseWriter http.Flusher http.CloseNotifier http.Hijacker interface{ EnableFullDuplex() error }")
		type combination_71_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			combination_71_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 72/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 72/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 72/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 72/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 72/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 72/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 72/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 72/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 72/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 73/128: http.ResponseWriter io.ReaderFrom interface{ EnableFullDuplex() error }")
		type combination_72_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			io.ReaderFrom
			combination_72_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 73/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 73/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 73/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 73/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 73/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 73/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 73/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 73/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 73/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 74/128: http.ResponseWriter http.Flusher io.ReaderFrom interface{ EnableFullDuplex() error }")
		type combination_73_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
			combination_73_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 74/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 74/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 74/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 74/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 74/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 74/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 74/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 74/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 74/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 75/128: http.ResponseWriter http.CloseNotifier io.ReaderFrom interface{ EnableFullDuplex() error }")
		type combination_74_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
			combination_74_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 75/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 75/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 75/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 75/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 75/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 75/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 75/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 75/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 75/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 76/128: http.ResponseWriter http.Flusher http.CloseNotifier io.ReaderFrom interface{ EnableFullDuplex() error }")
		type combination_75_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
			combination_75_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 76/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 76/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 76/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 76/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 76/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 76/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 76/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 76/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 76/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 77/128: http.ResponseWriter http.Hijacker io.ReaderFrom interface{ EnableFullDuplex() error }")
		type combination_76_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
			combination_76_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 77/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 77/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 77/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 77/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 77/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 77/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 77/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 77/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 77/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 78/128: http.ResponseWriter http.Flusher http.Hijacker io.ReaderFrom interface{ EnableFullDuplex() error }")
		type combination_77_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
			combination_77_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 78/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 78/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 78/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 78/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 78/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 78/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 78/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 78/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 78/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 79/128: http.ResponseWriter http.CloseNotifier http.Hijacker io.ReaderFrom interface{ EnableFullDuplex() error }")
		type combination_78_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			combination_78_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 79/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 79/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 79/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 79/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 79/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 79/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 79/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 79/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 79/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 80/128: http.ResponseWriter http.Flusher http.CloseNotifier http.Hijacker io.ReaderFrom interface{ EnableFullDuplex() error }")
		type combination_79_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			combination_79_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 80/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 80/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 80/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 80/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 80/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 80/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 80/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 80/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 80/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 81/128: http.ResponseWriter http.Pusher interface{ EnableFullDuplex() error }")
		type combination_80_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Pusher
			combination_80_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 81/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 81/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 81/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 81/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 81/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 81/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 81/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 81/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 81/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 82/128: http.ResponseWriter http.Flusher http.Pusher interface{ EnableFullDuplex() error }")
		type combination_81_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.Pusher
			combination_81_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 82/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 82/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 82/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 82/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 82/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 82/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 82/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 82/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 82/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 83/128: http.ResponseWriter http.CloseNotifier http.Pusher interface{ EnableFullDuplex() error }")
		type combination_82_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Pusher
			combination_82_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 83/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 83/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 83/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 83/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 83/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 83/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 83/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 83/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 83/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 84/128: http.ResponseWriter http.Flusher http.CloseNotifier http.Pusher interface{ EnableFullDuplex() error }")
		type combination_83_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Pusher
			combination_83_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 84/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 84/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 84/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 84/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 84/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 84/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 84/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 84/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 84/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 85/128: http.ResponseWriter http.Hijacker http.Pusher interface{ EnableFullDuplex() error }")
		type combination_84_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Hijacker
			http.Pusher
			combination_84_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 85/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 85/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 85/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 85/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 85/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 85/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 85/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 85/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 85/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 86/128: http.ResponseWriter http.Flusher http.Hijacker http.Pusher interface{ EnableFullDuplex() error }")
		type combination_85_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			http.Pusher
			combination_85_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 86/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 86/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 86/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 86/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 86/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 86/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 86/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 86/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 86/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 87/128: http.ResponseWriter http.CloseNotifier http.Hijacker http.Pusher interface{ EnableFullDuplex() error }")
		type combination_86_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			http.Pusher
			combination_86_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 87/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 87/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 87/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 87/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 87/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 87/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 87/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 87/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 87/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 88/128: http.ResponseWriter http.Flusher http.CloseNotifier http.Hijacker http.Pusher interface{ EnableFullDuplex() error }")
		type combination_87_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			http.Pusher
			combination_87_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 88/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 88/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 88/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 88/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 88/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 88/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 88/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 88/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 88/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 89/128: http.ResponseWriter io.ReaderFrom http.Pusher interface{ EnableFullDuplex() error }")
		type combination_88_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			io.ReaderFrom
			http.Pusher
			combination_88_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 89/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 89/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 89/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 89/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 89/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 89/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 89/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 89/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 89/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 90/128: http.ResponseWriter http.Flusher io.ReaderFrom http.Pusher interface{ EnableFullDuplex() error }")
		type combination_89_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
			http.Pusher
			combination_89_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 90/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 90/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 90/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 90/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 90/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 90/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 90/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 90/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 90/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 91/128: http.ResponseWriter http.CloseNotifier io.ReaderFrom http.Pusher interface{ EnableFullDuplex() error }")
		type combination_90_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
			http.Pusher
			combination_90_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 91/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 91/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 91/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 91/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 91/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 91/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 91/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 91/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 91/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 92/128: http.ResponseWriter http.Flusher http.CloseNotifier io.ReaderFrom http.Pusher interface{ EnableFullDuplex() error }")
		type combination_91_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
			http.Pusher
			combination_91_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 92/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 92/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 92/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 92/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 92/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 92/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 92/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 92/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 92/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 93/128: http.ResponseWriter http.Hijacker io.ReaderFrom http.Pusher interface{ EnableFullDuplex() error }")
		type combination_92_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
			http.Pusher
			combination_92_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 93/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 93/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 93/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 93/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 93/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 93/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 93/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 93/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 93/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 94/128: http.ResponseWriter http.Flusher http.Hijacker io.ReaderFrom http.Pusher interface{ EnableFullDuplex() error }")
		type combination_93_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
			http.Pusher
			combination_93_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 94/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 94/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 94/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 94/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 94/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 94/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 94/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 94/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 94/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 95/128: http.ResponseWriter http.CloseNotifier http.Hijacker io.ReaderFrom http.Pusher interface{ EnableFullDuplex() error }")
		type combination_94_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			http.Pusher
			combination_94_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 95/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 95/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 95/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 95/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 95/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 95/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 95/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 95/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 95/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 96/128: http.ResponseWriter http.Flusher http.CloseNotifier http.Hijacker io.ReaderFrom http.Pusher interface{ EnableFullDuplex() error }")
		type combination_95_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			http.Pusher
			combination_95_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 96/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 96/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 96/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 96/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 96/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != false {
			t.Errorf("combination 96/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 96/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 96/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 96/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 97/128: http.ResponseWriter interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_96_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_96_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			combination_96_5
			combination_96_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 97/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 97/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 97/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 97/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 97/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 97/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 97/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 97/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 97/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 98/128: http.ResponseWriter http.Flusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_97_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_97_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			combination_97_5
			combination_97_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 98/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 98/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 98/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 98/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 98/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 98/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 98/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 98/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 98/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 99/128: http.ResponseWriter http.CloseNotifier interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_98_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_98_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
			combination_98_5
			combination_98_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 99/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 99/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 99/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 99/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 99/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 99/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 99/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 99/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 99/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 100/128: http.ResponseWriter http.Flusher http.CloseNotifier interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_99_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_99_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			combination_99_5
			combination_99_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 100/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 100/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 100/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 100/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 100/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 100/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 100/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 100/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 100/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 101/128: http.ResponseWriter http.Hijacker interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_100_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_100_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Hijacker
			combination_100_5
			combination_100_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 101/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 101/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 101/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 101/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 101/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 101/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 101/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 101/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 101/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 102/128: http.ResponseWriter http.Flusher http.Hijacker interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_101_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_101_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			combination_101_5
			combination_101_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 102/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 102/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 102/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 102/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 102/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 102/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 102/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 102/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 102/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 103/128: http.ResponseWriter http.CloseNotifier http.Hijacker interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_102_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_102_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			combination_102_5
			combination_102_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 103/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 103/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 103/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 103/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 103/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 103/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 103/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 103/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 103/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 104/128: http.ResponseWriter http.Flusher http.CloseNotifier http.Hijacker interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_103_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_103_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			combination_103_5
			combination_103_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 104/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 104/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 104/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 104/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 104/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 104/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 104/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 104/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 104/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 105/128: http.ResponseWriter io.ReaderFrom interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_104_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_104_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			io.ReaderFrom
			combination_104_5
			combination_104_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 105/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 105/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 105/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 105/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 105/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 105/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 105/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 105/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 105/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 106/128: http.ResponseWriter http.Flusher io.ReaderFrom interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_105_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_105_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
			combination_105_5
			combination_105_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 106/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 106/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 106/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 106/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 106/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 106/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 106/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 106/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 106/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 107/128: http.ResponseWriter http.CloseNotifier io.ReaderFrom interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_106_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_106_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
			combination_106_5
			combination_106_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 107/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 107/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 107/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 107/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 107/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 107/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 107/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 107/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 107/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 108/128: http.ResponseWriter http.Flusher http.CloseNotifier io.ReaderFrom interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_107_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_107_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
			combination_107_5
			combination_107_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 108/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 108/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 108/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 108/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 108/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 108/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 108/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 108/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 108/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 109/128: http.ResponseWriter http.Hijacker io.ReaderFrom interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_108_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_108_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
			combination_108_5
			combination_108_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 109/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 109/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 109/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 109/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 109/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 109/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 109/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 109/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 109/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 110/128: http.ResponseWriter http.Flusher http.Hijacker io.ReaderFrom interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_109_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_109_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
			combination_109_5
			combination_109_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 110/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 110/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 110/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 110/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 110/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 110/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 110/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 110/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 110/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 111/128: http.ResponseWriter http.CloseNotifier http.Hijacker io.ReaderFrom interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_110_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_110_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			combination_110_5
			combination_110_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 111/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 111/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 111/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 111/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 111/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 111/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 111/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 111/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 111/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 112/128: http.ResponseWriter http.Flusher http.CloseNotifier http.Hijacker io.ReaderFrom interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_111_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_111_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			combination_111_5
			combination_111_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 112/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 112/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 112/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 112/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != false {
			t.Errorf("combination 112/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 112/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 112/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 112/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 112/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 113/128: http.ResponseWriter http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_112_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_112_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Pusher
			combination_112_5
			combination_112_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 113/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 113/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 113/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 113/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 113/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 113/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 113/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 113/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 113/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 114/128: http.ResponseWriter http.Flusher http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_113_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_113_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.Pusher
			combination_113_5
			combination_113_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 114/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 114/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 114/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 114/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 114/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 114/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 114/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 114/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 114/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 115/128: http.ResponseWriter http.CloseNotifier http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_114_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_114_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Pusher
			combination_114_5
			combination_114_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 115/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 115/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 115/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 115/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 115/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 115/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 115/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 115/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 115/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 116/128: http.ResponseWriter http.Flusher http.CloseNotifier http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_115_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_115_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Pusher
			combination_115_5
			combination_115_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 116/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 116/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 116/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 116/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 116/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 116/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 116/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 116/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 116/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 117/128: http.ResponseWriter http.Hijacker http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_116_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_116_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Hijacker
			http.Pusher
			combination_116_5
			combination_116_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 117/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 117/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 117/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 117/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 117/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 117/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 117/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 117/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 117/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 118/128: http.ResponseWriter http.Flusher http.Hijacker http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_117_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_117_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			http.Pusher
			combination_117_5
			combination_117_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 118/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 118/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 118/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 118/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 118/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 118/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 118/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 118/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 118/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 119/128: http.ResponseWriter http.CloseNotifier http.Hijacker http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_118_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_118_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			http.Pusher
			combination_118_5
			combination_118_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 119/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 119/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 119/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 119/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 119/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 119/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 119/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 119/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 119/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 120/128: http.ResponseWriter http.Flusher http.CloseNotifier http.Hijacker http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_119_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_119_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			http.Pusher
			combination_119_5
			combination_119_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 120/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 120/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 120/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 120/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 120/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 120/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 120/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 120/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 120/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 121/128: http.ResponseWriter io.ReaderFrom http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_120_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_120_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			io.ReaderFrom
			http.Pusher
			combination_120_5
			combination_120_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 121/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 121/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 121/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 121/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 121/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 121/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 121/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 121/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 121/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 122/128: http.ResponseWriter http.Flusher io.ReaderFrom http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_121_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_121_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			io.ReaderFrom
			http.Pusher
			combination_121_5
			combination_121_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 122/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 122/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 122/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 122/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 122/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 122/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 122/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 122/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 122/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 123/128: http.ResponseWriter http.CloseNotifier io.ReaderFrom http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_122_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_122_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
			io.ReaderFrom
			http.Pusher
			combination_122_5
			combination_122_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 123/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 123/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 123/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 123/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 123/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 123/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 123/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 123/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 123/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 124/128: http.ResponseWriter http.Flusher http.CloseNotifier io.ReaderFrom http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_123_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_123_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			io.ReaderFrom
			http.Pusher
			combination_123_5
			combination_123_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 124/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 124/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != false {
			t.Errorf("combination 124/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 124/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 124/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 124/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 124/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 124/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 124/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 125/128: http.ResponseWriter http.Hijacker io.ReaderFrom http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_124_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_124_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Hijacker
			io.ReaderFrom
			http.Pusher
			combination_124_5
			combination_124_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 125/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 125/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 125/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 125/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 125/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 125/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 125/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 125/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 125/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 126/128: http.ResponseWriter http.Flusher http.Hijacker io.ReaderFrom http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_125_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_125_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.Hijacker
			io.ReaderFrom
			http.Pusher
			combination_125_5
			combination_125_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 126/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != false {
			t.Errorf("combination 126/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 126/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 126/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 126/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 126/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 126/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 126/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 126/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 127/128: http.ResponseWriter http.CloseNotifier http.Hijacker io.ReaderFrom http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_126_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_126_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			http.Pusher
			combination_126_5
			combination_126_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != false {
			t.Errorf("combination 127/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 127/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 127/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 127/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 127/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 127/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 127/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 127/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 127/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 128/128: http.ResponseWriter http.Flusher http.CloseNotifier http.Hijacker io.ReaderFrom http.Pusher interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error } interface{ EnableFullDuplex() error }")
		type combination_127_5 interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}
		type combination_127_6 interface{ EnableFullDuplex() error }
		wrapped := struct {
			http.ResponseWriter
			http.Flusher
			http.CloseNotifier
			http.Hijacker
			io.ReaderFrom
			http.Pusher
			combination_127_5
			combination_127_6
		}{}
		w := WrapResponseWriter(wrapped, ResponseWriterInterceptor{})

		if _, ok := w.(http.Flusher); ok != true {
			t.Errorf("combination 128/128: unexpected interface http.Flusher")
		}
		if _, ok := w.(http.CloseNotifier); ok != true {
			t.Errorf("combination 128/128: unexpected interface http.CloseNotifier")
		}
		if _, ok := w.(http.Hijacker); ok != true {
			t.Errorf("combination 128/128: unexpected interface http.Hijacker")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 128/128: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(http.Pusher); ok != true {
			t.Errorf("combination 128/128: unexpected interface http.Pusher")
		}
		if _, ok := w.(interface {
			SetReadDeadline(time.Time) error
			SetWriteDeadline(time.Time) error
		}); ok != true {
			t.Errorf("combination 128/128: unexpected interface interface{ SetReadDeadline(time.Time) error; SetWriteDeadline(time.Time) error }")
		}
		if _, ok := w.(interface{ EnableFullDuplex() error }); ok != true {
			t.Errorf("combination 128/128: unexpected interface interface{ EnableFullDuplex() error }")
		}

		if w, ok := w.(ResponseWriterUnwrapper); ok {
			if w.UnwrapResponseWriter() != wrapped {
				t.Errorf("combination 128/128: UnwrapResponseWriter() failed")
			}
		} else {
			t.Errorf("combination 128/128: ResponseWriterUnwrapper interface not implemented")
		}
	}
}
