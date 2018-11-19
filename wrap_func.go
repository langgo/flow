package flow

type HandleFunc func(err error, ctx Context) (Context, error)

type wrapHandleFunc struct {
	fn HandleFunc
}

func (w *wrapHandleFunc) Handle(err error, ctx Context) (Context, error) {
	return w.fn(err, ctx)
}

func WrapFunc(f HandleFunc) Handler {
	return &wrapHandleFunc{fn: f}
}
