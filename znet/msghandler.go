package znet

import (
	"errors"
	"fmt"
	"zinx-learning/utils"
	"zinx-learning/ziface"
)

type MsgHandler struct {
	Routers          map[uint32]ziface.IRouter // Key is MsgID
	WorkerPoolSize   uint32
	MaxWorkerTaskLen uint32
	TaskQueues       []chan ziface.IRequest
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

func (msgHandler *MsgHandler) SendToTaskQueue(request ziface.IRequest) {
	i := request.GetConnection().GetConnID() % msgHandler.WorkerPoolSize
	msgHandler.TaskQueues[i] <- request
}

func (msgHandler *MsgHandler) StartWorkPool() {
	fmt.Println("[MsgHandler] Start work pool")

	for i := 0; i < len(msgHandler.TaskQueues); i++ {
		msgHandler.TaskQueues[i] = make(chan ziface.IRequest, msgHandler.MaxWorkerTaskLen)
		for 
	}
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Routers:          make(map[uint32]ziface.IRouter),
		TaskQueues:       make([]chan ziface.IRequest, utils.GlobalObj.WorkPoolSize),
		WorkerPoolSize:   utils.GlobalObj.WorkPoolSize,
		MaxWorkerTaskLen: utils.GlobalObj.MaxWorkerTaskLen,
	}
}
