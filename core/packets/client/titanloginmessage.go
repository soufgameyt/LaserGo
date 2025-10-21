package client

import (
	"LaserGo/core/network"
	AuthenticationResponseMessage "LaserGo/core/packets/server"
	PiranhaMessage "LaserGo/core/piranhamessage"
	"LaserGo/utils"
)

type TitanLoginMessage struct {
	*PiranhaMessage.PiranhaMessage
}

func NewTitanLoginMessage(buf []byte, client *network.Client) *TitanLoginMessage {
	msg := &TitanLoginMessage{
		PiranhaMessage: PiranhaMessage.NewPiranhaMessage(buf, client, "TitanLoginMessage"),
	}
	msg.SetMessageType(10101)
	return msg
}

func (t *TitanLoginMessage) Decode() *TitanLoginMessage {
	t.GetClient().SetHighId(int(t.ReadInt())) // wheelchaired code
	t.GetClient().SetLowId(int(t.ReadInt()))  // wheelchaired code
	_ = t.ReadString()
	return t
}

func (t *TitanLoginMessage) Process() *TitanLoginMessage {
	utils.DebuggerInst.Debug(
		"highid: ", t.GetClient().GetHighId(),
		". lowid: ", t.GetClient().GetLowId(),
	)
	AuthenticationResponseMessage.NewAuthenticationResponseMessage([]byte{}, t.GetClient()).Encode().Send()
	return t
}
