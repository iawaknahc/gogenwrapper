package stdlibio

import (
	io "io"
	testing "testing"
)

func TestWrap(t *testing.T) {
	{
		t.Log("combination 1/256: any")
		wrapped := struct {
			any
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 1/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 1/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 1/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 1/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 1/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 1/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 1/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 1/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 1/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 1/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 2/256: any io.Closer")
		wrapped := struct {
			any
			io.Closer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 2/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 2/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 2/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 2/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 2/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 2/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 2/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 2/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 2/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 2/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 3/256: any io.Reader")
		wrapped := struct {
			any
			io.Reader
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 3/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 3/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 3/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 3/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 3/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 3/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 3/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 3/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 3/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 3/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 4/256: any io.Closer io.Reader")
		wrapped := struct {
			any
			io.Closer
			io.Reader
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 4/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 4/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 4/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 4/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 4/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 4/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 4/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 4/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 4/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 4/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 5/256: any io.ReaderAt")
		wrapped := struct {
			any
			io.ReaderAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 5/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 5/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 5/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 5/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 5/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 5/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 5/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 5/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 5/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 5/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 6/256: any io.Closer io.ReaderAt")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 6/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 6/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 6/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 6/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 6/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 6/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 6/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 6/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 6/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 6/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 7/256: any io.Reader io.ReaderAt")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 7/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 7/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 7/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 7/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 7/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 7/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 7/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 7/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 7/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 7/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 8/256: any io.Closer io.Reader io.ReaderAt")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 8/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 8/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 8/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 8/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 8/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 8/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 8/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 8/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 8/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 8/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 9/256: any io.ReaderFrom")
		wrapped := struct {
			any
			io.ReaderFrom
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 9/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 9/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 9/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 9/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 9/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 9/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 9/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 9/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 9/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 9/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 10/256: any io.Closer io.ReaderFrom")
		wrapped := struct {
			any
			io.Closer
			io.ReaderFrom
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 10/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 10/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 10/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 10/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 10/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 10/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 10/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 10/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 10/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 10/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 11/256: any io.Reader io.ReaderFrom")
		wrapped := struct {
			any
			io.Reader
			io.ReaderFrom
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 11/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 11/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 11/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 11/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 11/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 11/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 11/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 11/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 11/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 11/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 12/256: any io.Closer io.Reader io.ReaderFrom")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderFrom
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 12/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 12/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 12/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 12/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 12/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 12/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 12/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 12/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 12/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 12/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 13/256: any io.ReaderAt io.ReaderFrom")
		wrapped := struct {
			any
			io.ReaderAt
			io.ReaderFrom
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 13/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 13/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 13/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 13/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 13/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 13/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 13/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 13/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 13/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 13/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 14/256: any io.Closer io.ReaderAt io.ReaderFrom")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 14/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 14/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 14/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 14/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 14/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 14/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 14/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 14/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 14/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 14/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 15/256: any io.Reader io.ReaderAt io.ReaderFrom")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 15/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 15/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 15/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 15/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 15/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 15/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 15/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 15/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 15/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 15/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 16/256: any io.Closer io.Reader io.ReaderAt io.ReaderFrom")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 16/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 16/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 16/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 16/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 16/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 16/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 16/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 16/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 16/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 16/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 17/256: any io.Seeker")
		wrapped := struct {
			any
			io.Seeker
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 17/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 17/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 17/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 17/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 17/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 17/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 17/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 17/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 17/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 17/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 18/256: any io.Closer io.Seeker")
		wrapped := struct {
			any
			io.Closer
			io.Seeker
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 18/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 18/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 18/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 18/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 18/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 18/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 18/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 18/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 18/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 18/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 19/256: any io.Reader io.Seeker")
		wrapped := struct {
			any
			io.Reader
			io.Seeker
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 19/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 19/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 19/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 19/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 19/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 19/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 19/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 19/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 19/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 19/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 20/256: any io.Closer io.Reader io.Seeker")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.Seeker
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 20/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 20/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 20/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 20/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 20/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 20/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 20/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 20/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 20/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 20/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 21/256: any io.ReaderAt io.Seeker")
		wrapped := struct {
			any
			io.ReaderAt
			io.Seeker
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 21/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 21/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 21/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 21/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 21/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 21/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 21/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 21/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 21/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 21/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 22/256: any io.Closer io.ReaderAt io.Seeker")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
			io.Seeker
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 22/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 22/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 22/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 22/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 22/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 22/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 22/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 22/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 22/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 22/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 23/256: any io.Reader io.ReaderAt io.Seeker")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
			io.Seeker
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 23/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 23/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 23/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 23/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 23/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 23/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 23/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 23/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 23/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 23/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 24/256: any io.Closer io.Reader io.ReaderAt io.Seeker")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.Seeker
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 24/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 24/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 24/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 24/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 24/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 24/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 24/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 24/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 24/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 24/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 25/256: any io.ReaderFrom io.Seeker")
		wrapped := struct {
			any
			io.ReaderFrom
			io.Seeker
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 25/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 25/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 25/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 25/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 25/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 25/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 25/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 25/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 25/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 25/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 26/256: any io.Closer io.ReaderFrom io.Seeker")
		wrapped := struct {
			any
			io.Closer
			io.ReaderFrom
			io.Seeker
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 26/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 26/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 26/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 26/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 26/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 26/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 26/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 26/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 26/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 26/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 27/256: any io.Reader io.ReaderFrom io.Seeker")
		wrapped := struct {
			any
			io.Reader
			io.ReaderFrom
			io.Seeker
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 27/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 27/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 27/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 27/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 27/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 27/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 27/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 27/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 27/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 27/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 28/256: any io.Closer io.Reader io.ReaderFrom io.Seeker")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderFrom
			io.Seeker
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 28/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 28/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 28/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 28/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 28/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 28/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 28/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 28/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 28/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 28/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 29/256: any io.ReaderAt io.ReaderFrom io.Seeker")
		wrapped := struct {
			any
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 29/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 29/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 29/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 29/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 29/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 29/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 29/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 29/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 29/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 29/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 30/256: any io.Closer io.ReaderAt io.ReaderFrom io.Seeker")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 30/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 30/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 30/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 30/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 30/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 30/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 30/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 30/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 30/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 30/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 31/256: any io.Reader io.ReaderAt io.ReaderFrom io.Seeker")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 31/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 31/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 31/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 31/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 31/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 31/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 31/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 31/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 31/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 31/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 32/256: any io.Closer io.Reader io.ReaderAt io.ReaderFrom io.Seeker")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 32/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 32/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 32/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 32/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 32/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 32/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 32/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 32/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 32/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 32/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 33/256: any io.Writer")
		wrapped := struct {
			any
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 33/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 33/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 33/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 33/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 33/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 33/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 33/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 33/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 33/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 33/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 34/256: any io.Closer io.Writer")
		wrapped := struct {
			any
			io.Closer
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 34/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 34/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 34/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 34/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 34/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 34/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 34/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 34/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 34/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 34/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 35/256: any io.Reader io.Writer")
		wrapped := struct {
			any
			io.Reader
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 35/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 35/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 35/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 35/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 35/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 35/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 35/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 35/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 35/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 35/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 36/256: any io.Closer io.Reader io.Writer")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 36/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 36/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 36/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 36/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 36/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 36/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 36/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 36/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 36/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 36/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 37/256: any io.ReaderAt io.Writer")
		wrapped := struct {
			any
			io.ReaderAt
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 37/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 37/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 37/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 37/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 37/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 37/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 37/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 37/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 37/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 37/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 38/256: any io.Closer io.ReaderAt io.Writer")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 38/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 38/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 38/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 38/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 38/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 38/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 38/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 38/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 38/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 38/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 39/256: any io.Reader io.ReaderAt io.Writer")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 39/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 39/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 39/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 39/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 39/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 39/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 39/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 39/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 39/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 39/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 40/256: any io.Closer io.Reader io.ReaderAt io.Writer")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 40/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 40/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 40/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 40/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 40/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 40/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 40/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 40/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 40/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 40/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 41/256: any io.ReaderFrom io.Writer")
		wrapped := struct {
			any
			io.ReaderFrom
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 41/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 41/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 41/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 41/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 41/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 41/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 41/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 41/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 41/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 41/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 42/256: any io.Closer io.ReaderFrom io.Writer")
		wrapped := struct {
			any
			io.Closer
			io.ReaderFrom
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 42/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 42/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 42/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 42/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 42/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 42/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 42/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 42/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 42/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 42/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 43/256: any io.Reader io.ReaderFrom io.Writer")
		wrapped := struct {
			any
			io.Reader
			io.ReaderFrom
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 43/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 43/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 43/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 43/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 43/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 43/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 43/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 43/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 43/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 43/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 44/256: any io.Closer io.Reader io.ReaderFrom io.Writer")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderFrom
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 44/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 44/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 44/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 44/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 44/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 44/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 44/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 44/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 44/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 44/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 45/256: any io.ReaderAt io.ReaderFrom io.Writer")
		wrapped := struct {
			any
			io.ReaderAt
			io.ReaderFrom
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 45/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 45/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 45/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 45/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 45/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 45/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 45/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 45/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 45/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 45/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 46/256: any io.Closer io.ReaderAt io.ReaderFrom io.Writer")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 46/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 46/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 46/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 46/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 46/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 46/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 46/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 46/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 46/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 46/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 47/256: any io.Reader io.ReaderAt io.ReaderFrom io.Writer")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 47/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 47/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 47/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 47/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 47/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 47/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 47/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 47/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 47/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 47/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 48/256: any io.Closer io.Reader io.ReaderAt io.ReaderFrom io.Writer")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 48/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 48/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 48/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 48/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 48/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 48/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 48/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 48/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 48/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 48/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 49/256: any io.Seeker io.Writer")
		wrapped := struct {
			any
			io.Seeker
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 49/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 49/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 49/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 49/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 49/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 49/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 49/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 49/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 49/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 49/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 50/256: any io.Closer io.Seeker io.Writer")
		wrapped := struct {
			any
			io.Closer
			io.Seeker
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 50/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 50/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 50/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 50/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 50/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 50/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 50/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 50/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 50/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 50/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 51/256: any io.Reader io.Seeker io.Writer")
		wrapped := struct {
			any
			io.Reader
			io.Seeker
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 51/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 51/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 51/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 51/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 51/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 51/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 51/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 51/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 51/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 51/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 52/256: any io.Closer io.Reader io.Seeker io.Writer")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.Seeker
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 52/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 52/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 52/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 52/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 52/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 52/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 52/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 52/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 52/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 52/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 53/256: any io.ReaderAt io.Seeker io.Writer")
		wrapped := struct {
			any
			io.ReaderAt
			io.Seeker
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 53/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 53/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 53/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 53/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 53/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 53/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 53/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 53/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 53/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 53/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 54/256: any io.Closer io.ReaderAt io.Seeker io.Writer")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
			io.Seeker
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 54/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 54/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 54/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 54/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 54/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 54/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 54/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 54/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 54/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 54/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 55/256: any io.Reader io.ReaderAt io.Seeker io.Writer")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
			io.Seeker
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 55/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 55/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 55/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 55/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 55/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 55/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 55/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 55/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 55/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 55/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 56/256: any io.Closer io.Reader io.ReaderAt io.Seeker io.Writer")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.Seeker
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 56/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 56/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 56/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 56/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 56/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 56/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 56/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 56/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 56/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 56/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 57/256: any io.ReaderFrom io.Seeker io.Writer")
		wrapped := struct {
			any
			io.ReaderFrom
			io.Seeker
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 57/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 57/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 57/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 57/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 57/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 57/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 57/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 57/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 57/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 57/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 58/256: any io.Closer io.ReaderFrom io.Seeker io.Writer")
		wrapped := struct {
			any
			io.Closer
			io.ReaderFrom
			io.Seeker
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 58/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 58/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 58/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 58/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 58/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 58/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 58/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 58/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 58/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 58/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 59/256: any io.Reader io.ReaderFrom io.Seeker io.Writer")
		wrapped := struct {
			any
			io.Reader
			io.ReaderFrom
			io.Seeker
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 59/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 59/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 59/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 59/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 59/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 59/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 59/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 59/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 59/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 59/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 60/256: any io.Closer io.Reader io.ReaderFrom io.Seeker io.Writer")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderFrom
			io.Seeker
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 60/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 60/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 60/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 60/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 60/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 60/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 60/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 60/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 60/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 60/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 61/256: any io.ReaderAt io.ReaderFrom io.Seeker io.Writer")
		wrapped := struct {
			any
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 61/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 61/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 61/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 61/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 61/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 61/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 61/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 61/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 61/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 61/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 62/256: any io.Closer io.ReaderAt io.ReaderFrom io.Seeker io.Writer")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 62/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 62/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 62/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 62/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 62/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 62/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 62/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 62/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 62/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 62/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 63/256: any io.Reader io.ReaderAt io.ReaderFrom io.Seeker io.Writer")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 63/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 63/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 63/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 63/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 63/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 63/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 63/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 63/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 63/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 63/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 64/256: any io.Closer io.Reader io.ReaderAt io.ReaderFrom io.Seeker io.Writer")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 64/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 64/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 64/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 64/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 64/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 64/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 64/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 64/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 64/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 64/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 65/256: any io.WriterAt")
		wrapped := struct {
			any
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 65/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 65/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 65/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 65/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 65/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 65/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 65/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 65/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 65/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 65/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 66/256: any io.Closer io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 66/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 66/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 66/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 66/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 66/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 66/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 66/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 66/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 66/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 66/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 67/256: any io.Reader io.WriterAt")
		wrapped := struct {
			any
			io.Reader
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 67/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 67/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 67/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 67/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 67/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 67/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 67/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 67/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 67/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 67/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 68/256: any io.Closer io.Reader io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 68/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 68/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 68/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 68/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 68/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 68/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 68/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 68/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 68/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 68/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 69/256: any io.ReaderAt io.WriterAt")
		wrapped := struct {
			any
			io.ReaderAt
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 69/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 69/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 69/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 69/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 69/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 69/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 69/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 69/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 69/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 69/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 70/256: any io.Closer io.ReaderAt io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 70/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 70/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 70/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 70/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 70/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 70/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 70/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 70/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 70/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 70/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 71/256: any io.Reader io.ReaderAt io.WriterAt")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 71/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 71/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 71/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 71/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 71/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 71/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 71/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 71/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 71/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 71/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 72/256: any io.Closer io.Reader io.ReaderAt io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 72/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 72/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 72/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 72/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 72/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 72/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 72/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 72/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 72/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 72/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 73/256: any io.ReaderFrom io.WriterAt")
		wrapped := struct {
			any
			io.ReaderFrom
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 73/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 73/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 73/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 73/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 73/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 73/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 73/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 73/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 73/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 73/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 74/256: any io.Closer io.ReaderFrom io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.ReaderFrom
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 74/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 74/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 74/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 74/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 74/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 74/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 74/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 74/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 74/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 74/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 75/256: any io.Reader io.ReaderFrom io.WriterAt")
		wrapped := struct {
			any
			io.Reader
			io.ReaderFrom
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 75/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 75/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 75/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 75/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 75/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 75/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 75/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 75/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 75/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 75/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 76/256: any io.Closer io.Reader io.ReaderFrom io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderFrom
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 76/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 76/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 76/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 76/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 76/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 76/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 76/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 76/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 76/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 76/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 77/256: any io.ReaderAt io.ReaderFrom io.WriterAt")
		wrapped := struct {
			any
			io.ReaderAt
			io.ReaderFrom
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 77/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 77/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 77/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 77/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 77/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 77/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 77/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 77/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 77/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 77/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 78/256: any io.Closer io.ReaderAt io.ReaderFrom io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 78/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 78/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 78/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 78/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 78/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 78/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 78/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 78/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 78/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 78/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 79/256: any io.Reader io.ReaderAt io.ReaderFrom io.WriterAt")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 79/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 79/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 79/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 79/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 79/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 79/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 79/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 79/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 79/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 79/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 80/256: any io.Closer io.Reader io.ReaderAt io.ReaderFrom io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 80/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 80/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 80/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 80/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 80/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 80/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 80/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 80/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 80/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 80/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 81/256: any io.Seeker io.WriterAt")
		wrapped := struct {
			any
			io.Seeker
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 81/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 81/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 81/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 81/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 81/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 81/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 81/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 81/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 81/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 81/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 82/256: any io.Closer io.Seeker io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.Seeker
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 82/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 82/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 82/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 82/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 82/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 82/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 82/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 82/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 82/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 82/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 83/256: any io.Reader io.Seeker io.WriterAt")
		wrapped := struct {
			any
			io.Reader
			io.Seeker
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 83/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 83/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 83/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 83/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 83/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 83/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 83/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 83/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 83/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 83/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 84/256: any io.Closer io.Reader io.Seeker io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.Seeker
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 84/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 84/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 84/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 84/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 84/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 84/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 84/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 84/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 84/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 84/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 85/256: any io.ReaderAt io.Seeker io.WriterAt")
		wrapped := struct {
			any
			io.ReaderAt
			io.Seeker
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 85/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 85/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 85/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 85/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 85/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 85/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 85/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 85/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 85/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 85/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 86/256: any io.Closer io.ReaderAt io.Seeker io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
			io.Seeker
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 86/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 86/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 86/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 86/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 86/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 86/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 86/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 86/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 86/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 86/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 87/256: any io.Reader io.ReaderAt io.Seeker io.WriterAt")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
			io.Seeker
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 87/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 87/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 87/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 87/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 87/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 87/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 87/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 87/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 87/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 87/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 88/256: any io.Closer io.Reader io.ReaderAt io.Seeker io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.Seeker
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 88/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 88/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 88/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 88/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 88/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 88/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 88/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 88/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 88/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 88/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 89/256: any io.ReaderFrom io.Seeker io.WriterAt")
		wrapped := struct {
			any
			io.ReaderFrom
			io.Seeker
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 89/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 89/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 89/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 89/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 89/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 89/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 89/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 89/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 89/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 89/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 90/256: any io.Closer io.ReaderFrom io.Seeker io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.ReaderFrom
			io.Seeker
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 90/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 90/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 90/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 90/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 90/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 90/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 90/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 90/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 90/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 90/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 91/256: any io.Reader io.ReaderFrom io.Seeker io.WriterAt")
		wrapped := struct {
			any
			io.Reader
			io.ReaderFrom
			io.Seeker
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 91/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 91/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 91/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 91/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 91/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 91/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 91/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 91/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 91/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 91/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 92/256: any io.Closer io.Reader io.ReaderFrom io.Seeker io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderFrom
			io.Seeker
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 92/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 92/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 92/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 92/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 92/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 92/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 92/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 92/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 92/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 92/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 93/256: any io.ReaderAt io.ReaderFrom io.Seeker io.WriterAt")
		wrapped := struct {
			any
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 93/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 93/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 93/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 93/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 93/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 93/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 93/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 93/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 93/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 93/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 94/256: any io.Closer io.ReaderAt io.ReaderFrom io.Seeker io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 94/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 94/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 94/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 94/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 94/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 94/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 94/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 94/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 94/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 94/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 95/256: any io.Reader io.ReaderAt io.ReaderFrom io.Seeker io.WriterAt")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 95/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 95/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 95/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 95/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 95/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 95/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 95/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 95/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 95/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 95/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 96/256: any io.Closer io.Reader io.ReaderAt io.ReaderFrom io.Seeker io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 96/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 96/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 96/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 96/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 96/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 96/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 96/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 96/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 96/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 96/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 97/256: any io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 97/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 97/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 97/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 97/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 97/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 97/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 97/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 97/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 97/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 97/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 98/256: any io.Closer io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 98/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 98/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 98/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 98/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 98/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 98/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 98/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 98/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 98/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 98/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 99/256: any io.Reader io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.Reader
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 99/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 99/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 99/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 99/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 99/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 99/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 99/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 99/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 99/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 99/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 100/256: any io.Closer io.Reader io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 100/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 100/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 100/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 100/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 100/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 100/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 100/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 100/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 100/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 100/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 101/256: any io.ReaderAt io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.ReaderAt
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 101/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 101/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 101/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 101/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 101/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 101/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 101/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 101/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 101/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 101/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 102/256: any io.Closer io.ReaderAt io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 102/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 102/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 102/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 102/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 102/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 102/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 102/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 102/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 102/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 102/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 103/256: any io.Reader io.ReaderAt io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 103/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 103/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 103/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 103/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 103/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 103/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 103/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 103/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 103/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 103/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 104/256: any io.Closer io.Reader io.ReaderAt io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 104/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 104/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 104/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 104/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 104/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 104/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 104/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 104/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 104/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 104/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 105/256: any io.ReaderFrom io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.ReaderFrom
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 105/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 105/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 105/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 105/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 105/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 105/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 105/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 105/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 105/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 105/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 106/256: any io.Closer io.ReaderFrom io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.ReaderFrom
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 106/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 106/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 106/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 106/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 106/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 106/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 106/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 106/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 106/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 106/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 107/256: any io.Reader io.ReaderFrom io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.Reader
			io.ReaderFrom
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 107/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 107/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 107/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 107/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 107/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 107/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 107/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 107/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 107/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 107/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 108/256: any io.Closer io.Reader io.ReaderFrom io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderFrom
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 108/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 108/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 108/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 108/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 108/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 108/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 108/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 108/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 108/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 108/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 109/256: any io.ReaderAt io.ReaderFrom io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.ReaderAt
			io.ReaderFrom
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 109/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 109/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 109/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 109/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 109/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 109/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 109/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 109/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 109/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 109/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 110/256: any io.Closer io.ReaderAt io.ReaderFrom io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 110/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 110/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 110/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 110/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 110/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 110/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 110/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 110/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 110/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 110/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 111/256: any io.Reader io.ReaderAt io.ReaderFrom io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 111/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 111/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 111/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 111/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 111/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 111/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 111/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 111/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 111/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 111/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 112/256: any io.Closer io.Reader io.ReaderAt io.ReaderFrom io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 112/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 112/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 112/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 112/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 112/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 112/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 112/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 112/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 112/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 112/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 113/256: any io.Seeker io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.Seeker
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 113/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 113/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 113/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 113/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 113/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 113/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 113/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 113/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 113/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 113/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 114/256: any io.Closer io.Seeker io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.Seeker
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 114/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 114/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 114/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 114/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 114/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 114/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 114/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 114/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 114/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 114/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 115/256: any io.Reader io.Seeker io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.Reader
			io.Seeker
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 115/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 115/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 115/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 115/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 115/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 115/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 115/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 115/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 115/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 115/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 116/256: any io.Closer io.Reader io.Seeker io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.Seeker
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 116/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 116/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 116/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 116/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 116/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 116/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 116/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 116/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 116/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 116/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 117/256: any io.ReaderAt io.Seeker io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.ReaderAt
			io.Seeker
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 117/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 117/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 117/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 117/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 117/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 117/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 117/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 117/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 117/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 117/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 118/256: any io.Closer io.ReaderAt io.Seeker io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
			io.Seeker
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 118/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 118/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 118/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 118/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 118/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 118/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 118/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 118/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 118/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 118/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 119/256: any io.Reader io.ReaderAt io.Seeker io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
			io.Seeker
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 119/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 119/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 119/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 119/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 119/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 119/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 119/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 119/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 119/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 119/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 120/256: any io.Closer io.Reader io.ReaderAt io.Seeker io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.Seeker
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 120/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 120/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 120/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 120/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 120/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 120/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 120/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 120/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 120/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 120/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 121/256: any io.ReaderFrom io.Seeker io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 121/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 121/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 121/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 121/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 121/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 121/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 121/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 121/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 121/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 121/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 122/256: any io.Closer io.ReaderFrom io.Seeker io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 122/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 122/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 122/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 122/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 122/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 122/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 122/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 122/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 122/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 122/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 123/256: any io.Reader io.ReaderFrom io.Seeker io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.Reader
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 123/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 123/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 123/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 123/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 123/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 123/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 123/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 123/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 123/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 123/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 124/256: any io.Closer io.Reader io.ReaderFrom io.Seeker io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 124/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 124/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 124/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 124/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 124/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 124/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 124/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 124/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 124/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 124/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 125/256: any io.ReaderAt io.ReaderFrom io.Seeker io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 125/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 125/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 125/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 125/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 125/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 125/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 125/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 125/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 125/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 125/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 126/256: any io.Closer io.ReaderAt io.ReaderFrom io.Seeker io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 126/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 126/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 126/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 126/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 126/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 126/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 126/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 126/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 126/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 126/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 127/256: any io.Reader io.ReaderAt io.ReaderFrom io.Seeker io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 127/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 127/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 127/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 127/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 127/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 127/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 127/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 127/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 127/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 127/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 128/256: any io.Closer io.Reader io.ReaderAt io.ReaderFrom io.Seeker io.Writer io.WriterAt")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 128/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 128/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 128/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 128/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 128/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 128/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 128/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != false {
			t.Errorf("combination 128/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 128/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 128/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 129/256: any io.WriterTo")
		wrapped := struct {
			any
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 129/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 129/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 129/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 129/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 129/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 129/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 129/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 129/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 129/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 129/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 130/256: any io.Closer io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 130/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 130/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 130/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 130/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 130/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 130/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 130/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 130/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 130/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 130/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 131/256: any io.Reader io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 131/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 131/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 131/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 131/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 131/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 131/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 131/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 131/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 131/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 131/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 132/256: any io.Closer io.Reader io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 132/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 132/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 132/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 132/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 132/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 132/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 132/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 132/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 132/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 132/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 133/256: any io.ReaderAt io.WriterTo")
		wrapped := struct {
			any
			io.ReaderAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 133/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 133/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 133/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 133/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 133/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 133/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 133/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 133/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 133/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 133/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 134/256: any io.Closer io.ReaderAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 134/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 134/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 134/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 134/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 134/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 134/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 134/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 134/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 134/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 134/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 135/256: any io.Reader io.ReaderAt io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 135/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 135/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 135/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 135/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 135/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 135/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 135/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 135/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 135/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 135/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 136/256: any io.Closer io.Reader io.ReaderAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 136/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 136/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 136/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 136/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 136/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 136/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 136/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 136/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 136/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 136/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 137/256: any io.ReaderFrom io.WriterTo")
		wrapped := struct {
			any
			io.ReaderFrom
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 137/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 137/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 137/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 137/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 137/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 137/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 137/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 137/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 137/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 137/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 138/256: any io.Closer io.ReaderFrom io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.ReaderFrom
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 138/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 138/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 138/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 138/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 138/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 138/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 138/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 138/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 138/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 138/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 139/256: any io.Reader io.ReaderFrom io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.ReaderFrom
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 139/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 139/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 139/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 139/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 139/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 139/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 139/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 139/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 139/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 139/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 140/256: any io.Closer io.Reader io.ReaderFrom io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderFrom
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 140/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 140/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 140/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 140/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 140/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 140/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 140/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 140/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 140/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 140/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 141/256: any io.ReaderAt io.ReaderFrom io.WriterTo")
		wrapped := struct {
			any
			io.ReaderAt
			io.ReaderFrom
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 141/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 141/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 141/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 141/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 141/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 141/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 141/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 141/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 141/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 141/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 142/256: any io.Closer io.ReaderAt io.ReaderFrom io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 142/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 142/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 142/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 142/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 142/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 142/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 142/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 142/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 142/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 142/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 143/256: any io.Reader io.ReaderAt io.ReaderFrom io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 143/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 143/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 143/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 143/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 143/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 143/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 143/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 143/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 143/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 143/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 144/256: any io.Closer io.Reader io.ReaderAt io.ReaderFrom io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 144/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 144/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 144/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 144/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 144/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 144/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 144/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 144/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 144/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 144/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 145/256: any io.Seeker io.WriterTo")
		wrapped := struct {
			any
			io.Seeker
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 145/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 145/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 145/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 145/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 145/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 145/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 145/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 145/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 145/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 145/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 146/256: any io.Closer io.Seeker io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Seeker
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 146/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 146/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 146/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 146/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 146/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 146/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 146/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 146/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 146/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 146/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 147/256: any io.Reader io.Seeker io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.Seeker
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 147/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 147/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 147/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 147/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 147/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 147/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 147/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 147/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 147/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 147/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 148/256: any io.Closer io.Reader io.Seeker io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.Seeker
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 148/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 148/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 148/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 148/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 148/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 148/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 148/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 148/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 148/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 148/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 149/256: any io.ReaderAt io.Seeker io.WriterTo")
		wrapped := struct {
			any
			io.ReaderAt
			io.Seeker
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 149/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 149/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 149/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 149/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 149/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 149/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 149/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 149/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 149/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 149/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 150/256: any io.Closer io.ReaderAt io.Seeker io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
			io.Seeker
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 150/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 150/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 150/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 150/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 150/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 150/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 150/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 150/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 150/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 150/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 151/256: any io.Reader io.ReaderAt io.Seeker io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
			io.Seeker
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 151/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 151/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 151/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 151/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 151/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 151/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 151/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 151/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 151/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 151/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 152/256: any io.Closer io.Reader io.ReaderAt io.Seeker io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.Seeker
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 152/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 152/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 152/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 152/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 152/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 152/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 152/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 152/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 152/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 152/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 153/256: any io.ReaderFrom io.Seeker io.WriterTo")
		wrapped := struct {
			any
			io.ReaderFrom
			io.Seeker
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 153/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 153/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 153/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 153/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 153/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 153/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 153/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 153/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 153/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 153/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 154/256: any io.Closer io.ReaderFrom io.Seeker io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.ReaderFrom
			io.Seeker
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 154/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 154/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 154/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 154/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 154/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 154/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 154/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 154/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 154/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 154/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 155/256: any io.Reader io.ReaderFrom io.Seeker io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.ReaderFrom
			io.Seeker
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 155/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 155/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 155/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 155/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 155/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 155/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 155/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 155/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 155/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 155/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 156/256: any io.Closer io.Reader io.ReaderFrom io.Seeker io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderFrom
			io.Seeker
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 156/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 156/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 156/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 156/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 156/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 156/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 156/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 156/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 156/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 156/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 157/256: any io.ReaderAt io.ReaderFrom io.Seeker io.WriterTo")
		wrapped := struct {
			any
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 157/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 157/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 157/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 157/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 157/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 157/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 157/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 157/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 157/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 157/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 158/256: any io.Closer io.ReaderAt io.ReaderFrom io.Seeker io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 158/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 158/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 158/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 158/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 158/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 158/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 158/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 158/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 158/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 158/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 159/256: any io.Reader io.ReaderAt io.ReaderFrom io.Seeker io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 159/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 159/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 159/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 159/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 159/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 159/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 159/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 159/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 159/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 159/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 160/256: any io.Closer io.Reader io.ReaderAt io.ReaderFrom io.Seeker io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 160/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 160/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 160/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 160/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 160/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 160/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 160/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 160/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 160/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 160/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 161/256: any io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 161/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 161/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 161/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 161/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 161/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 161/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 161/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 161/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 161/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 161/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 162/256: any io.Closer io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 162/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 162/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 162/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 162/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 162/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 162/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 162/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 162/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 162/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 162/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 163/256: any io.Reader io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 163/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 163/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 163/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 163/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 163/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 163/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 163/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 163/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 163/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 163/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 164/256: any io.Closer io.Reader io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 164/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 164/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 164/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 164/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 164/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 164/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 164/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 164/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 164/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 164/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 165/256: any io.ReaderAt io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.ReaderAt
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 165/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 165/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 165/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 165/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 165/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 165/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 165/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 165/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 165/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 165/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 166/256: any io.Closer io.ReaderAt io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 166/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 166/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 166/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 166/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 166/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 166/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 166/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 166/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 166/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 166/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 167/256: any io.Reader io.ReaderAt io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 167/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 167/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 167/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 167/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 167/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 167/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 167/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 167/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 167/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 167/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 168/256: any io.Closer io.Reader io.ReaderAt io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 168/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 168/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 168/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 168/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 168/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 168/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 168/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 168/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 168/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 168/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 169/256: any io.ReaderFrom io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.ReaderFrom
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 169/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 169/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 169/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 169/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 169/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 169/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 169/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 169/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 169/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 169/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 170/256: any io.Closer io.ReaderFrom io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.ReaderFrom
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 170/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 170/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 170/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 170/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 170/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 170/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 170/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 170/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 170/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 170/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 171/256: any io.Reader io.ReaderFrom io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.ReaderFrom
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 171/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 171/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 171/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 171/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 171/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 171/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 171/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 171/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 171/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 171/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 172/256: any io.Closer io.Reader io.ReaderFrom io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderFrom
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 172/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 172/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 172/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 172/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 172/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 172/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 172/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 172/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 172/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 172/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 173/256: any io.ReaderAt io.ReaderFrom io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.ReaderAt
			io.ReaderFrom
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 173/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 173/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 173/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 173/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 173/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 173/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 173/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 173/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 173/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 173/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 174/256: any io.Closer io.ReaderAt io.ReaderFrom io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 174/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 174/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 174/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 174/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 174/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 174/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 174/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 174/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 174/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 174/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 175/256: any io.Reader io.ReaderAt io.ReaderFrom io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 175/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 175/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 175/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 175/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 175/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 175/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 175/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 175/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 175/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 175/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 176/256: any io.Closer io.Reader io.ReaderAt io.ReaderFrom io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 176/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 176/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 176/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 176/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 176/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 176/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 176/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 176/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 176/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 176/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 177/256: any io.Seeker io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.Seeker
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 177/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 177/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 177/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 177/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 177/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 177/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 177/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 177/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 177/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 177/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 178/256: any io.Closer io.Seeker io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Seeker
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 178/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 178/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 178/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 178/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 178/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 178/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 178/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 178/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 178/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 178/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 179/256: any io.Reader io.Seeker io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.Seeker
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 179/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 179/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 179/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 179/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 179/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 179/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 179/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 179/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 179/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 179/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 180/256: any io.Closer io.Reader io.Seeker io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.Seeker
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 180/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 180/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 180/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 180/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 180/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 180/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 180/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 180/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 180/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 180/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 181/256: any io.ReaderAt io.Seeker io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.ReaderAt
			io.Seeker
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 181/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 181/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 181/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 181/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 181/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 181/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 181/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 181/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 181/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 181/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 182/256: any io.Closer io.ReaderAt io.Seeker io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
			io.Seeker
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 182/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 182/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 182/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 182/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 182/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 182/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 182/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 182/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 182/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 182/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 183/256: any io.Reader io.ReaderAt io.Seeker io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
			io.Seeker
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 183/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 183/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 183/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 183/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 183/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 183/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 183/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 183/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 183/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 183/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 184/256: any io.Closer io.Reader io.ReaderAt io.Seeker io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.Seeker
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 184/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 184/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 184/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 184/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 184/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 184/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 184/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 184/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 184/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 184/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 185/256: any io.ReaderFrom io.Seeker io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 185/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 185/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 185/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 185/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 185/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 185/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 185/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 185/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 185/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 185/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 186/256: any io.Closer io.ReaderFrom io.Seeker io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 186/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 186/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 186/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 186/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 186/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 186/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 186/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 186/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 186/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 186/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 187/256: any io.Reader io.ReaderFrom io.Seeker io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 187/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 187/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 187/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 187/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 187/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 187/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 187/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 187/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 187/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 187/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 188/256: any io.Closer io.Reader io.ReaderFrom io.Seeker io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 188/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 188/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 188/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 188/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 188/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 188/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 188/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 188/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 188/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 188/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 189/256: any io.ReaderAt io.ReaderFrom io.Seeker io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 189/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 189/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 189/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 189/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 189/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 189/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 189/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 189/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 189/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 189/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 190/256: any io.Closer io.ReaderAt io.ReaderFrom io.Seeker io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 190/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 190/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 190/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 190/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 190/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 190/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 190/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 190/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 190/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 190/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 191/256: any io.Reader io.ReaderAt io.ReaderFrom io.Seeker io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 191/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 191/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 191/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 191/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 191/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 191/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 191/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 191/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 191/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 191/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 192/256: any io.Closer io.Reader io.ReaderAt io.ReaderFrom io.Seeker io.Writer io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 192/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 192/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 192/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 192/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 192/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 192/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != false {
			t.Errorf("combination 192/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 192/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 192/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 192/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 193/256: any io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 193/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 193/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 193/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 193/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 193/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 193/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 193/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 193/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 193/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 193/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 194/256: any io.Closer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 194/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 194/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 194/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 194/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 194/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 194/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 194/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 194/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 194/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 194/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 195/256: any io.Reader io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 195/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 195/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 195/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 195/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 195/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 195/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 195/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 195/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 195/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 195/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 196/256: any io.Closer io.Reader io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 196/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 196/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 196/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 196/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 196/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 196/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 196/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 196/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 196/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 196/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 197/256: any io.ReaderAt io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.ReaderAt
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 197/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 197/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 197/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 197/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 197/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 197/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 197/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 197/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 197/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 197/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 198/256: any io.Closer io.ReaderAt io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 198/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 198/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 198/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 198/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 198/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 198/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 198/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 198/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 198/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 198/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 199/256: any io.Reader io.ReaderAt io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 199/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 199/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 199/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 199/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 199/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 199/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 199/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 199/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 199/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 199/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 200/256: any io.Closer io.Reader io.ReaderAt io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 200/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 200/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 200/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 200/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 200/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 200/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 200/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 200/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 200/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 200/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 201/256: any io.ReaderFrom io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.ReaderFrom
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 201/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 201/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 201/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 201/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 201/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 201/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 201/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 201/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 201/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 201/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 202/256: any io.Closer io.ReaderFrom io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.ReaderFrom
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 202/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 202/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 202/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 202/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 202/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 202/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 202/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 202/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 202/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 202/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 203/256: any io.Reader io.ReaderFrom io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.ReaderFrom
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 203/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 203/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 203/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 203/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 203/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 203/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 203/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 203/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 203/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 203/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 204/256: any io.Closer io.Reader io.ReaderFrom io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderFrom
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 204/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 204/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 204/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 204/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 204/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 204/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 204/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 204/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 204/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 204/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 205/256: any io.ReaderAt io.ReaderFrom io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.ReaderAt
			io.ReaderFrom
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 205/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 205/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 205/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 205/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 205/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 205/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 205/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 205/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 205/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 205/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 206/256: any io.Closer io.ReaderAt io.ReaderFrom io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 206/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 206/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 206/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 206/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 206/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 206/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 206/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 206/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 206/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 206/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 207/256: any io.Reader io.ReaderAt io.ReaderFrom io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 207/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 207/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 207/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 207/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 207/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 207/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 207/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 207/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 207/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 207/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 208/256: any io.Closer io.Reader io.ReaderAt io.ReaderFrom io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 208/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 208/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 208/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 208/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 208/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 208/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 208/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 208/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 208/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 208/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 209/256: any io.Seeker io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 209/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 209/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 209/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 209/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 209/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 209/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 209/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 209/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 209/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 209/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 210/256: any io.Closer io.Seeker io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 210/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 210/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 210/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 210/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 210/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 210/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 210/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 210/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 210/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 210/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 211/256: any io.Reader io.Seeker io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 211/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 211/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 211/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 211/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 211/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 211/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 211/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 211/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 211/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 211/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 212/256: any io.Closer io.Reader io.Seeker io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 212/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 212/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 212/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 212/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 212/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 212/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 212/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 212/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 212/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 212/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 213/256: any io.ReaderAt io.Seeker io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.ReaderAt
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 213/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 213/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 213/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 213/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 213/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 213/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 213/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 213/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 213/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 213/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 214/256: any io.Closer io.ReaderAt io.Seeker io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 214/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 214/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 214/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 214/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 214/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 214/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 214/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 214/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 214/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 214/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 215/256: any io.Reader io.ReaderAt io.Seeker io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 215/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 215/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 215/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 215/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 215/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 215/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 215/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 215/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 215/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 215/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 216/256: any io.Closer io.Reader io.ReaderAt io.Seeker io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 216/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 216/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 216/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 216/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 216/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 216/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 216/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 216/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 216/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 216/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 217/256: any io.ReaderFrom io.Seeker io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.ReaderFrom
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 217/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 217/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 217/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 217/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 217/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 217/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 217/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 217/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 217/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 217/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 218/256: any io.Closer io.ReaderFrom io.Seeker io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.ReaderFrom
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 218/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 218/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 218/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 218/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 218/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 218/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 218/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 218/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 218/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 218/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 219/256: any io.Reader io.ReaderFrom io.Seeker io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.ReaderFrom
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 219/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 219/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 219/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 219/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 219/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 219/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 219/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 219/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 219/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 219/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 220/256: any io.Closer io.Reader io.ReaderFrom io.Seeker io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderFrom
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 220/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 220/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 220/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 220/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 220/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 220/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 220/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 220/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 220/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 220/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 221/256: any io.ReaderAt io.ReaderFrom io.Seeker io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 221/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 221/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 221/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 221/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 221/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 221/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 221/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 221/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 221/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 221/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 222/256: any io.Closer io.ReaderAt io.ReaderFrom io.Seeker io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 222/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 222/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 222/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 222/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 222/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 222/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 222/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 222/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 222/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 222/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 223/256: any io.Reader io.ReaderAt io.ReaderFrom io.Seeker io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 223/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 223/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 223/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 223/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 223/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 223/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 223/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 223/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 223/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 223/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 224/256: any io.Closer io.Reader io.ReaderAt io.ReaderFrom io.Seeker io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 224/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 224/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 224/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 224/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 224/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != false {
			t.Errorf("combination 224/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 224/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 224/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 224/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 224/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 225/256: any io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 225/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 225/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 225/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 225/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 225/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 225/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 225/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 225/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 225/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 225/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 226/256: any io.Closer io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 226/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 226/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 226/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 226/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 226/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 226/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 226/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 226/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 226/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 226/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 227/256: any io.Reader io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 227/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 227/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 227/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 227/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 227/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 227/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 227/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 227/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 227/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 227/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 228/256: any io.Closer io.Reader io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 228/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 228/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 228/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 228/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 228/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 228/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 228/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 228/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 228/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 228/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 229/256: any io.ReaderAt io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.ReaderAt
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 229/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 229/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 229/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 229/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 229/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 229/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 229/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 229/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 229/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 229/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 230/256: any io.Closer io.ReaderAt io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 230/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 230/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 230/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 230/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 230/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 230/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 230/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 230/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 230/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 230/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 231/256: any io.Reader io.ReaderAt io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 231/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 231/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 231/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 231/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 231/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 231/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 231/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 231/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 231/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 231/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 232/256: any io.Closer io.Reader io.ReaderAt io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 232/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 232/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 232/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 232/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 232/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 232/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 232/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 232/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 232/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 232/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 233/256: any io.ReaderFrom io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.ReaderFrom
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 233/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 233/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 233/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 233/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 233/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 233/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 233/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 233/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 233/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 233/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 234/256: any io.Closer io.ReaderFrom io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.ReaderFrom
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 234/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 234/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 234/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 234/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 234/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 234/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 234/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 234/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 234/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 234/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 235/256: any io.Reader io.ReaderFrom io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.ReaderFrom
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 235/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 235/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 235/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 235/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 235/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 235/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 235/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 235/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 235/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 235/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 236/256: any io.Closer io.Reader io.ReaderFrom io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderFrom
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 236/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 236/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 236/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 236/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 236/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 236/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 236/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 236/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 236/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 236/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 237/256: any io.ReaderAt io.ReaderFrom io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.ReaderAt
			io.ReaderFrom
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 237/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 237/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 237/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 237/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 237/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 237/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 237/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 237/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 237/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 237/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 238/256: any io.Closer io.ReaderAt io.ReaderFrom io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 238/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 238/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 238/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 238/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 238/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 238/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 238/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 238/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 238/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 238/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 239/256: any io.Reader io.ReaderAt io.ReaderFrom io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 239/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 239/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 239/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 239/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 239/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 239/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 239/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 239/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 239/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 239/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 240/256: any io.Closer io.Reader io.ReaderAt io.ReaderFrom io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 240/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 240/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 240/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 240/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != false {
			t.Errorf("combination 240/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 240/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 240/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 240/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 240/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 240/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 241/256: any io.Seeker io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 241/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 241/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 241/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 241/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 241/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 241/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 241/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 241/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 241/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 241/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 242/256: any io.Closer io.Seeker io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 242/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 242/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 242/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 242/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 242/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 242/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 242/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 242/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 242/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 242/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 243/256: any io.Reader io.Seeker io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 243/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 243/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 243/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 243/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 243/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 243/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 243/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 243/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 243/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 243/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 244/256: any io.Closer io.Reader io.Seeker io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 244/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 244/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 244/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 244/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 244/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 244/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 244/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 244/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 244/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 244/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 245/256: any io.ReaderAt io.Seeker io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.ReaderAt
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 245/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 245/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 245/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 245/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 245/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 245/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 245/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 245/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 245/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 245/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 246/256: any io.Closer io.ReaderAt io.Seeker io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 246/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 246/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 246/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 246/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 246/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 246/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 246/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 246/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 246/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 246/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 247/256: any io.Reader io.ReaderAt io.Seeker io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 247/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 247/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 247/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 247/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 247/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 247/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 247/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 247/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 247/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 247/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 248/256: any io.Closer io.Reader io.ReaderAt io.Seeker io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 248/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 248/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 248/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != false {
			t.Errorf("combination 248/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 248/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 248/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 248/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 248/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 248/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 248/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 249/256: any io.ReaderFrom io.Seeker io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 249/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 249/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 249/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 249/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 249/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 249/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 249/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 249/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 249/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 249/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 250/256: any io.Closer io.ReaderFrom io.Seeker io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 250/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 250/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 250/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 250/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 250/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 250/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 250/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 250/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 250/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 250/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 251/256: any io.Reader io.ReaderFrom io.Seeker io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 251/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 251/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 251/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 251/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 251/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 251/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 251/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 251/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 251/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 251/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 252/256: any io.Closer io.Reader io.ReaderFrom io.Seeker io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 252/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 252/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != false {
			t.Errorf("combination 252/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 252/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 252/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 252/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 252/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 252/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 252/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 252/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 253/256: any io.ReaderAt io.ReaderFrom io.Seeker io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 253/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 253/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 253/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 253/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 253/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 253/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 253/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 253/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 253/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 253/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 254/256: any io.Closer io.ReaderAt io.ReaderFrom io.Seeker io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 254/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != false {
			t.Errorf("combination 254/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 254/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 254/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 254/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 254/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 254/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 254/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 254/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 254/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 255/256: any io.Reader io.ReaderAt io.ReaderFrom io.Seeker io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != false {
			t.Errorf("combination 255/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 255/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 255/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 255/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 255/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 255/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 255/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 255/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 255/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 255/256: Unwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 256/256: any io.Closer io.Reader io.ReaderAt io.ReaderFrom io.Seeker io.Writer io.WriterAt io.WriterTo")
		wrapped := struct {
			any
			io.Closer
			io.Reader
			io.ReaderAt
			io.ReaderFrom
			io.Seeker
			io.Writer
			io.WriterAt
			io.WriterTo
		}{}
		w := Wrap(wrapped, Interceptor{})

		if _, ok := w.(io.Closer); ok != true {
			t.Errorf("combination 256/256: unexpected interface io.Closer")
		}
		if _, ok := w.(io.Reader); ok != true {
			t.Errorf("combination 256/256: unexpected interface io.Reader")
		}
		if _, ok := w.(io.ReaderAt); ok != true {
			t.Errorf("combination 256/256: unexpected interface io.ReaderAt")
		}
		if _, ok := w.(io.ReaderFrom); ok != true {
			t.Errorf("combination 256/256: unexpected interface io.ReaderFrom")
		}
		if _, ok := w.(io.Seeker); ok != true {
			t.Errorf("combination 256/256: unexpected interface io.Seeker")
		}
		if _, ok := w.(io.Writer); ok != true {
			t.Errorf("combination 256/256: unexpected interface io.Writer")
		}
		if _, ok := w.(io.WriterAt); ok != true {
			t.Errorf("combination 256/256: unexpected interface io.WriterAt")
		}
		if _, ok := w.(io.WriterTo); ok != true {
			t.Errorf("combination 256/256: unexpected interface io.WriterTo")
		}

		if w, ok := w.(Unwrapper); ok {
			if w.Unwrap() != wrapped {
				t.Errorf("combination 256/256: Unwrap() failed")
			}
		} else {
			t.Errorf("combination 256/256: Unwrapper interface not implemented")
		}
	}
}
