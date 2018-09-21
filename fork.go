package flow

type SwitchHandler func(ctx Context) (Handler, error)

type ForkFlow struct {
	switchFunc SwitchHandler
}

func (ff *ForkFlow) Handle(err error, ctx Context) (Context, error) {
	handler, err := ff.switchFunc(ctx)

	return handler.Handle(err, ctx)
}
