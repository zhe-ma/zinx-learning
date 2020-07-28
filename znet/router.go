package znet

import "zinx-learning/ziface"

type BaseRouter struct {
}

func (br *BaseRouter) PreHandle(req ziface.IRequest) {
}

func (br *BaseRouter) Handle(req ziface.IRequest) {
}

func (br *BaseRouter) PostHandle(req ziface.IRequest) {
}
