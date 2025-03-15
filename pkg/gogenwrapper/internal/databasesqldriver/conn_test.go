package databasesqldriver

import (
	driver "database/sql/driver"
	testing "testing"
)

func TestWrapConn(t *testing.T) {
	{
		t.Log("combination 1/512: driver.Conn")
		wrapped := struct {
			driver.Conn
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 1/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 1/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 1/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 1/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 1/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 1/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 1/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 1/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 1/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 1/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 1/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 2/512: driver.Conn driver.ConnBeginTx")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 2/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 2/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 2/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 2/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 2/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 2/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 2/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 2/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 2/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 2/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 2/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 3/512: driver.Conn driver.Execer")
		wrapped := struct {
			driver.Conn
			driver.Execer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 3/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 3/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 3/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 3/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 3/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 3/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 3/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 3/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 3/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 3/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 3/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 4/512: driver.Conn driver.ConnBeginTx driver.Execer")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 4/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 4/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 4/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 4/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 4/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 4/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 4/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 4/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 4/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 4/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 4/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 5/512: driver.Conn driver.ExecerContext")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 5/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 5/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 5/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 5/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 5/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 5/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 5/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 5/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 5/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 5/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 5/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 6/512: driver.Conn driver.ConnBeginTx driver.ExecerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 6/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 6/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 6/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 6/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 6/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 6/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 6/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 6/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 6/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 6/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 6/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 7/512: driver.Conn driver.Execer driver.ExecerContext")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 7/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 7/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 7/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 7/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 7/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 7/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 7/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 7/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 7/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 7/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 7/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 8/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 8/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 8/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 8/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 8/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 8/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 8/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 8/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 8/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 8/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 8/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 8/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 9/512: driver.Conn driver.NamedValueChecker")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 9/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 9/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 9/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 9/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 9/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 9/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 9/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 9/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 9/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 9/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 9/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 10/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 10/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 10/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 10/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 10/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 10/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 10/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 10/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 10/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 10/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 10/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 10/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 11/512: driver.Conn driver.Execer driver.NamedValueChecker")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 11/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 11/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 11/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 11/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 11/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 11/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 11/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 11/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 11/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 11/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 11/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 12/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 12/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 12/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 12/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 12/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 12/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 12/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 12/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 12/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 12/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 12/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 12/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 13/512: driver.Conn driver.ExecerContext driver.NamedValueChecker")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 13/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 13/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 13/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 13/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 13/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 13/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 13/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 13/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 13/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 13/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 13/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 14/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 14/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 14/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 14/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 14/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 14/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 14/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 14/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 14/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 14/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 14/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 14/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 15/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 15/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 15/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 15/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 15/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 15/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 15/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 15/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 15/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 15/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 15/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 15/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 16/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 16/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 16/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 16/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 16/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 16/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 16/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 16/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 16/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 16/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 16/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 16/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 17/512: driver.Conn driver.Pinger")
		wrapped := struct {
			driver.Conn
			driver.Pinger
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 17/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 17/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 17/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 17/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 17/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 17/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 17/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 17/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 17/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 17/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 17/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 18/512: driver.Conn driver.ConnBeginTx driver.Pinger")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 18/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 18/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 18/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 18/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 18/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 18/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 18/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 18/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 18/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 18/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 18/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 19/512: driver.Conn driver.Execer driver.Pinger")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.Pinger
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 19/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 19/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 19/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 19/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 19/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 19/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 19/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 19/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 19/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 19/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 19/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 20/512: driver.Conn driver.ConnBeginTx driver.Execer driver.Pinger")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 20/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 20/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 20/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 20/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 20/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 20/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 20/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 20/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 20/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 20/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 20/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 21/512: driver.Conn driver.ExecerContext driver.Pinger")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.Pinger
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 21/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 21/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 21/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 21/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 21/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 21/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 21/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 21/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 21/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 21/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 21/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 22/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.Pinger")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 22/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 22/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 22/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 22/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 22/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 22/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 22/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 22/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 22/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 22/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 22/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 23/512: driver.Conn driver.Execer driver.ExecerContext driver.Pinger")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 23/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 23/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 23/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 23/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 23/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 23/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 23/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 23/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 23/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 23/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 23/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 24/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.Pinger")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 24/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 24/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 24/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 24/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 24/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 24/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 24/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 24/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 24/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 24/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 24/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 25/512: driver.Conn driver.NamedValueChecker driver.Pinger")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 25/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 25/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 25/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 25/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 25/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 25/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 25/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 25/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 25/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 25/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 25/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 26/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker driver.Pinger")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 26/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 26/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 26/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 26/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 26/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 26/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 26/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 26/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 26/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 26/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 26/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 27/512: driver.Conn driver.Execer driver.NamedValueChecker driver.Pinger")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 27/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 27/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 27/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 27/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 27/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 27/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 27/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 27/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 27/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 27/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 27/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 28/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker driver.Pinger")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 28/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 28/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 28/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 28/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 28/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 28/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 28/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 28/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 28/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 28/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 28/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 29/512: driver.Conn driver.ExecerContext driver.NamedValueChecker driver.Pinger")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 29/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 29/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 29/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 29/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 29/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 29/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 29/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 29/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 29/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 29/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 29/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 30/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker driver.Pinger")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 30/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 30/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 30/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 30/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 30/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 30/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 30/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 30/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 30/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 30/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 30/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 31/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 31/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 31/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 31/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 31/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 31/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 31/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 31/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 31/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 31/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 31/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 31/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 32/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 32/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 32/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 32/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 32/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 32/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 32/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 32/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 32/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 32/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 32/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 32/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 33/512: driver.Conn driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 33/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 33/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 33/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 33/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 33/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 33/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 33/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 33/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 33/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 33/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 33/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 34/512: driver.Conn driver.ConnBeginTx driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 34/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 34/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 34/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 34/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 34/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 34/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 34/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 34/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 34/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 34/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 34/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 35/512: driver.Conn driver.Execer driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 35/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 35/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 35/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 35/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 35/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 35/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 35/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 35/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 35/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 35/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 35/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 36/512: driver.Conn driver.ConnBeginTx driver.Execer driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 36/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 36/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 36/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 36/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 36/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 36/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 36/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 36/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 36/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 36/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 36/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 37/512: driver.Conn driver.ExecerContext driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 37/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 37/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 37/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 37/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 37/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 37/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 37/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 37/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 37/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 37/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 37/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 38/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 38/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 38/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 38/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 38/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 38/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 38/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 38/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 38/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 38/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 38/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 38/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 39/512: driver.Conn driver.Execer driver.ExecerContext driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 39/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 39/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 39/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 39/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 39/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 39/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 39/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 39/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 39/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 39/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 39/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 40/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 40/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 40/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 40/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 40/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 40/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 40/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 40/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 40/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 40/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 40/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 40/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 41/512: driver.Conn driver.NamedValueChecker driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 41/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 41/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 41/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 41/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 41/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 41/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 41/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 41/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 41/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 41/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 41/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 42/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 42/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 42/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 42/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 42/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 42/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 42/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 42/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 42/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 42/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 42/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 42/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 43/512: driver.Conn driver.Execer driver.NamedValueChecker driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 43/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 43/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 43/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 43/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 43/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 43/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 43/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 43/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 43/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 43/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 43/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 44/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 44/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 44/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 44/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 44/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 44/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 44/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 44/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 44/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 44/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 44/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 44/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 45/512: driver.Conn driver.ExecerContext driver.NamedValueChecker driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 45/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 45/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 45/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 45/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 45/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 45/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 45/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 45/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 45/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 45/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 45/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 46/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 46/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 46/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 46/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 46/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 46/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 46/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 46/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 46/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 46/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 46/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 46/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 47/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 47/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 47/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 47/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 47/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 47/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 47/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 47/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 47/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 47/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 47/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 47/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 48/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 48/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 48/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 48/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 48/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 48/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 48/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 48/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 48/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 48/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 48/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 48/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 49/512: driver.Conn driver.Pinger driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.Pinger
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 49/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 49/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 49/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 49/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 49/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 49/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 49/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 49/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 49/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 49/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 49/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 50/512: driver.Conn driver.ConnBeginTx driver.Pinger driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 50/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 50/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 50/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 50/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 50/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 50/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 50/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 50/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 50/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 50/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 50/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 51/512: driver.Conn driver.Execer driver.Pinger driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.Pinger
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 51/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 51/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 51/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 51/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 51/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 51/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 51/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 51/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 51/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 51/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 51/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 52/512: driver.Conn driver.ConnBeginTx driver.Execer driver.Pinger driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 52/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 52/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 52/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 52/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 52/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 52/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 52/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 52/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 52/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 52/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 52/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 53/512: driver.Conn driver.ExecerContext driver.Pinger driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 53/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 53/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 53/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 53/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 53/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 53/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 53/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 53/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 53/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 53/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 53/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 54/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.Pinger driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 54/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 54/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 54/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 54/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 54/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 54/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 54/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 54/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 54/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 54/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 54/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 55/512: driver.Conn driver.Execer driver.ExecerContext driver.Pinger driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 55/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 55/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 55/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 55/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 55/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 55/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 55/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 55/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 55/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 55/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 55/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 56/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.Pinger driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 56/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 56/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 56/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 56/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 56/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 56/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 56/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 56/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 56/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 56/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 56/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 57/512: driver.Conn driver.NamedValueChecker driver.Pinger driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 57/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 57/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 57/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 57/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 57/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 57/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 57/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 57/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 57/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 57/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 57/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 58/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker driver.Pinger driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 58/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 58/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 58/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 58/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 58/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 58/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 58/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 58/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 58/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 58/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 58/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 59/512: driver.Conn driver.Execer driver.NamedValueChecker driver.Pinger driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 59/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 59/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 59/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 59/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 59/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 59/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 59/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 59/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 59/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 59/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 59/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 60/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker driver.Pinger driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 60/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 60/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 60/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 60/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 60/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 60/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 60/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 60/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 60/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 60/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 60/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 61/512: driver.Conn driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 61/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 61/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 61/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 61/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 61/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 61/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 61/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 61/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 61/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 61/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 61/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 62/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 62/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 62/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 62/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 62/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 62/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 62/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 62/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 62/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 62/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 62/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 62/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 63/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 63/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 63/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 63/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 63/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 63/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 63/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 63/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 63/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 63/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 63/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 63/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 64/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 64/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 64/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 64/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 64/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 64/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 64/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 64/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 64/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 64/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 64/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 64/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 65/512: driver.Conn driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 65/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 65/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 65/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 65/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 65/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 65/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 65/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 65/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 65/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 65/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 65/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 66/512: driver.Conn driver.ConnBeginTx driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 66/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 66/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 66/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 66/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 66/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 66/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 66/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 66/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 66/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 66/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 66/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 67/512: driver.Conn driver.Execer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 67/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 67/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 67/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 67/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 67/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 67/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 67/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 67/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 67/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 67/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 67/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 68/512: driver.Conn driver.ConnBeginTx driver.Execer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 68/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 68/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 68/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 68/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 68/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 68/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 68/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 68/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 68/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 68/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 68/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 69/512: driver.Conn driver.ExecerContext driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 69/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 69/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 69/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 69/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 69/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 69/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 69/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 69/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 69/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 69/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 69/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 70/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 70/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 70/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 70/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 70/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 70/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 70/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 70/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 70/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 70/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 70/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 70/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 71/512: driver.Conn driver.Execer driver.ExecerContext driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 71/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 71/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 71/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 71/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 71/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 71/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 71/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 71/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 71/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 71/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 71/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 72/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 72/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 72/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 72/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 72/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 72/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 72/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 72/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 72/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 72/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 72/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 72/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 73/512: driver.Conn driver.NamedValueChecker driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 73/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 73/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 73/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 73/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 73/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 73/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 73/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 73/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 73/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 73/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 73/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 74/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 74/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 74/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 74/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 74/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 74/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 74/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 74/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 74/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 74/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 74/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 74/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 75/512: driver.Conn driver.Execer driver.NamedValueChecker driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 75/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 75/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 75/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 75/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 75/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 75/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 75/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 75/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 75/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 75/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 75/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 76/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 76/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 76/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 76/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 76/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 76/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 76/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 76/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 76/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 76/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 76/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 76/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 77/512: driver.Conn driver.ExecerContext driver.NamedValueChecker driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 77/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 77/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 77/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 77/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 77/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 77/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 77/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 77/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 77/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 77/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 77/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 78/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 78/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 78/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 78/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 78/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 78/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 78/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 78/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 78/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 78/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 78/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 78/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 79/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 79/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 79/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 79/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 79/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 79/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 79/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 79/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 79/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 79/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 79/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 79/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 80/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 80/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 80/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 80/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 80/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 80/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 80/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 80/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 80/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 80/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 80/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 80/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 81/512: driver.Conn driver.Pinger driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.Pinger
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 81/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 81/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 81/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 81/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 81/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 81/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 81/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 81/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 81/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 81/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 81/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 82/512: driver.Conn driver.ConnBeginTx driver.Pinger driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 82/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 82/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 82/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 82/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 82/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 82/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 82/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 82/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 82/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 82/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 82/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 83/512: driver.Conn driver.Execer driver.Pinger driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.Pinger
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 83/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 83/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 83/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 83/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 83/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 83/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 83/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 83/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 83/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 83/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 83/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 84/512: driver.Conn driver.ConnBeginTx driver.Execer driver.Pinger driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 84/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 84/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 84/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 84/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 84/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 84/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 84/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 84/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 84/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 84/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 84/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 85/512: driver.Conn driver.ExecerContext driver.Pinger driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 85/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 85/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 85/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 85/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 85/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 85/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 85/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 85/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 85/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 85/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 85/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 86/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.Pinger driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 86/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 86/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 86/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 86/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 86/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 86/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 86/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 86/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 86/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 86/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 86/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 87/512: driver.Conn driver.Execer driver.ExecerContext driver.Pinger driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 87/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 87/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 87/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 87/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 87/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 87/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 87/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 87/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 87/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 87/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 87/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 88/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.Pinger driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 88/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 88/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 88/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 88/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 88/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 88/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 88/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 88/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 88/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 88/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 88/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 89/512: driver.Conn driver.NamedValueChecker driver.Pinger driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 89/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 89/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 89/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 89/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 89/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 89/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 89/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 89/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 89/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 89/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 89/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 90/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker driver.Pinger driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 90/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 90/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 90/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 90/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 90/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 90/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 90/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 90/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 90/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 90/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 90/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 91/512: driver.Conn driver.Execer driver.NamedValueChecker driver.Pinger driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 91/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 91/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 91/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 91/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 91/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 91/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 91/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 91/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 91/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 91/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 91/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 92/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker driver.Pinger driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 92/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 92/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 92/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 92/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 92/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 92/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 92/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 92/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 92/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 92/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 92/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 93/512: driver.Conn driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 93/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 93/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 93/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 93/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 93/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 93/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 93/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 93/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 93/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 93/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 93/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 94/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 94/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 94/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 94/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 94/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 94/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 94/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 94/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 94/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 94/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 94/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 94/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 95/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 95/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 95/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 95/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 95/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 95/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 95/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 95/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 95/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 95/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 95/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 95/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 96/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 96/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 96/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 96/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 96/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 96/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 96/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 96/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 96/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 96/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 96/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 96/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 97/512: driver.Conn driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 97/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 97/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 97/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 97/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 97/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 97/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 97/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 97/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 97/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 97/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 97/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 98/512: driver.Conn driver.ConnBeginTx driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 98/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 98/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 98/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 98/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 98/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 98/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 98/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 98/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 98/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 98/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 98/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 99/512: driver.Conn driver.Execer driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 99/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 99/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 99/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 99/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 99/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 99/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 99/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 99/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 99/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 99/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 99/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 100/512: driver.Conn driver.ConnBeginTx driver.Execer driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 100/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 100/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 100/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 100/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 100/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 100/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 100/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 100/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 100/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 100/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 100/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 101/512: driver.Conn driver.ExecerContext driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 101/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 101/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 101/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 101/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 101/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 101/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 101/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 101/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 101/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 101/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 101/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 102/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 102/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 102/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 102/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 102/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 102/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 102/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 102/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 102/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 102/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 102/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 102/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 103/512: driver.Conn driver.Execer driver.ExecerContext driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 103/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 103/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 103/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 103/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 103/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 103/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 103/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 103/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 103/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 103/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 103/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 104/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 104/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 104/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 104/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 104/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 104/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 104/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 104/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 104/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 104/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 104/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 104/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 105/512: driver.Conn driver.NamedValueChecker driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 105/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 105/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 105/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 105/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 105/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 105/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 105/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 105/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 105/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 105/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 105/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 106/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 106/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 106/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 106/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 106/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 106/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 106/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 106/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 106/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 106/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 106/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 106/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 107/512: driver.Conn driver.Execer driver.NamedValueChecker driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 107/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 107/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 107/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 107/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 107/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 107/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 107/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 107/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 107/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 107/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 107/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 108/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 108/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 108/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 108/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 108/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 108/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 108/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 108/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 108/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 108/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 108/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 108/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 109/512: driver.Conn driver.ExecerContext driver.NamedValueChecker driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 109/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 109/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 109/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 109/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 109/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 109/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 109/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 109/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 109/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 109/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 109/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 110/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 110/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 110/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 110/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 110/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 110/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 110/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 110/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 110/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 110/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 110/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 110/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 111/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 111/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 111/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 111/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 111/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 111/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 111/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 111/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 111/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 111/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 111/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 111/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 112/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 112/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 112/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 112/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 112/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 112/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 112/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 112/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 112/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 112/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 112/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 112/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 113/512: driver.Conn driver.Pinger driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 113/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 113/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 113/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 113/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 113/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 113/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 113/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 113/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 113/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 113/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 113/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 114/512: driver.Conn driver.ConnBeginTx driver.Pinger driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 114/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 114/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 114/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 114/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 114/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 114/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 114/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 114/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 114/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 114/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 114/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 115/512: driver.Conn driver.Execer driver.Pinger driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 115/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 115/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 115/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 115/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 115/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 115/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 115/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 115/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 115/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 115/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 115/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 116/512: driver.Conn driver.ConnBeginTx driver.Execer driver.Pinger driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 116/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 116/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 116/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 116/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 116/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 116/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 116/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 116/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 116/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 116/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 116/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 117/512: driver.Conn driver.ExecerContext driver.Pinger driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 117/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 117/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 117/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 117/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 117/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 117/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 117/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 117/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 117/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 117/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 117/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 118/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.Pinger driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 118/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 118/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 118/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 118/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 118/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 118/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 118/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 118/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 118/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 118/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 118/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 119/512: driver.Conn driver.Execer driver.ExecerContext driver.Pinger driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 119/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 119/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 119/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 119/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 119/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 119/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 119/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 119/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 119/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 119/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 119/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 120/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.Pinger driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 120/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 120/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 120/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 120/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 120/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 120/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 120/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 120/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 120/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 120/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 120/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 121/512: driver.Conn driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 121/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 121/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 121/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 121/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 121/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 121/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 121/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 121/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 121/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 121/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 121/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 122/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 122/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 122/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 122/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 122/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 122/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 122/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 122/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 122/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 122/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 122/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 122/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 123/512: driver.Conn driver.Execer driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 123/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 123/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 123/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 123/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 123/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 123/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 123/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 123/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 123/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 123/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 123/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 124/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 124/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 124/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 124/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 124/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 124/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 124/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 124/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 124/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 124/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 124/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 124/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 125/512: driver.Conn driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 125/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 125/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 125/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 125/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 125/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 125/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 125/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 125/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 125/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 125/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 125/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 126/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 126/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 126/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 126/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 126/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 126/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 126/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 126/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 126/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 126/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 126/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 126/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 127/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 127/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 127/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 127/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 127/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 127/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 127/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 127/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 127/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 127/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 127/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 127/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 128/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 128/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 128/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 128/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 128/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 128/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 128/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 128/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 128/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 128/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 128/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 128/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 129/512: driver.Conn driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 129/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 129/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 129/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 129/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 129/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 129/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 129/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 129/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 129/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 129/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 129/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 130/512: driver.Conn driver.ConnBeginTx driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 130/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 130/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 130/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 130/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 130/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 130/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 130/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 130/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 130/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 130/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 130/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 131/512: driver.Conn driver.Execer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 131/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 131/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 131/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 131/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 131/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 131/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 131/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 131/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 131/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 131/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 131/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 132/512: driver.Conn driver.ConnBeginTx driver.Execer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 132/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 132/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 132/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 132/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 132/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 132/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 132/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 132/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 132/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 132/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 132/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 133/512: driver.Conn driver.ExecerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 133/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 133/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 133/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 133/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 133/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 133/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 133/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 133/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 133/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 133/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 133/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 134/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 134/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 134/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 134/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 134/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 134/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 134/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 134/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 134/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 134/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 134/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 134/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 135/512: driver.Conn driver.Execer driver.ExecerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 135/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 135/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 135/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 135/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 135/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 135/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 135/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 135/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 135/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 135/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 135/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 136/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 136/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 136/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 136/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 136/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 136/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 136/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 136/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 136/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 136/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 136/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 136/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 137/512: driver.Conn driver.NamedValueChecker driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 137/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 137/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 137/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 137/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 137/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 137/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 137/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 137/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 137/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 137/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 137/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 138/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 138/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 138/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 138/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 138/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 138/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 138/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 138/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 138/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 138/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 138/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 138/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 139/512: driver.Conn driver.Execer driver.NamedValueChecker driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 139/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 139/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 139/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 139/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 139/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 139/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 139/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 139/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 139/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 139/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 139/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 140/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 140/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 140/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 140/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 140/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 140/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 140/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 140/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 140/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 140/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 140/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 140/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 141/512: driver.Conn driver.ExecerContext driver.NamedValueChecker driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 141/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 141/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 141/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 141/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 141/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 141/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 141/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 141/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 141/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 141/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 141/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 142/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 142/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 142/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 142/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 142/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 142/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 142/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 142/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 142/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 142/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 142/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 142/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 143/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 143/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 143/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 143/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 143/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 143/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 143/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 143/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 143/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 143/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 143/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 143/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 144/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 144/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 144/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 144/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 144/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 144/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 144/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 144/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 144/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 144/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 144/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 144/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 145/512: driver.Conn driver.Pinger driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Pinger
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 145/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 145/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 145/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 145/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 145/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 145/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 145/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 145/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 145/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 145/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 145/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 146/512: driver.Conn driver.ConnBeginTx driver.Pinger driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 146/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 146/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 146/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 146/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 146/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 146/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 146/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 146/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 146/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 146/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 146/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 147/512: driver.Conn driver.Execer driver.Pinger driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.Pinger
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 147/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 147/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 147/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 147/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 147/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 147/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 147/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 147/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 147/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 147/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 147/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 148/512: driver.Conn driver.ConnBeginTx driver.Execer driver.Pinger driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 148/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 148/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 148/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 148/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 148/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 148/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 148/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 148/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 148/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 148/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 148/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 149/512: driver.Conn driver.ExecerContext driver.Pinger driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 149/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 149/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 149/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 149/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 149/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 149/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 149/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 149/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 149/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 149/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 149/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 150/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.Pinger driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 150/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 150/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 150/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 150/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 150/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 150/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 150/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 150/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 150/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 150/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 150/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 151/512: driver.Conn driver.Execer driver.ExecerContext driver.Pinger driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 151/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 151/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 151/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 151/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 151/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 151/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 151/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 151/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 151/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 151/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 151/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 152/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.Pinger driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 152/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 152/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 152/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 152/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 152/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 152/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 152/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 152/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 152/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 152/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 152/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 153/512: driver.Conn driver.NamedValueChecker driver.Pinger driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 153/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 153/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 153/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 153/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 153/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 153/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 153/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 153/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 153/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 153/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 153/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 154/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker driver.Pinger driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 154/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 154/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 154/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 154/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 154/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 154/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 154/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 154/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 154/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 154/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 154/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 155/512: driver.Conn driver.Execer driver.NamedValueChecker driver.Pinger driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 155/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 155/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 155/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 155/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 155/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 155/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 155/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 155/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 155/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 155/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 155/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 156/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker driver.Pinger driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 156/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 156/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 156/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 156/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 156/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 156/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 156/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 156/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 156/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 156/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 156/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 157/512: driver.Conn driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 157/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 157/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 157/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 157/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 157/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 157/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 157/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 157/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 157/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 157/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 157/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 158/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 158/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 158/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 158/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 158/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 158/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 158/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 158/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 158/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 158/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 158/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 158/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 159/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 159/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 159/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 159/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 159/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 159/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 159/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 159/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 159/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 159/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 159/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 159/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 160/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 160/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 160/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 160/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 160/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 160/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 160/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 160/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 160/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 160/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 160/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 160/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 161/512: driver.Conn driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 161/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 161/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 161/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 161/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 161/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 161/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 161/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 161/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 161/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 161/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 161/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 162/512: driver.Conn driver.ConnBeginTx driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 162/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 162/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 162/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 162/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 162/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 162/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 162/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 162/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 162/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 162/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 162/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 163/512: driver.Conn driver.Execer driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 163/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 163/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 163/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 163/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 163/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 163/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 163/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 163/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 163/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 163/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 163/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 164/512: driver.Conn driver.ConnBeginTx driver.Execer driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 164/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 164/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 164/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 164/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 164/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 164/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 164/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 164/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 164/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 164/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 164/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 165/512: driver.Conn driver.ExecerContext driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 165/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 165/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 165/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 165/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 165/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 165/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 165/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 165/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 165/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 165/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 165/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 166/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 166/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 166/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 166/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 166/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 166/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 166/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 166/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 166/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 166/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 166/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 166/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 167/512: driver.Conn driver.Execer driver.ExecerContext driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 167/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 167/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 167/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 167/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 167/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 167/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 167/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 167/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 167/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 167/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 167/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 168/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 168/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 168/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 168/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 168/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 168/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 168/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 168/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 168/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 168/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 168/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 168/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 169/512: driver.Conn driver.NamedValueChecker driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 169/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 169/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 169/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 169/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 169/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 169/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 169/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 169/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 169/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 169/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 169/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 170/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 170/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 170/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 170/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 170/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 170/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 170/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 170/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 170/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 170/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 170/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 170/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 171/512: driver.Conn driver.Execer driver.NamedValueChecker driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 171/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 171/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 171/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 171/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 171/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 171/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 171/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 171/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 171/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 171/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 171/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 172/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 172/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 172/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 172/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 172/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 172/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 172/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 172/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 172/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 172/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 172/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 172/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 173/512: driver.Conn driver.ExecerContext driver.NamedValueChecker driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 173/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 173/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 173/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 173/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 173/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 173/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 173/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 173/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 173/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 173/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 173/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 174/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 174/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 174/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 174/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 174/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 174/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 174/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 174/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 174/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 174/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 174/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 174/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 175/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 175/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 175/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 175/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 175/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 175/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 175/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 175/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 175/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 175/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 175/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 175/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 176/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 176/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 176/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 176/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 176/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 176/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 176/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 176/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 176/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 176/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 176/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 176/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 177/512: driver.Conn driver.Pinger driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 177/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 177/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 177/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 177/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 177/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 177/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 177/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 177/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 177/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 177/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 177/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 178/512: driver.Conn driver.ConnBeginTx driver.Pinger driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 178/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 178/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 178/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 178/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 178/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 178/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 178/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 178/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 178/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 178/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 178/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 179/512: driver.Conn driver.Execer driver.Pinger driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 179/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 179/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 179/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 179/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 179/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 179/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 179/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 179/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 179/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 179/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 179/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 180/512: driver.Conn driver.ConnBeginTx driver.Execer driver.Pinger driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 180/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 180/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 180/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 180/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 180/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 180/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 180/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 180/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 180/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 180/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 180/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 181/512: driver.Conn driver.ExecerContext driver.Pinger driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 181/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 181/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 181/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 181/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 181/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 181/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 181/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 181/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 181/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 181/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 181/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 182/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.Pinger driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 182/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 182/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 182/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 182/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 182/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 182/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 182/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 182/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 182/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 182/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 182/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 183/512: driver.Conn driver.Execer driver.ExecerContext driver.Pinger driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 183/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 183/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 183/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 183/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 183/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 183/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 183/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 183/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 183/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 183/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 183/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 184/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.Pinger driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 184/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 184/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 184/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 184/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 184/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 184/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 184/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 184/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 184/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 184/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 184/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 185/512: driver.Conn driver.NamedValueChecker driver.Pinger driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 185/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 185/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 185/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 185/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 185/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 185/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 185/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 185/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 185/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 185/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 185/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 186/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker driver.Pinger driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 186/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 186/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 186/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 186/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 186/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 186/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 186/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 186/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 186/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 186/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 186/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 187/512: driver.Conn driver.Execer driver.NamedValueChecker driver.Pinger driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 187/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 187/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 187/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 187/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 187/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 187/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 187/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 187/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 187/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 187/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 187/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 188/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker driver.Pinger driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 188/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 188/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 188/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 188/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 188/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 188/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 188/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 188/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 188/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 188/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 188/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 189/512: driver.Conn driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 189/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 189/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 189/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 189/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 189/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 189/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 189/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 189/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 189/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 189/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 189/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 190/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 190/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 190/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 190/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 190/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 190/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 190/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 190/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 190/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 190/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 190/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 190/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 191/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 191/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 191/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 191/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 191/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 191/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 191/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 191/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 191/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 191/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 191/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 191/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 192/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 192/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 192/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 192/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 192/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 192/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 192/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 192/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 192/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 192/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 192/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 192/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 193/512: driver.Conn driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 193/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 193/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 193/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 193/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 193/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 193/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 193/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 193/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 193/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 193/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 193/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 194/512: driver.Conn driver.ConnBeginTx driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 194/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 194/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 194/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 194/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 194/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 194/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 194/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 194/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 194/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 194/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 194/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 195/512: driver.Conn driver.Execer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 195/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 195/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 195/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 195/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 195/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 195/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 195/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 195/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 195/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 195/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 195/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 196/512: driver.Conn driver.ConnBeginTx driver.Execer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 196/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 196/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 196/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 196/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 196/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 196/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 196/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 196/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 196/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 196/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 196/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 197/512: driver.Conn driver.ExecerContext driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 197/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 197/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 197/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 197/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 197/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 197/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 197/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 197/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 197/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 197/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 197/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 198/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 198/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 198/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 198/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 198/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 198/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 198/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 198/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 198/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 198/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 198/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 198/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 199/512: driver.Conn driver.Execer driver.ExecerContext driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 199/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 199/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 199/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 199/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 199/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 199/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 199/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 199/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 199/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 199/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 199/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 200/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 200/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 200/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 200/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 200/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 200/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 200/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 200/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 200/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 200/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 200/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 200/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 201/512: driver.Conn driver.NamedValueChecker driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 201/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 201/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 201/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 201/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 201/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 201/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 201/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 201/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 201/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 201/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 201/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 202/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 202/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 202/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 202/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 202/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 202/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 202/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 202/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 202/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 202/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 202/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 202/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 203/512: driver.Conn driver.Execer driver.NamedValueChecker driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 203/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 203/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 203/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 203/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 203/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 203/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 203/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 203/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 203/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 203/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 203/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 204/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 204/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 204/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 204/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 204/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 204/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 204/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 204/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 204/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 204/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 204/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 204/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 205/512: driver.Conn driver.ExecerContext driver.NamedValueChecker driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 205/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 205/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 205/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 205/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 205/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 205/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 205/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 205/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 205/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 205/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 205/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 206/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 206/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 206/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 206/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 206/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 206/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 206/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 206/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 206/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 206/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 206/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 206/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 207/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 207/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 207/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 207/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 207/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 207/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 207/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 207/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 207/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 207/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 207/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 207/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 208/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 208/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 208/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 208/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 208/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 208/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 208/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 208/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 208/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 208/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 208/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 208/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 209/512: driver.Conn driver.Pinger driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 209/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 209/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 209/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 209/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 209/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 209/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 209/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 209/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 209/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 209/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 209/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 210/512: driver.Conn driver.ConnBeginTx driver.Pinger driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 210/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 210/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 210/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 210/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 210/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 210/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 210/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 210/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 210/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 210/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 210/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 211/512: driver.Conn driver.Execer driver.Pinger driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 211/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 211/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 211/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 211/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 211/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 211/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 211/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 211/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 211/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 211/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 211/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 212/512: driver.Conn driver.ConnBeginTx driver.Execer driver.Pinger driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 212/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 212/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 212/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 212/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 212/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 212/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 212/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 212/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 212/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 212/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 212/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 213/512: driver.Conn driver.ExecerContext driver.Pinger driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 213/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 213/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 213/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 213/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 213/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 213/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 213/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 213/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 213/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 213/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 213/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 214/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.Pinger driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 214/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 214/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 214/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 214/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 214/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 214/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 214/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 214/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 214/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 214/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 214/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 215/512: driver.Conn driver.Execer driver.ExecerContext driver.Pinger driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 215/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 215/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 215/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 215/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 215/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 215/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 215/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 215/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 215/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 215/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 215/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 216/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.Pinger driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 216/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 216/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 216/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 216/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 216/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 216/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 216/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 216/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 216/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 216/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 216/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 217/512: driver.Conn driver.NamedValueChecker driver.Pinger driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 217/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 217/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 217/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 217/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 217/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 217/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 217/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 217/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 217/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 217/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 217/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 218/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker driver.Pinger driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 218/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 218/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 218/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 218/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 218/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 218/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 218/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 218/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 218/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 218/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 218/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 219/512: driver.Conn driver.Execer driver.NamedValueChecker driver.Pinger driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 219/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 219/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 219/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 219/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 219/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 219/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 219/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 219/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 219/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 219/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 219/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 220/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker driver.Pinger driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 220/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 220/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 220/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 220/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 220/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 220/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 220/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 220/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 220/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 220/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 220/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 221/512: driver.Conn driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 221/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 221/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 221/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 221/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 221/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 221/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 221/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 221/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 221/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 221/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 221/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 222/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 222/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 222/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 222/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 222/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 222/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 222/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 222/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 222/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 222/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 222/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 222/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 223/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 223/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 223/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 223/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 223/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 223/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 223/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 223/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 223/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 223/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 223/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 223/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 224/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 224/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 224/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 224/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 224/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 224/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 224/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 224/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 224/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 224/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 224/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 224/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 225/512: driver.Conn driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 225/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 225/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 225/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 225/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 225/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 225/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 225/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 225/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 225/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 225/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 225/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 226/512: driver.Conn driver.ConnBeginTx driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 226/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 226/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 226/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 226/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 226/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 226/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 226/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 226/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 226/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 226/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 226/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 227/512: driver.Conn driver.Execer driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 227/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 227/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 227/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 227/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 227/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 227/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 227/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 227/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 227/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 227/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 227/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 228/512: driver.Conn driver.ConnBeginTx driver.Execer driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 228/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 228/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 228/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 228/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 228/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 228/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 228/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 228/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 228/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 228/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 228/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 229/512: driver.Conn driver.ExecerContext driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 229/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 229/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 229/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 229/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 229/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 229/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 229/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 229/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 229/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 229/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 229/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 230/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 230/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 230/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 230/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 230/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 230/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 230/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 230/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 230/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 230/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 230/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 230/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 231/512: driver.Conn driver.Execer driver.ExecerContext driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 231/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 231/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 231/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 231/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 231/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 231/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 231/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 231/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 231/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 231/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 231/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 232/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 232/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 232/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 232/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 232/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 232/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 232/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 232/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 232/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 232/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 232/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 232/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 233/512: driver.Conn driver.NamedValueChecker driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 233/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 233/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 233/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 233/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 233/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 233/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 233/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 233/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 233/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 233/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 233/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 234/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 234/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 234/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 234/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 234/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 234/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 234/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 234/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 234/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 234/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 234/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 234/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 235/512: driver.Conn driver.Execer driver.NamedValueChecker driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 235/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 235/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 235/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 235/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 235/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 235/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 235/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 235/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 235/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 235/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 235/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 236/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 236/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 236/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 236/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 236/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 236/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 236/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 236/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 236/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 236/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 236/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 236/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 237/512: driver.Conn driver.ExecerContext driver.NamedValueChecker driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 237/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 237/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 237/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 237/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 237/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 237/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 237/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 237/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 237/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 237/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 237/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 238/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 238/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 238/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 238/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 238/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 238/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 238/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 238/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 238/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 238/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 238/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 238/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 239/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 239/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 239/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 239/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 239/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 239/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 239/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 239/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 239/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 239/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 239/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 239/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 240/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 240/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 240/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 240/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 240/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 240/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 240/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 240/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 240/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 240/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 240/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 240/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 241/512: driver.Conn driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 241/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 241/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 241/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 241/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 241/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 241/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 241/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 241/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 241/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 241/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 241/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 242/512: driver.Conn driver.ConnBeginTx driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 242/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 242/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 242/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 242/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 242/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 242/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 242/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 242/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 242/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 242/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 242/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 243/512: driver.Conn driver.Execer driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 243/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 243/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 243/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 243/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 243/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 243/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 243/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 243/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 243/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 243/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 243/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 244/512: driver.Conn driver.ConnBeginTx driver.Execer driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 244/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 244/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 244/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 244/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 244/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 244/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 244/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 244/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 244/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 244/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 244/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 245/512: driver.Conn driver.ExecerContext driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 245/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 245/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 245/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 245/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 245/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 245/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 245/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 245/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 245/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 245/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 245/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 246/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 246/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 246/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 246/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 246/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 246/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 246/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 246/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 246/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 246/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 246/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 246/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 247/512: driver.Conn driver.Execer driver.ExecerContext driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 247/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 247/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 247/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 247/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 247/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 247/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 247/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 247/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 247/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 247/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 247/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 248/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 248/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 248/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 248/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 248/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 248/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 248/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 248/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 248/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 248/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 248/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 248/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 249/512: driver.Conn driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 249/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 249/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 249/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 249/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 249/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 249/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 249/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 249/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 249/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 249/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 249/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 250/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 250/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 250/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 250/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 250/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 250/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 250/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 250/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 250/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 250/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 250/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 250/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 251/512: driver.Conn driver.Execer driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 251/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 251/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 251/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 251/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 251/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 251/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 251/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 251/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 251/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 251/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 251/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 252/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 252/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 252/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 252/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 252/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 252/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 252/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 252/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 252/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 252/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 252/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 252/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 253/512: driver.Conn driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 253/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 253/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 253/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 253/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 253/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 253/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 253/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 253/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 253/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 253/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 253/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 254/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 254/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 254/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 254/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 254/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 254/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 254/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 254/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 254/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 254/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 254/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 254/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 255/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 255/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 255/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 255/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 255/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 255/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 255/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 255/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 255/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 255/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 255/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 255/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 256/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 256/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 256/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 256/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 256/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 256/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 256/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 256/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 256/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != false {
			t.Errorf("combination 256/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 256/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 256/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 257/512: driver.Conn driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 257/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 257/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 257/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 257/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 257/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 257/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 257/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 257/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 257/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 257/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 257/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 258/512: driver.Conn driver.ConnBeginTx driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 258/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 258/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 258/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 258/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 258/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 258/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 258/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 258/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 258/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 258/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 258/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 259/512: driver.Conn driver.Execer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 259/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 259/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 259/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 259/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 259/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 259/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 259/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 259/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 259/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 259/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 259/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 260/512: driver.Conn driver.ConnBeginTx driver.Execer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 260/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 260/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 260/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 260/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 260/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 260/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 260/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 260/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 260/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 260/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 260/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 261/512: driver.Conn driver.ExecerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 261/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 261/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 261/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 261/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 261/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 261/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 261/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 261/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 261/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 261/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 261/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 262/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 262/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 262/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 262/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 262/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 262/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 262/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 262/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 262/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 262/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 262/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 262/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 263/512: driver.Conn driver.Execer driver.ExecerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 263/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 263/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 263/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 263/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 263/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 263/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 263/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 263/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 263/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 263/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 263/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 264/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 264/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 264/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 264/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 264/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 264/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 264/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 264/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 264/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 264/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 264/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 264/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 265/512: driver.Conn driver.NamedValueChecker driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 265/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 265/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 265/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 265/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 265/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 265/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 265/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 265/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 265/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 265/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 265/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 266/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 266/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 266/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 266/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 266/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 266/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 266/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 266/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 266/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 266/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 266/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 266/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 267/512: driver.Conn driver.Execer driver.NamedValueChecker driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 267/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 267/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 267/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 267/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 267/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 267/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 267/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 267/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 267/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 267/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 267/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 268/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 268/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 268/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 268/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 268/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 268/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 268/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 268/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 268/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 268/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 268/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 268/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 269/512: driver.Conn driver.ExecerContext driver.NamedValueChecker driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 269/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 269/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 269/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 269/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 269/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 269/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 269/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 269/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 269/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 269/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 269/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 270/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 270/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 270/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 270/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 270/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 270/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 270/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 270/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 270/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 270/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 270/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 270/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 271/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 271/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 271/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 271/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 271/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 271/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 271/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 271/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 271/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 271/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 271/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 271/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 272/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 272/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 272/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 272/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 272/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 272/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 272/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 272/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 272/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 272/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 272/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 272/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 273/512: driver.Conn driver.Pinger driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Pinger
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 273/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 273/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 273/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 273/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 273/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 273/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 273/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 273/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 273/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 273/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 273/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 274/512: driver.Conn driver.ConnBeginTx driver.Pinger driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 274/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 274/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 274/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 274/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 274/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 274/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 274/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 274/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 274/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 274/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 274/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 275/512: driver.Conn driver.Execer driver.Pinger driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.Pinger
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 275/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 275/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 275/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 275/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 275/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 275/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 275/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 275/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 275/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 275/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 275/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 276/512: driver.Conn driver.ConnBeginTx driver.Execer driver.Pinger driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 276/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 276/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 276/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 276/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 276/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 276/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 276/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 276/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 276/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 276/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 276/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 277/512: driver.Conn driver.ExecerContext driver.Pinger driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 277/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 277/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 277/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 277/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 277/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 277/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 277/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 277/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 277/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 277/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 277/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 278/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.Pinger driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 278/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 278/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 278/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 278/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 278/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 278/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 278/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 278/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 278/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 278/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 278/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 279/512: driver.Conn driver.Execer driver.ExecerContext driver.Pinger driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 279/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 279/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 279/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 279/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 279/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 279/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 279/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 279/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 279/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 279/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 279/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 280/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.Pinger driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 280/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 280/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 280/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 280/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 280/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 280/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 280/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 280/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 280/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 280/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 280/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 281/512: driver.Conn driver.NamedValueChecker driver.Pinger driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 281/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 281/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 281/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 281/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 281/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 281/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 281/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 281/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 281/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 281/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 281/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 282/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker driver.Pinger driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 282/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 282/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 282/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 282/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 282/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 282/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 282/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 282/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 282/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 282/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 282/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 283/512: driver.Conn driver.Execer driver.NamedValueChecker driver.Pinger driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 283/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 283/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 283/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 283/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 283/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 283/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 283/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 283/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 283/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 283/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 283/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 284/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker driver.Pinger driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 284/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 284/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 284/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 284/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 284/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 284/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 284/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 284/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 284/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 284/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 284/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 285/512: driver.Conn driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 285/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 285/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 285/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 285/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 285/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 285/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 285/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 285/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 285/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 285/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 285/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 286/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 286/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 286/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 286/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 286/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 286/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 286/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 286/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 286/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 286/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 286/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 286/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 287/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 287/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 287/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 287/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 287/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 287/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 287/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 287/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 287/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 287/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 287/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 287/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 288/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 288/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 288/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 288/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 288/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 288/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 288/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 288/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 288/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 288/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 288/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 288/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 289/512: driver.Conn driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 289/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 289/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 289/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 289/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 289/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 289/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 289/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 289/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 289/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 289/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 289/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 290/512: driver.Conn driver.ConnBeginTx driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 290/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 290/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 290/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 290/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 290/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 290/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 290/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 290/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 290/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 290/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 290/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 291/512: driver.Conn driver.Execer driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 291/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 291/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 291/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 291/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 291/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 291/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 291/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 291/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 291/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 291/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 291/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 292/512: driver.Conn driver.ConnBeginTx driver.Execer driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 292/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 292/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 292/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 292/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 292/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 292/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 292/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 292/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 292/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 292/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 292/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 293/512: driver.Conn driver.ExecerContext driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 293/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 293/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 293/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 293/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 293/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 293/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 293/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 293/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 293/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 293/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 293/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 294/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 294/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 294/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 294/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 294/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 294/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 294/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 294/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 294/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 294/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 294/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 294/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 295/512: driver.Conn driver.Execer driver.ExecerContext driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 295/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 295/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 295/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 295/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 295/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 295/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 295/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 295/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 295/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 295/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 295/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 296/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 296/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 296/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 296/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 296/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 296/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 296/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 296/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 296/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 296/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 296/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 296/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 297/512: driver.Conn driver.NamedValueChecker driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 297/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 297/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 297/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 297/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 297/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 297/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 297/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 297/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 297/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 297/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 297/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 298/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 298/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 298/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 298/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 298/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 298/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 298/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 298/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 298/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 298/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 298/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 298/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 299/512: driver.Conn driver.Execer driver.NamedValueChecker driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 299/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 299/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 299/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 299/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 299/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 299/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 299/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 299/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 299/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 299/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 299/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 300/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 300/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 300/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 300/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 300/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 300/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 300/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 300/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 300/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 300/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 300/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 300/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 301/512: driver.Conn driver.ExecerContext driver.NamedValueChecker driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 301/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 301/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 301/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 301/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 301/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 301/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 301/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 301/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 301/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 301/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 301/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 302/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 302/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 302/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 302/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 302/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 302/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 302/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 302/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 302/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 302/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 302/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 302/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 303/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 303/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 303/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 303/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 303/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 303/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 303/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 303/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 303/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 303/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 303/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 303/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 304/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 304/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 304/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 304/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 304/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 304/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 304/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 304/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 304/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 304/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 304/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 304/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 305/512: driver.Conn driver.Pinger driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 305/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 305/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 305/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 305/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 305/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 305/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 305/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 305/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 305/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 305/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 305/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 306/512: driver.Conn driver.ConnBeginTx driver.Pinger driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 306/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 306/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 306/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 306/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 306/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 306/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 306/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 306/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 306/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 306/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 306/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 307/512: driver.Conn driver.Execer driver.Pinger driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 307/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 307/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 307/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 307/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 307/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 307/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 307/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 307/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 307/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 307/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 307/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 308/512: driver.Conn driver.ConnBeginTx driver.Execer driver.Pinger driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 308/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 308/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 308/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 308/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 308/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 308/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 308/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 308/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 308/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 308/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 308/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 309/512: driver.Conn driver.ExecerContext driver.Pinger driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 309/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 309/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 309/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 309/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 309/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 309/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 309/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 309/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 309/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 309/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 309/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 310/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.Pinger driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 310/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 310/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 310/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 310/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 310/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 310/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 310/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 310/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 310/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 310/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 310/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 311/512: driver.Conn driver.Execer driver.ExecerContext driver.Pinger driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 311/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 311/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 311/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 311/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 311/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 311/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 311/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 311/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 311/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 311/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 311/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 312/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.Pinger driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 312/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 312/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 312/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 312/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 312/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 312/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 312/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 312/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 312/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 312/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 312/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 313/512: driver.Conn driver.NamedValueChecker driver.Pinger driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 313/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 313/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 313/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 313/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 313/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 313/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 313/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 313/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 313/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 313/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 313/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 314/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker driver.Pinger driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 314/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 314/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 314/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 314/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 314/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 314/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 314/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 314/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 314/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 314/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 314/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 315/512: driver.Conn driver.Execer driver.NamedValueChecker driver.Pinger driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 315/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 315/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 315/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 315/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 315/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 315/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 315/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 315/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 315/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 315/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 315/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 316/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker driver.Pinger driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 316/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 316/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 316/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 316/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 316/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 316/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 316/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 316/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 316/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 316/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 316/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 317/512: driver.Conn driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 317/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 317/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 317/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 317/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 317/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 317/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 317/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 317/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 317/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 317/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 317/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 318/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 318/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 318/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 318/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 318/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 318/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 318/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 318/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 318/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 318/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 318/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 318/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 319/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 319/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 319/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 319/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 319/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 319/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 319/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 319/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 319/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 319/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 319/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 319/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 320/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 320/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 320/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 320/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 320/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 320/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 320/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 320/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 320/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 320/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 320/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 320/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 321/512: driver.Conn driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 321/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 321/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 321/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 321/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 321/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 321/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 321/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 321/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 321/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 321/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 321/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 322/512: driver.Conn driver.ConnBeginTx driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 322/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 322/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 322/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 322/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 322/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 322/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 322/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 322/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 322/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 322/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 322/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 323/512: driver.Conn driver.Execer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 323/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 323/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 323/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 323/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 323/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 323/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 323/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 323/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 323/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 323/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 323/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 324/512: driver.Conn driver.ConnBeginTx driver.Execer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 324/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 324/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 324/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 324/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 324/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 324/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 324/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 324/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 324/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 324/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 324/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 325/512: driver.Conn driver.ExecerContext driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 325/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 325/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 325/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 325/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 325/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 325/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 325/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 325/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 325/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 325/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 325/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 326/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 326/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 326/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 326/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 326/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 326/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 326/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 326/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 326/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 326/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 326/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 326/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 327/512: driver.Conn driver.Execer driver.ExecerContext driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 327/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 327/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 327/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 327/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 327/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 327/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 327/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 327/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 327/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 327/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 327/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 328/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 328/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 328/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 328/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 328/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 328/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 328/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 328/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 328/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 328/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 328/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 328/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 329/512: driver.Conn driver.NamedValueChecker driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 329/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 329/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 329/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 329/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 329/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 329/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 329/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 329/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 329/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 329/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 329/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 330/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 330/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 330/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 330/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 330/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 330/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 330/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 330/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 330/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 330/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 330/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 330/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 331/512: driver.Conn driver.Execer driver.NamedValueChecker driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 331/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 331/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 331/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 331/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 331/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 331/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 331/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 331/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 331/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 331/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 331/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 332/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 332/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 332/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 332/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 332/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 332/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 332/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 332/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 332/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 332/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 332/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 332/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 333/512: driver.Conn driver.ExecerContext driver.NamedValueChecker driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 333/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 333/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 333/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 333/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 333/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 333/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 333/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 333/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 333/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 333/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 333/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 334/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 334/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 334/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 334/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 334/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 334/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 334/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 334/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 334/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 334/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 334/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 334/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 335/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 335/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 335/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 335/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 335/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 335/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 335/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 335/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 335/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 335/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 335/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 335/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 336/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 336/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 336/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 336/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 336/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 336/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 336/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 336/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 336/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 336/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 336/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 336/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 337/512: driver.Conn driver.Pinger driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 337/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 337/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 337/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 337/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 337/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 337/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 337/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 337/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 337/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 337/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 337/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 338/512: driver.Conn driver.ConnBeginTx driver.Pinger driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 338/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 338/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 338/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 338/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 338/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 338/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 338/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 338/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 338/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 338/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 338/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 339/512: driver.Conn driver.Execer driver.Pinger driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 339/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 339/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 339/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 339/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 339/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 339/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 339/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 339/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 339/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 339/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 339/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 340/512: driver.Conn driver.ConnBeginTx driver.Execer driver.Pinger driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 340/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 340/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 340/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 340/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 340/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 340/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 340/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 340/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 340/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 340/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 340/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 341/512: driver.Conn driver.ExecerContext driver.Pinger driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 341/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 341/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 341/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 341/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 341/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 341/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 341/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 341/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 341/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 341/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 341/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 342/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.Pinger driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 342/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 342/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 342/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 342/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 342/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 342/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 342/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 342/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 342/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 342/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 342/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 343/512: driver.Conn driver.Execer driver.ExecerContext driver.Pinger driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 343/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 343/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 343/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 343/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 343/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 343/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 343/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 343/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 343/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 343/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 343/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 344/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.Pinger driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 344/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 344/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 344/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 344/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 344/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 344/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 344/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 344/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 344/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 344/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 344/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 345/512: driver.Conn driver.NamedValueChecker driver.Pinger driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 345/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 345/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 345/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 345/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 345/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 345/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 345/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 345/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 345/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 345/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 345/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 346/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker driver.Pinger driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 346/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 346/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 346/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 346/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 346/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 346/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 346/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 346/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 346/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 346/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 346/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 347/512: driver.Conn driver.Execer driver.NamedValueChecker driver.Pinger driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 347/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 347/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 347/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 347/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 347/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 347/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 347/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 347/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 347/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 347/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 347/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 348/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker driver.Pinger driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 348/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 348/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 348/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 348/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 348/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 348/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 348/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 348/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 348/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 348/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 348/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 349/512: driver.Conn driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 349/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 349/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 349/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 349/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 349/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 349/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 349/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 349/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 349/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 349/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 349/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 350/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 350/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 350/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 350/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 350/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 350/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 350/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 350/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 350/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 350/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 350/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 350/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 351/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 351/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 351/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 351/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 351/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 351/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 351/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 351/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 351/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 351/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 351/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 351/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 352/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 352/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 352/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 352/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 352/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 352/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 352/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 352/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 352/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 352/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 352/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 352/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 353/512: driver.Conn driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 353/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 353/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 353/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 353/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 353/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 353/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 353/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 353/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 353/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 353/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 353/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 354/512: driver.Conn driver.ConnBeginTx driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 354/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 354/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 354/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 354/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 354/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 354/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 354/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 354/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 354/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 354/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 354/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 355/512: driver.Conn driver.Execer driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 355/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 355/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 355/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 355/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 355/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 355/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 355/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 355/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 355/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 355/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 355/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 356/512: driver.Conn driver.ConnBeginTx driver.Execer driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 356/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 356/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 356/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 356/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 356/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 356/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 356/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 356/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 356/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 356/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 356/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 357/512: driver.Conn driver.ExecerContext driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 357/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 357/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 357/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 357/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 357/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 357/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 357/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 357/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 357/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 357/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 357/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 358/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 358/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 358/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 358/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 358/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 358/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 358/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 358/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 358/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 358/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 358/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 358/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 359/512: driver.Conn driver.Execer driver.ExecerContext driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 359/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 359/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 359/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 359/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 359/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 359/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 359/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 359/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 359/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 359/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 359/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 360/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 360/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 360/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 360/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 360/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 360/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 360/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 360/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 360/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 360/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 360/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 360/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 361/512: driver.Conn driver.NamedValueChecker driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 361/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 361/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 361/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 361/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 361/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 361/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 361/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 361/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 361/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 361/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 361/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 362/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 362/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 362/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 362/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 362/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 362/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 362/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 362/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 362/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 362/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 362/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 362/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 363/512: driver.Conn driver.Execer driver.NamedValueChecker driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 363/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 363/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 363/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 363/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 363/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 363/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 363/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 363/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 363/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 363/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 363/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 364/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 364/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 364/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 364/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 364/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 364/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 364/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 364/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 364/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 364/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 364/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 364/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 365/512: driver.Conn driver.ExecerContext driver.NamedValueChecker driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 365/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 365/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 365/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 365/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 365/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 365/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 365/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 365/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 365/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 365/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 365/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 366/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 366/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 366/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 366/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 366/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 366/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 366/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 366/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 366/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 366/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 366/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 366/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 367/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 367/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 367/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 367/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 367/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 367/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 367/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 367/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 367/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 367/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 367/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 367/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 368/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 368/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 368/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 368/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 368/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 368/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 368/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 368/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 368/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 368/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 368/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 368/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 369/512: driver.Conn driver.Pinger driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 369/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 369/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 369/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 369/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 369/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 369/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 369/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 369/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 369/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 369/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 369/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 370/512: driver.Conn driver.ConnBeginTx driver.Pinger driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 370/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 370/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 370/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 370/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 370/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 370/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 370/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 370/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 370/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 370/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 370/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 371/512: driver.Conn driver.Execer driver.Pinger driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 371/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 371/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 371/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 371/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 371/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 371/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 371/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 371/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 371/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 371/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 371/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 372/512: driver.Conn driver.ConnBeginTx driver.Execer driver.Pinger driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 372/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 372/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 372/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 372/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 372/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 372/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 372/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 372/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 372/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 372/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 372/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 373/512: driver.Conn driver.ExecerContext driver.Pinger driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 373/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 373/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 373/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 373/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 373/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 373/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 373/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 373/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 373/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 373/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 373/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 374/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.Pinger driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 374/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 374/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 374/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 374/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 374/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 374/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 374/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 374/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 374/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 374/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 374/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 375/512: driver.Conn driver.Execer driver.ExecerContext driver.Pinger driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 375/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 375/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 375/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 375/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 375/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 375/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 375/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 375/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 375/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 375/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 375/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 376/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.Pinger driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 376/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 376/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 376/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 376/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 376/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 376/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 376/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 376/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 376/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 376/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 376/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 377/512: driver.Conn driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 377/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 377/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 377/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 377/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 377/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 377/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 377/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 377/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 377/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 377/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 377/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 378/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 378/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 378/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 378/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 378/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 378/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 378/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 378/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 378/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 378/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 378/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 378/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 379/512: driver.Conn driver.Execer driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 379/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 379/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 379/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 379/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 379/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 379/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 379/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 379/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 379/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 379/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 379/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 380/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 380/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 380/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 380/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 380/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 380/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 380/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 380/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 380/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 380/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 380/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 380/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 381/512: driver.Conn driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 381/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 381/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 381/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 381/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 381/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 381/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 381/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 381/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 381/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 381/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 381/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 382/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 382/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 382/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 382/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 382/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 382/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 382/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 382/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 382/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 382/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 382/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 382/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 383/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 383/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 383/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 383/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 383/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 383/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 383/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 383/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 383/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 383/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 383/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 383/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 384/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 384/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 384/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 384/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 384/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 384/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 384/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 384/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != false {
			t.Errorf("combination 384/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 384/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 384/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 384/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 385/512: driver.Conn driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 385/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 385/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 385/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 385/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 385/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 385/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 385/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 385/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 385/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 385/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 385/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 386/512: driver.Conn driver.ConnBeginTx driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 386/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 386/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 386/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 386/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 386/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 386/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 386/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 386/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 386/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 386/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 386/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 387/512: driver.Conn driver.Execer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 387/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 387/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 387/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 387/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 387/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 387/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 387/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 387/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 387/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 387/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 387/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 388/512: driver.Conn driver.ConnBeginTx driver.Execer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 388/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 388/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 388/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 388/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 388/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 388/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 388/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 388/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 388/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 388/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 388/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 389/512: driver.Conn driver.ExecerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 389/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 389/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 389/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 389/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 389/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 389/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 389/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 389/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 389/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 389/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 389/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 390/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 390/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 390/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 390/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 390/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 390/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 390/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 390/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 390/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 390/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 390/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 390/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 391/512: driver.Conn driver.Execer driver.ExecerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 391/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 391/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 391/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 391/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 391/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 391/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 391/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 391/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 391/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 391/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 391/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 392/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 392/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 392/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 392/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 392/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 392/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 392/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 392/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 392/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 392/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 392/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 392/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 393/512: driver.Conn driver.NamedValueChecker driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 393/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 393/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 393/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 393/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 393/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 393/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 393/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 393/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 393/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 393/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 393/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 394/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 394/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 394/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 394/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 394/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 394/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 394/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 394/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 394/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 394/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 394/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 394/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 395/512: driver.Conn driver.Execer driver.NamedValueChecker driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 395/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 395/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 395/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 395/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 395/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 395/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 395/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 395/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 395/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 395/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 395/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 396/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 396/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 396/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 396/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 396/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 396/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 396/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 396/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 396/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 396/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 396/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 396/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 397/512: driver.Conn driver.ExecerContext driver.NamedValueChecker driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 397/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 397/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 397/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 397/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 397/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 397/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 397/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 397/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 397/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 397/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 397/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 398/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 398/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 398/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 398/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 398/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 398/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 398/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 398/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 398/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 398/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 398/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 398/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 399/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 399/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 399/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 399/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 399/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 399/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 399/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 399/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 399/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 399/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 399/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 399/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 400/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 400/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 400/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 400/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 400/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 400/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 400/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 400/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 400/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 400/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 400/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 400/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 401/512: driver.Conn driver.Pinger driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 401/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 401/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 401/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 401/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 401/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 401/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 401/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 401/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 401/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 401/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 401/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 402/512: driver.Conn driver.ConnBeginTx driver.Pinger driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 402/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 402/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 402/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 402/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 402/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 402/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 402/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 402/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 402/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 402/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 402/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 403/512: driver.Conn driver.Execer driver.Pinger driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 403/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 403/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 403/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 403/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 403/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 403/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 403/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 403/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 403/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 403/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 403/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 404/512: driver.Conn driver.ConnBeginTx driver.Execer driver.Pinger driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 404/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 404/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 404/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 404/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 404/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 404/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 404/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 404/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 404/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 404/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 404/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 405/512: driver.Conn driver.ExecerContext driver.Pinger driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 405/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 405/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 405/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 405/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 405/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 405/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 405/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 405/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 405/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 405/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 405/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 406/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.Pinger driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 406/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 406/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 406/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 406/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 406/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 406/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 406/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 406/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 406/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 406/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 406/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 407/512: driver.Conn driver.Execer driver.ExecerContext driver.Pinger driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 407/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 407/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 407/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 407/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 407/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 407/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 407/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 407/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 407/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 407/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 407/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 408/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.Pinger driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 408/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 408/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 408/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 408/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 408/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 408/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 408/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 408/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 408/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 408/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 408/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 409/512: driver.Conn driver.NamedValueChecker driver.Pinger driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 409/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 409/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 409/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 409/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 409/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 409/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 409/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 409/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 409/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 409/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 409/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 410/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker driver.Pinger driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 410/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 410/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 410/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 410/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 410/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 410/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 410/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 410/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 410/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 410/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 410/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 411/512: driver.Conn driver.Execer driver.NamedValueChecker driver.Pinger driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 411/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 411/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 411/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 411/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 411/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 411/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 411/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 411/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 411/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 411/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 411/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 412/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker driver.Pinger driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 412/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 412/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 412/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 412/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 412/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 412/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 412/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 412/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 412/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 412/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 412/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 413/512: driver.Conn driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 413/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 413/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 413/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 413/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 413/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 413/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 413/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 413/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 413/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 413/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 413/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 414/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 414/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 414/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 414/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 414/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 414/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 414/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 414/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 414/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 414/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 414/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 414/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 415/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 415/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 415/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 415/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 415/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 415/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 415/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 415/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 415/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 415/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 415/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 415/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 416/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 416/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 416/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 416/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 416/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 416/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 416/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 416/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 416/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 416/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 416/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 416/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 417/512: driver.Conn driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 417/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 417/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 417/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 417/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 417/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 417/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 417/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 417/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 417/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 417/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 417/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 418/512: driver.Conn driver.ConnBeginTx driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 418/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 418/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 418/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 418/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 418/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 418/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 418/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 418/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 418/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 418/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 418/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 419/512: driver.Conn driver.Execer driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 419/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 419/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 419/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 419/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 419/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 419/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 419/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 419/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 419/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 419/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 419/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 420/512: driver.Conn driver.ConnBeginTx driver.Execer driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 420/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 420/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 420/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 420/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 420/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 420/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 420/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 420/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 420/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 420/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 420/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 421/512: driver.Conn driver.ExecerContext driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 421/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 421/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 421/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 421/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 421/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 421/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 421/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 421/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 421/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 421/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 421/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 422/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 422/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 422/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 422/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 422/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 422/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 422/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 422/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 422/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 422/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 422/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 422/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 423/512: driver.Conn driver.Execer driver.ExecerContext driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 423/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 423/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 423/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 423/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 423/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 423/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 423/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 423/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 423/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 423/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 423/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 424/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 424/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 424/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 424/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 424/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 424/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 424/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 424/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 424/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 424/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 424/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 424/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 425/512: driver.Conn driver.NamedValueChecker driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 425/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 425/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 425/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 425/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 425/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 425/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 425/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 425/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 425/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 425/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 425/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 426/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 426/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 426/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 426/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 426/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 426/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 426/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 426/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 426/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 426/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 426/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 426/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 427/512: driver.Conn driver.Execer driver.NamedValueChecker driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 427/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 427/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 427/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 427/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 427/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 427/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 427/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 427/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 427/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 427/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 427/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 428/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 428/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 428/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 428/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 428/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 428/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 428/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 428/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 428/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 428/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 428/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 428/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 429/512: driver.Conn driver.ExecerContext driver.NamedValueChecker driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 429/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 429/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 429/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 429/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 429/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 429/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 429/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 429/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 429/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 429/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 429/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 430/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 430/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 430/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 430/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 430/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 430/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 430/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 430/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 430/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 430/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 430/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 430/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 431/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 431/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 431/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 431/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 431/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 431/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 431/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 431/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 431/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 431/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 431/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 431/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 432/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 432/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 432/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 432/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 432/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 432/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 432/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 432/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 432/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 432/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 432/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 432/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 433/512: driver.Conn driver.Pinger driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 433/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 433/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 433/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 433/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 433/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 433/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 433/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 433/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 433/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 433/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 433/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 434/512: driver.Conn driver.ConnBeginTx driver.Pinger driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 434/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 434/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 434/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 434/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 434/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 434/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 434/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 434/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 434/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 434/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 434/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 435/512: driver.Conn driver.Execer driver.Pinger driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 435/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 435/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 435/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 435/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 435/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 435/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 435/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 435/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 435/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 435/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 435/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 436/512: driver.Conn driver.ConnBeginTx driver.Execer driver.Pinger driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 436/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 436/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 436/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 436/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 436/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 436/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 436/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 436/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 436/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 436/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 436/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 437/512: driver.Conn driver.ExecerContext driver.Pinger driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 437/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 437/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 437/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 437/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 437/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 437/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 437/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 437/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 437/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 437/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 437/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 438/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.Pinger driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 438/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 438/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 438/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 438/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 438/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 438/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 438/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 438/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 438/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 438/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 438/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 439/512: driver.Conn driver.Execer driver.ExecerContext driver.Pinger driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 439/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 439/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 439/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 439/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 439/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 439/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 439/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 439/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 439/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 439/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 439/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 440/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.Pinger driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 440/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 440/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 440/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 440/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 440/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 440/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 440/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 440/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 440/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 440/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 440/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 441/512: driver.Conn driver.NamedValueChecker driver.Pinger driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 441/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 441/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 441/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 441/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 441/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 441/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 441/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 441/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 441/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 441/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 441/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 442/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker driver.Pinger driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 442/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 442/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 442/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 442/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 442/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 442/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 442/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 442/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 442/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 442/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 442/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 443/512: driver.Conn driver.Execer driver.NamedValueChecker driver.Pinger driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 443/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 443/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 443/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 443/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 443/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 443/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 443/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 443/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 443/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 443/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 443/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 444/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker driver.Pinger driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 444/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 444/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 444/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 444/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 444/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 444/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 444/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 444/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 444/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 444/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 444/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 445/512: driver.Conn driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 445/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 445/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 445/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 445/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 445/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 445/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 445/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 445/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 445/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 445/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 445/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 446/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 446/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 446/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 446/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 446/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 446/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 446/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 446/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 446/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 446/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 446/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 446/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 447/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 447/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 447/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 447/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 447/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 447/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 447/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 447/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 447/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 447/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 447/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 447/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 448/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 448/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 448/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 448/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 448/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 448/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 448/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != false {
			t.Errorf("combination 448/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 448/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 448/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 448/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 448/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 449/512: driver.Conn driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 449/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 449/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 449/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 449/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 449/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 449/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 449/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 449/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 449/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 449/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 449/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 450/512: driver.Conn driver.ConnBeginTx driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 450/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 450/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 450/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 450/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 450/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 450/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 450/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 450/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 450/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 450/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 450/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 451/512: driver.Conn driver.Execer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 451/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 451/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 451/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 451/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 451/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 451/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 451/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 451/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 451/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 451/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 451/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 452/512: driver.Conn driver.ConnBeginTx driver.Execer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 452/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 452/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 452/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 452/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 452/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 452/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 452/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 452/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 452/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 452/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 452/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 453/512: driver.Conn driver.ExecerContext driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 453/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 453/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 453/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 453/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 453/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 453/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 453/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 453/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 453/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 453/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 453/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 454/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 454/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 454/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 454/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 454/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 454/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 454/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 454/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 454/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 454/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 454/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 454/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 455/512: driver.Conn driver.Execer driver.ExecerContext driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 455/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 455/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 455/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 455/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 455/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 455/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 455/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 455/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 455/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 455/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 455/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 456/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 456/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 456/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 456/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 456/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 456/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 456/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 456/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 456/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 456/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 456/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 456/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 457/512: driver.Conn driver.NamedValueChecker driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 457/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 457/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 457/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 457/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 457/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 457/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 457/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 457/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 457/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 457/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 457/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 458/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 458/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 458/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 458/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 458/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 458/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 458/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 458/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 458/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 458/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 458/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 458/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 459/512: driver.Conn driver.Execer driver.NamedValueChecker driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 459/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 459/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 459/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 459/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 459/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 459/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 459/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 459/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 459/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 459/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 459/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 460/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 460/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 460/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 460/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 460/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 460/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 460/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 460/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 460/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 460/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 460/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 460/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 461/512: driver.Conn driver.ExecerContext driver.NamedValueChecker driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 461/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 461/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 461/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 461/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 461/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 461/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 461/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 461/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 461/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 461/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 461/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 462/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 462/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 462/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 462/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 462/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 462/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 462/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 462/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 462/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 462/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 462/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 462/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 463/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 463/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 463/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 463/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 463/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 463/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 463/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 463/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 463/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 463/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 463/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 463/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 464/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 464/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 464/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 464/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 464/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 464/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 464/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 464/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 464/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 464/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 464/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 464/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 465/512: driver.Conn driver.Pinger driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 465/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 465/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 465/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 465/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 465/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 465/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 465/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 465/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 465/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 465/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 465/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 466/512: driver.Conn driver.ConnBeginTx driver.Pinger driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 466/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 466/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 466/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 466/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 466/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 466/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 466/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 466/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 466/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 466/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 466/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 467/512: driver.Conn driver.Execer driver.Pinger driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 467/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 467/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 467/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 467/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 467/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 467/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 467/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 467/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 467/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 467/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 467/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 468/512: driver.Conn driver.ConnBeginTx driver.Execer driver.Pinger driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 468/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 468/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 468/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 468/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 468/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 468/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 468/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 468/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 468/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 468/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 468/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 469/512: driver.Conn driver.ExecerContext driver.Pinger driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 469/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 469/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 469/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 469/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 469/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 469/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 469/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 469/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 469/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 469/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 469/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 470/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.Pinger driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 470/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 470/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 470/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 470/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 470/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 470/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 470/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 470/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 470/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 470/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 470/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 471/512: driver.Conn driver.Execer driver.ExecerContext driver.Pinger driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 471/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 471/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 471/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 471/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 471/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 471/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 471/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 471/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 471/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 471/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 471/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 472/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.Pinger driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 472/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 472/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 472/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 472/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 472/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 472/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 472/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 472/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 472/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 472/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 472/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 473/512: driver.Conn driver.NamedValueChecker driver.Pinger driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 473/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 473/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 473/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 473/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 473/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 473/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 473/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 473/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 473/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 473/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 473/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 474/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker driver.Pinger driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 474/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 474/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 474/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 474/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 474/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 474/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 474/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 474/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 474/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 474/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 474/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 475/512: driver.Conn driver.Execer driver.NamedValueChecker driver.Pinger driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 475/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 475/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 475/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 475/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 475/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 475/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 475/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 475/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 475/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 475/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 475/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 476/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker driver.Pinger driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 476/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 476/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 476/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 476/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 476/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 476/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 476/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 476/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 476/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 476/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 476/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 477/512: driver.Conn driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 477/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 477/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 477/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 477/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 477/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 477/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 477/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 477/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 477/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 477/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 477/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 478/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 478/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 478/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 478/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 478/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 478/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 478/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 478/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 478/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 478/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 478/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 478/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 479/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 479/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 479/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 479/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 479/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 479/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 479/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 479/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 479/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 479/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 479/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 479/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 480/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 480/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 480/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 480/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 480/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 480/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != false {
			t.Errorf("combination 480/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 480/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 480/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 480/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 480/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 480/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 481/512: driver.Conn driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 481/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 481/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 481/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 481/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 481/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 481/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 481/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 481/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 481/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 481/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 481/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 482/512: driver.Conn driver.ConnBeginTx driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 482/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 482/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 482/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 482/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 482/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 482/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 482/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 482/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 482/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 482/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 482/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 483/512: driver.Conn driver.Execer driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 483/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 483/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 483/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 483/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 483/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 483/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 483/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 483/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 483/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 483/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 483/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 484/512: driver.Conn driver.ConnBeginTx driver.Execer driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 484/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 484/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 484/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 484/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 484/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 484/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 484/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 484/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 484/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 484/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 484/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 485/512: driver.Conn driver.ExecerContext driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 485/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 485/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 485/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 485/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 485/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 485/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 485/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 485/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 485/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 485/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 485/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 486/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 486/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 486/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 486/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 486/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 486/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 486/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 486/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 486/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 486/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 486/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 486/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 487/512: driver.Conn driver.Execer driver.ExecerContext driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 487/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 487/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 487/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 487/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 487/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 487/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 487/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 487/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 487/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 487/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 487/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 488/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 488/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 488/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 488/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 488/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 488/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 488/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 488/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 488/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 488/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 488/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 488/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 489/512: driver.Conn driver.NamedValueChecker driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 489/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 489/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 489/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 489/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 489/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 489/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 489/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 489/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 489/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 489/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 489/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 490/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 490/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 490/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 490/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 490/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 490/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 490/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 490/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 490/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 490/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 490/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 490/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 491/512: driver.Conn driver.Execer driver.NamedValueChecker driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 491/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 491/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 491/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 491/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 491/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 491/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 491/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 491/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 491/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 491/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 491/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 492/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 492/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 492/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 492/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 492/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 492/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 492/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 492/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 492/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 492/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 492/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 492/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 493/512: driver.Conn driver.ExecerContext driver.NamedValueChecker driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 493/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 493/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 493/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 493/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 493/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 493/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 493/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 493/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 493/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 493/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 493/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 494/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 494/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 494/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 494/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 494/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 494/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 494/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 494/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 494/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 494/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 494/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 494/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 495/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 495/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 495/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 495/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 495/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 495/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 495/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 495/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 495/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 495/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 495/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 495/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 496/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 496/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 496/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 496/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 496/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != false {
			t.Errorf("combination 496/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 496/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 496/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 496/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 496/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 496/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 496/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 497/512: driver.Conn driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 497/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 497/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 497/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 497/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 497/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 497/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 497/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 497/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 497/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 497/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 497/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 498/512: driver.Conn driver.ConnBeginTx driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 498/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 498/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 498/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 498/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 498/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 498/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 498/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 498/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 498/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 498/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 498/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 499/512: driver.Conn driver.Execer driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 499/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 499/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 499/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 499/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 499/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 499/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 499/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 499/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 499/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 499/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 499/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 500/512: driver.Conn driver.ConnBeginTx driver.Execer driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 500/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 500/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 500/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 500/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 500/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 500/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 500/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 500/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 500/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 500/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 500/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 501/512: driver.Conn driver.ExecerContext driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 501/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 501/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 501/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 501/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 501/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 501/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 501/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 501/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 501/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 501/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 501/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 502/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 502/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 502/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 502/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 502/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 502/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 502/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 502/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 502/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 502/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 502/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 502/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 503/512: driver.Conn driver.Execer driver.ExecerContext driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 503/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 503/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 503/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 503/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 503/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 503/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 503/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 503/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 503/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 503/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 503/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 504/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 504/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 504/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 504/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != false {
			t.Errorf("combination 504/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 504/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 504/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 504/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 504/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 504/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 504/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 504/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 505/512: driver.Conn driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 505/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 505/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 505/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 505/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 505/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 505/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 505/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 505/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 505/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 505/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 505/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 506/512: driver.Conn driver.ConnBeginTx driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 506/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 506/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 506/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 506/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 506/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 506/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 506/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 506/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 506/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 506/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 506/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 507/512: driver.Conn driver.Execer driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 507/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 507/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 507/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 507/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 507/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 507/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 507/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 507/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 507/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 507/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 507/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 508/512: driver.Conn driver.ConnBeginTx driver.Execer driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 508/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 508/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != false {
			t.Errorf("combination 508/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 508/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 508/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 508/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 508/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 508/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 508/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 508/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 508/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 509/512: driver.Conn driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 509/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 509/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 509/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 509/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 509/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 509/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 509/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 509/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 509/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 509/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 509/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 510/512: driver.Conn driver.ConnBeginTx driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 510/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != false {
			t.Errorf("combination 510/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 510/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 510/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 510/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 510/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 510/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 510/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 510/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 510/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 510/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 511/512: driver.Conn driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != false {
			t.Errorf("combination 511/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 511/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 511/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 511/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 511/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 511/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 511/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 511/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 511/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 511/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 511/512: ConnUnwrapper interface not implemented")
		}
	}
	{
		t.Log("combination 512/512: driver.Conn driver.ConnBeginTx driver.Execer driver.ExecerContext driver.NamedValueChecker driver.Pinger driver.Queryer driver.QueryerContext driver.SessionResetter driver.Validator")
		wrapped := struct {
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{}
		w := WrapConn(wrapped, ConnInterceptor{})

		if _, ok := w.(driver.ConnBeginTx); ok != true {
			t.Errorf("combination 512/512: unexpected interface driver.ConnBeginTx")
		}
		if _, ok := w.(driver.Execer); ok != true {
			t.Errorf("combination 512/512: unexpected interface driver.Execer")
		}
		if _, ok := w.(driver.ExecerContext); ok != true {
			t.Errorf("combination 512/512: unexpected interface driver.ExecerContext")
		}
		if _, ok := w.(driver.NamedValueChecker); ok != true {
			t.Errorf("combination 512/512: unexpected interface driver.NamedValueChecker")
		}
		if _, ok := w.(driver.Pinger); ok != true {
			t.Errorf("combination 512/512: unexpected interface driver.Pinger")
		}
		if _, ok := w.(driver.Queryer); ok != true {
			t.Errorf("combination 512/512: unexpected interface driver.Queryer")
		}
		if _, ok := w.(driver.QueryerContext); ok != true {
			t.Errorf("combination 512/512: unexpected interface driver.QueryerContext")
		}
		if _, ok := w.(driver.SessionResetter); ok != true {
			t.Errorf("combination 512/512: unexpected interface driver.SessionResetter")
		}
		if _, ok := w.(driver.Validator); ok != true {
			t.Errorf("combination 512/512: unexpected interface driver.Validator")
		}

		if w, ok := w.(ConnUnwrapper); ok {
			if w.UnwrapConn() != wrapped {
				t.Errorf("combination 512/512: UnwrapConn() failed")
			}
		} else {
			t.Errorf("combination 512/512: ConnUnwrapper interface not implemented")
		}
	}
}
