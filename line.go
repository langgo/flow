package flow

type LineFlow struct {
	handlers []Handler
}

func (lf *LineFlow) Handle(err error, ctx Context) (Context, error) {
	for _, handler := range lf.handlers {
		ctx, err = handler.Handle(err, ctx)
	}
	return ctx, err
}

func (lf *LineFlow) Push(handler Handler) {
	lf.handlers = append(lf.handlers, handler)
}
