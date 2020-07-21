package ziface

type IRouter interface {
	PreHandle(req IRequtest)
	Handle(req IRequtest)
	PostHandle(req IRequtest)
}
