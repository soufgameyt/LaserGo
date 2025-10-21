package core

import (
	"LaserGo/utils"
	"encoding/binary"
)

type Client struct{} // useless :grin:

func ReceiveMessage(data []byte, client *Client) {
	message_type := binary.BigEndian.Uint16(data[0:2])
	message_length := uint32(data[2])<<16 | uint32(data[3])<<8 | uint32(data[4])

	message_version := binary.BigEndian.Uint16(data[5:7])

	end := 7 + int(message_length)
	if end > len(data) {
		end = len(data)
	}

	message_buffer := data[7:end]

	utils.DebuggerInst.Debug("MessageManager.ReceiveMessage -> Received message of type:", message_type,
		"length:", message_length, "version:", message_version)

	createMessageByType(message_type, message_buffer, client)
}
