package databasesqldriver

import (
	context "context"
	driver "database/sql/driver"
)

type ConnUnwrapper interface {
	UnwrapConn() driver.Conn
}

func UnwrapConn(w driver.Conn) driver.Conn {
	if ww, ok := w.(ConnUnwrapper); ok {
		return UnwrapConn(ww.UnwrapConn())
	} else {
		return w
	}
}

type Conn_Begin func() (driver.Tx, error)
type Conn_Close func() error
type Conn_Prepare func(string) (driver.Stmt, error)
type Conn_BeginTx func(context.Context, driver.TxOptions) (driver.Tx, error)
type Conn_Exec func(string, []driver.Value) (driver.Result, error)
type Conn_ExecContext func(context.Context, string, []driver.NamedValue) (driver.Result, error)
type Conn_CheckNamedValue func(*driver.NamedValue) error
type Conn_Ping func(context.Context) error
type Conn_Query func(string, []driver.Value) (driver.Rows, error)
type Conn_QueryContext func(context.Context, string, []driver.NamedValue) (driver.Rows, error)
type Conn_ResetSession func(context.Context) error
type Conn_IsValid func() bool

type ConnInterceptor struct {
	Begin           func(Conn_Begin) Conn_Begin
	Close           func(Conn_Close) Conn_Close
	Prepare         func(Conn_Prepare) Conn_Prepare
	BeginTx         func(Conn_BeginTx) Conn_BeginTx
	Exec            func(Conn_Exec) Conn_Exec
	ExecContext     func(Conn_ExecContext) Conn_ExecContext
	CheckNamedValue func(Conn_CheckNamedValue) Conn_CheckNamedValue
	Ping            func(Conn_Ping) Conn_Ping
	Query           func(Conn_Query) Conn_Query
	QueryContext    func(Conn_QueryContext) Conn_QueryContext
	ResetSession    func(Conn_ResetSession) Conn_ResetSession
	IsValid         func(Conn_IsValid) Conn_IsValid
}

type ConnWrapper struct {
	wrapped     driver.Conn
	interceptor ConnInterceptor
}

func (w *ConnWrapper) UnwrapConn() driver.Conn {
	return w.wrapped
}

func (w *ConnWrapper) Begin() (driver.Tx, error) {
	f := w.wrapped.(driver.Conn).Begin
	if w.interceptor.Begin != nil {
		f = w.interceptor.Begin(f)
	}
	return f()
}

func (w *ConnWrapper) Close() error {
	f := w.wrapped.(driver.Conn).Close
	if w.interceptor.Close != nil {
		f = w.interceptor.Close(f)
	}
	return f()
}

func (w *ConnWrapper) Prepare(a0 string) (driver.Stmt, error) {
	f := w.wrapped.(driver.Conn).Prepare
	if w.interceptor.Prepare != nil {
		f = w.interceptor.Prepare(f)
	}
	return f(a0)
}

func (w *ConnWrapper) BeginTx(a0 context.Context, a1 driver.TxOptions) (driver.Tx, error) {
	f := w.wrapped.(driver.ConnBeginTx).BeginTx
	if w.interceptor.BeginTx != nil {
		f = w.interceptor.BeginTx(f)
	}
	return f(a0, a1)
}

func (w *ConnWrapper) Exec(a0 string, a1 []driver.Value) (driver.Result, error) {
	f := w.wrapped.(driver.Execer).Exec
	if w.interceptor.Exec != nil {
		f = w.interceptor.Exec(f)
	}
	return f(a0, a1)
}

func (w *ConnWrapper) ExecContext(a0 context.Context, a1 string, a2 []driver.NamedValue) (driver.Result, error) {
	f := w.wrapped.(driver.ExecerContext).ExecContext
	if w.interceptor.ExecContext != nil {
		f = w.interceptor.ExecContext(f)
	}
	return f(a0, a1, a2)
}

func (w *ConnWrapper) CheckNamedValue(a0 *driver.NamedValue) error {
	f := w.wrapped.(driver.NamedValueChecker).CheckNamedValue
	if w.interceptor.CheckNamedValue != nil {
		f = w.interceptor.CheckNamedValue(f)
	}
	return f(a0)
}

func (w *ConnWrapper) Ping(a0 context.Context) error {
	f := w.wrapped.(driver.Pinger).Ping
	if w.interceptor.Ping != nil {
		f = w.interceptor.Ping(f)
	}
	return f(a0)
}

func (w *ConnWrapper) Query(a0 string, a1 []driver.Value) (driver.Rows, error) {
	f := w.wrapped.(driver.Queryer).Query
	if w.interceptor.Query != nil {
		f = w.interceptor.Query(f)
	}
	return f(a0, a1)
}

func (w *ConnWrapper) QueryContext(a0 context.Context, a1 string, a2 []driver.NamedValue) (driver.Rows, error) {
	f := w.wrapped.(driver.QueryerContext).QueryContext
	if w.interceptor.QueryContext != nil {
		f = w.interceptor.QueryContext(f)
	}
	return f(a0, a1, a2)
}

func (w *ConnWrapper) ResetSession(a0 context.Context) error {
	f := w.wrapped.(driver.SessionResetter).ResetSession
	if w.interceptor.ResetSession != nil {
		f = w.interceptor.ResetSession(f)
	}
	return f(a0)
}

func (w *ConnWrapper) IsValid() bool {
	f := w.wrapped.(driver.Validator).IsValid
	if w.interceptor.IsValid != nil {
		f = w.interceptor.IsValid(f)
	}
	return f()
}

func WrapConn(wrapped driver.Conn, interceptor ConnInterceptor) driver.Conn {
	w := &ConnWrapper{wrapped: wrapped, interceptor: interceptor}
	_, ok0 := wrapped.(driver.ConnBeginTx)
	_, ok1 := wrapped.(driver.Execer)
	_, ok2 := wrapped.(driver.ExecerContext)
	_, ok3 := wrapped.(driver.NamedValueChecker)
	_, ok4 := wrapped.(driver.Pinger)
	_, ok5 := wrapped.(driver.Queryer)
	_, ok6 := wrapped.(driver.QueryerContext)
	_, ok7 := wrapped.(driver.SessionResetter)
	_, ok8 := wrapped.(driver.Validator)
	switch {
	// combination 1/512
	case !ok0 && !ok1 && !ok2 && !ok3 && !ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
		}{w, w}
	// combination 2/512
	case ok0 && !ok1 && !ok2 && !ok3 && !ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
		}{w, w, w}
	// combination 3/512
	case !ok0 && ok1 && !ok2 && !ok3 && !ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
		}{w, w, w}
	// combination 4/512
	case ok0 && ok1 && !ok2 && !ok3 && !ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
		}{w, w, w, w}
	// combination 5/512
	case !ok0 && !ok1 && ok2 && !ok3 && !ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
		}{w, w, w}
	// combination 6/512
	case ok0 && !ok1 && ok2 && !ok3 && !ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
		}{w, w, w, w}
	// combination 7/512
	case !ok0 && ok1 && ok2 && !ok3 && !ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
		}{w, w, w, w}
	// combination 8/512
	case ok0 && ok1 && ok2 && !ok3 && !ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
		}{w, w, w, w, w}
	// combination 9/512
	case !ok0 && !ok1 && !ok2 && ok3 && !ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
		}{w, w, w}
	// combination 10/512
	case ok0 && !ok1 && !ok2 && ok3 && !ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
		}{w, w, w, w}
	// combination 11/512
	case !ok0 && ok1 && !ok2 && ok3 && !ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
		}{w, w, w, w}
	// combination 12/512
	case ok0 && ok1 && !ok2 && ok3 && !ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
		}{w, w, w, w, w}
	// combination 13/512
	case !ok0 && !ok1 && ok2 && ok3 && !ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
		}{w, w, w, w}
	// combination 14/512
	case ok0 && !ok1 && ok2 && ok3 && !ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
		}{w, w, w, w, w}
	// combination 15/512
	case !ok0 && ok1 && ok2 && ok3 && !ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
		}{w, w, w, w, w}
	// combination 16/512
	case ok0 && ok1 && ok2 && ok3 && !ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
		}{w, w, w, w, w, w}
	// combination 17/512
	case !ok0 && !ok1 && !ok2 && !ok3 && ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Pinger
		}{w, w, w}
	// combination 18/512
	case ok0 && !ok1 && !ok2 && !ok3 && ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
		}{w, w, w, w}
	// combination 19/512
	case !ok0 && ok1 && !ok2 && !ok3 && ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.Pinger
		}{w, w, w, w}
	// combination 20/512
	case ok0 && ok1 && !ok2 && !ok3 && ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
		}{w, w, w, w, w}
	// combination 21/512
	case !ok0 && !ok1 && ok2 && !ok3 && ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.Pinger
		}{w, w, w, w}
	// combination 22/512
	case ok0 && !ok1 && ok2 && !ok3 && ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
		}{w, w, w, w, w}
	// combination 23/512
	case !ok0 && ok1 && ok2 && !ok3 && ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
		}{w, w, w, w, w}
	// combination 24/512
	case ok0 && ok1 && ok2 && !ok3 && ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
		}{w, w, w, w, w, w}
	// combination 25/512
	case !ok0 && !ok1 && !ok2 && ok3 && ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
		}{w, w, w, w}
	// combination 26/512
	case ok0 && !ok1 && !ok2 && ok3 && ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
		}{w, w, w, w, w}
	// combination 27/512
	case !ok0 && ok1 && !ok2 && ok3 && ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
		}{w, w, w, w, w}
	// combination 28/512
	case ok0 && ok1 && !ok2 && ok3 && ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
		}{w, w, w, w, w, w}
	// combination 29/512
	case !ok0 && !ok1 && ok2 && ok3 && ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
		}{w, w, w, w, w}
	// combination 30/512
	case ok0 && !ok1 && ok2 && ok3 && ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
		}{w, w, w, w, w, w}
	// combination 31/512
	case !ok0 && ok1 && ok2 && ok3 && ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
		}{w, w, w, w, w, w}
	// combination 32/512
	case ok0 && ok1 && ok2 && ok3 && ok4 && !ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
		}{w, w, w, w, w, w, w}
	// combination 33/512
	case !ok0 && !ok1 && !ok2 && !ok3 && !ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Queryer
		}{w, w, w}
	// combination 34/512
	case ok0 && !ok1 && !ok2 && !ok3 && !ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Queryer
		}{w, w, w, w}
	// combination 35/512
	case !ok0 && ok1 && !ok2 && !ok3 && !ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.Queryer
		}{w, w, w, w}
	// combination 36/512
	case ok0 && ok1 && !ok2 && !ok3 && !ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Queryer
		}{w, w, w, w, w}
	// combination 37/512
	case !ok0 && !ok1 && ok2 && !ok3 && !ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.Queryer
		}{w, w, w, w}
	// combination 38/512
	case ok0 && !ok1 && ok2 && !ok3 && !ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Queryer
		}{w, w, w, w, w}
	// combination 39/512
	case !ok0 && ok1 && ok2 && !ok3 && !ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Queryer
		}{w, w, w, w, w}
	// combination 40/512
	case ok0 && ok1 && ok2 && !ok3 && !ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Queryer
		}{w, w, w, w, w, w}
	// combination 41/512
	case !ok0 && !ok1 && !ok2 && ok3 && !ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
			driver.Queryer
		}{w, w, w, w}
	// combination 42/512
	case ok0 && !ok1 && !ok2 && ok3 && !ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Queryer
		}{w, w, w, w, w}
	// combination 43/512
	case !ok0 && ok1 && !ok2 && ok3 && !ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
		}{w, w, w, w, w}
	// combination 44/512
	case ok0 && ok1 && !ok2 && ok3 && !ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
		}{w, w, w, w, w, w}
	// combination 45/512
	case !ok0 && !ok1 && ok2 && ok3 && !ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
		}{w, w, w, w, w}
	// combination 46/512
	case ok0 && !ok1 && ok2 && ok3 && !ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
		}{w, w, w, w, w, w}
	// combination 47/512
	case !ok0 && ok1 && ok2 && ok3 && !ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
		}{w, w, w, w, w, w}
	// combination 48/512
	case ok0 && ok1 && ok2 && ok3 && !ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
		}{w, w, w, w, w, w, w}
	// combination 49/512
	case !ok0 && !ok1 && !ok2 && !ok3 && ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Pinger
			driver.Queryer
		}{w, w, w, w}
	// combination 50/512
	case ok0 && !ok1 && !ok2 && !ok3 && ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.Queryer
		}{w, w, w, w, w}
	// combination 51/512
	case !ok0 && ok1 && !ok2 && !ok3 && ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.Pinger
			driver.Queryer
		}{w, w, w, w, w}
	// combination 52/512
	case ok0 && ok1 && !ok2 && !ok3 && ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
			driver.Queryer
		}{w, w, w, w, w, w}
	// combination 53/512
	case !ok0 && !ok1 && ok2 && !ok3 && ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
		}{w, w, w, w, w}
	// combination 54/512
	case ok0 && !ok1 && ok2 && !ok3 && ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
		}{w, w, w, w, w, w}
	// combination 55/512
	case !ok0 && ok1 && ok2 && !ok3 && ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
		}{w, w, w, w, w, w}
	// combination 56/512
	case ok0 && ok1 && ok2 && !ok3 && ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
		}{w, w, w, w, w, w, w}
	// combination 57/512
	case !ok0 && !ok1 && !ok2 && ok3 && ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
		}{w, w, w, w, w}
	// combination 58/512
	case ok0 && !ok1 && !ok2 && ok3 && ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
		}{w, w, w, w, w, w}
	// combination 59/512
	case !ok0 && ok1 && !ok2 && ok3 && ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
		}{w, w, w, w, w, w}
	// combination 60/512
	case ok0 && ok1 && !ok2 && ok3 && ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
		}{w, w, w, w, w, w, w}
	// combination 61/512
	case !ok0 && !ok1 && ok2 && ok3 && ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
		}{w, w, w, w, w, w}
	// combination 62/512
	case ok0 && !ok1 && ok2 && ok3 && ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
		}{w, w, w, w, w, w, w}
	// combination 63/512
	case !ok0 && ok1 && ok2 && ok3 && ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
		}{w, w, w, w, w, w, w}
	// combination 64/512
	case ok0 && ok1 && ok2 && ok3 && ok4 && ok5 && !ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
		}{w, w, w, w, w, w, w, w}
	// combination 65/512
	case !ok0 && !ok1 && !ok2 && !ok3 && !ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.QueryerContext
		}{w, w, w}
	// combination 66/512
	case ok0 && !ok1 && !ok2 && !ok3 && !ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.QueryerContext
		}{w, w, w, w}
	// combination 67/512
	case !ok0 && ok1 && !ok2 && !ok3 && !ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.QueryerContext
		}{w, w, w, w}
	// combination 68/512
	case ok0 && ok1 && !ok2 && !ok3 && !ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.QueryerContext
		}{w, w, w, w, w}
	// combination 69/512
	case !ok0 && !ok1 && ok2 && !ok3 && !ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.QueryerContext
		}{w, w, w, w}
	// combination 70/512
	case ok0 && !ok1 && ok2 && !ok3 && !ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.QueryerContext
		}{w, w, w, w, w}
	// combination 71/512
	case !ok0 && ok1 && ok2 && !ok3 && !ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.QueryerContext
		}{w, w, w, w, w}
	// combination 72/512
	case ok0 && ok1 && ok2 && !ok3 && !ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.QueryerContext
		}{w, w, w, w, w, w}
	// combination 73/512
	case !ok0 && !ok1 && !ok2 && ok3 && !ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
			driver.QueryerContext
		}{w, w, w, w}
	// combination 74/512
	case ok0 && !ok1 && !ok2 && ok3 && !ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.QueryerContext
		}{w, w, w, w, w}
	// combination 75/512
	case !ok0 && ok1 && !ok2 && ok3 && !ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.QueryerContext
		}{w, w, w, w, w}
	// combination 76/512
	case ok0 && ok1 && !ok2 && ok3 && !ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.QueryerContext
		}{w, w, w, w, w, w}
	// combination 77/512
	case !ok0 && !ok1 && ok2 && ok3 && !ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
		}{w, w, w, w, w}
	// combination 78/512
	case ok0 && !ok1 && ok2 && ok3 && !ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
		}{w, w, w, w, w, w}
	// combination 79/512
	case !ok0 && ok1 && ok2 && ok3 && !ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
		}{w, w, w, w, w, w}
	// combination 80/512
	case ok0 && ok1 && ok2 && ok3 && !ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
		}{w, w, w, w, w, w, w}
	// combination 81/512
	case !ok0 && !ok1 && !ok2 && !ok3 && ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Pinger
			driver.QueryerContext
		}{w, w, w, w}
	// combination 82/512
	case ok0 && !ok1 && !ok2 && !ok3 && ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.QueryerContext
		}{w, w, w, w, w}
	// combination 83/512
	case !ok0 && ok1 && !ok2 && !ok3 && ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.Pinger
			driver.QueryerContext
		}{w, w, w, w, w}
	// combination 84/512
	case ok0 && ok1 && !ok2 && !ok3 && ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
			driver.QueryerContext
		}{w, w, w, w, w, w}
	// combination 85/512
	case !ok0 && !ok1 && ok2 && !ok3 && ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
		}{w, w, w, w, w}
	// combination 86/512
	case ok0 && !ok1 && ok2 && !ok3 && ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
		}{w, w, w, w, w, w}
	// combination 87/512
	case !ok0 && ok1 && ok2 && !ok3 && ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
		}{w, w, w, w, w, w}
	// combination 88/512
	case ok0 && ok1 && ok2 && !ok3 && ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
		}{w, w, w, w, w, w, w}
	// combination 89/512
	case !ok0 && !ok1 && !ok2 && ok3 && ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
		}{w, w, w, w, w}
	// combination 90/512
	case ok0 && !ok1 && !ok2 && ok3 && ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
		}{w, w, w, w, w, w}
	// combination 91/512
	case !ok0 && ok1 && !ok2 && ok3 && ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
		}{w, w, w, w, w, w}
	// combination 92/512
	case ok0 && ok1 && !ok2 && ok3 && ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
		}{w, w, w, w, w, w, w}
	// combination 93/512
	case !ok0 && !ok1 && ok2 && ok3 && ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
		}{w, w, w, w, w, w}
	// combination 94/512
	case ok0 && !ok1 && ok2 && ok3 && ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
		}{w, w, w, w, w, w, w}
	// combination 95/512
	case !ok0 && ok1 && ok2 && ok3 && ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
		}{w, w, w, w, w, w, w}
	// combination 96/512
	case ok0 && ok1 && ok2 && ok3 && ok4 && !ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
		}{w, w, w, w, w, w, w, w}
	// combination 97/512
	case !ok0 && !ok1 && !ok2 && !ok3 && !ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w}
	// combination 98/512
	case ok0 && !ok1 && !ok2 && !ok3 && !ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w, w}
	// combination 99/512
	case !ok0 && ok1 && !ok2 && !ok3 && !ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w, w}
	// combination 100/512
	case ok0 && ok1 && !ok2 && !ok3 && !ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w, w, w}
	// combination 101/512
	case !ok0 && !ok1 && ok2 && !ok3 && !ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w, w}
	// combination 102/512
	case ok0 && !ok1 && ok2 && !ok3 && !ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w, w, w}
	// combination 103/512
	case !ok0 && ok1 && ok2 && !ok3 && !ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w, w, w}
	// combination 104/512
	case ok0 && ok1 && ok2 && !ok3 && !ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w, w, w, w}
	// combination 105/512
	case !ok0 && !ok1 && !ok2 && ok3 && !ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w, w}
	// combination 106/512
	case ok0 && !ok1 && !ok2 && ok3 && !ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w, w, w}
	// combination 107/512
	case !ok0 && ok1 && !ok2 && ok3 && !ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w, w, w}
	// combination 108/512
	case ok0 && ok1 && !ok2 && ok3 && !ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w, w, w, w}
	// combination 109/512
	case !ok0 && !ok1 && ok2 && ok3 && !ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w, w, w}
	// combination 110/512
	case ok0 && !ok1 && ok2 && ok3 && !ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w, w, w, w}
	// combination 111/512
	case !ok0 && ok1 && ok2 && ok3 && !ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w, w, w, w}
	// combination 112/512
	case ok0 && ok1 && ok2 && ok3 && !ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w, w, w, w, w}
	// combination 113/512
	case !ok0 && !ok1 && !ok2 && !ok3 && ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w, w}
	// combination 114/512
	case ok0 && !ok1 && !ok2 && !ok3 && ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w, w, w}
	// combination 115/512
	case !ok0 && ok1 && !ok2 && !ok3 && ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w, w, w}
	// combination 116/512
	case ok0 && ok1 && !ok2 && !ok3 && ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w, w, w, w}
	// combination 117/512
	case !ok0 && !ok1 && ok2 && !ok3 && ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w, w, w}
	// combination 118/512
	case ok0 && !ok1 && ok2 && !ok3 && ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w, w, w, w}
	// combination 119/512
	case !ok0 && ok1 && ok2 && !ok3 && ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w, w, w, w}
	// combination 120/512
	case ok0 && ok1 && ok2 && !ok3 && ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w, w, w, w, w}
	// combination 121/512
	case !ok0 && !ok1 && !ok2 && ok3 && ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w, w, w}
	// combination 122/512
	case ok0 && !ok1 && !ok2 && ok3 && ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w, w, w, w}
	// combination 123/512
	case !ok0 && ok1 && !ok2 && ok3 && ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w, w, w, w}
	// combination 124/512
	case ok0 && ok1 && !ok2 && ok3 && ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w, w, w, w, w}
	// combination 125/512
	case !ok0 && !ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w, w, w, w}
	// combination 126/512
	case ok0 && !ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w, w, w, w, w}
	// combination 127/512
	case !ok0 && ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w, w, w, w, w}
	// combination 128/512
	case ok0 && ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && !ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
		}{w, w, w, w, w, w, w, w, w}
	// combination 129/512
	case !ok0 && !ok1 && !ok2 && !ok3 && !ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.SessionResetter
		}{w, w, w}
	// combination 130/512
	case ok0 && !ok1 && !ok2 && !ok3 && !ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.SessionResetter
		}{w, w, w, w}
	// combination 131/512
	case !ok0 && ok1 && !ok2 && !ok3 && !ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.SessionResetter
		}{w, w, w, w}
	// combination 132/512
	case ok0 && ok1 && !ok2 && !ok3 && !ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.SessionResetter
		}{w, w, w, w, w}
	// combination 133/512
	case !ok0 && !ok1 && ok2 && !ok3 && !ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.SessionResetter
		}{w, w, w, w}
	// combination 134/512
	case ok0 && !ok1 && ok2 && !ok3 && !ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.SessionResetter
		}{w, w, w, w, w}
	// combination 135/512
	case !ok0 && ok1 && ok2 && !ok3 && !ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.SessionResetter
		}{w, w, w, w, w}
	// combination 136/512
	case ok0 && ok1 && ok2 && !ok3 && !ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 137/512
	case !ok0 && !ok1 && !ok2 && ok3 && !ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
			driver.SessionResetter
		}{w, w, w, w}
	// combination 138/512
	case ok0 && !ok1 && !ok2 && ok3 && !ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.SessionResetter
		}{w, w, w, w, w}
	// combination 139/512
	case !ok0 && ok1 && !ok2 && ok3 && !ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.SessionResetter
		}{w, w, w, w, w}
	// combination 140/512
	case ok0 && ok1 && !ok2 && ok3 && !ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 141/512
	case !ok0 && !ok1 && ok2 && ok3 && !ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.SessionResetter
		}{w, w, w, w, w}
	// combination 142/512
	case ok0 && !ok1 && ok2 && ok3 && !ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 143/512
	case !ok0 && ok1 && ok2 && ok3 && !ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 144/512
	case ok0 && ok1 && ok2 && ok3 && !ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 145/512
	case !ok0 && !ok1 && !ok2 && !ok3 && ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Pinger
			driver.SessionResetter
		}{w, w, w, w}
	// combination 146/512
	case ok0 && !ok1 && !ok2 && !ok3 && ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.SessionResetter
		}{w, w, w, w, w}
	// combination 147/512
	case !ok0 && ok1 && !ok2 && !ok3 && ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.Pinger
			driver.SessionResetter
		}{w, w, w, w, w}
	// combination 148/512
	case ok0 && ok1 && !ok2 && !ok3 && ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 149/512
	case !ok0 && !ok1 && ok2 && !ok3 && ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.SessionResetter
		}{w, w, w, w, w}
	// combination 150/512
	case ok0 && !ok1 && ok2 && !ok3 && ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 151/512
	case !ok0 && ok1 && ok2 && !ok3 && ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 152/512
	case ok0 && ok1 && ok2 && !ok3 && ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 153/512
	case !ok0 && !ok1 && !ok2 && ok3 && ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
		}{w, w, w, w, w}
	// combination 154/512
	case ok0 && !ok1 && !ok2 && ok3 && ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 155/512
	case !ok0 && ok1 && !ok2 && ok3 && ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 156/512
	case ok0 && ok1 && !ok2 && ok3 && ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 157/512
	case !ok0 && !ok1 && ok2 && ok3 && ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 158/512
	case ok0 && !ok1 && ok2 && ok3 && ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 159/512
	case !ok0 && ok1 && ok2 && ok3 && ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 160/512
	case ok0 && ok1 && ok2 && ok3 && ok4 && !ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
		}{w, w, w, w, w, w, w, w}
	// combination 161/512
	case !ok0 && !ok1 && !ok2 && !ok3 && !ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w}
	// combination 162/512
	case ok0 && !ok1 && !ok2 && !ok3 && !ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w, w}
	// combination 163/512
	case !ok0 && ok1 && !ok2 && !ok3 && !ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w, w}
	// combination 164/512
	case ok0 && ok1 && !ok2 && !ok3 && !ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 165/512
	case !ok0 && !ok1 && ok2 && !ok3 && !ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w, w}
	// combination 166/512
	case ok0 && !ok1 && ok2 && !ok3 && !ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 167/512
	case !ok0 && ok1 && ok2 && !ok3 && !ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 168/512
	case ok0 && ok1 && ok2 && !ok3 && !ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 169/512
	case !ok0 && !ok1 && !ok2 && ok3 && !ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w, w}
	// combination 170/512
	case ok0 && !ok1 && !ok2 && ok3 && !ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 171/512
	case !ok0 && ok1 && !ok2 && ok3 && !ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 172/512
	case ok0 && ok1 && !ok2 && ok3 && !ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 173/512
	case !ok0 && !ok1 && ok2 && ok3 && !ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 174/512
	case ok0 && !ok1 && ok2 && ok3 && !ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 175/512
	case !ok0 && ok1 && ok2 && ok3 && !ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 176/512
	case ok0 && ok1 && ok2 && ok3 && !ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w, w, w, w, w}
	// combination 177/512
	case !ok0 && !ok1 && !ok2 && !ok3 && ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w, w}
	// combination 178/512
	case ok0 && !ok1 && !ok2 && !ok3 && ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 179/512
	case !ok0 && ok1 && !ok2 && !ok3 && ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 180/512
	case ok0 && ok1 && !ok2 && !ok3 && ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 181/512
	case !ok0 && !ok1 && ok2 && !ok3 && ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 182/512
	case ok0 && !ok1 && ok2 && !ok3 && ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 183/512
	case !ok0 && ok1 && ok2 && !ok3 && ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 184/512
	case ok0 && ok1 && ok2 && !ok3 && ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w, w, w, w, w}
	// combination 185/512
	case !ok0 && !ok1 && !ok2 && ok3 && ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 186/512
	case ok0 && !ok1 && !ok2 && ok3 && ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 187/512
	case !ok0 && ok1 && !ok2 && ok3 && ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 188/512
	case ok0 && ok1 && !ok2 && ok3 && ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w, w, w, w, w}
	// combination 189/512
	case !ok0 && !ok1 && ok2 && ok3 && ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 190/512
	case ok0 && !ok1 && ok2 && ok3 && ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w, w, w, w, w}
	// combination 191/512
	case !ok0 && ok1 && ok2 && ok3 && ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w, w, w, w, w}
	// combination 192/512
	case ok0 && ok1 && ok2 && ok3 && ok4 && ok5 && !ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
		}{w, w, w, w, w, w, w, w, w}
	// combination 193/512
	case !ok0 && !ok1 && !ok2 && !ok3 && !ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w}
	// combination 194/512
	case ok0 && !ok1 && !ok2 && !ok3 && !ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w}
	// combination 195/512
	case !ok0 && ok1 && !ok2 && !ok3 && !ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w}
	// combination 196/512
	case ok0 && ok1 && !ok2 && !ok3 && !ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 197/512
	case !ok0 && !ok1 && ok2 && !ok3 && !ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w}
	// combination 198/512
	case ok0 && !ok1 && ok2 && !ok3 && !ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 199/512
	case !ok0 && ok1 && ok2 && !ok3 && !ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 200/512
	case ok0 && ok1 && ok2 && !ok3 && !ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 201/512
	case !ok0 && !ok1 && !ok2 && ok3 && !ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w}
	// combination 202/512
	case ok0 && !ok1 && !ok2 && ok3 && !ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 203/512
	case !ok0 && ok1 && !ok2 && ok3 && !ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 204/512
	case ok0 && ok1 && !ok2 && ok3 && !ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 205/512
	case !ok0 && !ok1 && ok2 && ok3 && !ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 206/512
	case ok0 && !ok1 && ok2 && ok3 && !ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 207/512
	case !ok0 && ok1 && ok2 && ok3 && !ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 208/512
	case ok0 && ok1 && ok2 && ok3 && !ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w, w}
	// combination 209/512
	case !ok0 && !ok1 && !ok2 && !ok3 && ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w}
	// combination 210/512
	case ok0 && !ok1 && !ok2 && !ok3 && ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 211/512
	case !ok0 && ok1 && !ok2 && !ok3 && ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 212/512
	case ok0 && ok1 && !ok2 && !ok3 && ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 213/512
	case !ok0 && !ok1 && ok2 && !ok3 && ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 214/512
	case ok0 && !ok1 && ok2 && !ok3 && ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 215/512
	case !ok0 && ok1 && ok2 && !ok3 && ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 216/512
	case ok0 && ok1 && ok2 && !ok3 && ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w, w}
	// combination 217/512
	case !ok0 && !ok1 && !ok2 && ok3 && ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 218/512
	case ok0 && !ok1 && !ok2 && ok3 && ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 219/512
	case !ok0 && ok1 && !ok2 && ok3 && ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 220/512
	case ok0 && ok1 && !ok2 && ok3 && ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w, w}
	// combination 221/512
	case !ok0 && !ok1 && ok2 && ok3 && ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 222/512
	case ok0 && !ok1 && ok2 && ok3 && ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w, w}
	// combination 223/512
	case !ok0 && ok1 && ok2 && ok3 && ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w, w}
	// combination 224/512
	case ok0 && ok1 && ok2 && ok3 && ok4 && !ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w, w, w}
	// combination 225/512
	case !ok0 && !ok1 && !ok2 && !ok3 && !ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w}
	// combination 226/512
	case ok0 && !ok1 && !ok2 && !ok3 && !ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 227/512
	case !ok0 && ok1 && !ok2 && !ok3 && !ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 228/512
	case ok0 && ok1 && !ok2 && !ok3 && !ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 229/512
	case !ok0 && !ok1 && ok2 && !ok3 && !ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 230/512
	case ok0 && !ok1 && ok2 && !ok3 && !ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 231/512
	case !ok0 && ok1 && ok2 && !ok3 && !ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 232/512
	case ok0 && ok1 && ok2 && !ok3 && !ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w, w}
	// combination 233/512
	case !ok0 && !ok1 && !ok2 && ok3 && !ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 234/512
	case ok0 && !ok1 && !ok2 && ok3 && !ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 235/512
	case !ok0 && ok1 && !ok2 && ok3 && !ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 236/512
	case ok0 && ok1 && !ok2 && ok3 && !ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w, w}
	// combination 237/512
	case !ok0 && !ok1 && ok2 && ok3 && !ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 238/512
	case ok0 && !ok1 && ok2 && ok3 && !ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w, w}
	// combination 239/512
	case !ok0 && ok1 && ok2 && ok3 && !ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w, w}
	// combination 240/512
	case ok0 && ok1 && ok2 && ok3 && !ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w, w, w}
	// combination 241/512
	case !ok0 && !ok1 && !ok2 && !ok3 && ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w}
	// combination 242/512
	case ok0 && !ok1 && !ok2 && !ok3 && ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 243/512
	case !ok0 && ok1 && !ok2 && !ok3 && ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 244/512
	case ok0 && ok1 && !ok2 && !ok3 && ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w, w}
	// combination 245/512
	case !ok0 && !ok1 && ok2 && !ok3 && ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 246/512
	case ok0 && !ok1 && ok2 && !ok3 && ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w, w}
	// combination 247/512
	case !ok0 && ok1 && ok2 && !ok3 && ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w, w}
	// combination 248/512
	case ok0 && ok1 && ok2 && !ok3 && ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w, w, w}
	// combination 249/512
	case !ok0 && !ok1 && !ok2 && ok3 && ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w}
	// combination 250/512
	case ok0 && !ok1 && !ok2 && ok3 && ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w, w}
	// combination 251/512
	case !ok0 && ok1 && !ok2 && ok3 && ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w, w}
	// combination 252/512
	case ok0 && ok1 && !ok2 && ok3 && ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w, w, w}
	// combination 253/512
	case !ok0 && !ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w, w}
	// combination 254/512
	case ok0 && !ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w, w, w}
	// combination 255/512
	case !ok0 && ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w, w, w}
	// combination 256/512
	case ok0 && ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && ok7 && !ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
		}{w, w, w, w, w, w, w, w, w, w}
	// combination 257/512
	case !ok0 && !ok1 && !ok2 && !ok3 && !ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Validator
		}{w, w, w}
	// combination 258/512
	case ok0 && !ok1 && !ok2 && !ok3 && !ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Validator
		}{w, w, w, w}
	// combination 259/512
	case !ok0 && ok1 && !ok2 && !ok3 && !ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.Validator
		}{w, w, w, w}
	// combination 260/512
	case ok0 && ok1 && !ok2 && !ok3 && !ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Validator
		}{w, w, w, w, w}
	// combination 261/512
	case !ok0 && !ok1 && ok2 && !ok3 && !ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.Validator
		}{w, w, w, w}
	// combination 262/512
	case ok0 && !ok1 && ok2 && !ok3 && !ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Validator
		}{w, w, w, w, w}
	// combination 263/512
	case !ok0 && ok1 && ok2 && !ok3 && !ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Validator
		}{w, w, w, w, w}
	// combination 264/512
	case ok0 && ok1 && ok2 && !ok3 && !ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 265/512
	case !ok0 && !ok1 && !ok2 && ok3 && !ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
			driver.Validator
		}{w, w, w, w}
	// combination 266/512
	case ok0 && !ok1 && !ok2 && ok3 && !ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Validator
		}{w, w, w, w, w}
	// combination 267/512
	case !ok0 && ok1 && !ok2 && ok3 && !ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Validator
		}{w, w, w, w, w}
	// combination 268/512
	case ok0 && ok1 && !ok2 && ok3 && !ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 269/512
	case !ok0 && !ok1 && ok2 && ok3 && !ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Validator
		}{w, w, w, w, w}
	// combination 270/512
	case ok0 && !ok1 && ok2 && ok3 && !ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 271/512
	case !ok0 && ok1 && ok2 && ok3 && !ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 272/512
	case ok0 && ok1 && ok2 && ok3 && !ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 273/512
	case !ok0 && !ok1 && !ok2 && !ok3 && ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Pinger
			driver.Validator
		}{w, w, w, w}
	// combination 274/512
	case ok0 && !ok1 && !ok2 && !ok3 && ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.Validator
		}{w, w, w, w, w}
	// combination 275/512
	case !ok0 && ok1 && !ok2 && !ok3 && ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.Pinger
			driver.Validator
		}{w, w, w, w, w}
	// combination 276/512
	case ok0 && ok1 && !ok2 && !ok3 && ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 277/512
	case !ok0 && !ok1 && ok2 && !ok3 && ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.Validator
		}{w, w, w, w, w}
	// combination 278/512
	case ok0 && !ok1 && ok2 && !ok3 && ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 279/512
	case !ok0 && ok1 && ok2 && !ok3 && ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 280/512
	case ok0 && ok1 && ok2 && !ok3 && ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 281/512
	case !ok0 && !ok1 && !ok2 && ok3 && ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.Validator
		}{w, w, w, w, w}
	// combination 282/512
	case ok0 && !ok1 && !ok2 && ok3 && ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 283/512
	case !ok0 && ok1 && !ok2 && ok3 && ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 284/512
	case ok0 && ok1 && !ok2 && ok3 && ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 285/512
	case !ok0 && !ok1 && ok2 && ok3 && ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 286/512
	case ok0 && !ok1 && ok2 && ok3 && ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 287/512
	case !ok0 && ok1 && ok2 && ok3 && ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 288/512
	case ok0 && ok1 && ok2 && ok3 && ok4 && !ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 289/512
	case !ok0 && !ok1 && !ok2 && !ok3 && !ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Queryer
			driver.Validator
		}{w, w, w, w}
	// combination 290/512
	case ok0 && !ok1 && !ok2 && !ok3 && !ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Queryer
			driver.Validator
		}{w, w, w, w, w}
	// combination 291/512
	case !ok0 && ok1 && !ok2 && !ok3 && !ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.Queryer
			driver.Validator
		}{w, w, w, w, w}
	// combination 292/512
	case ok0 && ok1 && !ok2 && !ok3 && !ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Queryer
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 293/512
	case !ok0 && !ok1 && ok2 && !ok3 && !ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.Queryer
			driver.Validator
		}{w, w, w, w, w}
	// combination 294/512
	case ok0 && !ok1 && ok2 && !ok3 && !ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Queryer
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 295/512
	case !ok0 && ok1 && ok2 && !ok3 && !ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Queryer
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 296/512
	case ok0 && ok1 && ok2 && !ok3 && !ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Queryer
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 297/512
	case !ok0 && !ok1 && !ok2 && ok3 && !ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
			driver.Queryer
			driver.Validator
		}{w, w, w, w, w}
	// combination 298/512
	case ok0 && !ok1 && !ok2 && ok3 && !ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Queryer
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 299/512
	case !ok0 && ok1 && !ok2 && ok3 && !ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 300/512
	case ok0 && ok1 && !ok2 && ok3 && !ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 301/512
	case !ok0 && !ok1 && ok2 && ok3 && !ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 302/512
	case ok0 && !ok1 && ok2 && ok3 && !ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 303/512
	case !ok0 && ok1 && ok2 && ok3 && !ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 304/512
	case ok0 && ok1 && ok2 && ok3 && !ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 305/512
	case !ok0 && !ok1 && !ok2 && !ok3 && ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{w, w, w, w, w}
	// combination 306/512
	case ok0 && !ok1 && !ok2 && !ok3 && ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 307/512
	case !ok0 && ok1 && !ok2 && !ok3 && ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 308/512
	case ok0 && ok1 && !ok2 && !ok3 && ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 309/512
	case !ok0 && !ok1 && ok2 && !ok3 && ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 310/512
	case ok0 && !ok1 && ok2 && !ok3 && ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 311/512
	case !ok0 && ok1 && ok2 && !ok3 && ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 312/512
	case ok0 && ok1 && ok2 && !ok3 && ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 313/512
	case !ok0 && !ok1 && !ok2 && ok3 && ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 314/512
	case ok0 && !ok1 && !ok2 && ok3 && ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 315/512
	case !ok0 && ok1 && !ok2 && ok3 && ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 316/512
	case ok0 && ok1 && !ok2 && ok3 && ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 317/512
	case !ok0 && !ok1 && ok2 && ok3 && ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 318/512
	case ok0 && !ok1 && ok2 && ok3 && ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 319/512
	case !ok0 && ok1 && ok2 && ok3 && ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 320/512
	case ok0 && ok1 && ok2 && ok3 && ok4 && ok5 && !ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.Validator
		}{w, w, w, w, w, w, w, w, w}
	// combination 321/512
	case !ok0 && !ok1 && !ok2 && !ok3 && !ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w}
	// combination 322/512
	case ok0 && !ok1 && !ok2 && !ok3 && !ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w}
	// combination 323/512
	case !ok0 && ok1 && !ok2 && !ok3 && !ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w}
	// combination 324/512
	case ok0 && ok1 && !ok2 && !ok3 && !ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 325/512
	case !ok0 && !ok1 && ok2 && !ok3 && !ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w}
	// combination 326/512
	case ok0 && !ok1 && ok2 && !ok3 && !ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 327/512
	case !ok0 && ok1 && ok2 && !ok3 && !ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 328/512
	case ok0 && ok1 && ok2 && !ok3 && !ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 329/512
	case !ok0 && !ok1 && !ok2 && ok3 && !ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w}
	// combination 330/512
	case ok0 && !ok1 && !ok2 && ok3 && !ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 331/512
	case !ok0 && ok1 && !ok2 && ok3 && !ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 332/512
	case ok0 && ok1 && !ok2 && ok3 && !ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 333/512
	case !ok0 && !ok1 && ok2 && ok3 && !ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 334/512
	case ok0 && !ok1 && ok2 && ok3 && !ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 335/512
	case !ok0 && ok1 && ok2 && ok3 && !ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 336/512
	case ok0 && ok1 && ok2 && ok3 && !ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 337/512
	case !ok0 && !ok1 && !ok2 && !ok3 && ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w}
	// combination 338/512
	case ok0 && !ok1 && !ok2 && !ok3 && ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 339/512
	case !ok0 && ok1 && !ok2 && !ok3 && ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 340/512
	case ok0 && ok1 && !ok2 && !ok3 && ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 341/512
	case !ok0 && !ok1 && ok2 && !ok3 && ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 342/512
	case ok0 && !ok1 && ok2 && !ok3 && ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 343/512
	case !ok0 && ok1 && ok2 && !ok3 && ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 344/512
	case ok0 && ok1 && ok2 && !ok3 && ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 345/512
	case !ok0 && !ok1 && !ok2 && ok3 && ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 346/512
	case ok0 && !ok1 && !ok2 && ok3 && ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 347/512
	case !ok0 && ok1 && !ok2 && ok3 && ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 348/512
	case ok0 && ok1 && !ok2 && ok3 && ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 349/512
	case !ok0 && !ok1 && ok2 && ok3 && ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 350/512
	case ok0 && !ok1 && ok2 && ok3 && ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 351/512
	case !ok0 && ok1 && ok2 && ok3 && ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 352/512
	case ok0 && ok1 && ok2 && ok3 && ok4 && !ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w, w, w}
	// combination 353/512
	case !ok0 && !ok1 && !ok2 && !ok3 && !ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w}
	// combination 354/512
	case ok0 && !ok1 && !ok2 && !ok3 && !ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 355/512
	case !ok0 && ok1 && !ok2 && !ok3 && !ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 356/512
	case ok0 && ok1 && !ok2 && !ok3 && !ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 357/512
	case !ok0 && !ok1 && ok2 && !ok3 && !ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 358/512
	case ok0 && !ok1 && ok2 && !ok3 && !ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 359/512
	case !ok0 && ok1 && ok2 && !ok3 && !ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 360/512
	case ok0 && ok1 && ok2 && !ok3 && !ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 361/512
	case !ok0 && !ok1 && !ok2 && ok3 && !ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 362/512
	case ok0 && !ok1 && !ok2 && ok3 && !ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 363/512
	case !ok0 && ok1 && !ok2 && ok3 && !ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 364/512
	case ok0 && ok1 && !ok2 && ok3 && !ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 365/512
	case !ok0 && !ok1 && ok2 && ok3 && !ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 366/512
	case ok0 && !ok1 && ok2 && ok3 && !ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 367/512
	case !ok0 && ok1 && ok2 && ok3 && !ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 368/512
	case ok0 && ok1 && ok2 && ok3 && !ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w, w, w}
	// combination 369/512
	case !ok0 && !ok1 && !ok2 && !ok3 && ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 370/512
	case ok0 && !ok1 && !ok2 && !ok3 && ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 371/512
	case !ok0 && ok1 && !ok2 && !ok3 && ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 372/512
	case ok0 && ok1 && !ok2 && !ok3 && ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 373/512
	case !ok0 && !ok1 && ok2 && !ok3 && ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 374/512
	case ok0 && !ok1 && ok2 && !ok3 && ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 375/512
	case !ok0 && ok1 && ok2 && !ok3 && ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 376/512
	case ok0 && ok1 && ok2 && !ok3 && ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w, w, w}
	// combination 377/512
	case !ok0 && !ok1 && !ok2 && ok3 && ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 378/512
	case ok0 && !ok1 && !ok2 && ok3 && ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 379/512
	case !ok0 && ok1 && !ok2 && ok3 && ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 380/512
	case ok0 && ok1 && !ok2 && ok3 && ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w, w, w}
	// combination 381/512
	case !ok0 && !ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 382/512
	case ok0 && !ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w, w, w}
	// combination 383/512
	case !ok0 && ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w, w, w}
	// combination 384/512
	case ok0 && ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && !ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.Validator
		}{w, w, w, w, w, w, w, w, w, w}
	// combination 385/512
	case !ok0 && !ok1 && !ok2 && !ok3 && !ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w}
	// combination 386/512
	case ok0 && !ok1 && !ok2 && !ok3 && !ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w}
	// combination 387/512
	case !ok0 && ok1 && !ok2 && !ok3 && !ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w}
	// combination 388/512
	case ok0 && ok1 && !ok2 && !ok3 && !ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 389/512
	case !ok0 && !ok1 && ok2 && !ok3 && !ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w}
	// combination 390/512
	case ok0 && !ok1 && ok2 && !ok3 && !ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 391/512
	case !ok0 && ok1 && ok2 && !ok3 && !ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 392/512
	case ok0 && ok1 && ok2 && !ok3 && !ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 393/512
	case !ok0 && !ok1 && !ok2 && ok3 && !ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w}
	// combination 394/512
	case ok0 && !ok1 && !ok2 && ok3 && !ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 395/512
	case !ok0 && ok1 && !ok2 && ok3 && !ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 396/512
	case ok0 && ok1 && !ok2 && ok3 && !ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 397/512
	case !ok0 && !ok1 && ok2 && ok3 && !ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 398/512
	case ok0 && !ok1 && ok2 && ok3 && !ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 399/512
	case !ok0 && ok1 && ok2 && ok3 && !ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 400/512
	case ok0 && ok1 && ok2 && ok3 && !ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 401/512
	case !ok0 && !ok1 && !ok2 && !ok3 && ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w}
	// combination 402/512
	case ok0 && !ok1 && !ok2 && !ok3 && ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 403/512
	case !ok0 && ok1 && !ok2 && !ok3 && ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 404/512
	case ok0 && ok1 && !ok2 && !ok3 && ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 405/512
	case !ok0 && !ok1 && ok2 && !ok3 && ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 406/512
	case ok0 && !ok1 && ok2 && !ok3 && ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 407/512
	case !ok0 && ok1 && ok2 && !ok3 && ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 408/512
	case ok0 && ok1 && ok2 && !ok3 && ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 409/512
	case !ok0 && !ok1 && !ok2 && ok3 && ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 410/512
	case ok0 && !ok1 && !ok2 && ok3 && ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 411/512
	case !ok0 && ok1 && !ok2 && ok3 && ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 412/512
	case ok0 && ok1 && !ok2 && ok3 && ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 413/512
	case !ok0 && !ok1 && ok2 && ok3 && ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 414/512
	case ok0 && !ok1 && ok2 && ok3 && ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 415/512
	case !ok0 && ok1 && ok2 && ok3 && ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 416/512
	case ok0 && ok1 && ok2 && ok3 && ok4 && !ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w, w}
	// combination 417/512
	case !ok0 && !ok1 && !ok2 && !ok3 && !ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w}
	// combination 418/512
	case ok0 && !ok1 && !ok2 && !ok3 && !ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 419/512
	case !ok0 && ok1 && !ok2 && !ok3 && !ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 420/512
	case ok0 && ok1 && !ok2 && !ok3 && !ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 421/512
	case !ok0 && !ok1 && ok2 && !ok3 && !ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 422/512
	case ok0 && !ok1 && ok2 && !ok3 && !ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 423/512
	case !ok0 && ok1 && ok2 && !ok3 && !ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 424/512
	case ok0 && ok1 && ok2 && !ok3 && !ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 425/512
	case !ok0 && !ok1 && !ok2 && ok3 && !ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 426/512
	case ok0 && !ok1 && !ok2 && ok3 && !ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 427/512
	case !ok0 && ok1 && !ok2 && ok3 && !ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 428/512
	case ok0 && ok1 && !ok2 && ok3 && !ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 429/512
	case !ok0 && !ok1 && ok2 && ok3 && !ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 430/512
	case ok0 && !ok1 && ok2 && ok3 && !ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 431/512
	case !ok0 && ok1 && ok2 && ok3 && !ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 432/512
	case ok0 && ok1 && ok2 && ok3 && !ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w, w}
	// combination 433/512
	case !ok0 && !ok1 && !ok2 && !ok3 && ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 434/512
	case ok0 && !ok1 && !ok2 && !ok3 && ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 435/512
	case !ok0 && ok1 && !ok2 && !ok3 && ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 436/512
	case ok0 && ok1 && !ok2 && !ok3 && ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 437/512
	case !ok0 && !ok1 && ok2 && !ok3 && ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 438/512
	case ok0 && !ok1 && ok2 && !ok3 && ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 439/512
	case !ok0 && ok1 && ok2 && !ok3 && ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 440/512
	case ok0 && ok1 && ok2 && !ok3 && ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w, w}
	// combination 441/512
	case !ok0 && !ok1 && !ok2 && ok3 && ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 442/512
	case ok0 && !ok1 && !ok2 && ok3 && ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 443/512
	case !ok0 && ok1 && !ok2 && ok3 && ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 444/512
	case ok0 && ok1 && !ok2 && ok3 && ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w, w}
	// combination 445/512
	case !ok0 && !ok1 && ok2 && ok3 && ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 446/512
	case ok0 && !ok1 && ok2 && ok3 && ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w, w}
	// combination 447/512
	case !ok0 && ok1 && ok2 && ok3 && ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w, w}
	// combination 448/512
	case ok0 && ok1 && ok2 && ok3 && ok4 && ok5 && !ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w, w, w}
	// combination 449/512
	case !ok0 && !ok1 && !ok2 && !ok3 && !ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w}
	// combination 450/512
	case ok0 && !ok1 && !ok2 && !ok3 && !ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 451/512
	case !ok0 && ok1 && !ok2 && !ok3 && !ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 452/512
	case ok0 && ok1 && !ok2 && !ok3 && !ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 453/512
	case !ok0 && !ok1 && ok2 && !ok3 && !ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 454/512
	case ok0 && !ok1 && ok2 && !ok3 && !ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 455/512
	case !ok0 && ok1 && ok2 && !ok3 && !ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 456/512
	case ok0 && ok1 && ok2 && !ok3 && !ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 457/512
	case !ok0 && !ok1 && !ok2 && ok3 && !ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 458/512
	case ok0 && !ok1 && !ok2 && ok3 && !ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 459/512
	case !ok0 && ok1 && !ok2 && ok3 && !ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 460/512
	case ok0 && ok1 && !ok2 && ok3 && !ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 461/512
	case !ok0 && !ok1 && ok2 && ok3 && !ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 462/512
	case ok0 && !ok1 && ok2 && ok3 && !ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 463/512
	case !ok0 && ok1 && ok2 && ok3 && !ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 464/512
	case ok0 && ok1 && ok2 && ok3 && !ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w, w}
	// combination 465/512
	case !ok0 && !ok1 && !ok2 && !ok3 && ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 466/512
	case ok0 && !ok1 && !ok2 && !ok3 && ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 467/512
	case !ok0 && ok1 && !ok2 && !ok3 && ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 468/512
	case ok0 && ok1 && !ok2 && !ok3 && ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 469/512
	case !ok0 && !ok1 && ok2 && !ok3 && ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 470/512
	case ok0 && !ok1 && ok2 && !ok3 && ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 471/512
	case !ok0 && ok1 && ok2 && !ok3 && ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 472/512
	case ok0 && ok1 && ok2 && !ok3 && ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w, w}
	// combination 473/512
	case !ok0 && !ok1 && !ok2 && ok3 && ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 474/512
	case ok0 && !ok1 && !ok2 && ok3 && ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 475/512
	case !ok0 && ok1 && !ok2 && ok3 && ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 476/512
	case ok0 && ok1 && !ok2 && ok3 && ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w, w}
	// combination 477/512
	case !ok0 && !ok1 && ok2 && ok3 && ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 478/512
	case ok0 && !ok1 && ok2 && ok3 && ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w, w}
	// combination 479/512
	case !ok0 && ok1 && ok2 && ok3 && ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w, w}
	// combination 480/512
	case ok0 && ok1 && ok2 && ok3 && ok4 && !ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w, w, w}
	// combination 481/512
	case !ok0 && !ok1 && !ok2 && !ok3 && !ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w}
	// combination 482/512
	case ok0 && !ok1 && !ok2 && !ok3 && !ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 483/512
	case !ok0 && ok1 && !ok2 && !ok3 && !ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 484/512
	case ok0 && ok1 && !ok2 && !ok3 && !ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 485/512
	case !ok0 && !ok1 && ok2 && !ok3 && !ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 486/512
	case ok0 && !ok1 && ok2 && !ok3 && !ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 487/512
	case !ok0 && ok1 && ok2 && !ok3 && !ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 488/512
	case ok0 && ok1 && ok2 && !ok3 && !ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w, w}
	// combination 489/512
	case !ok0 && !ok1 && !ok2 && ok3 && !ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 490/512
	case ok0 && !ok1 && !ok2 && ok3 && !ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 491/512
	case !ok0 && ok1 && !ok2 && ok3 && !ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 492/512
	case ok0 && ok1 && !ok2 && ok3 && !ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w, w}
	// combination 493/512
	case !ok0 && !ok1 && ok2 && ok3 && !ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 494/512
	case ok0 && !ok1 && ok2 && ok3 && !ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w, w}
	// combination 495/512
	case !ok0 && ok1 && ok2 && ok3 && !ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w, w}
	// combination 496/512
	case ok0 && ok1 && ok2 && ok3 && !ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w, w, w}
	// combination 497/512
	case !ok0 && !ok1 && !ok2 && !ok3 && ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w}
	// combination 498/512
	case ok0 && !ok1 && !ok2 && !ok3 && ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 499/512
	case !ok0 && ok1 && !ok2 && !ok3 && ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 500/512
	case ok0 && ok1 && !ok2 && !ok3 && ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w, w}
	// combination 501/512
	case !ok0 && !ok1 && ok2 && !ok3 && ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 502/512
	case ok0 && !ok1 && ok2 && !ok3 && ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w, w}
	// combination 503/512
	case !ok0 && ok1 && ok2 && !ok3 && ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w, w}
	// combination 504/512
	case ok0 && ok1 && ok2 && !ok3 && ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.ExecerContext
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w, w, w}
	// combination 505/512
	case !ok0 && !ok1 && !ok2 && ok3 && ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w}
	// combination 506/512
	case ok0 && !ok1 && !ok2 && ok3 && ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w, w}
	// combination 507/512
	case !ok0 && ok1 && !ok2 && ok3 && ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w, w}
	// combination 508/512
	case ok0 && ok1 && !ok2 && ok3 && ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.Execer
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w, w, w}
	// combination 509/512
	case !ok0 && !ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w, w}
	// combination 510/512
	case ok0 && !ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.ConnBeginTx
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w, w, w}
	// combination 511/512
	case !ok0 && ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
			driver.Conn
			driver.Execer
			driver.ExecerContext
			driver.NamedValueChecker
			driver.Pinger
			driver.Queryer
			driver.QueryerContext
			driver.SessionResetter
			driver.Validator
		}{w, w, w, w, w, w, w, w, w, w}
	// combination 512/512
	case ok0 && ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && ok7 && ok8:
		return struct {
			ConnUnwrapper
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
		}{w, w, w, w, w, w, w, w, w, w, w}
	}
	panic("unreachable")
}
