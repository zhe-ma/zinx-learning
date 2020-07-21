package znet

import "zinx-learning/ziface"

type BaseRouter struct {
}

func (br *BaseRouter) PreHandle(req ziface.IRequtest) {
}

func (br *BaseRouter) Handle(req ziface.IRequtest) {
}

func (br *BaseRouter) PostHandle(req ziface.IRequtest) {
}
