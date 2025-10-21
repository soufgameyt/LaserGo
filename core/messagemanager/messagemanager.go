package messagemanager

import (
	"LaserGo/core/factory"
	"LaserGo/core/network"
	"LaserGo/utils"
	"encoding/binary"
)

func ReceiveMessage(data []byte, client *network.Client) {
	// Header -> *a2 = this[1] | (*this << 8); *a3 = (this[2] << 16) | (this[3] << 8) | this[4]; *a4 = this[6] | (this[5] << 8);
	// credit to Messaging::readHeader() xD
	message_type := binary.BigEndian.Uint16(data[0:2])
	message_length := uint32(data[2])<<16 | uint32(data[3])<<8 | uint32(data[4])

	message_version := binary.BigEndian.Uint16(data[5:7])

	end := 7 + int(message_length)
	if end > len(data) {
		end = len(data)
	}

	message_buffer := data[7:end]

	utils.DebuggerInst.Debug("MessageManager.ReceiveMessage -> Received message of type:", message_type, "length:", message_length, "version:", message_version)

	factory.CreateMessageByType(message_type, message_buffer, client)
}
