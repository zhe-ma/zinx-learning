package ziface

type IMsgHandler interface {
	HandleMsg(request IRequest)
	AddRouter(msgID uint32, router IRouter) error
}
