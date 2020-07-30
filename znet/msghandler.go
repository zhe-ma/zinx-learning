package znet

import (
	"errors"
	"fmt"
	"zinx-learning/ziface"
)

type MsgHandler struct {
	Routers map[uint32]ziface.IRouter // Key is MsgID
}

func (msgHandler *MsgHandler) HandleMsg(request ziface.IRequest) {
	router, ok := msgHandler.Routers[request.GetMsgID()]
	if !ok {
		fmt.Println("The msgID doesn't has router. MsgID:", request.GetMsgID())
		return
	}

	router.PreHandle(request)
	router.Handle(request)
	router.PostHandle(request)
}

func (msgHandler *MsgHandler) AddRouter(msgID uint32, router ziface.IRouter) error {
	if _, ok := msgHandler.Routers[msgID]; ok {
		fmt.Println("The msgID alreay has router.")
		return errors.New("The msgID alreay has router")
	}

	msgHandler.Routers[msgID] = router

	return nil
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Routers: make(map[uint32]ziface.IRouter),
	}
}
