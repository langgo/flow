package flow

type Context interface {
}

// 怎么终止 通过返回错误
// TODO 怎么直接结束流程
type Handler interface {
	Handle(err error, ctx Context) (Context, error)
}
