package main

import (
	"fmt"
	"io"
	"net"
	"zinx-learning/znet"
)

func main() {
	client, err := net.Dial("tcp4", "0.0.0.0:9547")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	for {
		dataPack := znet.DataPack{}
		data, err := dataPack.Pack(znet.NewMessage(1, []byte("Ping ping...")))
		if err != nil {
			fmt.Println("Failed to pack data:", err)
			break
		}

		if _, err = client.Write(data); err != nil {
			fmt.Println("Failed to write data:", err)
			break
		}

		headData := make([]byte, dataPack.GetHeadLen())
		if _, err = io.ReadFull(client, headData); err != nil {
			fmt.Println("Failed to read data:", err)
			break
		}

		msg, err := dataPack.Unpack(headData)
		if err != nil {
			fmt.Println("Failed to unpack data:", err)
			break
		}

		var receivedData []byte
		if msg.GetDataLen() > 0 {
			receivedData = make([]byte, msg.GetDataLen())
			if _, err = io.ReadFull(client, receivedData); err != nil {
				fmt.Println("Failed to read data:", err)
				break
			}
		}
		msg.SetData(receivedData)

		fmt.Println("MsgID:", msg.GetMsgID(), "MsgData:", string(msg.GetData()))
	}
}
