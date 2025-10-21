package factory

import (
	"LaserGo/core/network"
	"LaserGo/core/packets/client"
	"LaserGo/core/packets/server"
	"LaserGo/utils"
)

type LogicLaserMessageFactory struct{}

func CreateMessageByType(messageType uint16, buf []byte, clientinst *network.Client) {
	switch messageType {
	case 10101:
		clientMsg := client.NewTitanLoginMessage(buf, clientinst)
		clientMsg.Decode().Process()
	case 20104:
		resp := server.NewAuthenticationResponseMessage(buf, clientinst)
		resp.Encode().Send()
	default:
		utils.DebuggerInst.Warn("No message for type ", messageType, ". Skip.")
	}
}
